package main

import (
	"fmt"
	"sync"
)

const num = 500


func main() {

	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(num)

	value := 0
	for i := 0; i < num; i++ {
		go func(){
			defer wg.Done()

			mutex.Lock()
			value++
			mutex.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(value)

}
