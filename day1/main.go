package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("2022 Advent of Code Day 1")
	fmt.Println("--- Part 1 ---")

	if data, err := ioutil.ReadFile("input.txt"); err == nil {
		input := string(data)
		elves := strings.Split(input, "\n\n")
		max := -1
		for _, elv := range elves {
			current := 0
			for _, food := range strings.Split(elv, "\n") {
				f, _ := strconv.Atoi(food)
				current += f
			}
			max = int(math.Max(float64(max), float64(current)))
		}

		fmt.Println("Elv with most calories caried is carrying:", max, "calories.")
	}

	fmt.Println("--- Part 2 ---")

	if data, err := ioutil.ReadFile("input.txt"); err == nil {
		input := string(data)
		elves := strings.Split(input, "\n\n")
		values := []int{}
		for _, elv := range elves {
			current := 0
			for _, food := range strings.Split(elv, "\n") {
				f, _ := strconv.Atoi(food)
				current += f
			}
			values = append(values, current)
		}
		sort.Slice(values, func(i, j int) bool {
			return values[i] > values[j]
		})

		fmt.Println("Elv with most calories caried is carrying:",
			func(vals []int) int {
				s := 0
				for _, i := range vals {
					s += i
				}
				return s
			}(values[0:3]),
			"calories.")
	}

}
