package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println("task", n)
		}(i)
	}

	wg.Wait()
	fmt.Println("done")
}
