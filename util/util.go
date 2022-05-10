package util

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func GetInput() string {
	path, e := os.Getwd()
	if e != nil {
		panic(e)
	}
	
	pathf, e := os.ReadFile(filepath.Join(filepath.Dir(path), "input.txt"))
	if e != nil {
		panic(e)
	}

	return strings.TrimSpace(string(pathf))
}

func GetInputLines() []string {
	return strings.Split(GetInput(), "\n")
}

func Atoi(input string) int {
	out, e := strconv.Atoi(input)
	if(e != nil){
		panic(e)
	}
	return out
}

func ParseBinary(input string) int {
	i, e := strconv.ParseInt(input, 2, 64)
	if e != nil {
		panic(e)
	}
	return int(i)
}