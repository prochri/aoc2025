package main

import (
	"aoc2025/day01"
	"aoc2025/day02"
	"aoc2025/day03"
	"aoc2025/day04"
	"aoc2025/day05"
	"aoc2025/day06"
	"aoc2025/day07"
	"aoc2025/day08"
	"aoc2025/day09"
	"aoc2025/day10"
	"aoc2025/day11"
	"aoc2025/day12"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please specify a day to run (e.g., 'go run main.go 1')")
		return
	}

	day := os.Args[1]

	switch day {
	case "1":
		day01.Solve()
	case "2":
		day02.Solve()
	case "3":
		day03.Solve()
	case "4":
		day04.Solve()
	case "5":
		day05.Solve()
	case "6":
		day06.Solve()
	case "7":
		day07.Solve()
	case "8":
		day08.Solve()
	case "9":
		day09.Solve()
	case "10":
		day10.Solve()
	case "11":
		day11.Solve()
	case "12":
		day12.Solve()
	default:
		fmt.Printf("Unknown day: %s\n", day)
	}
}
