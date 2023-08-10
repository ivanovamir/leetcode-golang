package main

import (
	"fmt"
)

var phoneMap = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

// O(n) | O(n)
func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	result := []string{""}
	for i := 0; i < len(digits); i++ {
		digit := digits[i]
		letters := phoneMap[digit]
		if len(letters) == 0 {
			continue
		}
		var temp []string
		for _, prefix := range result {
			for _, letter := range letters {
				temp = append(temp, prefix+letter)
			}
		}
		result = temp
	}
	return result
}

func main() {
	fmt.Println(letterCombinations("234"))
}
