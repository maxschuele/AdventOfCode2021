package main

import (
	"AdventOfCode2021/util"
	"fmt"
)

//"AdventOfCode2021/util"
//"fmt"

func main () {
	input := util.GetInputLines()

	increases := 0
	for i, line := range input {
		if i>0 && util.Atoi(line) > util.Atoi(input[i-1]) {
			increases++
		}
	}
	fmt.Println(increases)
}

