package main

import (
	"log"
	"os"
	"text/template"
)

var temp *template.Template

func init() {
	temp = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	//ParseGlob should be preferred
	//Execute all templates
	err := temp.ExecuteTemplate(os.Stdout, "nestprime.gohtml", 42)
	if err != nil {
		log.Fatalf("Could not pase templates!, see error: %v\n", err)
	}
}
