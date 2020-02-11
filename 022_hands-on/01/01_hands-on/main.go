package main

import (
	"net/http"
)

func home(res http.ResponseWriter, req *http.Request) {
}

func dog(res http.ResponseWriter, req *http.Request) {
}

func me(res http.ResponseWriter, req *http.Request) {
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/dog", home)
	http.HandleFunc("/me", home)

	http.ListenAndServe(":8080", nil)

}
