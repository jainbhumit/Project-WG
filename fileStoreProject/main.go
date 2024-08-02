package main

import "fmt"

func main() {
	for {
		fmt.Println("------------------------------------------------------------")
		var choice string
		fmt.Println("To start or Continue press 1")
		fmt.Println("To End press 0")
		fmt.Println("Enter your Choice : ")
		fmt.Scanln(&choice)

		switch choice {
		case "0":
			return
		case "1":
			Doing()
		default:
			fmt.Println("Invalid choice")

		}

	}
}
