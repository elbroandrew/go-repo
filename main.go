package main

import "fmt"

func main() {
	var a = []int{2, 3, 4, 4, 3, 1, 4, 5, 4, 7, 6}
	var x = maxSubArray(a, 3)
	fmt.Println(x)
}

// sliding window
func maxSubArray(arr []int, num int) any {
	var maxSum = 0
	var tempSum = 0
	if len(arr) < num {
		return nil
	}
	//find sum of 3 elements
	for i := 0; i < num; i++ {
		maxSum += arr[i]
	}
	tempSum = maxSum
	// 1,2,3,4,5 -> -arr[0] 2+3+4 +arr[i];
	for i := num; i < len(arr); i++ {
		tempSum = -arr[i-num] + tempSum + arr[i]
		if maxSum < tempSum {
			maxSum = tempSum
		}
	}

	return maxSum
}
