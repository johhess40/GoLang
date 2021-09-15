package main

import (
	"fmt"
	"net/http"
)

type headers struct {
	HeaderOne map[string]string
	HeaderTwo map[string]string
}

var headless = headers{
	HeaderOne: map[string]string{
		"Johns-Key": "This is the first header sent from John",
	},
	HeaderTwo: map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	},
}

type resp int

func (z resp) ServeHTTP(res http.ResponseWriter, zip *http.Request) {

	for k, v := range headless.HeaderOne {
		res.Header().Set(k, v)
	}

	for k, v := range headless.HeaderTwo {
		res.Header().Set(k, v)
	}

	fmt.Fprintln(res, "<h1>This is how we know this code runs</h1>")
}

func main() {
	var x resp
	http.ListenAndServe(":8080", x)
}
