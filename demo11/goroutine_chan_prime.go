// goroutine结合channel实现统计1-120000的数字中那些是素数
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func putNum(intChan chan int) {
	for i := 2; i <= 1000; i++ {
		intChan <- i
	}
	close(intChan)
	wg.Done()
}

func primeNum(exitChan chan bool, intChan, primeChan chan int) {
	for num := range intChan {
		var flag bool = true
		for i := 2; i < num; i++ {
			if num % i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	exitChan <- true
	wg.Done()
}

func printPrime (primeChan chan int) {
	for val := range primeChan {
		fmt.Println(val)
	}
	wg.Done()
}

func main() {
	start := time.Now().Unix()
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 1000)

	// 标识退出管道
	exitChan := make(chan bool, 8)
	wg.Add(1)
	go putNum(intChan)

	for i := 0; i < 8; i++ {
		wg.Add(1)
		go primeNum(exitChan, intChan, primeChan)
	}

	wg.Add(1)
	go printPrime(primeChan)

	wg.Add(1)
	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
		}
		close(primeChan)
		wg.Done()
	}()
	wg.Wait()
	end := time.Now().Unix()

	fmt.Println("time:",end - start)
}