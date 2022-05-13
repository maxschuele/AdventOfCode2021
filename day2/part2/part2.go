package main

import (
	"AdventOfCode2021/util"
	"fmt"
	"strings"
)

func main() {
	horizontal := 0
	depth := 0
	aim := 0

	input := util.GetInputLines()

	for _, v := range input {
		line := strings.Split(v, " ")

		command := line[0]
		value := util.Atoi(line[1])

		if command == "forward" {
			horizontal += value
			depth += (value*aim)
		}

		if command == "up" {
			aim -= value
		}

		if line[0] == "down" {
			aim += value
		
		}
	}

	fmt.Printf("horizontal: %v\n",horizontal)
	fmt.Printf("depth: %v\n", depth)
	fmt.Printf("horizontal * depth: %v\n", horizontal*depth)
}