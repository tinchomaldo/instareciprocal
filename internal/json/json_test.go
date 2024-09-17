package json

import (
	"os"
	"reflect"
	"testing"
)

func TestReadJSON(t *testing.T) {
	// create a temporary JSON file for testing
	content := `[
		{"username": "user1"},
		{"username": "user2"},
		{"username": "user3"}
	]`
	tmpfile, err := os.CreateTemp("", "test*.json")
	if err != nil {
		t.Fatalf("cannot create temporary file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatalf("failed to write to temporary file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("failed to close temporary file: %v", err)
	}

	// test ReadJSON function
	expected := []string{"user1", "user2", "user3"}
	result, err := ReadJSON(tmpfile.Name())
	if err != nil {
		t.Fatalf("ReadJSON failed: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ReadJSON() = %v, want %v", result, expected)
	}
}

func TestReadJSON_InvalidJSON(t *testing.T) {
	// create a temporary file with invalid JSON
	content := `invalid json`
	tmpfile, err := os.CreateTemp("", "test*.json")
	if err != nil {
		t.Fatalf("cannot create temporary file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatalf("failed to write to temporary file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("failed to close temporary file: %v", err)
	}

	// test ReadJSON function with invalid JSON
	_, err = ReadJSON(tmpfile.Name())
	if err == nil {
		t.Errorf("ReadJSON() with invalid JSON should return an error")
	}
}
