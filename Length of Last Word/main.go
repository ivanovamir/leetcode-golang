package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLastWord("a "))
}

func lengthOfLastWord(s string) int {
	topLen := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			topLen += 1
			if i != 0 {
				if s[i-1] == ' ' {
					break
				}
			}
		}
	}
	return topLen
}
