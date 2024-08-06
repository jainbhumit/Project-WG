package component

import (
	"encoding/json"
	"file/models"
	"io/ioutil"
)

func WriteUsers(users map[string]models.User) error {
	var userList []models.User
	for _, user := range users {
		userList = append(userList, user)
	}

	data, err := json.MarshalIndent(userList, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fileName, data, 0644)
}
