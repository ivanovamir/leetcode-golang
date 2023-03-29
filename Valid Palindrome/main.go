package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(isPalindrome("race a car"))
}

func isPalindrome(s string) bool {
	var strB []byte
	for _, val := range strings.ToLower(s) {
		if unicode.IsLetter(val) || unicode.IsNumber(val) {
			strB = append(strB, byte(val))
		}
	}
	strArr := string(strB)
	for i := 0; i < len(strArr); i++ {
		if strArr[i] != strArr[len(strArr)-i-1] {
			return false
		}
	}
	return true
}

// Second solution:
//func isPalindrome(s string) bool {
//	strB := strings.Builder{}
//	for _, val := range strings.ToLower(s) {
//		if unicode.IsLetter(val) || unicode.IsNumber(val) {
//			strB.WriteString(string(val))
//		}
//	}
//	xxx := strB.String()
//	for i := 0; i < len(xxx); i++ {
//		if xxx[i] != xxx[len(xxx)-i-1] {
//			return false
//		}
//	}
//	return true
//}
