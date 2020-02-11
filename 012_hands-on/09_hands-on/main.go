package main

import (
    "bufio"
    "encoding/csv"
    "io"
    "log"
	"os"
	"text/template"

)

type data struct{
Date string
Volume string
}

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	var datas []data
	csvFile, _ := os.Open("table.csv")
    reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
        line, error := reader.Read()
        if error == io.EOF {
            break
        } else if error != nil {
            log.Fatal(error)
        }
        datas = append(datas, data{
            Date: line[0],
            Volume:  line[1],
            },
        )
	}
	
	err := tpl.Execute(os.Stdout, datas)
	if err != nil {
		log.Fatalln(err)
	}


}