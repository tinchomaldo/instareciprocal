package json

import (
	"encoding/json"
	"os"
)

type user struct {
	Username string `json:"username"`
}

func ReadJSON(filename string) ([]string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var users []user
	err = json.Unmarshal(file, &users)
	if err != nil {
		return nil, err
	}

	usernames := make([]string, len(users))
	for i, u := range users {
		usernames[i] = u.Username
	}

	return usernames, nil
}
