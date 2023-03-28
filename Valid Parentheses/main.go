package main

import (
	"fmt"
)

func main() {
	str := "){"
	fmt.Println(isValid(str))
}

func isValid(s string) bool {

	if len(s)%2 != 0 {
		return false
	}

	mp := []rune{}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			mp = append(mp, char)
		} else if len(mp) == 0 {
			return false
		} else if char == ')' && mp[len(mp)-1] == '(' {
			mp = mp[:len(mp)-1]
		} else if char == '}' && mp[len(mp)-1] == '{' {
			mp = mp[:len(mp)-1]
		} else if char == ']' && mp[len(mp)-1] == '[' {
			mp = mp[:len(mp)-1]
		} else {
			return false
		}
	}
	return len(mp) == 0
}
