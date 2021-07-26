package main

import (
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
	err := temp.ExecuteTemplate(os.Stdout, "more.gohtml", nil)
	if err != nil {
		return
	}

	err = temp.ExecuteTemplate(os.Stdout, "again.gohtml", nil)
	if err != nil {
		return
	}

	err = temp.ExecuteTemplate(os.Stdout, "package.gohtml", nil)
	if err != nil {
		return
	}

}
