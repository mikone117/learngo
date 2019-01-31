package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Args[1]
	fmt.Println("Number of items inside or.Items: ", len(os.Args))
	fmt.Println("Hey mate", name, "!!")

	name, age := "Dummy guy", 1102

	fmt.Println("Name: " + name)
	fmt.Println("Age : ", age)

}
