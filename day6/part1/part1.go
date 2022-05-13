package main

import (
	"AdventOfCode2021/util"
	"fmt"
	"strings"
)

var fish []int

func main() {
	fish = util.ArrAtoi(strings.Split(util.GetInput(), ","))
	PassDays(256)
	fmt.Println("fish count:", len(fish))
}

func PassDays(days int) {
	for i := 0; i < days; i++ {
		newFish := 0
		for i := range fish {
			if fish[i] == 0 {
				fish[i] = 6
				newFish++
			} else {
				fish[i]--
			}
		}
		for i := 0; i < newFish; i++ {
			fish = append(fish, 8)
		}
		fmt.Println(fish)
	}
}
