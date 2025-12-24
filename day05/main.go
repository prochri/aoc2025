package day05

import (
	"aoc2025/utils"
	"fmt"
	"strings"

	"github.com/hij1nx/go-indexof"
)

func Solve() {
	lines, err := utils.ReadLines("day05/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 05:")
	fmt.Println("Part 1:", Part1(lines))
	fmt.Println("Part 2:", Part2(lines))
}

type Range struct {
	start int
	end   int
}

func (r Range) overlapCount(other Range) int {
	if r.start >= other.end {
		return 0
	}
	if r.end <= other.start {
		return 0
	}
	// my range completely contains the other range
	if r.start <= other.start && r.end >= other.end {
		return other.itemCount()
	}
	if other.start <= r.start && other.end >= r.end {
		return r.itemCount()
	}
	// I start before the other range starts, but end after before the other range ends
	if r.start <= other.start {
		return Range{start: other.start, end: r.end}.itemCount()
	}
	if other.start <= r.start {
		return Range{start: r.start, end: other.end}.itemCount()
	}
	fmt.Println("I might have missed something :(", r, other)

	return 0
}

func (r Range) itemCount() int {
	return r.end - r.start + 1
}

func parseInput(lines []string) (ranges []Range, ids []int) {
	index := indexof.IndexOf("", lines)
	rangesUnparsed := lines[:index]
	idsUnparsed := lines[index+1:]

	for _, line := range rangesUnparsed {
		rangeParts := strings.Split(line, "-")
		ranges = append(ranges, Range{start: utils.ParseIntUnsafe(rangeParts[0]), end: utils.ParseIntUnsafe(rangeParts[1])})
	}

	for _, line := range idsUnparsed {
		ids = append(ids, utils.ParseIntUnsafe(line))
	}
	return
}

func Part1(lines []string) int {
	ranges, ids := parseInput(lines)
	count := 0
	for _, id := range ids {
		isFresh := false
		for _, r := range ranges {
			if id >= r.start && id <= r.end {
				isFresh = true
				break
			}
		}
		if isFresh {
			count++
		}
	}
	return count
}

func addNonOverlappingRangeSingle(r Range, other Range) []Range {
	// no overlap
	if (r.start < other.start && r.end < other.start) ||
		(r.start > other.end && r.end > other.end) {
		return []Range{r}
	}
	ranges := make([]Range, 0)
	// take the first part
	if r.start < other.start && r.end >= other.start {
		ranges = append(ranges, Range{start: r.start, end: other.start - 1})
	}
	if r.start <= other.end && r.end > other.end {
		ranges = append(ranges, Range{start: other.end + 1, end: r.end})
	}
	return ranges
}

func addNonOverlappingRange(r Range, rangesClean []Range) []Range {
	newRanges := make([]Range, 0)
	// start with my new range
	newRanges = append(newRanges, r)
	for _, rr := range rangesClean {
		nextNewRanges := make([]Range, 0)
		for _, newRange := range newRanges {
			nextNewRanges = append(nextNewRanges, addNonOverlappingRangeSingle(newRange, rr)...)
		}
		newRanges = nextNewRanges
	}
	// assert that the new ranges are not overlapping anything else
	for _, newRange := range newRanges {
		for _, otherRange := range rangesClean {
			if newRange.overlapCount(otherRange) > 0 {
				fmt.Println("overlapping", newRange, otherRange)
				panic("overlapping")
			}
		}
	}
	newRanges = append(newRanges, rangesClean...)

	return newRanges
}

func Part2(lines []string) int {
	ranges, _ := parseInput(lines)
	count := 0
	rangesClean := make([]Range, 0)
	for _, r := range ranges {
		rangesClean = addNonOverlappingRange(r, rangesClean)
	}
	for _, r := range rangesClean {
		count += r.itemCount()
	}
	// Placeholder for Part 2
	return count
}
