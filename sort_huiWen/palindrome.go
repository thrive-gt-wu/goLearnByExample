package main

import "fmt"

func isPalindrome(s *string) bool {
	str := *s
	n := len(str)

	for i := 0; i < n; i++ {
		if str[i] != str[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	s := "abac"
	if isPalindrome(&s) {
		fmt.Println("是回文")
	} else {
		fmt.Println("不是")
	}
}