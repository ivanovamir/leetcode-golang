package main

import "fmt"

// Функция для разделения массива
func partition(arr []int, from, to int) int {
	pivot := from // нижний указатель

	// j верхний указатель
	for j := from; j < to; j++ { // итерация по массиву
		if arr[j] > arr[to] {
			arr[pivot], arr[j] = arr[j], arr[pivot]
			pivot++
		}
	}

	arr[pivot], arr[to] = arr[to], arr[pivot]
	return pivot
}

func findKthLargest(nums []int, k int) int {
	idx := partition(nums, 0, len(nums)-1)

	if idx == k-1 {
		return nums[idx]
	}

	if idx < k {
		return findKthLargest(nums[idx+1:], k-idx-1)
	}

	return findKthLargest(nums[:idx], k)
}

func main() {
	fmt.Printf("Результат: %d", findKthLargest(GenerateWorstScenario()))
}

func GenerateWorstScenario() ([]int, int) {
	size := 10
	arr := make([]int, 0, size)
	for i := size - 1; i >= 0; i-- {
		arr = append(arr, i+1)
	}
	return arr, 1
}
