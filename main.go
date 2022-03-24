package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Book 示範的書本結構
type Book struct {
	Title  string
	Author string
	Year   int `json:"Released Year"`
	Position string `json:"Sale Position,omitempty"`
}

func main() {
	var books = getExampleBooks()

	// 將 go 結構轉換為 json 格式
	data, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	fmt.Printf("%s\n", data)
}

// getBooks 取得範例 Book 結構 slice
func getExampleBooks() []Book {
	return []Book{
		{
			Title:  "Hello Golang",
			Author: "Jack",
			Year:   1998,
			Position: "Book world - F2",
		},
		{
			Title:  "Time is money",
			Author: "Bob",
			Year:   1970,
		},
		{
			Title:  "The three little pig",
			Author: "Sam",
			Year:   1965,
		},
	}
}
