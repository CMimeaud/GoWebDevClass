package main

	import (
		"log"
		"os"
		"text/template"
	)
type item struct {
Name, Description, Ingredients string
Price float64
}

type meal struct {
Name string
Items []item
}

type Meals []meal

	var tpl *template.Template


func init(){
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}



func main() {

h := Meals{
	meal{
Name: "lunch",
Items: []item{
item{
Name: "Frites",
Description: "Des frites francaises",
Ingredients: "Pommes de terre",
Price: 15,
},
item{
	Name: "Poulet fris",
	Description: "Du Poulet cuit au four",
	Ingredients: "un poulet entier",
	Price: 10,
	},
},
},
meal{
	Name: "lunch",
	Items: []item{
	item{
	Name: "Frites",
	Description: "Des frites francaises",
	Ingredients: "Pommes de terre",
	Price: 15,
	},
	item{
		Name: "Poulet fris",
		Description: "Du Poulet cuit au four",
		Ingredients: "un poulet entier",
		Price: 10,
		},
	},
},
}

err := tpl.Execute(os.Stdout, h)
if err != nil {
	log.Fatalln(err)
}
}