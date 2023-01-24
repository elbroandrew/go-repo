package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	fmt.Println(math.MinInt8)
}

func isPalindrome(s string) bool {

	var skips = ",.!?_-:;' "
	var newString = ""
	for i := 0; i < len(s); i++ {
		if !strings.ContainsRune(skips, rune(s[i])) {
			newString += string(s[i])
		}
	}
	newString = strings.ToLower(newString)

	var p1 = 0
	var p2 = len(newString) - 1

	for p1 < p2 {
		if newString[p1] == newString[p2] {
			p1++
			p2--
		} else {
			return false
		}
	}
	return true
}
