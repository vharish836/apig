package util

import (
	"os"
)

// FileExists ...
func FileExists(dir string) bool {
	_, err := os.Stat(dir)
	return err == nil
}

// Mkdir ...
func Mkdir(dir string) error {
	return os.MkdirAll(dir, os.ModePerm)
}
