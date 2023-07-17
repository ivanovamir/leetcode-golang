package main

import "fmt"

// O(n+m) | O(1)
func merge(nums1 []int, m int, nums2 []int, n int) {
	var ptr1, ptr2, ptr3 = m - 1, n - 1, m + n - 1
	for ptr1 >= 0 && ptr2 >= 0 {
		if nums2[ptr2] >= nums1[ptr1] {
			nums1[ptr3] = nums2[ptr2]
			ptr2--
		} else {
			nums1[ptr3] = nums1[ptr1]
			ptr1--
		}
		ptr3--
	}

	if ptr2 >= 0 {
		copy(nums1[:ptr2+1], nums2[:ptr2+1])
	}

	fmt.Println(nums1)
}

func main() {
	merge([]int{4, 0, 0, 0, 0, 0}, 1, []int{1, 2, 3, 5, 6}, 5)
}
