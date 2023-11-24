package fs

import (
	"os"
)

// FileExists checks if the file exists.
func FileExists(fileName string) bool {
	if _, err := os.Stat(fileName); err != nil {
		return false
	}
	return true
}
