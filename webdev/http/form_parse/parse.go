package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type agent int

func (m agent) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method        string
		URL           *url.URL
		Header        http.Header
		Subs          map[string][]string
	}{
		req.Method,
		req.URL,
		req.Header,
		req.Form,
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var agents agent
	http.ListenAndServe(":8081", agents)
}
