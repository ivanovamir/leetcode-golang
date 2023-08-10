package main

import "fmt"

// O(n) | O(n)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var res []float64

	i := 0
	j := 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			res = append(res, float64(nums1[i]))
			i++
		} else {
			res = append(res, float64(nums2[j]))
			j++
		}
	}

	if i != len(nums1) {
		for _, val := range nums1[i:] {
			res = append(res, float64(val))
		}
	}

	if j != len(nums2) {
		for _, val := range nums2[j:] {
			res = append(res, float64(val))
		}
	}

	if len(res)%2 != 0 {
		return res[len(res)/2]
	} else {
		return (res[len(res)/2] + res[(len(res)/2)-1]) / 2
	}
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}

	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
