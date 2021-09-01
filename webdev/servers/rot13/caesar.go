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
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := strings.ToLower(scanner.Text())
		bs := []byte(ln)
		r := rot13(bs)
		fmt.Fprintf(conn, "%s -- %s\n\n", ln, r)
	}
	defer conn.Close()

	//How does my code actually get here?
	//I don't believe it is supposed to...
	fmt.Println("Our code actually reached this!")
}

func rot13(bs []byte) []byte {
	var r13 = make([]byte, len(bs))
	for i, v := range bs {
		if v <= 109 {
			r13[i] = v + 13
		} else {
			r13[i] = v - 13
		}
	}

	return r13
}
