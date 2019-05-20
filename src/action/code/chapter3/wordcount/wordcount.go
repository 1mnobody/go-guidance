package main

import (
	"../words"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// edit configuration中增加了 Program arguments
	filename := os.Args[1]
	fmt.Println("filename: ", filename)
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("There was an error opening the file:", err)
		return
	}

	text := string(content)

	count := words.CountWords(text)
	fmt.Printf("There are %d words in your text. \n", count)
}
