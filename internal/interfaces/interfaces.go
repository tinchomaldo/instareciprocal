package interfaces

import "github.com/tinchomaldo/instareciprocal/internal/domain"

type UserRepository interface {
	GetFollowers() ([]domain.User, error)
	GetFollowing() ([]domain.User, error)
}
