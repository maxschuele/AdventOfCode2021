package main

import (
	"AdventOfCode2021/util"
	"fmt"
	"strings"
)

func main() {
	crabs := util.ArrAtoi(strings.Split(util.GetInput(), ","))

	maxHeight := 0
	for i := 1; i < len(crabs); i++ {
		if crabs[i-1] < crabs[i] {
			maxHeight = crabs[i]
		}
	}

	costs := make([]int, maxHeight+1)
	n := 0
	for i := range costs {
		for j := range crabs {
			if i > crabs[j] {
				n = i - crabs[j]
				costs[i] += n*(n+1)/2
			} else if  i < crabs[j] {
				n = crabs[j] - i
				costs[i] += n*(n+1)/2
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