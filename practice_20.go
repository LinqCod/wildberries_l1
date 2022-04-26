package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error while reading from terminal: %s", err)
	}
	str = strings.TrimSuffix(str, "\n")
	fmt.Println(reverseWordsOrder(str))
}

func reverseWordsOrder(str string) string {
	var result string
	words := strings.Fields(str)

	for i := len(words) - 1; i >= 0; i-- {
		result += words[i] + " "
	}

	return strings.TrimSuffix(result, " ")
}
