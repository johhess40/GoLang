package main

import (
	"os"
	"strings"
	"text/template"
)

var temp *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

type Cars struct {
	Name         string
	Manufacturer string
	Year         string
}

type Inventory struct {
	Items []Cars
}

func init() {
	temp = template.Must(template.New("").Funcs(fm).ParseFiles("./templates/funky.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
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
	err := temp.ExecuteTemplate(os.Stdout, "funky.gohtml", data)
	if err != nil {
		return
	}
}
