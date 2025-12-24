package day12

import (
	"aoc2025/utils"
	"fmt"
	"strings"
)

func Solve() {
	lines, err := utils.ReadLines("day12/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 12:")
	fmt.Println("Part 1:", Part1(lines))
	fmt.Println("Part 2:", Part2(lines))
}

// assume: all gifts are 3x3
type Present struct {
	index         int
	shape         [][]bool
	occupiedSlots int
}

type Region struct {
	width         int
	height        int
	presentCounts []int
}

func nextEmptyLineIndex(lines []string) int {
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			return i
		}
	}
	return -1
}

func parse(lines []string) ([]Present, []Region) {
	presents := make([]Present, 0)
	regions := make([]Region, 0)

	// parse presents
	for {
		nextEmptyIndex := nextEmptyLineIndex(lines)
		if nextEmptyIndex == -1 {
			break
		}
		index := utils.ParseIntUnsafe(strings.Trim(lines[0], ":"))
		shape := make([][]bool, 0)
		occupiedSlots := 0
		for i := 1; i < nextEmptyIndex; i++ {
			shapeLine := make([]bool, 0)
			for _, char := range strings.Trim(lines[i], " ") {
				occupied := char == '#'
				if occupied {
					occupiedSlots++
				}
				shapeLine = append(shapeLine, occupied)
			}
			shape = append(shape, shapeLine)
		}
		presents = append(presents, Present{
			index:         index,
			shape:         shape,
			occupiedSlots: occupiedSlots,
		})

		lines = lines[nextEmptyIndex+1:]
	}

	for _, line := range lines {
		lineSplit := strings.Split(line, ": ")
		dimensions := strings.Split(lineSplit[0], "x")
		width := utils.ParseIntUnsafe(dimensions[0])
		height := utils.ParseIntUnsafe(dimensions[1])
		presentCounts := make([]int, 0)
		for count := range strings.SplitSeq(strings.Trim(lineSplit[1], " "), " ") {
			presentCounts = append(presentCounts, utils.ParseIntUnsafe(count))
		}
		regions = append(regions, Region{
			width:         width,
			height:        height,
			presentCounts: presentCounts,
		})
	}

	return presents, regions
}

func definitlyDoesNotFit(region Region, presents []Present) bool {
	regionSize := region.width * region.height
	presentSize := 0
	for i, presentCount := range region.presentCounts {
		presentSize += presentCount * presents[i].occupiedSlots
	}
	return regionSize < presentSize
}

func Part1(lines []string) int {
	presents, regions := parse(lines)
	fmt.Println(presents)
	fmt.Println(regions)
	count := 0
	for _, region := range regions {
		if definitlyDoesNotFit(region, presents) {
			continue
		}
		count++

	}
	return count
}

func Part2(lines []string) int {
	// Placeholder for Part 2
	return 0
}
