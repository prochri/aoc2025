package day03

import (
	"aoc2025/utils"
	"fmt"
	"math"
	"strconv"
)

func Solve() {
	lines, err := utils.ReadLines("day03/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 02:")
	fmt.Println("Part 1:", Part1(lines))
	fmt.Println("Part 2:", Part2(lines))
}

func Part1(lines []string) int {
	joltage := 0
	for _, line := range lines {
		joltage += findLineJoltage(line)
	}
	return joltage
}

func findLineJoltage(line string) int {
	maxJoltage := 0
	for i, c := range line {

		remainder := line[i+1:]
		if len(remainder) == 0 {
			continue
		}
		localMaxJoltage := '0'
		for _, localC := range remainder {
			localMaxJoltage = max(localMaxJoltage, localC)
		}

		maxJoltageValue := string(c) + string(localMaxJoltage)
		localMaxJoltageFull, err := strconv.Atoi(maxJoltageValue)
		if err != nil {
			panic(err)
		}

		if localMaxJoltageFull > maxJoltage {
			maxJoltage = localMaxJoltageFull
		}
	}
	return maxJoltage
}

func findLineJoltage2(line string, batteryCount int) int {
	if batteryCount == 0 {
		return 0
	}
	if len(line) == 0 && batteryCount == 0 {
		return 0
	}
	if len(line) < batteryCount {
		return -99999999999999999
	}
	maxJoltage := 0
	for i, c := range line {
		remainder := line[i+1:]
		myVoltageNumber, err := strconv.Atoi(string(c))
		if err != nil {
			panic(err)
		}
		myVoltage := int(float64(myVoltageNumber) * math.Pow10(batteryCount-1))
		if myVoltage < maxJoltage {
			continue
		}
		remainderJoltage := findLineJoltage2(remainder, batteryCount-1)
		totalVoltage := myVoltage + remainderJoltage
		if totalVoltage > maxJoltage {
			maxJoltage = totalVoltage
		}
	}
	return maxJoltage

}

func Part2Old(lines []string) int {
	joltage := 0
	for i, line := range lines {
		fmt.Println("line", i)
		joltage += findLineJoltage2(line, 12)
	}
	return joltage
}

func Part2(lines []string) int {
	joltage := 0
	channel := make(chan int)
	for _, line := range lines {
		go func(line string) {
			channel <- findLineJoltage2(line, 12)
		}(line)
	}
	for i := 0; i != len(lines); i++ {
		joltage += <-channel
	}
	return joltage
}
