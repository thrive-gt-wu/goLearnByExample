package main

import "fmt"

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	num1 := 48
	num2 := 18
	result := gcd(num1, num2)
	fmt.Printf("%d和%d的最大公约数是：%d\n", num1, num2, result)
}