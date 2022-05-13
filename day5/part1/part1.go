package main

import (
	"AdventOfCode2021/util"
	"fmt"
	"strings"
)

func main() {
	diagram := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		diagram[i] = make([]int, 1000)
	}
	input := util.GetInputLines()
	for _, v := range input {
		points := strings.Split(v, " -> ")
		x1, y1 := getPoint(points[0])
		x2, y2 := getPoint(points[1])

		if x1 == x2 {
			start, finish := getDirection(y1, y2)
			for start <= finish {
				diagram[x1][start]++
				start++
			}
		} else if y1 == y2 {
			start, finish := getDirection(x1, x2)
			for start <= finish {
				diagram[start][y1]++
				start++
			}
		}
		overlap := 0
		for _, row := range diagram {
			for _, p := range row {
				if p > 1 {
					overlap++
				}
			}
		}
		fmt.Println("overlaps:", overlap)
	}
}

func getDirection(x1 int, x2 int) (int, int) {
	if x1 > x2 {
		return x2, x1
	} else {
		return x1, x2
	}
}

func getPoint(s string) (int, int) {
	split := strings.Split(s, ",")
	return util.Atoi(split[0]), util.Atoi(split[1])
}
