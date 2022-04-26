package main

import "fmt"

func main() {
	//Массив обязан быть упорядоченным
	data := []int{-124, -65, -10, 0, 7, 19, 65, 99, 111, 1234}
	target := 0
	resultIndex := binarySearch(data, target)
	if resultIndex >= 0 {
		fmt.Printf("Value: %d, index: %d", target, resultIndex)
	} else {
		fmt.Printf("Value %d not found", target)
	}
}

func binarySearch(data []int, target int) int {
	left := 0
	right := len(data) - 1
	var middle int

	for left <= right {
		middle = left + (right-left)/2
		if target == data[middle] {
			return middle
		}
		if target < data[middle] {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return -1
}
