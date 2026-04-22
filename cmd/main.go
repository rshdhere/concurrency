package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("A")
	go fmt.Println("B")
	go fmt.Println("C")

	time.Sleep(time.Second)
}
