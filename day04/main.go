package day04

import (
	"aoc2025/utils"
	"fmt"
)

func Solve() {
	lines, err := utils.ReadLines("day04/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 04:")
	fmt.Println("Part 1:", Part1(lines))
	fmt.Println("Part 2:", Part2(lines))
}

func parseLines(lines []string) [][]rune {
	rows := make([][]rune, 0)
	for _, line := range lines {
		row := make([]rune, 0)
		for _, c := range line {
			row = append(row, c)
		}
		rows = append(rows, row)
	}
	return rows
}

type Position struct {
	row    int
	column int
}

func findFreeRolls(rows [][]rune) []Position {
	freeRolls := make([]Position, 0)
	for rowNumber, row := range rows {
		for columnNumber, c := range row {
			if c != '@' {
				continue
			}
			if neigbourOccupations(rows, rowNumber, columnNumber) < 4 {
				freeRolls = append(freeRolls, Position{rowNumber, columnNumber})
			}

		}
	}
	return freeRolls
}

func neigbourOccupations(rows [][]rune, row, column int) int {
	neighbours := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if isOccupiedWithRoll(rows, row+i, column+j) {
				neighbours++
			}
		}
	}
	return neighbours
}

func isOccupiedWithRoll(rows [][]rune, row, column int) bool {
	if row < 0 || row >= len(rows) {
		return false
	}
	if column < 0 || column >= len(rows[row]) {
		return false
	}
	return rows[row][column] == '@'
}

func Part1(lines []string) int {
	rows := parseLines(lines)
	return len(findFreeRolls(rows))
}

func removeOccupiedCells(rows [][]rune, positions []Position) [][]rune {
	for _, position := range positions {
		rows[position.row][position.column] = '.'
	}
	return rows
}

func Part2(lines []string) int {
	rows := parseLines(lines)
	freeRolls := findFreeRolls(rows)
	total := len(freeRolls)
	for true {
		rows = removeOccupiedCells(rows, freeRolls)
		freeRolls = findFreeRolls(rows)
		total += len(freeRolls)
		if len(freeRolls) == 0 {
			break
		}
	}

	// Placeholder for Part 2
	return total
}
