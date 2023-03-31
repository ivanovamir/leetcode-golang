package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(Move(10, []int{1, 2, 3, 4}))
}

func Move(moneyCount int, slice []int) int {
	var attempts int
	var costArr []int
	for i := 0; i < len(slice); i++ {
		costArr = append(costArr, i+1+slice[i])
	}

	sort.Ints(slice)
	for _, val := range costArr {
		if moneyCount > 0 {
			if moneyCount >= val {
				attempts += 1
				moneyCount -= val
			}
		}
	}

	return attempts
}
