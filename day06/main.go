package day06

import (
	"aoc2025/utils"
	"fmt"
	"regexp"
)

func Solve() {
	lines, err := utils.ReadLines("day06/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 06:")
	fmt.Println("Part 1:", Part1(lines))
	fmt.Println("Part 2:", Part2(lines))
}

type Equation struct {
	numbers    []int
	startIndex int
	endIndex   int
	operator   string
}

// Source - https://stackoverflow.com/a
// Posted by icza, modified by community. See post 'Timeline' for change history
// Retrieved 2025-12-06, License - CC BY-SA 4.0

func filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func notIsEmpty(s string) bool {
	return len(s) > 0
}

func parseInput(lines []string) (equations []Equation) {
	regex := regexp.MustCompile(`\s+`)
	lastLine := lines[len(lines)-1]
	operators := filter(regex.Split(lastLine, -1), notIsEmpty)
	equations = make([]Equation, 0)
	for _, operator := range operators {
		equations = append(equations, Equation{operator: operator})
	}

	for _, line := range lines[:len(lines)-1] {
		numberStrs := filter(regex.Split(line, -1), notIsEmpty)
		for i, numberStr := range numberStrs {
			number := utils.ParseIntUnsafe(numberStr)
			equations[i].numbers = append(equations[i].numbers, number)
		}
	}
	return
}

func applyEquation(equation Equation) int {
	result := 0
	if equation.operator == "*" {
		result = 1
	}
	for _, number := range equation.numbers {
		if equation.operator == "+" {
			result += number
		} else if equation.operator == "*" {
			result *= number
		}
	}
	return result
}

func Part1(lines []string) int {
	equations := parseInput(lines)
	total := 0
	for _, equation := range equations {
		total += applyEquation(equation)
	}
	return total
}

func parseInput2(lines []string) (equations []Equation) {
	lastLine := lines[len(lines)-1]
	otherLines := lines[:len(lines)-1]
	equations = make([]Equation, 0)
	for i, char := range lastLine {
		if char == '+' || char == '*' {
			if len(equations) > 0 {
				equations[len(equations)-1].endIndex = i - 1
			}
			equations = append(equations, Equation{numbers: make([]int, 0),
				startIndex: i,
				endIndex:   -1,

				operator: string(char),
			})
		}
	}
	equations[len(equations)-1].endIndex = len(lastLine) - 1
	for equationIndex, equation := range equations {
		for i := equation.startIndex; i <= equation.endIndex; i++ {
			newNumber := ""
			for _, line := range otherLines {
				char := line[i]
				if char == ' ' {
					continue
				}
				newNumber += string(char)
			}
			if newNumber != "" {
				equations[equationIndex].numbers = append(equations[equationIndex].numbers, utils.ParseIntUnsafe(newNumber))
			}
		}
	}

	return equations
}

func Part2(lines []string) int {
	equations := parseInput2(lines)
	total := 0
	for _, equation := range equations {
		total += applyEquation(equation)
	}
	return total
}
