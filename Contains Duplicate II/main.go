package main

import (
	"fmt"
)

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	mp := make(map[int]int)

	for i, val := range nums {

		if val, ok := mp[val]; ok && (i-val) <= k {
			return true
		}

		mp[val] = i
	}

	return false
}
