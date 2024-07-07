package main

import (
	"fmt"
	"math/big"
)

func collectionsModule() {
	for {
		clear()
		fmt.Println("Counting Collections")
		fmt.Println("-> Calculates the number of ways to choose k objects without repetition from a set of n objects (n choose k).")
		fmt.Println("-> This is equivalent to the number of ways to choose (n - k) objects without repetition.")
		fmt.Println("\nFormula:\nn! / (k!(n-k)!)")
		lineSingleDecoration()

		var n, k int64

		fmt.Print("\nn: ")
		fmt.Scan(&n)
		fmt.Print("k: ")
		fmt.Scan(&k)
		lineSingleDecoration()

		// Input validation
		if n < 0 {
			fmt.Println("Invalid input: n should be >= 0")
		} else if k <= 0 {
			fmt.Println("Invalid input: k should be >= 0")
		} else if k > n {
			fmt.Println("Invalid input: k should be <= n")
		} else {
			result := collections(n, k)
			fmt.Printf("Result: %v\n", result)
		}

		returnToMain := subMenu()
		if returnToMain {
			break
		}
	}
}

// main formula
func collections(n, k int64) *big.Int {
	nFact := factorial(n)
	nkFact := factorial(n - k)
	kFact := factorial(k)

	nkkFact := new(big.Int).Mul(nkFact, kFact)
	result := new(big.Int).Div(nFact, nkkFact)

	return result
}
