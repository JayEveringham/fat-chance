package main

import "fmt"

func main() {
	clear()
	fmt.Print("Welcome to Fat Chance CLI companion!\n\n")
	var choice int

mainloop:
	for {

		fmt.Println("1. Counting")
		fmt.Println("2. Sequences")
		fmt.Println("3. Collections/Binomials")
		fmt.Println("4. Multinomials")
		fmt.Println("0. Exit")
		fmt.Print("\n\nChoose: ")
		fmt.Scan(&choice)

		switch choice {
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
		default:
			{
				fmt.Print("Invalid choice, please choose:\n")
			}
		}
	}

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