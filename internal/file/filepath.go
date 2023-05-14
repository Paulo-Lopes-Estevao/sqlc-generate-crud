package file

import (
	"os"
	"path/filepath"
)

func checkFileExists(dir, filename string) (bool, error) {
	filePath := filepath.Join(dir, filename)
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}
