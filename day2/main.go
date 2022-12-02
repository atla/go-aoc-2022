package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println("2022 Advent of Code Day 2")
	fmt.Println("--- Part 1 ---")

	lookup := map[string]int{
		"A X": 4, "A Y": 8, "A Z": 3, // 1+3, 2+6, 3+0
		"B X": 1, "B Y": 5, "B Z": 9, // 1+0, 2+3, 3+6
		"C X": 7, "C Y": 2, "C Z": 6, // 1+6, 2+6, 3+3
	}
	if data, err := ioutil.ReadFile("input.txt"); err == nil {
		points, rounds := 0, strings.Split(string(data), "\n")
		for _, round := range rounds {
			points += lookup[round]
		}
		fmt.Println("My total score according to the given strategy guide is: ", points)
	}

	fmt.Println("--- Part 2 ---")

	prune := map[string]string{
		"A X": "A Z", "A Y": "A X", "A Z": "A Y",
		"B X": "B X", "B Y": "B Y", "B Z": "B Z",
		"C X": "C Y", "C Y": "C Z", "C Z": "C X",
	}
	if data, err := ioutil.ReadFile("input.txt"); err == nil {
		points, rounds := 0, strings.Split(string(data), "\n")
		for _, round := range rounds {
			points += lookup[prune[round]]
		}
		fmt.Println("My total score according to NEW rules is: ", points)
	}
}
