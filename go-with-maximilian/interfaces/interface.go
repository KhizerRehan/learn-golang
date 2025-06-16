package interfaces

import "fmt"

// Saveable interface defines the contract for saving
type Saveable interface {
	SaveFile() error
}


func SaveData(s Saveable) error {
	fmt.Println("Saving data...")
	err := s.SaveFile()
	if err != nil {
		fmt.Println("Error saving file:", err)
		return err
	}

	return nil
}

