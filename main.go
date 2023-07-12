package main

import (
	"fmt"
)

func main() {
	var src []int
	for x := 1; x <= 100; x++ {
		src = append(src, x)
	}

	t := append(src[:10], src[(len(src)-10):]...)

	var v []int
	for i := len(t); i >= t[0]; i-- {
		v = append(v, t[i-1])
	}

	fmt.Println(t)
	fmt.Println(v)
}
