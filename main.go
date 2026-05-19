package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var reader = bufio.NewReader(os.Stdin)

func main() {

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
		errorHandler(err)
	}
	os.WriteFile(filename, bytes, 0644)
	return nil
}

func createUser() error {
	var users []User
	fmt.Println("__ Username __")
	username, err := reader.ReadString('\n')
	if err != nil {
		errorHandler(err)
	}
	username = strings.TrimSpace(username)
	fmt.Println("__ Password __")
	password, err := reader.ReadString('\n')
	if err != nil {
		errorHandler(err)
	}
	password = strings.TrimSpace(password)
	response, err := readFile("users.json")
	if err != nil {
		errorHandler(err)
	}
	err = writeFile("users.json", response)
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		errorHandler(err)
	}
	user := User{Username: username, Password: string(hash)}
	users = append(response, user)
	err = writeFile("users.json", users)
	if err != nil {
		errorHandler(err)
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
