package main

import (
	"fmt"
)

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
}

func containsDuplicate(nums []int) bool {

	newMap := make(map[int]struct{})

	for _, x := range nums {

		if _, ok := newMap[x]; ok {
			return true
		}
		newMap[x] = struct{}{}
	}
	return false
}

// Solution 2:
//func containsDuplicate(nums []int) bool {
//	if len(nums) <= 1 {
//		return false
//	}
//
//
//	sort.Ints(nums)
//
//
//	for i := 0; i < len(nums)-1; i++ {
//		if nums[i] == nums[i+1] {
//			return true
//		}
//	}
//
//	return false
//}
