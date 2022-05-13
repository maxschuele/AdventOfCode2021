package main

import (
	"AdventOfCode2021/util"
	"fmt"
)

type number struct {
	value  int
	marked bool
}

func main() {
	ox, co2, mask := 0, 0, 2048
	numbers := []number{}
	for _, v := range util.GetInputLines() {
		n := number{value: util.ParseBinary(v), marked: false}
		numbers = append(numbers, n)
	}
	numbers2 := make([]number, len(numbers))
	copy(numbers2, numbers)

	for mask > 0 {
		//oxygen generator
		zero, one := 0, 0
		for _, v := range numbers {
			if !v.marked && v.value&mask == 0 {
				zero++
			}
			if !v.marked && v.value&mask != 0 {
				one++
			}
		}
		for i := 0; i < len(numbers); i++ {
			n := &numbers[i]

			if one >= zero {
				if !n.marked && n.value&mask == 0 {
					n.marked = true
				}
			} else {
				if !n.marked && n.value&mask != 0 {
					n.marked = true
				}
			}
		}
		for _, v := range numbers {
			if !v.marked {
				ox = v.value
			}
		}
		//co2 scrubber
		zero, one = 0, 0
		for _, v := range numbers2 {
			if !v.marked && v.value&mask == 0 {
				zero++
			}
			if !v.marked && v.value&mask != 0 {
				one++
			}
		}
		for i := 0; i < len(numbers2); i++ {
			n := &numbers2[i]
			if one >= zero {
				if !n.marked && n.value&mask != 0 {
					n.marked = true
				}
			} else {
				if !n.marked && n.value&mask == 0 {
					n.marked = true
				}
			}
		}
		for _, v := range numbers2 {
			if !v.marked {
				co2 = v.value
			}
		}

		//shift mask right one bit
		mask = mask >> 1
	}

	

	fmt.Println(ox)
	fmt.Println(co2)
	fmt.Printf("Result: %v\n", ox*co2)
}
