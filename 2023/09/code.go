package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func predictPrevious(list []int) int {
	zeroDiff := true
	for _, element := range list {
		if element != 0 {
			zeroDiff = false
		}
	}

	if zeroDiff {
		return list[0]
	}

	var newList []int
	for i, element := range list[:len(list)-1] {
		newList = append(newList, list[i+1]-element)
	}

	return list[0] - predictPrevious(newList)
}

func predictNext(list []int) int {
	zeroDiff := true
	for _, element := range list {
		if element != 0 {
			zeroDiff = false
		}
	}

	if zeroDiff {
		return list[0]
	}

	var newList []int
	for i, element := range list[:len(list)-1] {
		newList = append(newList, list[i+1]-element)
	}

	return list[len(list)-1] + predictNext(newList)
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
	inputs = inputs[:len(inputs)-1]

	var result int

	for _, data := range inputs {
		var history []int
		for _, readings := range strings.Fields(data) {
			val, err := strconv.Atoi(readings)
			if err != nil {
				return err
			}
			history = append(history, val)
		}
		if part2 {
			fmt.Println(history, predictPrevious(history))
			result += predictPrevious(history)
		} else {
			fmt.Println(history, predictNext(history))
			result += predictNext(history)
		}
	}
	return result
}
