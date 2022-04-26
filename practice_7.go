package main

import (
	"fmt"
	"strconv"
	"sync"
)

type concurrentMap struct {
	mx sync.RWMutex
	m  map[string]int
}

func (s *concurrentMap) addValue(key string, value int) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.m[key] = value
}

func (s *concurrentMap) readValue(key string) (int, bool) {
	s.mx.RLock()
	defer s.mx.RUnlock()
	val, ok := s.m[key]
	return val, ok
}

func NewConcurrentMap() *concurrentMap {
	return &concurrentMap{
		m: make(map[string]int),
	}
}

func main() {
	myMap := NewConcurrentMap()
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write(i, myMap, wg)
	}
	wg.Done()
	fmt.Println(myMap)
}

func write(i int, myMap *concurrentMap, wg *sync.WaitGroup) {
	res := multByTwo(i)
	myMap.addValue(strconv.Itoa(i), res)
	wg.Done()
}

func multByTwo(i int) int {
	return i * 2
}
