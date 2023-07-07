package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	/*
		В таком случае угорутины запускаются в случайном порядке.
	*/

	for i := 0; i < 4; i++ {
		wg.Add(1)

		go func(x int) {
			defer wg.Done()
			fmt.Printf("START %d \n", x)
			time.Sleep(1 * time.Second)
			fmt.Printf("END %d\n", x)
		}(i)
	}

	wg.Wait()
}
