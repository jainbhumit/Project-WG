package main

import (
	"file/functionality"
	"fmt"
)

func main() {
	for {
		fmt.Println("------------------------------------------------------------")
		var choice string
		fmt.Println("For SingUp press:  1")
		fmt.Println("For Login press :  2")
		fmt.Println("For Exit press	 :  3")
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			functionality.SignUp()
		case "2":
			functionality.Login()
		case "3":
			return
		default:
			fmt.Println("Invalid choice")

		}

	}
}
