package main

import "fmt"

func main() {
	arr := []int{2, 4, 6, 8, 10}
	c := make(chan int)

	//Запуск горутины с переданными массивом чисел
	//и каналом для транспортировки данных
	go calcSquares(c, arr)

	//Получение данных из канала и их вывод
	for v := range c {
		fmt.Println(v)
	}
}

func calcSquares(c chan int, arr []int) {
	//Возведение в квадрат всех чисел из массива
	//и их запись в канал
	for _, v := range arr {
		c <- v * v
	}

	//закрытие канала по завершению работы
	close(c)
}
