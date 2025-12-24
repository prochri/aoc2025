package day07

import (
	"aoc2025/utils"
	"fmt"
)

func Solve() {
	lines, err := utils.ReadLines("day07/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 07:")
	fmt.Println("Part 1:", Part1(lines))
	fmt.Println("Part 2:", Part2(lines))
}

func Part1(lines []string) int {
	currentSet := make(map[int](bool))
	for i, char := range lines[0] {
		if char == 'S' {
			currentSet[i] = true
		}
	}
	count := 0
	for _, line := range lines[1:] {
		// fmt.Println(currentSet)
		nextSet := make(map[int](bool))
		for i, _ := range currentSet {
			if line[i] == '^' {
				nextSet[i+1] = true
				nextSet[i-1] = true
				if i < 0 || i >= len(line) {
					fmt.Println("oops out of bounds")
				}
				if line[i-1] == '^' {
					fmt.Println("oops left neighbor is also splitter")
				}
				if line[i+1] == '^' {
					fmt.Println("oops right neighbor is also splitter")
				}
				count++
			} else {
				nextSet[i] = true
			}
		}
		currentSet = nextSet
	}

	return count
}

type TimelineKey struct {
	currentParticlePosition int
	remainingLines          int
}

var timelineCache = make(map[TimelineKey]int)

func timelines(currentParticlePosition int, lines []string) int {
	cacheKey := TimelineKey{currentParticlePosition, len(lines)}
	if timelineCache[cacheKey] != 0 {
		return timelineCache[cacheKey]
	}
	if len(lines) == 0 {
		return 1
	}
	nextLine := lines[0]
	result := 0
	if nextLine[currentParticlePosition] == '^' {
		// channel := make(chan int)
		// channel <- timelines(currentParticlePosition-1, lines[1:])
		// channel <- timelines(currentParticlePosition+1, lines[1:])
		result = timelines(currentParticlePosition-1, lines[1:]) + timelines(currentParticlePosition+1, lines[1:])
	} else {
		result = timelines(currentParticlePosition, lines[1:])
	}
	timelineCache[cacheKey] = result
	return result
}

func Part2(lines []string) int {
	startTimeline := 0
	for i, char := range lines[0] {
		if char == 'S' {
			startTimeline = i
		}
	}
	// Placeholder for Part 2
	return timelines(startTimeline, lines)
}
