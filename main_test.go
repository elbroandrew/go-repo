package main

import "testing"
import "fmt"

func TestTableIsPalindrome(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"A man, a plan, a canal: Panama.", true},
		{"race a car", false},
		{" ", true},
	}

	for _, test := range tests {
		testname := fmt.Sprintf("%v", test.input)
		t.Run(testname, func(t *testing.T) {
			answer := isPalindrome(test.input)
			if answer != test.expected {
				t.Errorf("got %v, want %v", answer, test.expected)
			}
		})
	}
}
