package userRepo

import (
	"encoding/json"
	"file/functionality"
	"file/models"
	"io/ioutil"
)

func WriteUsers(users map[string]string) error {
	var userList []models.User
	for username, password := range users {
		userList = append(userList, models.User{Username: username, Password: password})
	}

	data, err := json.MarshalIndent(userList, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(functionality.fileName, data, 0644)
}
