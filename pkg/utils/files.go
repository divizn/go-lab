package utils

import (
	"os"
	"path/filepath"
)

// GetCWD returns the current working directory
func GetCWD() (string, error) {
	return os.Getwd()
}

// JoinWithCWD joins a relative path with the current working directory
func JoinWithCWD(relativePath string) (string, error) {
	cwd, err := GetCWD()
	if err != nil {
		return "", err
	}
	return filepath.Join(cwd, relativePath), nil
}
