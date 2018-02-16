package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func() {
			fmt.Print(i) // HL
			wg.Done()
		}()
	}

	wg.Wait()
}
