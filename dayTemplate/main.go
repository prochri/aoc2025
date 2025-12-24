package daytemplate

import (
	"aoc2025/utils"
	"fmt"
	"strings"
)

func Solve() {
	lines, err := utils.ReadLines("day02/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 02:")
	fmt.Println("Part 1:", Part1(lines))
	fmt.Println("Part 2:", Part2(lines))
}

func Part1(lines []string) int {
	count := 0
	for _, line := range lines {
		if strings.Contains(line, "a") {
			count++
		}
	}
	return count
}

func Part2(lines []string) int {
	// Placeholder for Part 2
	return 0
}
