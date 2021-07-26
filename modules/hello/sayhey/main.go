package main

import (
	"fmt"
	"github.com/sayhello"
)

func main (){
	message := sayhello.Hello()

	fmt.Println(message)
}