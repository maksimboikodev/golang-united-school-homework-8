package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func makeUsers(filename string) ([]User, error) {
	data, err := getFileData(filename)

	if err != nil {
		return []User{}, err
	}

	if len(data) == 0 {
		return []User{}, nil
	}

	var users []User

	if err := json.Unmarshal(data, &users); err != nil {

		return []User{}, fmt.Errorf("Error decoding: %w", err)
	}

	return users, nil
}

func getFileData(filename string) ([]byte, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0644)

	if err != nil {
		if os.IsNotExist(err) {
			return []byte{}, nil
		}
		return []byte{}, fmt.Errorf("Error openning %q: %w", filename, err)
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		return []byte{}, fmt.Errorf("Error file reading: %w", err)
	}

	return data, nil
}

func findById(id string, users []User) []byte {
	for _, user := range users {
		if user.Id == id {
			data, _ := json.Marshal(user)
			return data
		}
	}
	return []byte{}
}

func addUser(item string, users []User) ([]User, error) {
	var user User

	if err := json.Unmarshal([]byte(item), &user); err != nil {
		return users, fmt.Errorf("Error unmarshaling %q: %w", item, err)
	}

	if idExists(user.Id, users) {
		return users, fmt.Errorf("Item with id %v already exists", user.Id)
	}

	users = append(users, user)

	return users, nil
}

func idExists(id string, users []User) bool {
	for _, user := range users {
		if user.Id == id {
			return true
		}
	}
	return false
}

func saveUsers(filename string, users []User) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)

	if err != nil {
		return fmt.Errorf("Error openning %q: %w", filename, err)
	}

	defer file.Close()

	data, err := json.Marshal(users)

	if err != nil {
		return fmt.Errorf("Error marshalling: %w", err)
	}

	if _, err := file.Write(data); err != nil {
		return fmt.Errorf("Error file writing %q: %w", filename, err)
	}

	return nil
}

func deleteUserById(id string, users []User) ([]User, error) {
	var newUsers []User

	for _, user := range users {
		if user.Id != id {
			newUsers = append(newUsers, user)
		}
	}

	if !idExists(id, users) {
		return newUsers, fmt.Errorf("Item with id %v not found", id)
	}

	return newUsers, nil
}
