package day10

import (
	"aoc2025/utils"
	"fmt"
	"maps"
	"math"
	"slices"
	"strings"
)

func Solve() {
	lines, err := utils.ReadLines("day10/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 10:")
	// fmt.Println("Part 1:", Part1(lines))
	fmt.Println("Part 2:", Part2(lines))
}

type Machine struct {
	targetConfig        string
	buttons             [][]int
	joltageRequirements []int
}

func parseLine(line string) Machine {
	parts := strings.Split(line, " ")
	configString := parts[0]
	buttonStrings := parts[1 : len(parts)-1]
	joltageString := parts[len(parts)-1]

	buttons := make([][]int, len(buttonStrings))
	for i, buttonString := range buttonStrings {
		numberStrs := strings.Split(buttonString[1:len(buttonString)-1], ",")
		numbers := make([]int, len(numberStrs))
		for j, num := range numberStrs {
			numbers[j] = utils.ParseIntUnsafe(num)
		}
		buttons[i] = numbers
	}

	joltageStrs := strings.Split(joltageString[1:len(joltageString)-1], ",")
	joltages := make([]int, len(joltageStrs))
	for i, joltageStr := range joltageStrs {
		joltages[i] = utils.ParseIntUnsafe(joltageStr)
	}

	return Machine{
		targetConfig:        configString[1 : len(configString)-1],
		buttons:             buttons,
		joltageRequirements: joltages,
	}
}

func parse(lines []string) []Machine {
	machines := make([]Machine, len(lines))
	for i, line := range lines {
		machines[i] = parseLine(line)
	}

	return machines
}

func applyButton(config string, button []int) string {
	newString := make([]rune, len(config))
	for i, char := range config {
		if slices.Contains(button, i) {
			if char == '#' {
				newString[i] = '.'
			} else {
				newString[i] = '#'
			}
		} else {
			newString[i] = char
		}
	}
	return string(newString)
}

func emptyConfig(config string) string {
	newString := make([]rune, len(config))
	for i := range config {
		newString[i] = '.'
	}
	return string(newString)
}

func recursiveF(machine Machine, current string, visitedSet map[string](bool), presses int, maxPresses int) (int, bool, []int) {
	if current == machine.targetConfig {
		return presses, true, []int{}
	}
	if visitedSet[current] {
		return -1, false, []int{}
	}
	if maxPresses == presses {
		return -1, false, []int{}
	}

	myVisitedSet := maps.Clone(visitedSet)
	myVisitedSet[current] = true

	for j, button := range machine.buttons {
		newString := applyButton(current, button)
		buttonPresses, succ, sequence := recursiveF(machine, newString, myVisitedSet, presses+1, maxPresses)
		if succ {
			return buttonPresses, true, append(sequence, j)
		}
	}
	return -1, false, []int{}
}

func assertSequence(machine Machine, sequence []int) {
	currentConfig := emptyConfig(machine.targetConfig)
	for _, buttonIndex := range sequence {
		currentConfig = applyButton(currentConfig, machine.buttons[buttonIndex])
	}
	if currentConfig != machine.targetConfig {
		fmt.Println("failed sequence!", currentConfig, machine.targetConfig)
		panic("failed sequence")
	}
}

func minimumPressesForMachine(machine Machine) int {
	visitedSet := make(map[string](bool))
	startConfig := emptyConfig(machine.targetConfig)

	i := 0
	for {
		presses, success, sequence := recursiveF(machine, startConfig, visitedSet, 0, i)
		if success {
			slices.Reverse(sequence)
			assertSequence(machine, sequence)
			return presses
		}
		i++
	}
}

func Part1(lines []string) int {
	machines := parse(lines)
	count := 0
	for _, machine := range machines {
		presses := minimumPressesForMachine(machine)
		count += presses
	}
	return count
}

func applyJoltageButton(joltageState []int, button []int, count int) []int {
	newJoltageState := slices.Clone(joltageState)
	for _, buttonIndex := range button {
		newJoltageState[buttonIndex] += count
	}
	return newJoltageState
}

func emptyJoltage(joltageState []int) []int {
	newJoltageState := make([]int, len(joltageState))
	for i, _ := range joltageState {
		newJoltageState[i] = 0
	}
	return newJoltageState
}

//
// b * b_0 + b * b_1 + ... + b * b_n = l1
// ...

func (machine *Machine) toGaussianElimation() [][]float64 {
	matrix := make([][]float64, len(machine.joltageRequirements))
	for i := range machine.joltageRequirements {
		matrix[i] = make([]float64, len(machine.buttons)+1)
		matrix[i][len(machine.buttons)] = float64(machine.joltageRequirements[i])
		for j, buttonIndexes := range machine.buttons {
			for _, buttonIndex := range buttonIndexes {
				if buttonIndex == i {
					matrix[i][j] = 1.0
				}
			}
		}
		// fmt.Println(matrix[i])
	}
	return matrix
}

// see https://en.wikipedia.org/wiki/Gaussian_elimination
func applyGuassianElimination(matrix [][]float64) [][]float64 {
	h := 0
	k := 0
	for h < len(matrix) && k < len(matrix[0]) {
		i_max := h
		for i := h; i < len(matrix); i++ {
			current_max := math.Abs(matrix[i_max][k])
			next_max := math.Abs(matrix[i][k])
			if current_max < next_max {
				i_max = i
			}
		}

		if matrix[i_max][k] == 0 {
			k++
			continue
		}
		// swap rows
		matrix[h], matrix[i_max] = matrix[i_max], matrix[h]
		for i := h + 1; i < len(matrix); i++ {
			f := matrix[i][k] / matrix[h][k]
			matrix[i][k] = 0
			for j := k + 1; j < len(matrix[0]); j++ {
				matrix[i][j] -= f * matrix[h][j]
			}
		}

		h++
		k++
	}

	// fmt.Println("done")
	// for i := range matrix {
	// 	fmt.Println(matrix[i])
	// }
	// fmt.Println("free variables", len(matrix[0])-len(matrix)-1)
	matrix = removeEmptyRows(matrix)
	// for i := range matrix {
	// 	fmt.Println(matrix[i])
	// }
	return matrix
}

var smallFloat = 1e-8

func removeEmptyRows(matrix [][]float64) [][]float64 {
	for rowIndex, row := range matrix {
		empty := true
		for _, value := range row {
			if math.Abs(value) > smallFloat {
				empty = false
				break
			}
		}
		if empty {
			return matrix[:rowIndex]
		}
	}
	return matrix
}

func applyLastValue(matrix [][]float64, value float64) [][]float64 {
	newMatrix := make([][]float64, len(matrix))
	for i := range matrix {
		newMatrix[i] = make([]float64, len(matrix[i])-1)
		// copy old values over
		for j := 0; j < len(newMatrix[i])-1; j++ {
			newMatrix[i][j] = matrix[i][j]
		}
		newMatrix[i][len(newMatrix[i])-1] = matrix[i][len(matrix[i])-1] - value*matrix[i][len(matrix[i])-2]
	}
	return newMatrix
}

func hasSingleSolution(lastRow []float64) bool {
	// check if last row solveable like this
	firstNotNullIndex := len(lastRow)
	for i, value := range lastRow {
		if math.Abs(value) > smallFloat {
			firstNotNullIndex = i
			break
		}
	}
	return firstNotNullIndex >= len(lastRow)-2
}

func (machine *Machine) solveSingleSolution(matrix [][]float64) ([]int, bool) {
	values := make([]int, 0)
	for len(matrix) > 0 {
		lastRow := matrix[len(matrix)-1]
		if !hasSingleSolution(lastRow) {
			bruteForceValues, success := machine.bruteForcePresses(matrix)
			if !success {
				return values, false
			}
			values = append(values, bruteForceValues...)
			return values, true
		}

		lastValue := lastRow[len(lastRow)-1] / lastRow[len(lastRow)-2]
		lastValueInt := int(math.Round(lastValue))
		if math.Abs(float64(lastValueInt)-lastValue) > smallFloat || lastValueInt < 0 {
			// fmt.Println("called solve single solution, without valid solution", lastRow, lastValue)
			return values, false
		}
		matrix = applyLastValue(matrix, lastValue)
		// eliminate last row
		matrix = matrix[:len(matrix)-1]
		values = append(values, lastValueInt)
	}
	return values, true
}

func (machine *Machine) maxViablePressesOfButton(buttonIndex int) int {
	maxPresses := 0
	for _, buttonTarget := range machine.buttons[buttonIndex] {
		if maxPresses < machine.joltageRequirements[buttonTarget] {
			maxPresses = machine.joltageRequirements[buttonTarget]
		}
	}
	return maxPresses + 1
}

func pressSum(values []int) int {
	totalPresses := 0
	for _, value := range values {
		totalPresses += value
	}
	return totalPresses
}

func (machine *Machine) bruteForcePresses(matrix [][]float64) ([]int, bool) {
	remainingButtonLength := len(matrix[0]) - 1
	// implicitly get already considered buttons via matrix
	if hasSingleSolution(matrix[len(matrix)-1]) {
		// fmt.Println("brute force approach", len(matrix[0])-1)
		return machine.solveSingleSolution(matrix)
	}

	minPressSum := 2_000_000_000
	minPresses := []int{}
	globalSuccess := false
	for buttonPresses := range machine.maxViablePressesOfButton(remainingButtonLength - 1) {
		localMatrix := applyLastValue(matrix, float64(buttonPresses))
		presses, success := machine.bruteForcePresses(localMatrix)
		if !success {
			continue
		}
		sum := pressSum(presses) + buttonPresses
		if sum < minPressSum {
			minPressSum = sum
			slices.Reverse(presses)
			minPresses = append(presses, buttonPresses)
			slices.Reverse(minPresses)
			// fmt.Println("found brute force approach", buttonPresses, minPresses, success)
			globalSuccess = true
		}
	}

	return minPresses, globalSuccess && minPressSum != 2_000_000_000
}

func (machine *Machine) verifyButtonPresses(buttonPresses []int) bool {
	buttonPressesOrdered := slices.Clone(buttonPresses)
	slices.Reverse(buttonPressesOrdered)
	current := emptyJoltage(machine.joltageRequirements)
	for i, buttonPress := range buttonPressesOrdered {
		current = applyJoltageButton(current, machine.buttons[i], buttonPress)
	}
	success := slices.Equal(current, machine.joltageRequirements)
	if !success {
		fmt.Println("buttonPressed ordered", buttonPressesOrdered, "result", current, machine.joltageRequirements)
	}
	return success
}

func Part2(lines []string) int {
	machines := parse(lines)
	count := 0
	for i, machine := range machines {
		if i == 50000 {
			continue
		}
		matrix := machine.toGaussianElimation()
		matrix = applyGuassianElimination(matrix)
		values, success := machine.bruteForcePresses(matrix)
		if !success {
			panic("brute force failed")
		}
		if !machine.verifyButtonPresses(values) {
			fmt.Println(values, machine, i)
			panic("button press verification failed")
		}
		count += pressSum(values)
		slices.Reverse(values)
		fmt.Println(i, pressSum(values), values, success)
	}
	// Placeholder for Part 2
	return count
}

// failed: 13914 too low
// failed: 19080 too high
