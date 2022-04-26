package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Задержка в 3 секунды")
	sleep(3)
	fmt.Println("3 секунды прошли")
}

//Функция sleep принимает кол-во секунд для сна
//Бесконечным циклом проверяем разницу между стартовым временем и текущим
//Если разница становится больше или равна установленному времени
//-> Выходим из цикла и продолжаем работу
func sleep(seconds int) {
	from := time.Now().Unix()
	for time.Now().Unix()-from < int64(seconds) {
	}
}
