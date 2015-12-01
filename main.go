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
	Title string
	ISBN  string
	Lib   string
	Reqs  int
}

// Is it more efficient to use map[string]bool with search .Books method?
type Catalogue struct {
	Name   string
	Count  int
	Books  []*Book
	Titles map[string][]*Book
	ISBNs  map[string][]*Book
	Libs   map[string][]*Book
}

// save marshals .Books then writes to .Name + ".json"
func (c *Catalogue) save() error {
	data, err := json.MarshalIndent(c.Books, "", "	")
	if err != nil {
		return fmt.Errorf("JSON marshing failed: %s", err)
	}
	filename := "./cats/" + c.Name + ".json"
	return ioutil.WriteFile(filename, data, 0600)
}

func makeCatalogue(name string, books []*Book) *Catalogue {
	titles := make(map[string][]*Book)
	isbns := make(map[string][]*Book)
	libs := make(map[string][]*Book)
	for _, b := range books {
		titles[b.Title] = append(titles[b.Title], b)
		isbns[b.ISBN] = append(isbns[b.ISBN], b)
		libs[b.Lib] = append(libs[b.Lib], b)
	}
	return &Catalogue{Name: name, Count: len(books), Books: books, Titles: titles, ISBNs: isbns, Libs: libs}
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

	return makeCatalogue(name, books), nil
}

// templates cache
var templates = template.Must(template.ParseFiles("./tmpl/editBook.html", "./tmpl/viewBook.html", "./tmpl/viewCatalogue.html", "./tmpl/find.html"))

func renderBookTemplate(w http.ResponseWriter, tmpl string, b *Book) {
	err := templates.ExecuteTemplate(w, tmpl+".html", b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewBookHandler(w http.ResponseWriter, r *http.Request, b *Book) {
	renderBookTemplate(w, "viewBook", b) // place holder
}

func editBookHandler(fn func(http.ResponseWriter, *http.Request, *Book)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get book from catalogue from url
		// or return findBook (or refactor into find)
		fn(w, r, UniCat.Books[1]) // place holder
	}
}

func renderCatalogueTemplate(w http.ResponseWriter, tmpl string, c *Catalogue) {
	err := templates.ExecuteTemplate(w, tmpl+".html", c)
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

// findHandler presents the user with search options
func findHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "find.html", nil) // placeholder
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	isbn := r.FormValue("isbn")
	fmt.Println("\""+title+"\"", "\""+isbn+"\"")
	c := &UniCat // placeholder
	fmt.Println(*c)
	if title != "" {
		fmt.Println(c.Titles[title])
		c = makeCatalogue("Results", c.Titles[title])
	}
	if isbn != "" {
		c = makeCatalogue("Results", c.ISBNs[isbn])
	}
	fmt.Println(*c)
	renderCatalogueTemplate(w, "viewCatalogue", c)
}



var validPath = regexp.MustCompile("^/view?.*$") // placeholder

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[0]) // placeholder: send whole url onward
	}
}

var UniCat Catalogue

func main() {
	b0 := Book{Title: "Authority and the Individual", ISBN: "9781134812271", Lib: "Pembrook Public Library"}
	b1 := Book{Title: "The Principles of Mathematics", ISBN: "9780203864760", Lib: "Pembrook Public Library"}
	books := []*Book{&b0, &b1}
	UniCat = *makeCatalogue("UniCat", books)
	UniCat.save()
	http.HandleFunc("/view", makeHandler(viewHandler))
	http.HandleFunc("/find", findHandler)
	http.HandleFunc("/search", searchHandler)

	http.ListenAndServe(":8080", nil)
}
