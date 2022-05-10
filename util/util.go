package util

import (
	"os"
	"path/filepath"
	"strings"
)

func GetInput(day string) string {
	path, _ := os.Getwd()
	pathf, _ := os.ReadFile(filepath.Join(path, day, "input.txt"))
	return strings.TrimSpace(string(pathf))
}