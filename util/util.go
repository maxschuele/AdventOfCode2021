package util

import (
	"os"
	"path/filepath"
	"sort"
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

func ArrAtoi(input []string) []int {
	var out []int
	for _, v := range input {
		out = append(out, Atoi(v))
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

func SortString(s string) string {
	sa := strings.Split(s, "")
	sort.Strings(sa)
	return strings.Join(sa,"")
}

func SplitToString(a []int, sep string) string {
	if len(a) == 0 {
		return ""
	}
	
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}
	return strings.Join(b, sep)
}