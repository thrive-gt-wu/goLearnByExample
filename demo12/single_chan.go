package main

import "fmt"

func main() {
	var writeChan chan<- int
	writeChan = make(chan int, 10)
	writeChan <- 20
	
	var readChan <-chan int
	readChan = make(chan int, 10)
	x := <-readChan
	fmt.Println("x", x)
}