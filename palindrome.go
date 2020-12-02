package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	var prep = strings.ToLower(s)
	for i := 0; i < len(prep)/2; i++ {
		if prep[i] != prep[len(prep)-(i+1)] {
			return false
		}
	}
	return true
}

func main() {
	var words []string = []string{"Mom", "Dad", "Anything", "qwertytrewq", "00001000", "000010000"}

	for _, w := range words {
		if isPalindrome(w) {
			fmt.Printf("%v: Palindrome\n", w)
			continue
		}
		fmt.Printf("%v: NOT Palindrome\n", w)
	}
}
