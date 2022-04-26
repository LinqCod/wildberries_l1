package main

import (
	"fmt"
)

func main() {
	var t interface{}
	t = "Hello World"
	checkType(t)
	t = 1
	checkType(t)

}

func checkType(i interface{}) {

	switch i := i.(type) {
	default:
		fmt.Printf("unexpected type %T\n", i)
	case int:
		fmt.Printf("integer type %T\n", i)
	case string:
		fmt.Printf("string type %T\n", i)
	case bool:
		fmt.Printf("bool type %T\n", i)
	case chan interface{}:
		fmt.Printf("channel type %T\n", i)
	}

}
