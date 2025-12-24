package day11

import (
	"aoc2025/utils"
	"fmt"
	"strings"
)

func Solve() {
	lines, err := utils.ReadLines("day11/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Day 11:")
	fmt.Println("Part 1:", Part1(lines))
	fmt.Println("Part 2:", Part2(lines))
}

type Node struct {
	name      string
	neighbors []string
}

type Graph struct {
	nodes map[string]Node
}

func parse(lines []string) Graph {
	graph := Graph{
		nodes: make(map[string]Node),
	}
	for _, line := range lines {
		node := strings.Split(line, ":")
		nodeName := node[0]
		neighbors := strings.Split(strings.Trim(node[1], " "), " ")
		graph.nodes[nodeName] = Node{
			name:      nodeName,
			neighbors: neighbors,
		}
	}
	return graph
}

// assume non-cyclic graph
// NOTE: could use DP
func (g *Graph) numberOfPaths(start string) int {
	if start == "out" {
		return 1
	}
	node := g.nodes[start]
	count := 0
	for _, neighbor := range node.neighbors {
		newPaths := g.numberOfPaths(neighbor)
		count += newPaths
	}
	return count
}

func Part1(lines []string) int {
	graph := parse(lines)
	return graph.numberOfPaths("you")
}

type Stuff struct {
	hasFft      bool
	hasDac      bool
	currentNode string
	depth       int
	path        string
}

func (s *Stuff) String() string {
	return fmt.Sprintf("%d,%d,%s", s.hasFft, s.hasDac, s.currentNode)
}

var dpCache = make(map[string]int)

func (g *Graph) numberOfPathsWithStuff(start string, stuff Stuff) int {
	if start == "out" {
		if stuff.hasFft && stuff.hasDac {
			return 1
		}
		return 0
	}
	if start == "fft" {
		stuff.hasFft = true
	}
	if start == "dac" {
		stuff.hasDac = true
	}
	stuff.depth++
	stuff.path = stuff.path + "->" + start
	stuff.currentNode = start

	cacheKey := stuff.String()
	if val, ok := dpCache[cacheKey]; ok {
		return val
	}

	node := g.nodes[start]
	count := 0
	for _, neighbor := range node.neighbors {
		newPaths := g.numberOfPathsWithStuff(neighbor, stuff)
		count += newPaths
	}
	dpCache[cacheKey] = count
	return count
}

func Part2(lines []string) int {
	graph := parse(lines)
	paths := graph.numberOfPathsWithStuff("svr", Stuff{
		hasFft: false,
		hasDac: false,
	})
	// Placeholder for Part 2
	return paths
}
