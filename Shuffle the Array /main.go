package main

import "fmt"

func shuffle(nums []int, n int) []int {
	var newArr []int

	for i := 0; i != n; i++ {
		newArr = append(newArr, nums[i], nums[n+i])
	}

	return newArr
}

func main() {
	fmt.Println(shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
}
