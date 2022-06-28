package main

import (
	"fmt"
	"io"
)

func listOperation(args Arguments, writer io.Writer) error {
	data, err := getFileData(args["fileName"])

	if err != nil {
		return fmt.Errorf("Error fetching data from %q: %w", args["fileName"], err)
	}

	_, err = writer.Write(data)

	if err != nil {
		return err
	}

	return nil
}

func findByIdOperation(args Arguments, writer io.Writer) error {
	if args["id"] == "" {
		return errorNotSpecifiedId
	}

	users, err := makeUsers(args["fileName"])

	if err != nil {
		return fmt.Errorf("Error fetching users: %w", err)
	}

	findedId := findById(args["id"], users)

	writer.Write(findedId)

	return nil
}

func addOperation(args Arguments, writer io.Writer) error {
	if args["item"] == "" {
		return errorNotSpecifiedItem
	}

	users, err := makeUsers(args["fileName"])

	if err != nil {
		return fmt.Errorf("Error fetching users: %w", err)
	}

	newUsers, err := addUser(args["item"], users)

	if err != nil {
		writer.Write([]byte(err.Error()))
	}

	if err := saveUsers(args["fileName"], newUsers); err != nil {
		return fmt.Errorf("Error saving users: %w", err)
	}

	return nil
}

func removeOperation(args Arguments, writer io.Writer) error {
	if args["id"] == "" {
		return errorNotSpecifiedId
	}

	users, err := makeUsers(args["fileName"])

	if err != nil {
		return fmt.Errorf("Error fetching users: %w", err)
	}

	newUsers, err := deleteUserById(args["id"], users)

	if err != nil {
		writer.Write([]byte(err.Error()))
	}

	if err := saveUsers(args["fileName"], newUsers); err != nil {
		return fmt.Errorf("Error saving users: %w", err)
	}

	return nil
}
