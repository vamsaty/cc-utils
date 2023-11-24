package cc_utils

import (
	"io"
	"os"
)

// FileExists checks if the file exists.
func FileExists(fileName string) bool {
	if _, err := os.Stat(fileName); err != nil {
		return false
	}
	return true
}

func ReadFile(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return io.ReadAll(file)
}
