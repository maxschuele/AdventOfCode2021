package main

import (
	"AdventOfCode2021/util"
	"fmt"
	"sort"
	"strings"
)

func main() {
	total := 0
	_ = total
	for _, line := range util.GetInputLines() {
		patternsCut, valuesCut, _ := strings.Cut(line, "|")
		patterns := strings.Split(strings.TrimSpace(patternsCut), " ")
		values := strings.Split(strings.TrimSpace(valuesCut), " ")
		//sort patterns and values
		for i := 0; i < len(patterns); i++ {
			patterns[i] = util.SortString(patterns[i])
		}
		for j := 0; j < len(values); j++ {
			values[j] = util.SortString(values[j])
		}
		sort.Slice(patterns, func(i, j int) bool {
			return len(patterns[i]) < len(patterns[j])
		})
		numbers := make(map[string]int, 10)
		numbers[patterns[0]] = 1
		numbers[patterns[1]] = 7
		numbers[patterns[2]] = 4
		numbers[patterns[9]] = 8

		segments := make([]byte, 7)
		//identify segments 0 and 1
		candidates := [2]byte{patterns[0][0], patterns[0][1]}
		for i := 6; i < 9; i++ {
			if Contains(patterns[i], candidates[0]) && !Contains(patterns[i], candidates[1]) {
				segments[0] = candidates[1]
				segments[1] = candidates[0]
			} else if Contains(patterns[i], candidates[1]) && !Contains(patterns[i], candidates[0]) {
				segments[0] = candidates[0]
				segments[1] = candidates[1]
			}
		}
		//identify segment 2
		patterns[1] = strings.Replace(patterns[1], string(segments[0]), "", 1)
		patterns[1] = strings.Replace(patterns[1],string(segments[1]), "", 1)
		segments[2] = patterns[1][0]
		
		//identify segments 3 and 4
		patterns[2] = strings.Replace(patterns[2], string(segments[0]), "", 1)
		patterns[2] = strings.Replace(patterns[2],string(segments[1]), "", 1)
		candidates = [2]byte{patterns[2][0], patterns[2][1]}
		for i := 3; i < 6; i++ {
			if Contains(patterns[i], candidates[0]) && !Contains(patterns[i], candidates[1]) {
				segments[3] = candidates[0]
				segments[4] = candidates[1]
				if Contains(patterns[i], segments[1]) {
					numbers[patterns[i]] = 3
				} else {
					numbers[patterns[i]] = 2
				}
			} else if Contains(patterns[i], candidates[1]) && !Contains(patterns[i], candidates[0]) {
				segments[3] = candidates[1]
				segments[4] = candidates[0]
				if Contains(patterns[i], segments[1]) {
					numbers[patterns[i]] = 3
				} else {
					numbers[patterns[i]] = 2
				}
			} else {
				numbers[patterns[i]] = 5
			}
		}
		
		//identify segments 5 and 6
		for i:= 6; i < 9; i++ {
			if !Contains(patterns[i], segments[3]) {
				numbers[patterns[i]] = 0
			} else if Contains(patterns[i], segments[0]) {
				numbers[patterns[i]] = 9
			} else {
				numbers[patterns[i]] = 6
			}
		}
		value,mult := 0, 1000
		for _, v := range values {
			value += numbers[v] * mult
			mult = mult / 10
		}
		fmt.Println(value)
		total += value
	}
	fmt.Println("total:",total)
}

func Contains(s string, b byte) bool {
	return strings.Contains(s, string(b))
}
