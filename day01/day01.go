package day01

import (
	"aoc2025/utils"
	"fmt"
	"strconv"
)

func Solve() {
	lines, err := utils.ReadLines("day01/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 01:")
	fmt.Println("Part 1:", Part1(lines))
	fmt.Println("Part 2:", Part2(lines))
}

func Part1(lines []string) int {
	currentPosition := 50
	countZeroes := 0

	for _, line := range lines {
		direction := line[0]
		clicks, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		switch direction {
		case 'L':
			currentPosition += clicks
		case 'R':
			currentPosition -= clicks
		default:
			panic("Invalid direction")
		}
		currentPosition %= 100
		if currentPosition == 0 {
			countZeroes++
		}
	}
	return countZeroes
}

func Part2(lines []string) int {
	currentPosition := 50
	countZeroes := 0

	for _, line := range lines {
		direction := line[0]
		clicks, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		sign := 1
		switch direction {
		case 'L':
			sign = -1
		case 'R':
			sign = 1
		default:
			panic("Invalid direction")
		}
		fmt.Println("Line", line)

		for i := 0; i != clicks; i++ {
			currentPosition += sign
			currentPosition %= 100
			if currentPosition == 0 {
				countZeroes++
			}
		}

	}
	return countZeroes
}
