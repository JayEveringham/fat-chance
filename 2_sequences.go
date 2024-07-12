package main

import (
	"fmt"
	"math/big"
	"time"

	"github.com/JayEveringham/fat-chance/scientific"
)

func sequencesModule() {

	for {
		clear()
		displayMem()
		fmt.Println("Multiplication Principle and Factorials")
		fmt.Println("-> Calculates the number of sequences of k objects chosen without repetition from a collection of n objects.")
		fmt.Println("\nFormula:\nn! / (n - k)!")
		lineSingleDecoration()
		lineSingleDecoration()

		var n, k int64
		var result = new(big.Int)

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
			result = sequences(n, k)
			fmt.Printf("Result: %v\n", scientific.BigIntToScientific(result))
		}

		returnToMain := subMenu(result)
		if returnToMain {
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

// iterative factorialisation
func factorial(n int64) *big.Int {
    if n < 0 {
        return big.NewInt(0) 
    }

    result := big.NewInt(1)
    start := time.Now()
    updateInterval := time.Millisecond * 500 // 500ms update interval
    firstUpdateDone := false

    for i := int64(2); i <= n; i++ {
        result.Mul(result, big.NewInt(i))

        if !firstUpdateDone && time.Since(start) > updateInterval {
            firstUpdateDone = true
        }

		// Checks if time taken > updateInterval, then starts outputting updates if so
        if firstUpdateDone && time.Since(start) > updateInterval {
            fmt.Printf("\rCalculating factorial: %d/%d (%.2f%%)", i, n, float64(i)/float64(n)*100)
            updateInterval += time.Millisecond * 500 
        }
    }

    if firstUpdateDone {
        fmt.Printf("\rCalculation complete: %d/%d (100.00%%)\n", n, n)
    }

    return result
}
