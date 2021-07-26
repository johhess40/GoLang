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

	playas := []string{"Michael Jordan", "Magic Johnson", "Larry Bird"}
	err := temp.ExecuteTemplate(os.Stdout, "data.gohtml", playas)
	if err != nil {
		return
	}
}
