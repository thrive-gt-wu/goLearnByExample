// （Goroutine）有三个函数，分别打印"cat", "fish","dog"要求每一个函数都用一个goroutine，按照顺序打印100次。

package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
var dog = make(chan struct{}, 1)
var cat = make(chan struct{}, 1)
var fish = make(chan struct{}, 1)

func Dog() {
	<-fish
	fmt.Println("dog")
	dog <- struct{}{}
	wg.Done()
}

func Cat() {
	<-dog
	fmt.Println("cat")
	cat <- struct{}{}
	wg.Done()
}

func Fish() {
	<-cat
	fmt.Println("fish")
	fish <- struct{}{}
	wg.Done()
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(3)
		go Dog()
		go Cat()
		go Fish()
	}
	dog <- struct{}{}

	wg.Wait()
}


// 两个协程交替打印10个字母和数字
// package main

// import (
// 	"fmt"
// 	"sync"
// )
// var wg sync.WaitGroup
// var word = make(chan struct{}, 1)
// var num = make(chan struct{}, 1)

// func printNum() {
// 	for i := 0; i < 10; i++ {
// 		<-word
// 		fmt.Println(1)
// 		num <- struct{}{}
// 	}
// 	wg.Done()
// }
// func printWord() {
// 	for i := 0; i < 10; i++ {
// 		<-num
// 		fmt.Println("a")
// 		word <- struct{}{}
// 	}
// 	wg.Done()
// }

// func main() {
// 	num <- struct{}{}
// 	wg.Add(1)
// 	go printNum()
// 	wg.Add(1)
// 	go printWord()
// 	wg.Wait()
// }

