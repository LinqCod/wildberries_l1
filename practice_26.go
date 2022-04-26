package main

import (
	"fmt"
	"strings"
)

func main() {
	data := []string{
		"abcd",
		"abCdefAaf",
		"aabcd",
		"Ccab",
	}

	for _, str := range data {
		fmt.Printf("%s - %v\n", str, isStringAlphasUnique(str))
	}
}

//В качестве алгоритма используем работу с мапой
//Проверяем каждый символ строки на наличие в словаре
//Если такой символ уже имеется, возвращаем false
//Иначе - добавляем символ в словарь и проходимся по оставшимся символам
func isStringAlphasUnique(str string) bool {
	runes := []rune(strings.ToLower(str))
	alphasSet := make(map[rune]bool)
	for _, alpha := range runes {
		if _, found := alphasSet[alpha]; found {
			return false
		} else {
			alphasSet[alpha] = true
		}
	}

	return true
}
