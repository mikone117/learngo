package main

import "fmt"

var global = 288

//global2 := 100 ERROR

func main() {
	var myAge, yourAge int
	fmt.Println(myAge, yourAge)

	var (
		var1 string
		var2 string
		var3 string
	)

	_, _, _ = var1, var2, var3

	var number1 = 1
	number2 := 2
	number3point3 := 3.3

	fmt.Println(number1, number2, number3point3)

}
