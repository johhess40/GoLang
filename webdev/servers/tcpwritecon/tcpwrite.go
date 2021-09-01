package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
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
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say %s\n", ln)
	}
	defer conn.Close()

	//How does my code actually get here?
	//I don't believe it is supposed to...
	fmt.Println("Our code actually reached this!")
}
