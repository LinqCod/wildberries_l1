package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	//Чтение и запись кол-ва секунд, через которое нужно остановить выполнение
	var n time.Duration
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatalf("error while reading value from console: %v", err)
	}

	//канал для данных
	c := make(chan int)
	defer close(c)

	//Горутина для чтения данных из канала
	go func() {
		for v := range c {
			fmt.Printf("Value received: %d\n", v)
			time.Sleep(time.Second)
		}
	}()

	begin := time.Now()
	end := time.Second * n

	//Считывание данных до того момента, пока не пройдет n секунд
	for time.Since(begin) < end {
		c <- rand.Int()
	}
}
