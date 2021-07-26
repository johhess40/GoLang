package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("Starting goroutine demonstration")

	primechan := make(chan string)
	secdchan := make(chan string)
	lastchan := make(chan string)

	wg.Add(3)
	go echoSomething(primechan)
	go echoAnother(secdchan)
	go echoAgain(lastchan)

	fmt.Println(<-primechan)
	fmt.Println(<-secdchan)
	fmt.Println(<-lastchan)
	wg.Wait()

}

func echoSomething(ch chan<- string) {

	fmt.Println("Executing the first function")
	time.Sleep(3 * time.Second)
	defer wg.Done()

	ch <- "First function has finished"
}

func echoAnother(ch chan<- string) {

	fmt.Println("Executing the second function")
	time.Sleep(7 * time.Second)
	defer wg.Done()

	ch <- "Second function has finished"
}

func echoAgain(ch chan<- string) {

	fmt.Println("Executing the third function")
	time.Sleep(3 * time.Second)
	defer wg.Done()

	ch <- "Third function has finished"
}
