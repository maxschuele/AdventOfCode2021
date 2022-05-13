package main

import (
	"AdventOfCode2021/util"
	"fmt"
	"strings"
)

func main() {
	crabs := util.ArrAtoi(strings.Split(util.GetInput(), ","))

	max := 0
	for i := 1; i < len(crabs); i++ {
		if crabs[i-1] < crabs[i] {
			max = crabs[i]
		}
	}

	costs := make([]int, max+1)
	for i := range costs {
		for j := range crabs {
			if i > crabs[j] {
				costs[i] += i - crabs[j]
			} else if  i < crabs[j] {
				costs[i] += crabs[j] - i
			}
		}
	}
	
	minCost := 0
	for i := 1; i < len(costs); i++ {
		if costs[i-1] > costs[i] {
			minCost = costs[i]
		}
	}
	fmt.Println("min cost:",minCost)

}
