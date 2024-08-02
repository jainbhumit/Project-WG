package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"strings"
)

const fileName = "users.txt"

func readUsers() (map[string]string, error) {
	users := make(map[string]string)
	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return users, nil // If file doesn't  return empty map
		}
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // reading the line from the file return false if there is no line
		line := scanner.Text() // return the string of the current line by Scan
		parts := strings.Split(line, ":")
		if len(parts) == 2 {
			users[parts[0]] = parts[1]
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func writeUser(username, password string) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%s:%s\n", username, password))
	return err
}
func Doing() {
	var username, password string

	fmt.Print("Enter username: ")
	fmt.Scanln(&username)
	fmt.Print("Enter password: ")
	fmt.Scanln(&password)

	users, err := readUsers()
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
		fmt.Println("You are not registered To confirm that you are authorized")
		fmt.Println("Enter your age : ")
		var age int
		fmt.Scanln(&age)
		if age >= 18 {
			newPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				fmt.Println("Error generating password:", err)
			}
			err = writeUser(username, string(newPassword))
			if err != nil {
				fmt.Println("Error writing user:", err)
				return
			}
			fmt.Println("User added successfully")
		} else {
			log.Println("Not Applicable for authorization")
		}

	}
}
