package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("hello from go-routine")
	time.Sleep(time.Second)
}
