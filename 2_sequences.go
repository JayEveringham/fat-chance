package main

import (
	"fmt"
	"math/big"
)

func sequencesModule() {

	for {
		clear()
		fmt.Println("Multiplication Principle and Factorials")
		fmt.Println("-> Calculates the number of sequences of k objects chosen without repetition from a collection of n objects.")
		fmt.Println("\nFormula:\nn! / (n - k)!")
		lineSingleDecoration()
		lineSingleDecoration()

		var n, k int64
		var choice int

		fmt.Print("\nn: ")
		fmt.Scan(&n)
		fmt.Print("k: ")
		fmt.Scan(&k)
		lineSingleDecoration()

		// Input validation
		if n < 0 {
			fmt.Println("Invalid input: n should be >= 0")
		} else if k < 0 {
			fmt.Println("Invalid input: k should be >= 0")
		} else if k > n {
			fmt.Println("Invalid input: k should be <= n")
		} else if k == 0 {
			fmt.Println("Invalid input: k should be > 0")
		} else {
			result := sequences(n, k)
			fmt.Printf("Result: %v\n", result)
		}

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
func sequences(n, k int64) *big.Int {
	nFact := factorial(n)
	nkFact := factorial(n - k)

	result := new(big.Int).Div(nFact, nkFact)

	return result
}

// recursive factorialisation
func factorial(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}
	result := big.NewInt(n)
	return result.Mul(result, factorial(n-1))
}
