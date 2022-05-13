package main

import (
	"AdventOfCode2021/util"
	"fmt"
	"strings"
)

var fish []int

func main() {
	fish = make([]int, 9)
	input := strings.Split(util.GetInput(),",")
	for _, v := range input {
		fish[util.Atoi(v)]++
	}
	PassDays(256)
	fmt.Println("",fish)
	fishCount := 0
	for i:= 0; i < 9; i++ {
		fishCount += fish[i]
	}
	fmt.Println("fish count:",fishCount)
}

func PassDays(days int) {
	for i := 0; i < days; i++ {
		newFish := fish[0]
		for j := 0; j < 8; j++ {
			fish[j] = fish[j+1]
		}
		fish[6] += newFish
		fish[8] = newFish
	}
}
