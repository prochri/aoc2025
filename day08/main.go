package day08

import (
	"aoc2025/utils"
	"container/heap"
	"fmt"
	"math"
	"slices"
	"sort"
	"strings"
)

func Solve() {
	lines, err := utils.ReadLines("day08/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 08:")
	fmt.Println("Part 1:", Part1(lines))
	fmt.Println("Part 2:", Part2(lines))
}

type Vec3 struct {
	x, y, z float64
}

type Distance struct {
	from, to       Vec3
	distance       float64
	fromIdx, toIdx int
}

func (v Vec3) distanceTo(other Vec3) float64 {
	return math.Sqrt(math.Pow(v.x-other.x, 2) + math.Pow(v.y-other.y, 2) + math.Pow(v.z-other.z, 2))
}

func parseData(lines []string) []Vec3 {
	data := make([]Vec3, len(lines))
	for i, line := range lines {
		split := strings.Split(line, ",")
		data[i] = Vec3{
			x: float64(utils.ParseIntUnsafe(split[0])),
			y: float64(utils.ParseIntUnsafe(split[1])),
			z: float64(utils.ParseIntUnsafe(split[2])),
		}
	}
	return data
}

func pointsToDistanceMatrix(points []Vec3) [][]float64 {
	matrix := make([][]float64, len(points))
	for i := range points {
		matrix[i] = make([]float64, len(points))
		for j := range points {
			matrix[i][j] = points[i].distanceTo(points[j])
		}
	}
	return matrix
}

func distanceMatrixToDistances(matrix [][]float64, points []Vec3) []Distance {
	distances := make([]Distance, 0)
	for i := range matrix {
		for j := range matrix {
			if i == j {
				continue
			}
			distances = append(distances, Distance{
				from:     points[i],
				to:       points[j],
				distance: matrix[i][j],
				fromIdx:  i,
				toIdx:    j,
			})
		}
	}
	return distances
}

type PriorityQueue []*Distance

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Distance)
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func sliceUnique(slice []*map[int](bool)) []*map[int](bool) {
	keys := make(map[*map[int](bool)]bool)
	newSlice := make([]*map[int](bool), 0)
	for _, entry := range slice {
		if keys[entry] {
			continue
		}
		keys[entry] = true
		newSlice = append(newSlice, entry)
	}
	return newSlice
}

func Part1(lines []string) int {
	parsed := parseData(lines)
	distanceMatrix := pointsToDistanceMatrix(parsed)
	distances := distanceMatrixToDistances(distanceMatrix, parsed)
	distancePq := make(PriorityQueue, len(distances))
	for i := range distances {
		distancePq[i] = &distances[i]
	}
	heap.Init(&distancePq)

	vecToCircuit := make([]*map[int](bool), len(parsed))
	for i := range vecToCircuit {
		newCircuit := make(map[int](bool))
		newCircuit[i] = true
		vecToCircuit[i] = &newCircuit
	}

	for range 1000 {
		best := heap.Pop(&distancePq).(*Distance)
		bestSwapped := heap.Pop(&distancePq).(*Distance)
		if bestSwapped.fromIdx != best.toIdx || bestSwapped.toIdx != best.fromIdx {
			panic("Distances are not symmetric")
		}
		circuitA := vecToCircuit[best.fromIdx]
		circuitB := vecToCircuit[best.toIdx]

		if circuitA == circuitB {
			continue
		}
		// now merge the smaller circuit into the larger one
		smallerCircuit := circuitA
		largerCircuit := circuitB
		if len(*circuitA) > len(*circuitB) {
			smallerCircuit = circuitB
			largerCircuit = circuitA
		}
		for k := range *smallerCircuit {
			if (*largerCircuit)[k] {
				fmt.Println("Unexpected: Circuit is not disjoint", k, smallerCircuit, largerCircuit)
				panic("Unexpected: Circuit is not disjoint")
			}
			(*largerCircuit)[k] = true
			vecToCircuit[k] = largerCircuit
		}
	}

	uniqueCircuits := sliceUnique(vecToCircuit)
	circuitLengths := make([]int, len(uniqueCircuits))
	for i := range uniqueCircuits {
		circuitLengths[i] = len(*uniqueCircuits[i])
	}
	sort.Ints(circuitLengths)
	slices.Reverse(circuitLengths)

	return circuitLengths[0] * circuitLengths[1] * circuitLengths[2]
}

func Part2(lines []string) int {
	parsed := parseData(lines)
	distanceMatrix := pointsToDistanceMatrix(parsed)
	distances := distanceMatrixToDistances(distanceMatrix, parsed)
	distancePq := make(PriorityQueue, len(distances))
	for i := range distances {
		distancePq[i] = &distances[i]
	}
	heap.Init(&distancePq)

	vecToCircuit := make([]*map[int](bool), len(parsed))
	for i := range vecToCircuit {
		newCircuit := make(map[int](bool))
		newCircuit[i] = true
		vecToCircuit[i] = &newCircuit
	}

	uniqueCircuits := len(vecToCircuit)
	for uniqueCircuits > 1 {
		best := heap.Pop(&distancePq).(*Distance)
		bestSwapped := heap.Pop(&distancePq).(*Distance)
		if bestSwapped.fromIdx != best.toIdx || bestSwapped.toIdx != best.fromIdx {
			panic("Distances are not symmetric")
		}
		circuitA := vecToCircuit[best.fromIdx]
		circuitB := vecToCircuit[best.toIdx]

		if circuitA == circuitB {
			continue
		}
		// now merge the smaller circuit into the larger one
		smallerCircuit := circuitA
		largerCircuit := circuitB
		if len(*circuitA) > len(*circuitB) {
			smallerCircuit = circuitB
			largerCircuit = circuitA
		}
		for k := range *smallerCircuit {
			if (*largerCircuit)[k] {
				fmt.Println("Unexpected: Circuit is not disjoint", k, smallerCircuit, largerCircuit)
				panic("Unexpected: Circuit is not disjoint")
			}
			(*largerCircuit)[k] = true
			vecToCircuit[k] = largerCircuit
		}
		uniqueCircuits--
		if uniqueCircuits == 1 {
			return int(best.from.x * best.to.x)
		}
	}

	return 0
}
