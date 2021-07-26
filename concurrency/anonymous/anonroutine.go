package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
func main(){

	wg.Add(2)

	go func() {
		fmt.Println("--First anonymous go routine running--")
		time.Sleep(5 * time.Second)
		defer wg.Done()
		fmt.Println("--First anonymous go routine finished--")
	}()
	go func() {
		fmt.Println("--Second anonymous go routine running--")
		time.Sleep(5 * time.Second)
		wg.Done()
		fmt.Println("--Second anonymous go routine finished--")
	}()
	wg.Wait()

	fmt.Println("Anonymous go routine demo finished...exiting program...")
}
