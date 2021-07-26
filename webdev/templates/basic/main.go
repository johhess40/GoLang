package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main(){
	str := fmt.Sprint(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="utf-8">
		</head>
		<body>
		<h1>Welcome to the most basic of web pages!</h1>
		</body>
		</html>
`)
	newf, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer newf.Close()

	io.Copy(newf, strings.NewReader(str))
}
