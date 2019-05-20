package main

// sql包的使用
import (
	_ "./postgres"
	"database/sql"
)

func main() {
	sql.Open("postgres", "mydb")
}
