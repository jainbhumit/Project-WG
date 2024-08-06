package ui

import (
	"errors"
	"file/component"
	"file/models"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"sync"

	//"golang.org/x/tools/go/ssa"
	"strconv"
	"unicode"
)

var wg sync.WaitGroup

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

func isValidMobile(s string) bool {

	// Check if the phone number is numeric
	if _, err := strconv.Atoi(s); err != nil {
		return false
	}

	// Check the length of the phone number
	if len(s) < 10 || len(s) > 10 { // Assuming a range of valid lengths
		return false
	}

	return true

}

func AddUser(users map[string]models.User, username string, password string) error {
	var mobile string
	fmt.Print("Enter your age: ")
	var age int
	fmt.Scanln(&age)

	if age < 18 {
		fmt.Println("Age must be greater than 18")
		return errors.New("SignUp Fail")
	}
	for {
		fmt.Println("Enter Mobile Number : ")

		fmt.Scanln(&mobile)
		if !isValidMobile(mobile) {
			fmt.Println("Invalid Mobile Number")
			continue
		} else {
			break
		}
	}
	newPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error generating password:", err)
		return errors.New("SignUp Fail")
	}
	// Add the new user to the map
	users[username] = models.User{
		Username: username,
		Password: string(newPassword),
		Age:      age,
		Mobile:   mobile,
	}

	err = component.WriteUsers(users)
	if err != nil {
		fmt.Println("Error writing user:", err)
		return errors.New("SignUp Fail")
	}

	wg.Add(1)
	// Initialize user progress with assigned course
	func() {
		progress := models.Progress{
			Username: username,
			Courses: []models.CourseProgress{
				{
					CourseID:         1, // Assuming course ID 1 is assigned
					CompletedLessons: []float32{},
					TotalLessons:     12,
				},
			},
		}

		err = component.WriteProgress(progress)
		if err != nil {
			fmt.Println("Error initializing progress:", err)
			return
		}
		wg.Done()
	}()

	fmt.Println("User added successfully")
	return nil
}

func SignUp() {
	var username, password string

	fmt.Print("Enter username: ")
	fmt.Scanln(&username)

	users, err := component.ReadUsers()
	if err != nil {
		fmt.Println("Error reading users:", err)
		return
	}

	if _, ok := users[username]; ok {
		fmt.Println("Username alrady exist")
		return
	} else {
		for {
			fmt.Println(`Enter password: 
(must contain 1 small, 1 capital,1 numeric and symbol)`)
			fmt.Scanln(&password)
			if !IsValidPassword(password) {
				fmt.Println("Enter the strong password as it does not meet our credential")
				continue
			} else {
				break
			}
		}
		err := AddUser(users, username, password)
		if err != nil {
			fmt.Println("Error adding user:", err)
			return
		}

		fmt.Println("Plz login To continue : ")
		Login()

	}

}

func Login() {
	var username, password string

	fmt.Print("Enter username: ")
	fmt.Scanln(&username)
	fmt.Print("Enter password: ")
	fmt.Scanln(&password)

	users, err := component.ReadUsers()
	if err != nil {
		fmt.Println("Error reading users:", err)
		return
	}

	if storedPassword, ok := users[username]; ok {
		err := bcrypt.CompareHashAndPassword([]byte(storedPassword.Password), []byte(password))
		if err == nil {
			fmt.Println("Authentication successful")
			wg.Wait()
			DashBoard(username)
		} else {
			fmt.Println("Invalid password")
		}
	} else {
		fmt.Println("You are not registered Plz register yourself")
		SignUp()

	}
}
