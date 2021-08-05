package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var temp *template.Template

func init() {
	temp = template.Must(template.New("").Funcs(fum).ParseFiles("./templates/pipe.gohtml"))
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

var fum = template.FuncMap{
	"fdateMDY": monthDayYear,
}

func main() {

	err := temp.ExecuteTemplate(os.Stdout, "pipe.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
