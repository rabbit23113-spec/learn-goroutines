package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	chooseAction()
}

func readFile(filename string) ([]User, error) {
	var response []User
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, errorHandler(err)
	}
	err = json.Unmarshal(content, &response)
	if err != nil {
		return nil, errorHandler(err)
	}
	return response, nil
}

func writeFile(filename string, content any) error {
	bytes, err := json.Marshal(content)
	if err != nil {
		return errorHandler(err)
	}
	os.WriteFile(filename, bytes, 0644)
	return nil
}

func createUser() error {
	var users []User
	fmt.Println("__ Username __")
	username, err := reader.ReadString('\n')
	if err != nil {
		return errorHandler(err)
	}
	username = strings.TrimSpace(username)
	fmt.Println("__ Password __")
	password, err := reader.ReadString('\n')
	if err != nil {
		return errorHandler(err)
	}
	password = strings.TrimSpace(password)
	response, err := readFile("users.json")
	if err != nil {
		return errorHandler(err)
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errorHandler(err)
	}
	user := User{Username: username, Password: string(hash)}
	users = append(response, user)
	err = writeFile("users.json", users)
	if err != nil {
		return errorHandler(err)
	}
	return nil
}

func chooseAction() error {
	var choice int = 0
	for choice == 0 {
		fmt.Println("__ Action __")
		fmt.Println("__ 1) Create a new account __")
		fmt.Println("__ 2) Sign in __")
		fmt.Println("__ 3) Exit __")
		action, err := reader.ReadString('\n')
		if err != nil {
			return errorHandler(err)
		}
		action = strings.TrimSpace(action)
		choice, err = strconv.Atoi(action)
		if err != nil {
			return errorHandler(err)
		}
		switch choice {
		case 1:
			createUser()
			choice = 0
		case 2:
			auth()
		case 3:
			os.Exit(0)
		}
	}
	return nil
}

func auth() error {
	fmt.Println("__ Username __")
	username, err := reader.ReadString('\n')
	if err != nil {
		return errorHandler(err)
	}
	username = strings.TrimSpace(username)
	fmt.Println("__ Password __")
	password, err := reader.ReadString('\n')
	if err != nil {
		return errorHandler(err)
	}
	password = strings.TrimSpace(password)
	response, err := readFile("users.json")
	if err != nil {
		return errorHandler(err)
	}
	for i := range response {
		if response[i].Username == username {
			err := bcrypt.CompareHashAndPassword([]byte(response[i].Password), []byte(password))
			if err != nil {
				return errorHandler(err)
			}
			fmt.Println("Success")
			return nil
		}
	}
	return nil
}

func errorHandler(err error) error {
	msg := err.Error()
	return fmt.Errorf("Error: %s", msg)
}

type User struct {
	Username string
	Password string
}
