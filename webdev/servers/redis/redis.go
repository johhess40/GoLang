package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
		}
		//conn.Close()
		go handleScanner(conn)
	}
}

func handleScanner(conn net.Conn) {
	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)

		switch fs[0] {
		case "GET":
			k := fs[1]
			v := data[k]
			fmt.Fprintf(conn, "%s\n", v)
		case "SET":
			if len(fs) != 3 {
				fmt.Fprintln(conn, "Expected Value!")
			}
			k := fs[1]
			v := fs[2]
			data[k] = v
		case "DEL":
			k := fs[1]
			delete(data, k)
		default:
			fmt.Fprintln(conn, "Invalid command try again...")
		}
	}
	defer conn.Close()

	//How does my code actually get here?
	//I don't believe it is supposed to...
	fmt.Println("Our code actually reached this!")
}
