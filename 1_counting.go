package main

import (
	"fmt"
)

func countingModule() {

	for {
		clear()
		fmt.Println("Counting Numbers")
		fmt.Println("-> Counts the numbers (or multiples) between two integers, n and k, inclusive.")
		fmt.Println("-> If a multiple (m) is provided, it counts the multiples of m between n and k.")
		fmt.Println("\nFormula:")
		fmt.Println("For counting numbers: n - k + 1")
		fmt.Println("For counting multiples of m between n and k:")
		fmt.Println("1. Find the largest multiple of m less than or equal to n.")
		fmt.Println("2. Find the smallest multiple of m greater than or equal to k.")
		fmt.Println("3. Divide both bounds by m.")
		fmt.Println("4. Apply the formula (n - k + 1) on the adjusted bounds.")
		lineSingleDecoration()

		var n, k, m int
		var choice int

		fmt.Print("\nFirst integer: ")
		fmt.Scan(&n)
		fmt.Print("Second integer: ")
		fmt.Scan(&k)
		fmt.Print("Enter a multiple to count: ")
		fmt.Scan(&m)
		lineSingleDecoration()

		checkNK(&n, &k) // checks for upper and lower values to conform with formula

		// Need to find multiples between upper and lower bounds
		if m != 1 {
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

// Counts multiples of m between n and k
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
