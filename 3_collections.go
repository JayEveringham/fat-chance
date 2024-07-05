package main

import (
	"fmt"
	"math/big"
)

func collectionsModule() {
	for {
		clear()
		fmt.Println("Counting collections")
		fmt.Println("-> The number of collections of k objects without repetition from a set of n objects (n choose k).")
		fmt.Println("-> This is the same as the number of ways of choosing a collection of (n - k) objects without repetition.")
		fmt.Println("\nFormula:\nn! / ( (n-k)! × k! )")
		lineSingleDecoration()

		var n, k int64
		var choice int

		fmt.Print("\nn: ")
		fmt.Scan(&n)
		fmt.Print("k: ")
		fmt.Scan(&k)
		lineSingleDecoration()

		if n < k {
			fmt.Println(("Invalid: k should be <= n"))
		} else if k <= 0 {
			fmt.Println("Invalid: k should be > 0")
		} else {
			result := collections(n, k)
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
func collections(n, k int64) *big.Int {
	nFact := factorial(n)
	nkFact := factorial(n-k)
	kFact :=  factorial(k)

	nkkFact := new(big.Int).Mul(nkFact, kFact)
	result := new(big.Int).Div(nFact, nkkFact)
	
	return result
}
