package main

import "fmt"

func main() {
	data := make([]int, 15)
	for i := 0; i < len(data); i++ {
		data[i] = i + 456
	}

	fmt.Println(data)
	fmt.Printf("len: %d, cap: %d\n", len(data), cap(data))
	indexToDelete := 3
	data = removeByIndex(data, indexToDelete)
	fmt.Println(data)
	fmt.Printf("len: %d, cap: %d\n", len(data), cap(data))
}

//Быстрое удаление элемента, но с изменением порядка
func removeByIndex(data []int, index int) []int {
	//замена элемента по индексу на последний элемент последовательности
	data[index] = data[len(data)-1]
	//Отсекание последнего элемента, поскольку мы его переместили
	return data[:len(data)-1]
}
