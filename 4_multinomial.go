package main

import (
	"fmt"
	"math/big"
	"github.com/JayEveringham/fat-chance/scientific"
)

func multinomialsModule() {
	for {
		clear()
		displayMem()
		fmt.Println("Multinomials")
		fmt.Println("-> Calculates the number of ways to divide n objects into k groups of sizes a1, a2, ..., ak.")
		fmt.Println("\nFormula:")
		fmt.Println("(n choose a1, a2, ..., ak) = n! / (a1! * a2! * ... * ak!)")
		lineSingleDecoration()

		var n, aInput, aSum int64
		var a []int64

		countA := 1

		fmt.Print("\nn: ")
		fmt.Scan(&n)
		for {
			fmt.Printf("%v of n objects remain\n", n-aSum)
			fmt.Printf("> Group a%v size: ", countA)
			_, err := fmt.Scan(&aInput)
			if err != nil {
				fmt.Println("Invalid - please enter an integer")
				continue
			}

			if n-(aInput+aSum) > 0 {
				a = append(a, aInput)
				aSum += aInput
				countA++
			} else if n-(aInput+aSum) == 0 {
				a = append(a, aInput)
				break
			} else if n-(aInput+aSum) < 0 {
				fmt.Println("Invalid - must be <= ", n-aSum)
			}

		}

		lineSingleDecoration()

		result := multinomials(n, a)
		fmt.Printf("%v\n", scientific.BigIntToScientific(result))

		returnToMain := subMenu(result)
		if returnToMain {
			break
		}
	}
}

// main formula
func multinomials(n int64, a []int64) *big.Int {
	nFact := factorial(n)

	denominator := big.NewInt(1)
	for _, ai := range a {
		aiFact := factorial(ai)
		denominator.Mul(denominator, aiFact)
	}

	result := new(big.Int).Div(nFact, denominator)
	return result
}
