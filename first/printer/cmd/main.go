package main

import (
	"fmt"

	"github.com/mikone117/learngo/first/printer"
)

func main() {
	printer.Hello()
	version := printer.Version()
	fmt.Println(version)
}
