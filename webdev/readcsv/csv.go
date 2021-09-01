package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("csv.gohtml"))
}

type CsvMadness struct {
	Data [][]string
	Length int
}
func main() {
	csvGrab, err1 := readCSV("table.csv")
	if err1 != nil {
		log.Fatalf("func csvData failed with error: %v", err1)
	}

	csvLen := len(csvGrab)

	csvGrab = csvGrab[1:csvLen]

	csvData := CsvMadness{
		Data: csvGrab,
		Length: csvLen,
	}



	fmt.Println(csvLen)

	err2 := tpl.Execute(os.Stdout, csvData)
	if err2 != nil {
		log.Fatalln(err2)
	}

}

func readCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
	log.Fatalf("Either incorrect filename or file does not exist...\n%v\nCheck current path?", err)
 }
	r := csv.NewReader(file)
	r.Comma = ','
	r.Comment = '#'
	readAll, err := r.ReadAll()
	if err != nil {
		log.Fatalf("Unable to parse CSV, check delimeters and formatting?\nSee new error: %v", err)
	}

	return readAll, nil
}