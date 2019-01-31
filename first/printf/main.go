package main

import "fmt"

func main() {
	var (
		planet   = "Venus"
		distance = 261
		orbital  = 224.701
		hasLife  = false
	)

	fmt.Printf("Planet: %s\n", planet)
	fmt.Printf("Distance: %d millions kms\n", distance)
	fmt.Printf("Orbital period: %.4f days\n", orbital)
	fmt.Printf("Does %s have life? %t\n", planet, hasLife)
	fmt.Printf(
		"%v is %v away. Think! %[2]v kms! %[1]v OMG! \n",
		planet, distance,
	)
}
