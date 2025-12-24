package day09

import (
	"aoc2025/utils"
	"fmt"
	"math"
	"strings"
)

func Solve() {
	lines, err := utils.ReadLines("day09/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 09:")
	fmt.Println("Part 1:", Part1(lines))
	fmt.Println("Part 2:", Part2(lines))
}

type Vec2 struct {
	x, y int
}

type Line struct {
	from, to *Vec2
}

func area(a Vec2, b Vec2) int {
	deltaX := math.Abs(float64(a.x-b.x)) + 1
	deltaY := math.Abs(float64(a.y-b.y)) + 1
	return int(deltaX * deltaY)
}

func (a Vec2) distanceTo(b Vec2) float64 {
	return math.Sqrt(math.Pow(float64(a.x-b.x), 2) + math.Pow(float64(a.y-b.y), 2))
}

func (recLine Line) horizontal() bool {
	return recLine.from.y == recLine.to.y
}
func (recLine Line) vertical() bool {
	return recLine.from.x == recLine.to.x
}

// if the rec line just touches the other line, this is fine.
// if it goes through it, this is problematic

func (recLine Line) intersect(other Line) bool {
	if recLine.horizontal() == other.horizontal() {
		return false
	}
	if !recLine.horizontal() {
		return other.intersect(recLine)
	}
	myY := recLine.from.y
	otherX := other.from.y

	if (myY >= other.from.y && myY >= other.to.y) || (myY <= other.from.y && myY <= other.to.y) {
		return false
	}
	if (otherX >= recLine.from.x && otherX >= recLine.to.x) || (otherX <= recLine.from.x && otherX <= recLine.to.x) {
		return false
	}

	return true
}

func parseData(lines []string) []Vec2 {
	data := make([]Vec2, len(lines))
	for i, line := range lines {
		split := strings.Split(line, ",")
		data[i] = Vec2{
			x: utils.ParseIntUnsafe(split[0]),
			y: utils.ParseIntUnsafe(split[1]),
		}
	}
	return data
}

func Part1(lines []string) int {
	parsed := parseData(lines)
	maxval := 0
	for i := range parsed {
		for j := range parsed {
			val := area(parsed[i], parsed[j])
			if val > maxval {
				maxval = val
			}
		}
	}
	return maxval
}

func cornerIsRight(l Line, p Vec2) bool {
	if l.horizontal() {
		// left to right -> p below
		if l.from.x < l.to.x {
			return p.y >= l.from.y
		} else { // right to left -> p above
			return p.y <= l.from.y
		}
	} else {
		// up to down  -> p below
		if l.from.y < l.to.y {
			return p.x >= l.from.x
		} else {
			return p.y <= l.from.y
		}
	}
}

func minMax(a, b int) (int, int) {
	if a > b {
		return b, a
	} else {
		return a, b
	}

}

func onLine(l Line, p Vec2) bool {
	if l.horizontal() {
		minX, maxX := minMax(l.from.x, l.to.x)
		return p.y == l.from.y && minX <= p.x && p.x <= maxX
	} else {
		minY, maxY := minMax(l.from.y, l.to.y)
		return p.x == l.from.x && minY <= p.y && p.y <= maxY
	}
}

func Part2(lines []string) int {
	parsed := parseData(lines)
	vecLines := make([]Line, len(parsed))

	for i := range parsed {
		vecLines[i] = Line{
			from: &parsed[i],
			to:   &parsed[(i+1)%len(parsed)],
		}
	}

	for i := range parsed {
		for j := range parsed {
			if i != j && parsed[i].distanceTo(parsed[j]) <= 1 {
				fmt.Println("direct neighbors found")
			}
		}
		if !vecLines[i].horizontal() && !vecLines[i].vertical() {
			fmt.Println("unexpected input")
		}
	}

	isValid2 := func(i, j int) bool {
		if i == j {
			return true
		}
		minX := parsed[i].x
		maxX := parsed[j].x
		if minX > maxX {
			tmp := minX
			minX = maxX
			maxX = tmp
		}

		minY := parsed[i].y
		maxY := parsed[j].y
		if minY > maxY {
			tmp := minY
			minY = maxY
			maxY = tmp
		}
		// fmt.Println(minX, maxX, minY, maxY)

		// can't contain any other corner
		for _, p := range parsed {
			// fmt.Println("p", p, minX, maxX, minY, maxY)
			if minX < p.x && p.x < maxX && minY < p.y && p.y < maxY {
				// fmt.Println("failed at", p, minX, maxX, minY, maxY)
				return false
			}
		}

		// lines cant cross the container
		for _, l := range vecLines {
			if l.horizontal() {
				lineMaxX := l.from.x
				lineMinX := l.to.x
				if lineMaxX < lineMinX {
					lineMaxX, lineMinX = lineMinX, lineMaxX
				}
				if l.from.y > minY && l.from.y < maxY && ((lineMinX <= minX && lineMaxX > minX) || (lineMaxX >= maxX && lineMinX < maxX)) {
					return false
				}
			} else {
				lineMaxY := l.from.y
				lineMinY := l.to.y
				if lineMaxY < lineMinY {
					lineMaxY, lineMinY = lineMinY, lineMaxY
				}
				if l.from.x > minX && l.from.x < maxX && ((lineMinY <= minY && lineMaxY > minY) || (lineMaxY >= maxY && lineMinY < maxY)) {
					return false
				}
			}
		}

		corners := []Vec2{
			{x: minX, y: minY},
			{x: minX, y: maxY},
			{x: maxX, y: minY},
			{x: maxX, y: maxY},
		}
		recLines := make([]Line, 4)
		for i := range corners {
			recLines[i] = Line{
				from: &corners[i],
				to:   &corners[(i+1)%len(corners)],
			}
		}

		return true
	}

	maxval := 0
	for i := range parsed {
		for j := range parsed {
			if !isValid2(i, j) {
				continue
			}
			val := area(parsed[i], parsed[j])
			if val > maxval {
				fmt.Println(parsed[i], parsed[j], i, j)
				maxval = val
			}
		}
	}
	// example
	// fmt.Println(isValid2(4, 6))
	// myexample
	// fmt.Println(isValid2(9, 1))

	// Placeholder for Part 2
	return maxval
}
