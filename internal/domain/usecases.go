package domain

func (c *Comparator) NotFollowingBack() []User {
	return Difference(c.Following, c.Followers)
}

func (c *Comparator) NotFollowedBack() []User {
	return Difference(c.Followers, c.Following)
}

func Difference(slice1, slice2 []User) []User {
	m := make(map[string]bool)
	for _, item := range slice2 {
		m[item.Username] = true
	}

	var diff []User
	for _, item := range slice1 {
		if _, found := m[item.Username]; !found {
			diff = append(diff, item)
		}
	}
	return diff
}
