package domain

import (
	"reflect"
	"testing"
)

func TestComparator_NotFollowingBack(t *testing.T) {
	c := NewComparator(
		[]User{{Username: "user1"}, {Username: "user2"}, {Username: "user3"}},
		[]User{{Username: "user2"}, {Username: "user3"}, {Username: "user4"}, {Username: "user5"}},
	)

	expected := []User{{Username: "user4"}, {Username: "user5"}}
	result := c.NotFollowingBack()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("NotFollowingBack() = %v, want %v", result, expected)
	}
}

func TestComparator_NotFollowedBack(t *testing.T) {
	c := NewComparator(
		[]User{{Username: "user1"}, {Username: "user2"}, {Username: "user3"}, {Username: "user4"}},
		[]User{{Username: "user2"}, {Username: "user3"}, {Username: "user5"}},
	)

	expected := []User{{Username: "user1"}, {Username: "user4"}}
	result := c.NotFollowedBack()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("NotFollowedBack() = %v, want %v", result, expected)
	}
}

func TestDifference(t *testing.T) {
	tests := []struct {
		name     string
		a        []User
		b        []User
		expected []User
	}{
		{
			name:     "simple difference",
			a:        []User{{Username: "a"}, {Username: "b"}, {Username: "c"}},
			b:        []User{{Username: "b"}, {Username: "c"}, {Username: "d"}},
			expected: []User{{Username: "a"}},
		},
		{
			name:     "no difference",
			a:        []User{{Username: "a"}, {Username: "b"}, {Username: "c"}},
			b:        []User{{Username: "a"}, {Username: "b"}, {Username: "c"}},
			expected: []User{},
		},
		{
			name:     "completely different",
			a:        []User{{Username: "a"}, {Username: "b"}, {Username: "c"}},
			b:        []User{{Username: "d"}, {Username: "e"}, {Username: "f"}},
			expected: []User{{Username: "a"}, {Username: "b"}, {Username: "c"}},
		},
		{
			name:     "both slices empty",
			a:        []User{},
			b:        []User{},
			expected: []User{},
		},
		{
			name:     "first slice empty",
			a:        []User{},
			b:        []User{{Username: "a"}, {Username: "b"}, {Username: "c"}},
			expected: []User{},
		},
		{
			name:     "second slice empty",
			a:        []User{{Username: "a"}, {Username: "b"}, {Username: "c"}},
			b:        []User{},
			expected: []User{{Username: "a"}, {Username: "b"}, {Username: "c"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Difference(tt.a, tt.b)
			for i := 0; i < len(result); i++ {
				if result[i].Username != tt.expected[i].Username {
					t.Errorf("Difference() = %v, want %v", result, tt.expected)
				}
			}
		})
	}
}
