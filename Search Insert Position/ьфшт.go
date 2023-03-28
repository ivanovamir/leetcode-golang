package main

import "fmt"

func main() {
	//fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
	fmt.Println(searchInsert([]int{1, 3}, 2))
}

func searchInsert(nums []int, target int) int {
	if len(nums) == 1 {
		if target < nums[0] {
			return 0
		} else if target > nums[0] {
			return 1
		}
	}
	for i, num := range nums {
		if num == target {
			return i
		} else if i == 0 {
			if target < num {
				return i
			}
		} else if i != 0 && i <= len(nums)-1 {
			if target < num {
				return i
			}
		}
	}
	return len(nums)
}
