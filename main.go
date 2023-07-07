package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	/*
		Канал ялвяется блокиратором и с ним можно выстроить так работу, чтоб горутина выполнялась в порядке как мне надо
	*/

	go func() {
		fmt.Println("START G")

	}()

	<-time.After(2 * time.Second)

}
