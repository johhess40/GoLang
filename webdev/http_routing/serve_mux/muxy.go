package main

import (
	"html/template"
	"io"
	"net/http"
)

var tpl *template.Template

type muxies struct {
	Paths        []string
	FirstReturn  func(w http.ResponseWriter, req *http.Request)
	SecondReturn func(w http.ResponseWriter, req *http.Request)
}

var mux muxies

func main() {

	mux = muxies{
		Paths: []string{
			"/states",
			"/countries",
			"/counties",
			"/population",
		},
		FirstReturn:  sendOne,
		SecondReturn: sendTwo,
	}
	//ranging over paths and matching those to what handlers we want to use
	for _, v := range mux.Paths {
		switch v {
		case mux.Paths[0]:
			http.HandleFunc(v, mux.FirstReturn)
		case mux.Paths[1]:
			http.HandleFunc(v, mux.FirstReturn)
		case mux.Paths[2]:
			http.HandleFunc(v, mux.SecondReturn)
		case mux.Paths[3]:
			http.HandleFunc(v, mux.SecondReturn)
		}
	}
	//http.HandleFunc("/states", x)
	//http.HandleFunc("/countries", x)
	//http.HandleFunc("/counties", y)
	//http.HandleFunc("/population", y)

	http.ListenAndServe(":8080", nil)

}

//func init() {
//	tpl = template.Must(template.ParseFiles("index.gohtml"))
//}

func sendOne(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/states":
		io.WriteString(w, "MA WA CA TX NY")
	case "/countries":
		io.WriteString(w, "USA AFG ENG GER RUS")
	}
	//using switch or if statements we can show how we do routing
	//at this path etc and this method etc lets run this code etc
}

func sendTwo(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/counties":
		io.WriteString(w, "Essex Milton Johnson Texarcana")
	case "/population":
		io.WriteString(w, "1Mil 2Mil 3Mil 4mil")
	}
}
