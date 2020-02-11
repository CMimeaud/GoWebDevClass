package main

import (
	"log"
	"os"
	"text/template"
)

type Hotel struct {
Name, Address, City, Zip, Region string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
hotels := [] Hotel {
Hotel {
	Name : "Hotel California",
	Address: "5 rue des Tamaris",
	City: "San Francisco",
	Zip: "7878",
	Region: "Santa Fe",
},
Hotel {
	Name : "Ritz",
	Address: "Champs Elysee",
	City: "LA",
	Zip: "9678",
	Region: "South California",
},

}


	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}

}