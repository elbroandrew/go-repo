package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		message := <-ch
		fmt.Println(message)
		ch <- "pong"
	}()

	ch <- "ping"
	fmt.Println(<-ch)
}
