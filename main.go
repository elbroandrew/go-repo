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

	// file, err := os.Open("file.txt") // open in read-only format
	file, err := os.OpenFile("file.txt", os.O_RDWR|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println("Unable to open file:", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(file)

	var n string
	fmt.Fscan(reader, &n)
	// fmt.Fprintln(writer, n)
	_, err = writer.WriteString(n)
	if err != nil {
		fmt.Println("Unable to write file:", err)
	}
	fmt.Printf("wrote %s\n", n)

	writer.Flush()

}
