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
}

func request(conn net.Conn) {
	i := 0

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			//hitting first line in the request
			muxy(conn, ln)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}

}

func muxy(conn net.Conn, st string) {
	//the request line
	m := strings.Fields(st)[0] // method
	u := strings.Fields(st)[1] // url/uri

	fmt.Println("****METHOD", m)
	fmt.Println("****URL", u)

	if m == "GET" && u == "/" {
		index(conn)
	} else if m == "GET" && u == "/aboutus" {
		aboutus(conn)
	} else if m == "GET" && u == "/contactus" {
		contactus(conn)
	}else if m == "GET" && u == "/requests" {
		requests(conn)
	}else if m == "POST" && u == "/requests" {
		requestsSubmit(conn)
	}
}

//func respond(conn net.Conn) {
//	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World, Grab the URL!</strong></body></html>`
//
//	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
//	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
//	fmt.Fprint(conn, "Content-Type: text/html\r\n")
//	fmt.Fprint(conn, "\r\n")
//	fmt.Fprint(conn, body)
//}

func index(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>INDEX</strong><br>
	<a href="/">index</a><br>
	<a href="/aboutus">aboutus</a><br>
	<a href="/contactus">contactus</a><br>
	<a href="/requests">requests</a><br>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func aboutus(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>ABOUT</strong><br>
	<a href="/">index</a><br>
	<a href="/aboutus">aboutus</a><br>
	<a href="/contactus">contactus</a><br>
	<a href="/requests">requests</a><br>
	</body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func contactus(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>CONTACT</strong><br>
	<a href="/">index</a><br>
	<a href="/aboutus">aboutus</a><br>
	<a href="/contactus">contactus</a><br>
	<a href="/requests">requests</a><br>
	</body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
func requestsSubmit(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>REQUEST SUBMISSION</strong><br>
	<a href="/">index</a><br>
	<a href="/aboutus">aboutus</a><br>
	<a href="/contactus">contactus</a><br>
	<a href="/requests">requests</a><br>
	</body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func requests(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body>
	<strong>REQUESTS</strong><br>
	<a href="/">index</a><br>
	<a href="/aboutus">aboutus</a><br>
	<a href="/contactus">contactus</a><br>
	<a href="/requests">requests</a><br>
	<form method="POST" action="/requests">
	<input type="submit" value="requests">
	</form>
	</body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
