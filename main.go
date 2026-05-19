package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

}

func readFile(filename string) ([]User, error) {
	var response []User
	content, err := os.ReadFile(filename)
	if err != nil {
		msg := err.Error()
		return []User{}, fmt.Errorf("File reading error: %s", msg)
	}
	err = json.Unmarshal(content, &response)
	if err != nil {
		msg := err.Error()
		return []User{}, fmt.Errorf("JSON unmarshal error: %s", msg)
	}
	return response, nil
}

func writeFile(filename string, content any) error {
	bytes, err := json.Marshal(content)
	if err != nil {
		msg := err.Error()
		return fmt.Errorf("JSON marshal error: %s", msg)
	}
	os.WriteFile(filename, bytes, 0644)
	return nil
}

type User struct {
	Username string
	Password string
}
