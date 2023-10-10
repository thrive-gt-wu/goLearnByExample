package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

//1、定义全局的 WaitGroup
var wg sync.WaitGroup

func syncTest() {
	for i := 1; i <= 10; i++ {
		fmt.Println("test() hello, world", strconv.Itoa(i))
		time.Sleep(time.Second)
	}
	// 4、goroutine 结束就登记
	wg.Done()
}

func main() {
	//2、启动一个 goroutine 就登记+1
	wg.Add(1)
	go syncTest()

	for i := 1; i <= 10; i++ {
		fmt.Println(" main() 你好 golang" + strconv.Itoa(i))
		time.Sleep(time.Millisecond * 50)
	}
	// 3、等待所有登记的 goroutine 都结束
	wg.Wait()
}