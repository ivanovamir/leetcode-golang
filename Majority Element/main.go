package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}

func majorityElement(nums []int) int {
	dict := make(map[int]int)
	maxCount := 0
	mostFrequent := 0
	for _, num := range nums {
		dict[num]++
		if dict[num] > maxCount {
			maxCount = dict[num]
			mostFrequent = num
		}
	}
	return mostFrequent
}
