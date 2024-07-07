package main

import (
	"fmt"
	"math/big"
)

var GlobalMem *big.Int = new(big.Int)

// MAIN MENU
func main() {
	clear()

	fmt.Print("Welcome to Fat Chance CLI companion!\n\n")
	var choice int

mainloop:
	for {
		if displayMem() {
			fmt.Println("-1. Clear memory")
		}
		fmt.Println("1. Counting")
		fmt.Println("2. Sequences")
		fmt.Println("3. Collections/Binomials")
		fmt.Println("4. Multinomials")
		fmt.Println("5. Collections with repetition")
		fmt.Println("0. Exit")
		fmt.Print("\n\n>")
		fmt.Scan(&choice)

		switch choice {
		case -1:
			{
				GlobalMem.SetInt64(0)
				clear()
			}
		case 0:
			{
				clear()
				fmt.Print("Farewell!\n")
				break mainloop
			}
		case 1:
			{
				countingModule()
			}
		case 2:
			{
				sequencesModule()
			}
		case 3:
			{
				collectionsModule()
			}
		case 4:
			{
				multinomialsModule()
			}
		case 5:
			{
				collectionsRepModule()
			}
		default:
			{
				fmt.Print("Invalid choice, please choose:\n")
			}
		}
	}

}

// SUB MENU
func subMenu(result *big.Int) bool {
	var choice int
	lineDoubleDecoration()
	fmt.Println("0. Store to memory")
	fmt.Println("1. Continue")
	fmt.Println("2. Return to main menu")
	fmt.Print(">")
	fmt.Scan(&choice)
	if choice == 0 {
		GlobalMem.Set(result)
	} else if choice == 2 {
		clear()
		return true
	}
	return false
}

func displayMem() bool {
	if GlobalMem.Cmp(big.NewInt(0)) != 0 {
		fmt.Printf("M: %s\n", GlobalMem.String())
		lineSingleDecoration()
		return true
	}
	return false
}

// TEXT DECORATIONS -------------------------
func clear() {
	fmt.Println("\033[H\033[J")
}

func lineSingleDecoration() {
	fmt.Println("___________________________")
}

func lineDoubleDecoration() {
	fmt.Println("===========================")
}
