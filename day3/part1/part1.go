package main

import (
	"AdventOfCode2021/util"
	"fmt"
)

func main() {
	gamma, epsilon, mask := 0, 0, 2048

	for mask > 0 {
		zero, one := 0, 0
		for _, v := range util.GetInputLines() {
			if util.ParseBinary(v)&mask == 0 {
				zero++
			} else {
				one++
			}
		}
		if one > zero {
			gamma = gamma | mask
		} else {
			epsilon = epsilon | mask
		}
		mask = mask >> 1
	}
	fmt.Printf("Result: %v\n", gamma*epsilon)
}
