package main

import (
	"fmt"
)

// O(n) | O(n)
func firstMissingPositive(nums []int) int {
	list := make(map[int]struct{})

	for _, val := range nums {
		list[val] = struct{}{}
	}

	i := 0
	for true {

		_, ok := list[i+1]

		if !ok {
			return i + 1
		}

		i++
	}

	return 0
}

func main() {
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
}
