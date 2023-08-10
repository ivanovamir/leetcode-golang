package main

import (
	"fmt"
)

//func longestPalindrome(s string) string {
//
//	if len(s) < 2 {
//		return s
//	}
//
//	var maxPalindrome string
//	arr := []byte(s)
//
//	var isPalindrome bool
//
//	//arrOfPalindromes := make(map[int]string)
//
//	for i := 0; i < len(arr); i++ {
//		if i < len(arr)-i {
//			newSubString := arr[i : len(arr)-i]
//			fmt.Println(string(newSubString))
//
//			for j, _ := range newSubString {
//				if newSubString[i] == newSubString[len(arr)-j] {
//					isPalindrome = true
//				}
//			}
//		}
//	}
//
//	return maxPalindrome
//}

func longestPalindrome(s string) string {

	if len(s) < 2 {
		return s
	}

	var maxPalindrome string
	arr := []byte(s)

	arrOfPalindromes := make(map[int]string)

	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; j > i; j-- {
			if i == j-1 {
				continue
			}

			if arr[i] != arr[j] {
				delete(arrOfPalindromes, i)
				continue
			}

			val, ok := arrOfPalindromes[i]
			if ok {
				if len(val) < len(string(arr[i:j+1])) {
					arrOfPalindromes[i] = string(arr[i : j+1])
				}
			} else {
				arrOfPalindromes[i] = string(arr[i : j+1])
			}
		}
	}

	if len(arrOfPalindromes) != 0 {
		for _, val := range arrOfPalindromes {
			if len(val) > len(maxPalindrome) {
				maxPalindrome = val
			}
		}
	} else {
		return string(arr[0])
	}

	return maxPalindrome
}

func main() {
	fmt.Println(longestPalindrome("aacabdkacaa"))
}
