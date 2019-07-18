package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	msg := os.Args[1]

	length := len(msg)
	exclamations := strings.Repeat("!", length)
	output := exclamations + strings.ToUpper(msg) + exclamations

	fmt.Println(output)
}
