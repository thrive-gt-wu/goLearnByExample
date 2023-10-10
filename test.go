package main

import (
	"fmt"
	"sort"
)

func scanFunc () {
var n int
	fmt.Print("请输入要输入的个数：")
	fmt.Scan(&n)

	// arr := make([]int, n)
	// for i:= 0; i < n; i++ {
	// 	fmt.Printf("请输入第%d个整数: ", i+1)
	// 	fmt.Scan(&arr[i])
	// }

	// sort.Ints(arr)
	// fmt.Println(arr)
	nums := make([]int, 0)
	for i:= 0; i < n; i++ {
		var num int
		
		fmt.Scan(&num)
		nums = append(nums, num)
	}

	sort.Ints(nums)
	fmt.Print(nums)
}

func reverseString(s string) string {
	strBytes := []byte(s)
	length := len(strBytes)

	for i := 0; i < length/2; i++ {
		strBytes[i], strBytes[length-i-1] = strBytes[length-i-1], strBytes[i]
	}
	return string(strBytes)
}

func main() {
	s := "hello world"
	reverseStr := reverseString(s)
	fmt.Println(reverseStr)
}