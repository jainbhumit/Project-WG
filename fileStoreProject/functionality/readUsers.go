package functionality

import (
	"encoding/json"
	"file/models"
	"os"
)

const fileName = "users.json"

func ReadUsers() (map[string]models.User, error) {
	users := make(map[string]models.User)
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
		users[user.Username] = user
	}

	return users, nil
}
