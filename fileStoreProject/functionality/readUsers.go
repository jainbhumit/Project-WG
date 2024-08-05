package functionality

import (
	"encoding/json"
	"file/models"
	"os"
)

func ReadUsers() (map[string]string, error) {
	users := make(map[string]string)
	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return users, nil // File doesn't exist, return empty map
		}
		return nil, err
	}
	defer file.Close()

	var userList []models.User
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&userList); err != nil {
		return nil, err
	}

	for _, user := range userList {
		users[user.Username] = user.Password
	}

	return users, nil
}
