package main

import (
	"html/template"
	"net/http"
)

var testTemplate *template.Template

type widget struct {
	Name  string
	Price int
}

type viewData struct {
	Name    string
	Widgets []widget
}

func main() {
	var err error
	testTemplate, err = template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	vd := viewData{
		Name: "Christian",
		Widgets: []widget{
			{
				Name:  "Butter",
				Price: 1,
			},
			{
				Name:  "Eistee",
				Price: 2,
			},
			{
				Name:  "Honig",
				Price: 5,
			},
		},
	}
	err := testTemplate.Execute(w, vd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
