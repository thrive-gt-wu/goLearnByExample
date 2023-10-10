package main

import (
	"fmt"
)

func SelectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		// 放到最后面的肯定是最大的数，进行下一次的时候排 -i-1 后剩余的部分
		for j := 0; j < len(arr) -i -1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func InsertSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1;
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func QuickSort(arr []int) []int {
	if(len(arr) <= 1) {
		return arr
	}
	pivot := arr[0]
	var left, right []int

	for _, value := range arr[1:] {
		if value < pivot {
			left = append(left, value)
		} else {
			right = append(right, value)
		}
	}
	left = QuickSort(left)
	right = QuickSort(right)

	return append(append(left, pivot), right...)
}

// 归并排序
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]

	left = MergeSort(left)
	right = MergeSort(right)

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0)

	for len(left) > 0 || len(right) > 0 {
		if len(left) > 0 && len(right) > 0 {
			if left[0] <= right[0] {
				result = append(result, left[0])
				left = left[1:]
			} else {
				result = append(result, right[0])
				right = right[1:]
			}
		} else if len(left) > 0 {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	return result
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("原始数组:", arr)

	// // 冒泡排序
	// bubbleSortArr := make([]int, len(arr))
	// copy(bubbleSortArr, arr)
	// BubbleSort(bubbleSortArr)
	// fmt.Println("冒泡排序结果:", bubbleSortArr)

	// // 选择排序
	// selectionSortArr := make([]int, len(arr))
	// copy(selectionSortArr, arr)
	// SelectSort(selectionSortArr)
	// fmt.Println("选择排序结果:", selectionSortArr)

	// // 插入排序
	// insertionSortArr := make([]int, len(arr))
	// copy(insertionSortArr, arr)
	// InsertSort(insertionSortArr)
	// fmt.Println("插入排序结果:", insertionSortArr)

	// 快速排序
	quickSortArr := make([]int, len(arr))
	copy(quickSortArr, arr)
	quickSortArr = QuickSort(quickSortArr)
	fmt.Println("快速排序结果:", quickSortArr)

	// // 归并排序
	// mergeSortArr := make([]int, len(arr))
	// copy(mergeSortArr, arr)
	// mergeSortArr = MergeSort(mergeSortArr)
	// fmt.Println("归并排序结果:", mergeSortArr)
}