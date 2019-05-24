package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	Contact struct {
		Name    string `json:"name"`
		Title   string `json:"title"`
		Contact struct {
			Home string `json:"home"`
			Cell string `json:"cell"`
		} `json:"contact"`
	}

	// *********************************** 定义了elasticsearch.json对应的json数据 ***********************************
	Shards struct {
		Failed     int `json:"failed"`
		Skipped    int `json:"skipped"`
		Successful int `json:"successful"`
		Total      int `json:"total"`
	}

	Entity struct {
		AppCode    string `json:"appCode"`
		DeployEnv  string `json:"deployEnv"`
		InstanceIp string `json:"instanceIp"`
		Level      string `json:"level"`
		LogDate    string `json:"logDate"`
	}

	SearchResult struct {
		Took    int    `json:"took"`
		Timeout bool   `json:"time_out"`
		Shards  Shards `json:"_shards"`
		Hits    struct {
			MaxScore float32 `json:"max_score"`
			Total    int     `json:"total"`
			HitList  []struct {
				Id     string  `json:"_id"`
				Index  string  `json:"_index"`
				Score  float32 `json:"_score"`
				Source Entity  `json:"_source"`
				Type   string  `json:"_type"`
			} `json:"hits"`
		}
	}
	// ************************************************************************************************************

)

var v = ` {
	"name" : "Gopher",
	"title": "Programmer",
	"contact": {
		"home": "415.333.3333",
		"cell": "415.555.3333"
	}
}
`

func main() {
	// decode the json string to struct
	var c Contact
	err := json.Unmarshal([]byte(v), &c)
	if err != nil {
		log.Println("Unmarshal json error:", err)
		return
	}
	fmt.Println(c)

	// decode the json string to map
	var m map[string]interface{}
	err = json.Unmarshal([]byte(v), &m)
	if err != nil {
		log.Println("Unmarshal json error:", err)
		return
	}
	fmt.Println("Name:", m["name"])
	fmt.Println("Home:", m["contact"].(map[string]interface{})["home"])
	fmt.Println("Cell:", m["contact"].(map[string]interface{})["cell"])

	// 通过搭建的es获取日志数据，decode 成 SearchResult struct
	uri := "http://10.142.232.116:8324/devops-log-test-2019-05-17/_search"
	resp, err := http.Get(uri)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	var sr SearchResult
	err = json.NewDecoder(resp.Body).Decode(&sr)
	if err != nil {
		log.Println("Error:", err)
		return
	}

	fmt.Println(sr)

	//pretty, err := json.MarshalIndent(sr, "", "  ")
	//if err != nil {
	//	log.Println("Error:", err)
	//	return
	//}

	//fmt.Println(string(pretty))

	// *************************************** 将struct转化成string ***************************************
	m2 := make(map[string]interface{})
	m2["name"] = "Java"
	m2["title"] = "Programmer"
	m2["contact"] = map[string]interface{}{
		"home": "123.1221.212121",
		"cell": "331.2132.132344",
	}
	data, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		fmt.Println("Marshal error:", err)
		return
	}
	fmt.Println(string(data))

}
