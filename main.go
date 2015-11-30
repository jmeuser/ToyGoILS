// Package main implements a toy Integrated Library System.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Book struct {
	Title    string
	ISBN     string
	Library  string
	ReqCount int
}

type Catalogue struct {
	Name  string
	Count int
	Books []*Book
}

// save marshals .Books then writes to .Name + ".json"
func (c *Catalogue) save() error {
	data, err := json.Marshal(c.Books)
	if err != nil {
		return fmt.Errorf("JSON marshing failed: %s", err)
	}
	filename := c.Name + ".json"
	return ioutil.WriteFile(filename, data, 0600)
}

// loadCatalogue unmarshals data read from name.json into a Catalogue
func loadCatalogue(name string) (*Catalogue, error) {
	filename := name + ".json"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var books []*Book
	if err := json.Unmarshal(data, &books); err != nil {
		return nil, fmt.Errorf("JSON unmarshling failed: %s", err)
	}
	return &Catalogue{Name: name, Count: len(books), Books: books}, nil
}

func main() {
	b0 := Book{Title: "Authority and the Individual", ISBN: "9781134812271", Library: "Pembrook Public Library"}
	b1 := Book{Title: "The Principles of Mathematics", ISBN: "9780203864760", Library: "Pembrook Public Library"}
	books := []*Book{&b0, &b1}
	UniCat := Catalogue{Name: "UniCat"}
	UniCat.Books = books
	UniCat.save()
	for _, x := range UniCat.Books {
		fmt.Println(x)
	}
	b, _ := loadCatalogue("UniCat")
	for _, x := range b.Books {
		fmt.Println(x)
	}
}
