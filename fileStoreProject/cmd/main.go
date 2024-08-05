package main

import (
	"file/ui"
	"fmt"
)

func main() {
	for {
		fmt.Println("------------------------------------------------------------")
		var choice string
		fmt.Println("For SignUp press:  1")
		fmt.Println("For Login press :  2")
		fmt.Println("For Exit press	 :  3")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			ui.SignUp()
		case "2":
			ui.Login()
		case "3":
			return
		default:
			fmt.Println("Invalid choice")

		}

	}
}
