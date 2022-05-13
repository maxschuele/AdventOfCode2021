package main

import (
	"AdventOfCode2021/util"
	"fmt"
	"strings"
)

func main() {
	horizontal := 0
	depth := 0

	input := util.GetInputLines()

	for _, v := range input {
		line := strings.Split(v, " ")
		command := line[0]
		value := util.Atoi(line[1])

		if command == "forward" {
			horizontal += value
		}

		if command == "up" {
			depth -= value
		}

		if command == "down" {
			depth += value
		
		}
	}

	fmt.Printf("horizontal: %v\n",horizontal)
	fmt.Printf("depth: %v\n", depth)
	fmt.Printf("horizontal * depth: %v\n", horizontal*depth)
}