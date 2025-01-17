package main

import (
	"fmt"
	"math/big"
	"github.com/JayEveringham/fat-chance/scientific"
)

func collectionsRepModule() {
	for {
		clear()
		displayMem()

		fmt.Println("Counting collections with repetition")
		fmt.Println("-> The number of collections of k objects, chosen from a pool of n objects with repetition allowed")
		fmt.Println("\nFormula:\n(n + k - 1) choose k (or n - 1)")
		fmt.Println("=\n(n + k - 1)! / k!(n - 1)!")
		lineSingleDecoration()

		var n, k int64
		var result = new(big.Int)

		fmt.Print("\nn: ")
		fmt.Scan(&n)
		fmt.Print("k: ")
		fmt.Scan(&k)
		lineSingleDecoration()

		if n == 0 {
			fmt.Println("n cannot be 0")
		} else {
			result = collectionsRep(n, k)
			fmt.Printf("%v\n", scientific.BigIntToScientific(result))
		}

		returnToMain := subMenu(result)
		if returnToMain {
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
