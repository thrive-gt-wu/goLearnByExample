package main

import "fmt"

func main() {
	// 无缓冲，会发生死锁
	// ch := make(chan int)
	ch := make (chan int, 1) // 缓冲为1
	ch <- 10
	fmt.Println("发送成功")
}