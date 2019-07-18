package main

import "fmt"

type gram float64
type ounce float64

// GramToOunceFactor Factor to convert grams into ounces
const GramToOunceFactor = 0.035274

func main() {
	var g gram = 1000
	var o ounce

	o = ounce(g) * GramToOunceFactor

	fmt.Printf("%g grams is %.2f ounces\n", g, o)
}
