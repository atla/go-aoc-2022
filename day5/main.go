package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Stack []rune
type Field struct {
	Columns []Stack
	Size    int
}

func makeField(cols int) *Field {
	columns := []Stack{}
	for i := 0; i < cols; i++ {
		columns = append(columns, Stack{})
	}
	return &Field{
		Columns: columns,
		Size:    cols,
	}
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
func (s *Stack) Push(r rune) {
	*s = append(*s, r)
}

func (s *Stack) PushMany(r ...rune) {
	*s = append(*s, r...)
}

func (s *Stack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return rune(' '), false
	} else {
		idx := len(*s) - 1
		r := (*s)[idx]
		*s = (*s)[:idx]
		return r, true
	}
}

func (s *Stack) PopMany(count int) ([]rune, bool) {
	if s.IsEmpty() {
		return []rune{}, false
	} else {
		idx := len(*s)
		r := (*s)[idx-count : idx]
		*s = (*s)[:idx-count]
		return r, true
	}
}

func main() {
	fmt.Println("2022 Advent of Code Day 5")

	fmt.Println("--- Part 1 ---")

	findTopElements(func(field *Field, amount int, from int, to int) *Field {
		for p := 0; p < amount; p++ {
			r, _ := field.Columns[from-1].Pop()
			field.Columns[to-1].Push(r)
		}
		return field
	}, 9)

	fmt.Println("--- Part 2 ---")

	findTopElements(func(field *Field, amount int, from int, to int) *Field {
		r, _ := field.Columns[from-1].PopMany(amount)
		field.Columns[to-1].PushMany(r...)
		return field
	}, 9)
}

func findTopElements(move func(field *Field, amount int, from int, to int) *Field, cols int) {

	if data, err := ioutil.ReadFile("input.txt"); err == nil {
		input := strings.Split(string(data), "\n\n")		
		field := parseInitialStack(input[0], cols)		

		for _, ins := range strings.Split(input[1], "\n") {
			amount, from, to := parseInstruction(ins)
			field = move(field, amount, from, to)
		}
		field.printTopElements()
	}
}

func parseInstruction(ins string) (int, int, int) {
	split := strings.Split(ins, " ")
	v1, _ := strconv.Atoi(split[1])
	v2, _ := strconv.Atoi(split[3])
	v3, _ := strconv.Atoi(split[5])
	return v1, v2, v3
}

func (field *Field) printTopElements() {
	for _, e := range field.Columns {
		fmt.Print(string(e[len(e)-1]))
	}
	fmt.Println("")
}

func parseInitialStack(initialStack string, cols int) *Field {

	f := makeField(cols)

	lines := strings.Split(initialStack, "\n")
	for line := len(lines) - 2; line > -1; line-- {
		for i := 0; i < cols; i++ {
			r := rune(lines[line][i*4+1])
			if r != ' ' {
				f.Columns[i].Push(r)
			}

			//	fmt.Print(strconv.QuoteRune(r), " - ")
		}
		//		fmt.Print("\n")
	}
	return f
}
