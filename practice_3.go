package main

import "fmt"

func main() {
	arr := []int{2, 4, 6, 8, 10}
	c := make(chan int)

	go sumOfSquares(c, arr)

	result := <-c
	fmt.Printf("Result: %d", result)
}

// sumOfSquares функция суммирования квадратов всех чисел массива
// и последующей передачей в канал для вывода в главном потоке
func sumOfSquares(c chan int, arr []int) {
	result := 0
	for _, v := range arr {
		result += v * v
	}

	c <- result
}
