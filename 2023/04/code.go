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
	// when you're ready to do part 2, remove this "not implemented" block

	// solve part 1 here
	cards := strings.Split(input, "\n")
	cards = cards[:len(cards)-1]
	var result int

	card_count := make([]int, len(cards))

	for i, card := range cards {

		var match int

		if part2 {
			fmt.Println(card)
			card_count[i] += 1
		}
		set := make(map[int]struct{})

		if card == "" {
			continue
		}

		info := strings.Split(card[strings.Index(card, ":")+1:], "|")

		wc_str := strings.Split(info[0], " ")
		cc_str := strings.Split(info[1], " ")

		for _, str := range wc_str {
			if str != "" {
				val, err := strconv.Atoi(str)
				if err != nil {
					return err
				}
				if _, ok := set[val]; !ok {
					set[val] = struct{}{}
				}
			}
		}
		for _, str := range cc_str {
			if str != "" {
				val, err := strconv.Atoi(str)
				if err != nil {
					return err
				}

				if _, ok := set[val]; ok {
					// fmt.Println(val)
					match += 1
				}
			}
		}

		if part2 {
			if match > len(card_count)-i {
				match = len(card_count) - i
			}

			fmt.Println("Card", i, "Match", match)

			count := card_count[i]

			for j := 0; j < count; j++ {
				for k := 1; k <= match; k++ {
					card_count[i+k] += 1
				}
			}

		} else {
			if match != 0 {
				// fmt.Println(1 << (match - 1))
				result += 1 << (match - 1)
			}
		}
	}

	if part2 {
		for _, val := range card_count {
			result += val
		}
	}

	return result
}
