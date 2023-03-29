package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(hammingWeight(00000000000000000000000000001011))
}

func hammingWeight(num uint32) int {
	totalCount := 0
	for _, val := range strconv.FormatUint(uint64(num), 2) {
		if val == '1' {
			totalCount += 1
		}
	}
	return totalCount
}
