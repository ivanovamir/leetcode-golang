package main

import (
	"fmt"
)

// O(n) | O(n)
func maxArea(height []int) int {
	mxArea := 0

	p1 := 0
	p2 := len(height) - 1

	for p1 != p2 {
		if height[p1] < height[p2] {
			newArea := height[p1] * (p2 - p1)
			if mxArea < newArea {
				mxArea = newArea
			}
			p1++
		} else {
			newArea := height[p2] * (p2 - p1)
			if mxArea < newArea {
				mxArea = newArea
			}
			p2--
		}
	}

	return mxArea
}

func main() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	fmt.Println(maxArea(arr))
}
