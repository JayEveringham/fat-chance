package main

import "fmt"

func countingNumbers() {

	for {
		clear()
		fmt.Println("Counting Numbers\n")
		fmt.Println("Counts the numbers between n and k")
		fmt.Println("\nFormula:\nn - k + 1")

		var n, k int
		var choice int

		fmt.Print("\nn: ")
		fmt.Scan(&n)
		fmt.Print("k: ")
		fmt.Scan(&k)

		result := countNK(n, k)
		fmt.Printf("\n%v - %v + 1 \n=\n%v\n\n", n, k, result)

		fmt.Println("1. Continue")
		fmt.Println("2. Return to main menu")
		fmt.Scan(&choice)
		if choice == 2 {
			clear()
			break
		}
	}
}

func countNK(n, k int) int {
	return n - k + 1

}
