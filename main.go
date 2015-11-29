// Package main implements a toy Integrated Library System.
package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Book struct {
	Title string
	ISBN string
	Library string
	ReqCount uint
}

// loadBooks unmarshals data read from books.json
func loadBooks() (*[]Book, error) {
	data, err := ioutil.ReadFile("books.json")
	if err != nil {
		return nil, err
	}
	var books []Book
	if err := json.Unmarshal(data,&books); err != nil {
		return nil, fmt.Errorf("JSON unmarshling failed: %s", err)
	}
	return &books, nil
}

// saveBooks marshals books then writes to books.json
func saveBooks(books *[]Book) error {
	data, err := json.Marshal(books)
	if err != nil {
		return fmt.Errorf("JSON marshing failed: %s", err)
	}
	return ioutil.WriteFile("books.json", data, 0600)
}

func main() {
	b0 := Book{Title: "Authority and the Individual", ISBN: "9781134812271", Library: "Pembrook Public Library"}
	b1 := Book{Title: "The Principles of Mathematics", ISBN: "9780203864760", Library: "Pembrook Public Library"}
	books := []Book{b0, b1}
	saveBooks(&books)
	b, _ := loadBooks()
	fmt.Println(books)
	fmt.Println(b)
}
