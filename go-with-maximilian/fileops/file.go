package fileops

import (
	"os"
)

func ReadFile(fileName string) (string, error) {
	// Open file
	file, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	return string(file), nil
}
