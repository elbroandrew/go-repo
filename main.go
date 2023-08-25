package main

import (
	"bufio"
	"fmt"
	"os"
)

func Add(x int, y int) int {
	return x + y
}

func main() {

	// file, err := os.Open("data1.txt")
	// if err != nil {
	// 	fmt.Println("Unable to open file:", err)
	// 	return
	// }

	// defer file.Close()

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var testCount int
	fmt.Fscan(reader, &testCount)

	for i := 0; i < testCount; i++ {
		var n, m int
		fmt.Fscan(reader, &n, &m)
		fmt.Fprintln(writer, Add(n, m))

	}

	writer.Flush()

}
