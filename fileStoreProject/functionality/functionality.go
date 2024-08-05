package functionality

import (
	"file/userRepo"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"unicode"
)

const fileName = "users.json"

func IsValidPassword(password string) bool {
	if len(password) < 8 {
		return false // Password is too short
	}

	var hasUpper, hasLower, hasDigit, hasSpecial bool

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	// Return true if all criteria are met
	return hasUpper && hasLower && hasDigit && hasSpecial
}
func SignUp() {
	var username, password string

	fmt.Print("Enter username: ")
	fmt.Scanln(&username)
	fmt.Println(`Enter password: 
(must contain 1 small, 1 capital,1 numeric and symbol)`)

	fmt.Scanln(&password)
	if !IsValidPassword(password) {
		fmt.Println("Enter the strong password as it does not meet our credential")
		return
	}
	users, err := userRepo.ReadUsers()
	if err != nil {
		fmt.Println("Error reading users:", err)
		return
	} else {
		if _, ok := users[username]; ok {
			fmt.Println("Username can't be used ")
			return
		} else {

			fmt.Print("Enter your age: ")
			var age int
			fmt.Scanln(&age)
			if age >= 18 {
				newPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
				if err != nil {
					fmt.Println("Error generating password:", err)
					return
				}
				// Add the new user to the map
				users[username] = string(newPassword)

				err = userRepo.WriteUsers(users)
				if err != nil {
					fmt.Println("Error writing user:", err)
					return
				}

				fmt.Println("User added successfully")
			} else {
				fmt.Println("Not Applicable for authorization")
			}
		}
	}
}

func Login() {
	var username, password string

	fmt.Print("Enter username: ")
	fmt.Scanln(&username)
	fmt.Print("Enter password: ")
	fmt.Scanln(&password)

	users, err := userRepo.ReadUsers()
	if err != nil {
		fmt.Println("Error reading users:", err)
		return
	}

	if storedPassword, ok := users[username]; ok {
		err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
		if err == nil {
			fmt.Println("Authentication successful")
		} else {
			fmt.Println("Invalid password")
		}
	} else {
		fmt.Println("You are not registered Plz register yourself")
		SignUp()

	}
}
