package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func isPrime(n int) {
	for num := (n - 1) * 3000 + 1; num <= n * 3000; num++ {
		if num == 1 {
			continue
		}

		flag := true
		for i := 2; i < num; i++ {
			if num % i == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(num)
		}
	}
	wg.Done()
}
func main() {
	start := time.Now().Unix()
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go isPrime(i)
	}
	wg.Wait()
	end := time.Now().Unix()
	result := end - start
	fmt.Println("耗时：", result)
}