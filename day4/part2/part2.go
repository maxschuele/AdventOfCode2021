package main

import (
	"AdventOfCode2021/util"
	"fmt"
	"strings"
)

type field struct {
	val    int
	marked bool
}

type board struct {
	fields [5][5]field
}

var boards []board
var numbers []int

func main() {
	numbers = util.ArrAtoi(strings.Split(util.GetInputLines()[0], ","))
	GetBoards()
	fmt.Println("", PlayBingo())

}
func PlayBingo() int {
	sum,index := 0, 0
	_ = index
	for _, v := range numbers {
		for h := 0; h < len(boards); h++ {
			for j := 0; j < 5; j++ {
				for k := 0; k < 5; k++ {
					if boards[h].fields[j][k].val == v {
						boards[h].fields[j][k].marked = true
					}
				}
			}
		}
		sum = CheckBingo()
		if sum != 0 {
			//return v * CalcResult(sum)
			fmt.Println(RemoveBoard(sum))
		}
	}
	return 0
}

func RemoveBoard(b int) board{
	removed := boards[b]
	boards = append(boards[:b], boards[b+1:]...)
	return removed
}

func CheckBingo() int {
	//check horizontal
	for b, v := range boards {
		for i := 0; i < 5; i++ {
			if v.fields[i][0].marked {
				for j := 1; j < 5; j++ {
					if !v.fields[i][j].marked {
						break
					}
					if j == 4 {
						return b
					}
				}
			}
		}
	}
	//check vertical
	for b, v := range boards {
		for i := 0; i < 5; i++ {
			if v.fields[0][i].marked {
				for j := 1; j < 5; j++ {
					if !v.fields[j][i].marked {
						break
					}
					if j == 4 {
						return b
					}
				}
			}
		}
	}
	return 0
}

func CalcResult(b int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !boards[b].fields[i][j].marked {
				sum += boards[b].fields[i][j].val
			}
		}
	}
	return sum
}

func GetBoards() {
	_, boardString, _ := strings.Cut(util.GetInput(), "\n")
	for boardString != "" {
		boardString = strings.TrimSpace(boardString)
		b := board{}
		for i := 0; i < 5; i++ {
			l, r, _ := strings.Cut(boardString, "\n")
			boardString = r
			row := util.ArrAtoi(strings.Fields(l))
			for j := 0; j < 5; j++ {
				b.fields[i][j] = field{row[j], false}
			}
		}
		boards = append(boards, b)
	}
}
