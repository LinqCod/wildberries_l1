package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

//TODO: fix bags and add shutdown ctrl+c

type controller struct {
	sig  chan os.Signal
	data chan int
	stop chan interface{}
	term chan interface{}
}

func NewController(workerPoolSize int) *controller {
	return &controller{
		sig:  make(chan os.Signal, 1),
		data: make(chan int, workerPoolSize),
		stop: make(chan interface{}),
		term: make(chan interface{}),
	}
}

func (c *controller) wait() {
	<-c.term
	close(c.term)
}

func main() {
	// Считываем с консоли кол-во рабочих
	var workerPoolSize int
	_, err := fmt.Scan(&workerPoolSize)
	if err != nil {
		log.Fatalf("error while reading value from console: %v", err)
	}

	//Создаем контроллер для всех каналов
	controller := NewController(workerPoolSize)

	//Горутина для остановки программы по нажатию на ^C
	signal.Notify(controller.sig, syscall.SIGINT)
	go stopController(controller)

	//Запуск рабочих и записи данных
	go writeData(controller)
	go startWorkers(workerPoolSize, controller)

	controller.wait()
}

func stopController(c *controller) {
	for {
		switch <-c.sig {
		case syscall.SIGINT:
			fmt.Println("^C pressed")
			c.stop <- 0
			close(c.stop)
			close(c.sig)
			return
		default:
			fmt.Println("Unknown signal")
		}
	}
}

func startWorkers(workerPoolSize int, c *controller) {
	wg := new(sync.WaitGroup)
	for i := 0; i < workerPoolSize; i++ {
		workerID := i
		wg.Add(1)
		go worker(workerID, c.data, wg)
	}
	wg.Wait()
	fmt.Println("Work finished")
	c.term <- 0
}

func worker(id int, dataChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Чтение и вывод данных из канала
	for v := range dataChan {
		fmt.Printf("Worker #%d, Value: %d\n", id, v)
		time.Sleep(time.Second)
	}
}

func writeData(c *controller) {
	// Запись данных в канал
	for i := 0; i < 1000000; i++ {
		select {
		case <-c.stop:
			fmt.Println("Stop signal detected")
			close(c.data)
			return
		case c.data <- i:
			time.Sleep(time.Second)
		}
	}
}
