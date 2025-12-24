package day02

import (
	"aoc2025/utils"
	"fmt"
	"slices"
	"strconv"
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

type Range struct {
	start    int
	startStr string
	end      int
	endStr   string
}

func GetRangesLines(line []string) []Range {
	return GetRanges(line[0])
}

func GetRanges(line string) []Range {
	rangeStrs := strings.Split(line, ",")
	ranges := make([]Range, len(rangeStrs))

	for i, r := range rangeStrs {
		rangeParts := strings.SplitN(r, "-", 2)
		start, err := strconv.Atoi(rangeParts[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(rangeParts[1])
		if err != nil {
			panic(err)
		}
		ranges[i] = Range{
			startStr: rangeParts[0],
			start:    start,
			endStr:   rangeParts[1],
			end:      end,
		}
	}
	return ranges
}

func (r Range) FindIds() (ids []int) {
	digitsStart := len(r.startStr)
	digitsEnd := len(r.endStr)
	if digitsStart%2 != 0 && digitsEnd%2 != 0 {
		return
	}
	digits := digitsStart
	if digitsStart%2 == 0 {
		digits = digitsStart
	} else if digitsEnd%2 == 0 {
		digits = digitsEnd
	}
	digitsHalf := digits / 2
	for _, sequence := range generateSequence(digitsHalf, false) {
		if r.isInRange(sequence + sequence) {
			ids = append(ids, sequenceToInt(sequence+sequence))
		}
	}
	return
}

func (r Range) FindIds2() (ids []int) {
	digitsStart := len(r.startStr)
	digitsEnd := len(r.endStr)

	newIds := r.FindIdsForDigits(digitsStart)
	ids = append(ids, newIds...)
	if digitsEnd != digitsStart {
		newIds = r.FindIdsForDigits(digitsEnd)
		ids = append(ids, newIds...)
	}
	return
}

func (r Range) FindIdsForDigits(digits int) (ids []int) {
	for repeats := 2; repeats <= digits; repeats++ {
		if digits%repeats != 0 {
			continue
		}

		sequences := generateSequence(digits/repeats, false)
		for _, sequence := range sequences {
			sequenceFull := strings.Repeat(sequence, repeats)
			sequenceInt := sequenceToInt(sequenceFull)
			if r.isInRange(sequenceFull) && !slices.Contains(ids, sequenceInt) {
				ids = append(ids, sequenceToInt(sequenceFull))
			}
		}
	}
	return
}

func generateSequence(digits int, includeZero bool) []string {
	if digits == 0 {
		return []string{""}
	}
	nextSequences := generateSequence(digits-1, true)
	sequences := make([]string, 0)
	for i := 0; i != 10; i++ {
		if !includeZero && i == 0 {
			continue
		}
		for _, nextSequence := range nextSequences {
			sequences = append(sequences, fmt.Sprintf("%d%s", i, nextSequence))
		}
	}
	return sequences
}

func sequenceToInt(sequence string) int {
	number, err := strconv.Atoi(sequence)
	if err != nil {
		panic(err)
	}
	return number
}

func (r Range) isInRange(sequence string) bool {
	number := sequenceToInt(sequence)
	return number >= r.start && number <= r.end
}

func Part1(lines []string) int {
	ranges := GetRangesLines(lines)
	sum := 0
	for _, r := range ranges {
		ids := r.FindIds()
		for _, id := range ids {
			sum += id
		}
	}
	return sum
}

func Part2(lines []string) int {
	ranges := GetRangesLines(lines)
	sum := 0
	for _, r := range ranges {
		ids := r.FindIds2()
		for _, id := range ids {
			sum += id
		}
	}
	return sum
}
