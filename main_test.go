package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)


func TestSum(t *testing.T) {

	fileInputs, err := os.Open("inputs.txt")
	if err != nil {
		fmt.Println("Unable to open inputs file:", err)
		return
	}

	fileOutputs, err := os.Open("outputs.txt")
	if err != nil {
		fmt.Println("Unable to open outputs file:", err)
		return
	}

	defer fileInputs.Close()
	defer fileOutputs.Close()

	reader := bufio.NewReader(fileInputs)
	reader2 := bufio.NewReader(fileOutputs)

	var testCount int
	fmt.Fscan(reader, &testCount)

	

	for i := 0; i < testCount; i++ {
		
		inputs := make(map[int]int)
		var expectedSum int

		var goodsCount, price int
		fmt.Fscan(reader, &goodsCount)

		for j := 0; j < goodsCount; j++ {

			fmt.Fscan(reader, &price)
			if _, ok := inputs[price]; ok {
				inputs[price] += 1
			} else {
				inputs[price] = 1
			}
		}

		fmt.Fscan(reader2, &expectedSum)

		fmt.Printf("%d\n", expectedSum)

		if output := PaymentSum(inputs); output != expectedSum {
			t.Errorf("Output %d not equal to expected %d", output, expectedSum)
		}

	}
}
