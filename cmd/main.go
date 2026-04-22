package main

import (
	"fmt"
	"time"
)

func main() {
	for index := range 5 {
		go func() {
			fmt.Println(index)
		}()
	}
	time.Sleep(time.Second)
}
