package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "MCMXCIV"
	fmt.Println(romanToInt(s))
}

func romanToInt(s string) int {
	var romanNumbers = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	var number int
	soloNum := strings.Split(s, "")

	for i := 0; i < len(soloNum); i++ {
		if i != len(soloNum)-1 {
			if romanNumbers[soloNum[i]] < romanNumbers[soloNum[i+1]] {
				number += romanNumbers[soloNum[i+1]] - romanNumbers[soloNum[i]]
				i++
				continue
			}
		}
		number += romanNumbers[soloNum[i]]
	}
	return number
}
