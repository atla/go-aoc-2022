package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

func priority(r rune) int32 {
	if unicode.IsUpper(r) {
		return r - 'A' + 27
	}
	return r - 'a' + 1
}

func main() {
	fmt.Println("2022 Advent of Code Day 3")
	fmt.Println("--- Part 1 ---")

	if data, err := ioutil.ReadFile("input.txt"); err == nil {
		sum, backpacks := 0, strings.Split(string(data), "\n")
		for _, backpack := range backpacks {
			items := map[rune]bool{}
			for i, r := range backpack {
				if i < len(backpack)/2 {
					items[r] = true
				} else {
					if _, ok := items[r]; ok {
						sum += int(priority(r))
						//fmt.Println("Duplicate is ", strconv.QuoteRune(r), " (", priority(r), ")")
						break
					}
				}
			}
		}
		fmt.Println("Sum is ", sum)
	}

	fmt.Println("--- Part 2 ---")

	if data, err := ioutil.ReadFile("input.txt"); err == nil {
		sum, backpacks := 0, strings.Split(string(data), "\n")

		for i := 0; i < len(backpacks)/3; i++ {
			group, items1, items2 := backpacks[i*3:(i+1)*3], map[rune]bool{}, map[rune]bool{}
			for _, r := range group[0] {
				items1[r] = true
			}
			for _, r := range group[1] {
				items2[r] = true
			}
			for _, r := range group[2] {
				_, ok1 := items1[r]
				_, ok2 := items2[r]
				if ok1 && ok2 {
					sum += int(priority(r))
					//fmt.Println("Duplicate is ", strconv.QuoteRune(r), " (", priority(r), ")")
					break
				}
			}
		}
		fmt.Println("Sum is ", sum)
	}

}
