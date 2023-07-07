package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	for i := 0; i < 4; i++ {
		wg.Add(2) // группа будет состоять из 2 горутин
		go func(x int) {
			defer wg.Done()
			fmt.Printf("START %d \n", x)
			time.Sleep(2 * time.Second)
			fmt.Printf("END %d\n", x)
		}(i)

		go func(x int) {
			defer wg.Done()
			fmt.Printf("START %d \n", x)
			time.Sleep(2 * time.Second)
			fmt.Printf("END %d\n", x)
		}(i)
	}

	wg.Wait() // это блокирует одну горутину, пока не выполнится первая 4 раза в (моём случае)
}
