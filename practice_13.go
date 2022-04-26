package main

import "fmt"

func main() {
	a := 1234
	b := 13526

	//Используя встроенные возможности golang
	fmt.Printf("a: %d, b: %d", a, b)
	a, b = b, a
	fmt.Printf("\na: %d, b: %d", a, b)

	//Используя сложение/вычетание
	fmt.Printf("\na: %d, b: %d", a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("\na: %d, b: %d", a, b)
}
