// Package main implements a toy Integrated Library System.
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
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
	filename := "./cats/" + c.Name + ".json"
	return ioutil.WriteFile(filename, data, 0600)
}

// loadCatalogue unmarshals data read from name.json into a Catalogue
func loadCatalogue(name string) (*Catalogue, error) {
	filename := "./cats/" + name + ".json"
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

// templates cache
var templates = template.Must(template.ParseFiles("./tmpl/editBook.html", "./tmpl/viewBook.html", "./tmpl/viewCatalogue.html"))

func renderBookTemplate(w http.ResponseWriter, tmpl string, b *Book) {
	err := templates.ExecuteTemplate(w, tmpl + ".html", b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewBookHandler(w http.ResponseWriter, r *http.Request, b *Book) {
	renderBookTemplate(w, "viewBook", b) // place holder
}

func makeBookHandler(fn func(http.ResponseWriter, *http.Request, *Book)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get book from catalogue from url
		// or return findBook (or refactor into find)
		fn(w, r, UniCat.Books[1]) // place holder
	}
}

func renderCatalogueTemplate(w http.ResponseWriter, tmpl string, c *Catalogue) {
	err := templates.ExecuteTemplate(w, tmpl + ".html", c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewCatalogueHandler(w http.ResponseWriter, r *http.Request, name string) {
	c, err := loadCatalogue(name)
	if err != nil {
		http.NotFound(w, r) // place holder: redirect to "makeCatalogue"?
		return
	}
	renderCatalogueTemplate(w, "viewCatalogue", c)
}

func viewHandler(w http.ResponseWriter, r *http.Request, url string) {
	v := r.URL.Query()
	b := v.Get("b") // get book query parameters
	c := v.Get("c") // get catalogue query parameters
	if b == "" && c != "" {
		viewCatalogueHandler(w, r, c)
	}
}

var validPath = regexp.MustCompile("^/view?.*$") // placeholder

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w,r)
			return
		}
		fn(w, r, m[0]) // placeholder: send whole url onward
	}
}


var UniCat Catalogue

func main() {
	b0 := Book{Title: "Authority and the Individual", ISBN: "9781134812271", Library: "Pembrook Public Library"}
	b1 := Book{Title: "The Principles of Mathematics", ISBN: "9780203864760", Library: "Pembrook Public Library"}
	books := []*Book{&b0, &b1}
	UniCat = Catalogue{Name: "UniCat"}
	UniCat.Books = books
	UniCat.save()
	http.HandleFunc("/view", makeHandler(viewHandler))

	http.ListenAndServe(":8080", nil)
}
