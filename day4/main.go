package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type elf struct {
	L int
	U int
}

func makeElf(input []string) elf {
	v1, _ := strconv.Atoi(input[0])
	v2, _ := strconv.Atoi(input[1])
	return elf{
		v1,
		v2,
	}
}

func (e1 elf) contains(e2 elf) bool {
	return e1.L <= e2.L && e1.U >= e2.U
}

func (e1 elf) intersects(e2 elf) bool {
	return (e1.L <= e2.L && e1.U >= e2.L) ||
		(e2.U >= e1.L && e2.U <= e1.U)
}

func main() {
	fmt.Println("2022 Advent of Code Day 4")
	fmt.Println("--- Part 1&2 ---")

	if data, err := ioutil.ReadFile("input.txt"); err == nil {
		duplicates, d2, pairs := 0, 0, strings.Split(string(data), "\n")

		for _, p := range pairs {
			e1, e2 := strings.Split(p, ",")[0], strings.Split(p, ",")[1]
			elf1, elf2 := makeElf(strings.Split(e1, "-")), makeElf(strings.Split(e2, "-"))
			if elf1.contains(elf2) || elf2.contains(elf1) {
				duplicates++
			}
			if elf1.intersects(elf2) || elf2.intersects(elf1) {
				d2++
			}
		}
		fmt.Println("Duplicate assignments: ", duplicates)
		fmt.Println("intersectign assignments: ", d2)
	}
}
