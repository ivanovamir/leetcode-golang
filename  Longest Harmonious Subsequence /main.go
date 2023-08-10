package main

import (
	"fmt"
	"sort"
)

func findLHS(nums []int) int {
	sort.Ints(nums)
	if len(nums) == 0 {
		return 0
	}

	var maxValue int
	var minValue int
	var maxLen int
	var currentLen int

	p1 := 0
	p2 := 1

	for p1 < len(nums)-1 {
		if nums[p1]-1 == nums[p2] || nums[p1]+1 == nums[p2] || nums[p1] == nums[p2] && ((maxValue <= nums[p1]-1 || maxValue <= nums[p1]+1) && (minValue >= nums[p1]+1 || minValue >= nums[p1]-1)) {

			if maxValue <= nums[p1]-1 {
				maxValue = nums[p1]
			}

			if maxValue <= nums[p1]+1 {
				maxValue = nums[p1]
			}

			if minValue >= nums[p1]+1 {
				minValue = nums[p1]
			}

			if minValue >= nums[p1]-1 {
				minValue = nums[p1]
			}

			currentLen++
		} else {
			if currentLen > maxLen {
				maxLen = currentLen
			}
			currentLen = 1
		}
		p1++
		p2++
	}

	return maxLen
}

func main() {
	fmt.Println(findLHS([]int{1, 2, 3, 4}))
}
