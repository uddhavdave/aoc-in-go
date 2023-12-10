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

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	inputs := strings.Split(input, "\n")

	splitFn := func(c rune) bool {
		return c == ' '
	}
	stimes := strings.FieldsFunc(inputs[0], splitFn)[1:]
	srecords := strings.FieldsFunc(inputs[1], splitFn)[1:]

	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {

		var stime, srecord string

		for _, val := range stimes {
			stime += val
		}

		for _, val := range srecords {
			srecord += val
		}

		time, err := strconv.Atoi(stime)
		if err != nil {
			return err
		}

		record, err := strconv.Atoi(srecord)
		if err != nil {
			return err
		}

		var result int

		fmt.Println(record, time)
		for j := 0; j <= time/2; j++ {
			if j*(time-j) > record {
				// fmt.Println(j)
				result = ((time/2)-j)*2 + 1 // one for the middle case
				break
			}
		}

		return result
	}

	var records, times []int

	for _, stime := range stimes {
		val, err := strconv.Atoi(stime)
		if err != nil {
			return err
		}
		times = append(times, val)
	}

	for _, srecord := range srecords {
		val, err := strconv.Atoi(srecord)
		if err != nil {
			return err
		}
		records = append(records, val)
	}

	var TotalPossibleWays []int

	for i, time := range times {
		var PossibleWays int
		for j := 0; j < time; j++ {
			if j*(time-j) > records[i] {
				PossibleWays += 1
			}
		}
		TotalPossibleWays = append(TotalPossibleWays, PossibleWays)
	}

	fmt.Println(TotalPossibleWays)

	result := 1

	for _, ways := range TotalPossibleWays {
		result *= ways
	}

	// solve part 1 here
	return result
}
