// forrange从管道循环取值

package main

import "fmt"

func main() {
	var ch1 = make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch1 <- i
	}
	// 	使用for range遍历管道，需要再遍历前关闭管道，如果没有关闭管道就会报错
	close(ch1)

	for val := range ch1 {
		fmt.Println(val)
	}
	// 管道被关闭的时候就会退出for range
}