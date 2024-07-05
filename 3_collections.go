package main

import "fmt"

func collectionsModule() {
	for {
		clear()
		fmt.Println("Counting collections")
		fmt.Println("-> The number of collections of k objects without repetition from a set of n objects")
		fmt.Println("\nFormula:\nn! / ( (n-k)! × k! )")
		lineSingleDecoration()

		var n, k int
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
			fmt.Printf("%v\n", float64(collections(n, k)))
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
func collections(n, k int) int {
	return factorial(n) / (factorial(n-k) * factorial(k))
}
