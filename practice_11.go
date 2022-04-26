package main

import "fmt"

func main() {

	arr1 := []int{4, 1, 5, 2, 9, 3}
	arr2 := []int{3, 1, 2, 5}

	var result []int

	myMap := make(map[int]int)

	for _, key := range arr1 {
		//_, found := myMap[key]
		//if found {
		//	myMap[key] += 1
		//} else {
		//	myMap[key] = 1
		//}
		myMap[key] = 1
	}

	for _, key := range arr2 {
		if _, found := myMap[key]; found /*&& v > 0*/ {
			result = append(result, key)
			myMap[key] -= 1
		}
	}

	fmt.Println(result)
}
