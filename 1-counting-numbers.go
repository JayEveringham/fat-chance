package main

import (
	"fmt"
	"math"
)

func countingNumbers() {

	for {
		clear()
		fmt.Println("Counting Numbers")
		fmt.Println("-> Counts the numbers (or multiples) between n and k")
		fmt.Println("\nFormula:\nn - k + 1")
		lineSingleDecoration()

		var n, k, m int
		var choice int

		fmt.Print("\nFirst integer: ")
		fmt.Scan(&n)
		fmt.Print("Second integer: ")
		fmt.Scan(&k)
		fmt.Print("Enter a multiple to count (default: 1): ")
		fmt.Scan(&m)
		lineSingleDecoration()

		checkNK(&n, &k) // checks for upper and lower values to conform with formula

		// Need to find multiples between upper and lower bounds
		if m != 1 {
			
			// first method
			countAndDivide(n, k, m)
			// second method
			findMultipleNK(&n, &k, m)
			lineSingleDecoration()
		}

		result := countNK(n, k)
		fmt.Printf("\n%v - %v + 1 \n=\n%v\n\n", n, k, result)

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

// First (very quick) method
func countAndDivide(n, k, m int) {
	var countedAndDivided float64 = float64(countNK(n, k) / m)
	floor := math.Floor(countedAndDivided)
	fmt.Printf("There are %v multiples of %v between %v and %v\n", floor, m, n, k)
	lineSingleDecoration()
}

// Second (long) method
func findMultipleNK(n *int, k *int, m int) {

	// finds initial remainders
	nmodm := *n % m
	kmodm := *k % m

	// purely for debugging
	fmt.Printf("%v mod %v = %v\n", *n, m, nmodm)
	fmt.Printf("%v mod %v = %v\n", *k, m, kmodm)

	fmt.Printf("\nMultiple of %v <= %v:\n", m, *n)

	// find upper multiple below n
	if *n < 0 {
		*n -= m + (*n % m)
	} else {
		*n -= (*n % m)
	}

	fmt.Printf("%v", *n)
	fmt.Printf("\nMultiple of %v >= %v:\n", m, *k)

	// find lower multiple above k
	if *k < 0 {
		*k -= *k % m
	} else {
		*k += m - (*k % m)
	}
	fmt.Printf("%v\n", *k)

	*n /= m
	*k /= m
	fmt.Printf("\nDividing upper and lower bounds by %v to get: \nn = %v\nk = %v\n", m, *n, *k)
}

// Order to n > k
func checkNK(n, k *int) {
	if *n < *k {
		// Need to swap as n is less than k
		// Bitwise XOR!
		*n = *n ^ *k
		*k = *n ^ *k
		*n = *n ^ *k
	}
}

// Main function
func countNK(n, k int) int {
	return n - k + 1

}
