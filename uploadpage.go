package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func (r *Page) save() error {
	filename := r.Title + ".txt"
	fmt.Println("Saving the file with the name:", filename)
	return ioutil.WriteFile(filename, r.Body, 0600)
}

func loadPage(title string) (*Page, error) {

	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p := &Page{Title: "MyPage", Body: []byte("This is the page content")}
	fmt.Println(p.Title)
	fmt.Println(string(p.Body))
	p.Body = []byte("Adding more content to the body")
	fmt.Println(string(p.Body))
	filename := p.Title + ".txt"
	fmt.Println("Filename is :", filename)
	p.save()
	p2, _ := loadPage("MyPage")
	fmt.Println("Reading the file ....\n", string(p2.Body))
}
