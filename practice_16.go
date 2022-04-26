package main

import (
	"fmt"
	"math/rand"
)

func main() {
	data := []int{9, 1, 124, 4, 78, 3, 84, -150, 1, 2}
	fmt.Println(quickSort(data))
}

func quickSort(data []int) []int {
	if len(data) < 2 {
		return data
	}
	// берем индексы крайних элементов
	left := 0
	right := len(data) - 1
	// выбираем случайный индекс элемента, который будет опорным
	pivotIndex := rand.Uint64() % uint64(len(data))
	// меняем опорный и крайний правый элементы местами
	data[pivotIndex], data[right] = data[right], data[pivotIndex]
	// отбрасываем элементы, левее опорного, которые меньше него
	for i := range data {
		if data[i] < data[right] {
			data[i], data[left] = data[left], data[i]
			left++
		}
	}
	// опорник ставим сразу после последнего элемента, который оказался меньше опорного
	data[left], data[right] = data[right], data[left]
	quickSort(data[:left])
	quickSort(data[left+1:])
	return data
}
