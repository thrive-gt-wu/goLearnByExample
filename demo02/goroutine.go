package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("test() hello, world", strconv.Itoa(i))
		time.Sleep(time.Second * 2)
	}
}
func main() {

	// 主线程运行完之后，协程即使没有执行完毕也不会在执行下一次
	go test()

	for i := 1; i <= 10; i++ {
		fmt.Println(" main() hello,golang" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}