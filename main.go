package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s)
	}

}

func main() {

	/*
		Sleep() передает контроль
	*/

	go say("hello")
	say("world")

	/*
		result: hello, world 5 раз выведет, но если уберу Sleep, то выведет только world
	*/

}
