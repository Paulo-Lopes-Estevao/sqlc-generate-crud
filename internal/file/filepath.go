package file

import (
	"fmt"
	"os"
	"path/filepath"
)

func CheckFileIfExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false, nil
	}

	if err != nil {
		return false, fmt.Errorf("failed to check file: %v", err)
	}

	return true, nil
}

func ReadFile(filePath string) ([]byte, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func JoinPath(dir, filename string) string {
	return filepath.Join(dir, filename)
}
