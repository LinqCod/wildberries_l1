package main

import "fmt"

// Human Главная структура
type Human struct {
	Name   string
	Gender string
	Age    int
}

// PrintInfo Метод вывода данных объекта Human
func (h *Human) PrintInfo() {
	fmt.Printf("Gender: %s, Name: %s, Age: %d", h.Gender, h.Name, h.Age)
}

// Action дочерняя структура, перенимающая состояние и поведение от Human
type Action struct {
	Human
}

func main() {
	action := Action{Human{Age: 18, Gender: "Male", Name: "Maxim"}}
	action.PrintInfo()
}
