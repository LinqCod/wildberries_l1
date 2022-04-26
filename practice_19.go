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

	fmt.Println(reverseString(str))
}

func reverseString(str string) string {
	runeStr := []rune(str)
	resultBytes := make([]rune, len(str))
	for i := len(runeStr) - 1; i >= 0; i-- {
		resultBytes = append(resultBytes, runeStr[i])
	}

	return string(resultBytes)
}