package main

import "fmt"

// O(n) | O(n)
func twoSum(nums []int, target int) []int {
	mp := make(map[int]int, len(nums))

	for i, num := range nums {
		sum := target - num

		val, ok := mp[sum]

		if ok {
			return []int{val, i}
		}
		mp[num] = i
	}
	return nil
}

// O(2n) = O(n) | O(n)
//func twoSum(nums []int, target int) []int {
//	numsMap := make(map[int]int, len(nums))
//	for i, val := range nums {
//		numsMap[val] = i
//	}
//
//	for i, val := range nums {
//		c := target - val
//		idx, ok := numsMap[c]
//
//		if ok && i != idx {
//			return []int{idx, i}
//		}
//	}
//
//	return nil
//}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
