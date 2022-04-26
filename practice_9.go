package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c1 := make(chan int, 10)
	c2 := make(chan int, 10)

	for _, x := range arr {
		c1 <- x
	}
	close(c1)

	wg.Add(1)
	go func() {
		for x := range c1 {
			c2 <- x * x
		}
		close(c2)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for v := range c2 {
			fmt.Println(v)
		}
		wg.Done()
	}()

	wg.Wait()
}
