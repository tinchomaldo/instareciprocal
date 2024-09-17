package domain

type User struct {
	Username string
}

type Comparator struct {
	Followers []User
	Following []User
}

func NewComparator(followers, following []User) *Comparator {
	return &Comparator{
		Followers: followers,
		Following: following,
	}
}
