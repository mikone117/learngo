package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const errorMessage = "Usage: go run main.go <celsius number>"

func main() {
	if len(os.Args) < 2 {
		fmt.Println(strings.TrimSpace(errorMessage))
		return
	}
	arg := os.Args[1]
	celsius, err := strconv.ParseFloat(arg, 64)

	if err != nil {
		fmt.Printf("error: %q is not a number\n", arg)
		return
	}

	fahrenheit := celsius*1.8 + 32

	fmt.Printf("%g celsius is %g fahrenheit.\n", celsius, fahrenheit)

}
