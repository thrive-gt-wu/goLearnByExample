/* 定义两个方法，一个方法给管道里面写数据，一个给管道里面读取数据。要求同步
进行。
	1、开启一个fn1的的协程给向管道inChan中写入100条数据
	2、开启一个fn2的协程读取inChan中写入的数据
	3、注意：fn1和fn2同时操作一个管道
	4、主线程必须等待操作完成后才可以退出
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func fn1(intChan chan int) {
	for i:= 0; i < 100; i++ {
		intChan <- i
		fmt.Printf("writeData: %v\n", i)
		time.Sleep(time.Microsecond * 50)
	}
	close(intChan)	// 写完数据后要关闭管道，否则会发生死锁
	wg.Done()
}

func fn2(intChan chan int) {
	for v := range intChan {
		fmt.Printf("readData: %v\n", v)
		time.Sleep(time.Microsecond * 50)
	}
	wg.Done()
}
func main() {
	allChan := make(chan int, 100)
	wg.Add(1)
	go fn1(allChan)
	wg.Add(1)
	go fn2(allChan)
	wg.Wait()
	fmt.Println("数据读取完毕")
 }