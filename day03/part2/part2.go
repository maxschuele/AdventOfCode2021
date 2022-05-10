package main

import (
	"AdventOfCode2021/util"
	"fmt"
)

func main() {
	ox, co2, mask := 0, 0, 2048
	var numbers []int
	for _, v := range util.GetInputLines() {
		numbers = append(numbers, util.ParseBinary(v))
	}

	//check for common bit
	for mask > 0 {
		zero, one := 0, 0
		for _, v := range numbers {
			if v&mask == 0 {
				zero++
			} else {
				one++
			}
		}
		fmt.Printf("%d %d\n", one, zero)
		for i, v := range numbers {
			if one >= zero {
				if v&mask == 0 {
					numbers = append(numbers[:i], numbers[i+1:]...)
				}
			} else {
				if v&mask == 1 {
					numbers = append(numbers[:i], numbers[i+1:]...)
				}
			}
		}
		mask = mask >> 1
	}
	fmt.Println(numbers)
	fmt.Printf("Result: %v\n", ox*co2)
}
