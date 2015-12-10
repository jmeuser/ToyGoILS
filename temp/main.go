package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os"
)

type Book struct {
	Lib, Title, ISBN string
	copies, requests int
}

func makeBook(lib, title, isbn string) *Book {
	return &Book{lib, title, isbn, 0, 0}
}

func (b *Book) String() string {
	s := "Title:    "+b.Title+"\n"
	s += "ISBN:     "+b.ISBN+"\n"
	s += "Library:  "+b.Lib+"\n"
	s += "Requests: " + fmt.Sprint(b.requests)+"\n"
	return s
}

type Catalog struct {
	fieldName string
	index     map[string][]*Book
}

func makeCatalog(fieldName string) *Catalog {
	return &Catalog{fieldName, make(map[string][]*Book)}
}

func (c *Catalog) addBook(b *Book) {
	switch c.fieldName {
		case "Lib":
			c.index[b.Lib] = append(c.index[b.Lib], b)
		case "Title":
			c.index[b.Title] = append(c.index[b.Title], b)
		case "ISBN":
			c.index[b.ISBN] = append(c.index[b.ISBN], b)
	}
}

func (c *Catalog) save() error {
	for fileName, books := range c.index {
		data, err := json.Marshal(books)
		if err != nil {
			return fmt.Errorf("Marshal fail: %s", err)
		}
		 filePath := "./" + c.fieldName + "/" + fileName + ".json"
		err = os.MkdirAll(c.fieldName, 0600)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(filePath, data, 0600)
		if err != nil {
			return err
		}
	}
	return nil
}

func loadCatalog(fieldName string) (*Catalog, error) {
	fileInfoList, err := ioutil.ReadDir(fieldName)
	if err != nil {
		return nil, err
	}
	c := makeCatalog(fieldName)
	for _, f := range fileInfoList {
		fileName := f.Name()
		filePath := "./" + fieldName + "/" + fileName + ".json"
		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			return nil, err
		}
		var books []*Book
		err = json.Unmarshal(data, &books)
		if err != nil {
			return nil, fmt.Errorf("Unmarshal fail: %s", err)
		}
		for _, book := range books {
			c.addBook(book)
		}
	}
	return c, nil
}

func (c *Catalog) String() string {
	s := "Cataloged by " + c.fieldName + "\n\n"
	for key, books := range c.index {
		s += key + "\n"
		s += "------------\n"
		for _, b := range books {
			s += b.String() + "\n"
		}
		s += "\n"
	}
	return s
}
func main() {
	b0 := makeBook("Pembrook Public Library", "Authority and the Individual", "9781134812271")
	b1 := makeBook("Pembrook Public Library", "The Principles of Mathematics", "9780203864760")
	c := makeCatalog("Lib")
	c.addBook(b0)
	c.addBook(b1)
	fmt.Println(c)
	fmt.Println(c.save())
}
