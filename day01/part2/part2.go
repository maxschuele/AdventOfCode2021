package main

import (
	"AdventOfCode2021/util"
	"fmt"
)

func main() {
	input := util.GetInputLines()
	fmt.Println(input)
	var triples []int
	lc := len(input)

	for i := 0; i < lc -2; i++ {
		triples = append(triples, util.Atoi(input[i]) + util.Atoi(input[i+1]) + util.Atoi(input[i+2]))
	}
	//fmt.Println(triples)

	increases := 0
	for i := range triples {
		if i>0 && triples[i] > triples[i-1] {
			increases++
		}
	}
	fmt.Println(increases)
}