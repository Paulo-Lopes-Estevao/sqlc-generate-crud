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

func GetPath(filePath string) (string, error) {
	abspath, err := filepath.Abs(filePath)
	if err != nil {
		return "", fmt.Errorf("error parsing config: absolute file path lookup failed: %s\n", err)
	}
	base := filepath.Base(abspath)

	dir := filepath.Dir(base)

	return dir, nil
}

func JoinPath(dir, filename string) string {
	return filepath.Join(dir, filename)
}

func CreateDirAndFile(filePath string, content []byte) error {
	dir := filepath.Dir(filePath)
	err := os.MkdirAll(dir, 0750)
	if err != nil && !os.IsExist(err) {
		return err
	}

	if err := os.WriteFile(filePath, content, 0644); err != nil {
		return err
	}
	return nil
}
