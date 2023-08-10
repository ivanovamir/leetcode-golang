package main

import (
	"fmt"
)

// O(n) | O(n)
func lengthOfLongestSubstring(s string) int {
	repeat := make(map[byte]struct{})
	var uniqueCounter []int

	iterByteStr := []byte(s)
	counter := 0

	for i := 0; i < len(iterByteStr); {
		_, ok := repeat[iterByteStr[i]]

		if !ok {
			repeat[iterByteStr[i]] = struct{}{}
			counter++
			i++
			continue
		}

		uniqueCounter = append(uniqueCounter, counter)
		repeat = map[byte]struct{}{}
		counter = 0

		iterByteStr = iterByteStr[1:]
		i = 0
	}

	for _, val := range uniqueCounter {
		if val > counter {
			counter = val
		}
	}

	return counter
}

func main() {
	fmt.Println(lengthOfLongestSubstring("anviaj"))
}
