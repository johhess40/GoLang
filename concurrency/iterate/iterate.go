package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var messages []string

func main() {

	messages = []string{"First Message", "Second Message", "Third Message"}

	for _, v := range messages {
		wg.Add(1)

		go func(v string) {

			fmt.Printf("%v\n", v)

			time.Sleep(1 * time.Second)

			wg.Done()
		}(v)

		wg.Wait()
	}
}
