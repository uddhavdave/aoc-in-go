package main

import (
	"fmt"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type Graph struct {
	adjList map[string][]string
}

func NewGraph() *Graph {
	return &Graph{
		adjList: make(map[string][]string),
	}
}

func (graph Graph) AddEdge(u, v string) {
	graph.adjList[u] = append(graph.adjList[u], v)
}

func AtFinalPosition(pos string) bool {
	finalPositionSuffix := 'Z'
	if rune(pos[len(pos)-1]) != finalPositionSuffix {
		return false
	}
	return true
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// solve part 1 here
	inputs := strings.Split(input, "\n")

	pattern := inputs[0]

	inputs = inputs[2 : len(inputs)-1]

	graph := NewGraph()

	for _, node := range inputs {
		fields := strings.Fields(node)

		from := fields[0]

		for _, to := range fields[2:] {
			to = strings.Trim(to, "(),")
			graph.AddEdge(from, to)
		}
	}

	var totalSteps int
	if part2 {
		var startingPositions []string
		startingPositionSuffix := 'A'

		for key := range graph.adjList {
			if key[len(key)-1] == byte(startingPositionSuffix) {
				startingPositions = append(startingPositions, key)
			}
		}

		currPositions := startingPositions

		fmt.Println(currPositions)

		var stepsPerPosition []int

		for _, pos := range currPositions {
			var steps int
			for !AtFinalPosition(pos) {
				move := pattern[steps%len(pattern)]
				if move == 'R' {
					pos = graph.adjList[pos][1]
				} else {
					pos = graph.adjList[pos][0]
				}
				steps++
			}
			stepsPerPosition = append(stepsPerPosition, steps)
		}

		if len(stepsPerPosition) > 1 {
			fmt.Println(LCM(stepsPerPosition[0], stepsPerPosition[1], stepsPerPosition...))
		} else {

			fmt.Println(stepsPerPosition[0])
		}
		return totalSteps
	}

	startingPosition := "AAA"
	finalPosition := "ZZZ"
	currPosition := startingPosition

	for currPosition != finalPosition {
		move := pattern[totalSteps%len(pattern)]
		if move == 'R' {
			currPosition = graph.adjList[currPosition][1]
		} else {
			currPosition = graph.adjList[currPosition][0]
		}
		totalSteps++
	}

	return totalSteps
}
