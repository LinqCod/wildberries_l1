package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "cat", "dog", "cat", "tree"}
	//Поскольку из коробки в golang set отсутствует,
	//В качестве замены возьмем мапу
	set := make(map[string]bool)

	//Записть уникальных значений в мапу
	for _, v := range arr {
		set[v] = true
	}

	//Вывод результата
	for key := range set {
		fmt.Println(key)
	}
}
