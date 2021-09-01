/*
Demo for how to grab fields from tcp connection using go and net package
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go connect(conn)
	}
}

func connect(conn net.Conn) {
	defer conn.Close()

	//here we will read the request for the server
	request(conn)

	//here we will return the request for the server
	respond(conn)
}

func request(conn net.Conn) {
	i := 0

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			//hitting first line in the request
			url := strings.Fields(ln)[1]
			fmt.Println("***URL:", url)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}

}

func respond(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World, Grab the URL!</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
