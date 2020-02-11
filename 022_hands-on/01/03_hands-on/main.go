package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", bar)
	http.HandleFunc("/me/", myName)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	data := "This is foo"
	err := tpl.ExecuteTemplate(w, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}

func bar(w http.ResponseWriter, req *http.Request) {
	data := "This is bar"
	err := tpl.ExecuteTemplate(w, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}

func myName(w http.ResponseWriter, req *http.Request) {
	data := "This is MyName"
	err := tpl.ExecuteTemplate(w, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
