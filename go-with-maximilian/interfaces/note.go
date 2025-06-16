package interfaces

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Note struct {
	Text string `json:"text"`
}

// Constructor Function
func NewNote(text string) (Note, error) {

	if text == "" {
		return Note{}, errors.New("text is required")
	}

	return Note{
		Text: text,
	}, nil
}

// Getter Function
func (t Note) GetText() string {
	return t.Text
}

// Reciever Function
func (t Note) Print() {
	fmt.Println(t.GetText())
}

// Reciever Function
func (t Note) SaveFile() error {
	fileName := "note.json"

	jsonContent, err := json.Marshal(t)

	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, jsonContent, 0644)

	if err != nil {
		return err
	}

	return nil
}

// Exported Function
func CallSaveNote() {
	todo, err := NewNote("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.")

	fmt.Println(todo.GetText())

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	

	SaveData(todo)
}
