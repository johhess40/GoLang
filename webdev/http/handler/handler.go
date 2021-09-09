package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

//type Handler interface {
//	ServeHTTP(ResponseWriter, *request)
//}

type HandleType struct {
	Method string
	URL    string
	Data   string
}

func (meth HandleType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for {
		if meth.Method == "GET" {
			fmt.Fprintln(w, "The Method is GET")
		} else if meth.Method == "POST" {
			fmt.Fprintln(w, "The Method is POST")
		} else if meth.Method == "PUT" {
			fmt.Fprintln(w, "The Method is PUT")
		}
		fmt.Fprintln(w, meth.URL)
		fmt.Fprintln(w, meth.Data)
		return
		}
}

func main() {
	var send HandleType
	send = HandleType{
		Method: os.Args[1],
		URL:    os.Args[2],
		Data:   os.Args[3],
	}

	err := http.ListenAndServe(":8081", send)
	if err != nil {
		log.Fatal(err)
	}
}
