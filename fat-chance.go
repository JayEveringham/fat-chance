package main

import "fmt"

func main() {
	clear()
	fmt.Println("Welcome to Fat Chance CLI companion!\n\n")
	var choice int

mainloop:
	for {

		fmt.Println("1. Counting")
		fmt.Println("2. [placeholder]")
		fmt.Println("0. Exit")
		fmt.Print("\n\nChoose: ")
		fmt.Scan(&choice)

		switch choice {
		case 0:
			{
				clear()
				fmt.Println("Farewell!\n")
				break mainloop
			}
		case 1:
			{
				countingNumbers()
			}
		case 2:
			{
				fmt.Println("You have selected 2:")
			}
		default:
			{
				fmt.Println("Invalid choice, please choose:\n")
			}
		}
	}

}

func clear() {
	fmt.Println("\033[H\033[J")
}
