package main

import (
	"fmt"
)

func multiplicationAndFactorials() {

	for {
		clear()
		fmt.Println("Multiplication principle and factorials")
		fmt.Println("-> The number of sequences of k objects chosen without repetition from a collection of n objects")
		fmt.Println("\nFormula:\nn! / (n-k)!")
		lineSingleDecoration()

		var n, k int
		var choice int

		fmt.Print("\nn: ")
		fmt.Scan(&n)
		fmt.Print("k: ")
		fmt.Scan(&k)
		lineSingleDecoration()

		fmt.Printf("%e\n", float64(sequenceWithoutRepetition(n, k)))

		lineDoubleDecoration()
		fmt.Println("1. Continue")
		fmt.Println("2. Return to main menu")
		fmt.Scan(&choice)
		if choice == 2 {
			clear()
			break
		}
	}
}

// main formula
func sequenceWithoutRepetition(n, k int) int {
	return factorial(n) / factorial(n-k)
}

// recursive factorialisation
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
