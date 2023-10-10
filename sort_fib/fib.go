package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
		fmt.Println(b)
	}
	return b
}

func main() {
	n := 10
	result := fibonacci(n)
	fmt.Println("结果为：", result)
}