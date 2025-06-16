package interfaces

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

// Constructor Function
func NewTodo(text string) (Todo, error) {

	if text == "" {
		return Todo{}, errors.New("text is required")
	}

	return Todo{
		Text: text,
	}, nil
}

// Getter Function
func (t Todo) GetText() string {
	return t.Text
}

// Reciever Function
func (t Todo) Print() {
	fmt.Println(t.GetText())
}

// Reciever Function
func (t Todo) SaveFile() error {
	fileName := "todo.json"

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
func CallSaveTodo() {
	todo, err := NewTodo("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.")

	fmt.Println(todo.GetText())

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	

	SaveData(todo)
}
