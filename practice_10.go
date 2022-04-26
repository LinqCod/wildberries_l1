package main

import "fmt"

func main() {
	data := []float32{-24.5, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	tempGroups := make(map[int][]float32)

	//Для каждого числа в исходных данных получаем первую цифру и знак,
	//После чего домножаем на 10 дабы получить ключ группы;
	//Добавляем значение по ключу в соответствующую группу
	for _, v := range data {
		multiplier := 10
		multiplier *= (int(v) / 10) % 10

		tempGroups[multiplier] = append(tempGroups[multiplier], v)
	}

	fmt.Println(tempGroups)
}
