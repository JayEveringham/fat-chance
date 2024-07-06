package main

import (
	"fmt"
	"math/big"
)

func collectionsRepModule() {
	for {
		clear()
		fmt.Println("Counting collections with repetition")
		fmt.Println("-> The number of collections of k objects, chosen from a pool of n objects with repetition allowed")
		fmt.Println("\nFormula:\n(n + k - 1) choose k (or n - 1)")
		fmt.Println("=\n(n + k - 1)! / k!(n - 1)!")
		lineSingleDecoration()

		var n, k int64
		var choice int

		fmt.Print("\nn: ")
		fmt.Scan(&n)
		fmt.Print("k: ")
		fmt.Scan(&k)
		lineSingleDecoration()

		if n == 0 {
			fmt.Println("n cannot be 0")
		} else {
			result := collectionsRep(n, k)
			fmt.Printf("%v\n", result)
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
func collectionsRep(n, k int64) *big.Int {
	numeratorFact := factorial(n + k - 1)
	kFact := factorial(k)
	nMinOneFact := factorial(n - 1)

	denominatorFact := new(big.Int).Mul(kFact, nMinOneFact)
	result := new(big.Int).Div(numeratorFact, denominatorFact)

	return result
}
