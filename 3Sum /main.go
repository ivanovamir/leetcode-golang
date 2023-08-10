package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var res [][]int

	mp := make(map[[3]int]struct{})
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if i != j && i != k && k != j {
					if nums[i]+nums[j]+nums[k] == 0 {

						subArr := []int{nums[i], nums[j], nums[k]}

						_, ok := mp[[3]int{subArr[0], subArr[1], subArr[2]}]

						if !ok {
							res = append(res, subArr)
							mp[[3]int{subArr[0], subArr[1], subArr[2]}] = struct{}{}
						}
					}
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
