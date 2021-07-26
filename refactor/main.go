package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)
var files string

func main() {
	readDir()
}

func readDir() {
	err := filepath.Walk("./test", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		files = path
		//fmt.Println(files)

		//test2 := filepath.Ext(".txt")
		//fmt.Println(test2)

		test := strings.Contains(files, filepath.Ext(".txt"))
		//fmt.Println(test)

		if test == true {
			fmt.Println(files)
		}else if test == false {
			fmt.Println("No files match")
		}
		//fmt.Println(path, info.Mode())
		return nil

	})
	if err != nil {
		log.Fatalf("Unable to process directory, see error: %v", err)
	}
}