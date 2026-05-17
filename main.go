package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"syscall"
)

func main() {
	var action int
	for action <= 0 {
		fmt.Print(`
	Choose an action:
	1) Read
	2) Create
	3) Exit

	`)
		fmt.Scan(&action)
		reader.ReadString('\n')
		switch action {
		case 1:
			notes := readNotes()
			for _, note := range notes {
				fmt.Println(note.Title, " - ", note.Description)
			}
			action = 0
		case 2:
			createNote()
			action = 0
		case 3:
			syscall.Exit(0)
		}
	}
}

var reader = bufio.NewReader(os.Stdin)

func readNotes() []Note {
	var response []Note
	file, _ := os.ReadFile("notes.json")
	json.Unmarshal(file, &response)
	var notes []Note = response
	return notes
}

func createNote() {

	fmt.Println("Title")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Println("Description")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)
	var notes []Note = readNotes()
	var note Note = Note{Title: title, Description: description}
	notes = append(notes, note)
	notesBytes, _ := json.Marshal(notes)
	os.WriteFile("notes.json", notesBytes, 0644)
}

type Note struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
