package infrastructure

import (
	"encoding/json"
	"os"

	"github.com/tinchomaldo/instareciprocal/internal/domain"
)

type JSONRepository struct {
	FollowersFile string
	FollowingFile string
}

func (r *JSONRepository) GetFollowers() ([]domain.User, error) {
	return r.readJSON(r.FollowersFile)
}

func (r *JSONRepository) GetFollowing() ([]domain.User, error) {
	return r.readJSON(r.FollowingFile)
}

func (r *JSONRepository) readJSON(filename string) ([]domain.User, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var users []domain.User
	err = json.Unmarshal(file, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
