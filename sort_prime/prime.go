package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	n := 11
	if isPrime(n) {
		fmt.Println("是素数")
	}
}