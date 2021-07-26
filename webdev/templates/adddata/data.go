package main

import (
	"os"
	"text/template"
)

var temp *template.Template

type Cars struct {
	Name         string
	Manufacturer string
	Year         string
}

type Inventory struct {
	Items []Cars
}

func init() {
	temp = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	//ParseGlob should be preferred
	//Execute all templates

	a := Cars{
		Name:         "Mustang",
		Manufacturer: "Ford Motors",
		Year:         "1967",
	}
	b := Cars{
		Name:         "Delorean",
		Manufacturer: "Who cares",
		Year:         "1986",
	}

	favcars := []Cars{a, b}

	data := Inventory{
		Items: favcars,
	}
	err := temp.ExecuteTemplate(os.Stdout, "data.gohtml", data)
	if err != nil {
		return
	}
}
