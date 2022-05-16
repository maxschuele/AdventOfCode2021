package main

import (
	"AdventOfCode2021/util"
	"fmt"
	"strings"
)

func main() {
	counter := 0
	for _, line := range util.GetInputLines() {
		_, lineCut, _ := strings.Cut(line,"|")
		values := strings.Split(strings.TrimSpace(lineCut), " ")
		for _, value := range values{
			length := len(value)
			if length == 2 || length == 3 || length == 4 || length == 7 {
				counter++
			}
		}
	}
	fmt.Println(counter)	
}