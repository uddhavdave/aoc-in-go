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

type hand struct {
	hand string
	bid  int
}

var cardMap = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

func (self hand) isStronger(other hand) bool {
	for i := 0; i < len(self.hand); i++ {
		if cardMap[rune(self.hand[i])] < cardMap[rune(other.hand[i])] {
			return false
		} else if cardMap[rune(self.hand[i])] > cardMap[rune(other.hand[i])] {
			return true
		}
	}
	return true
}

func mergeSort(hands []hand) []hand {
	if len(hands) < 2 {
		return hands
	}

	first := mergeSort(hands[:len(hands)/2])
	second := mergeSort(hands[len(hands)/2:])
	return merge(first, second)
}

func merge(left []hand, right []hand) []hand {
	var res []hand
	var i, j int

	for i < len(left) && j < len(right) {
		// fmt.Println(left[i], right[j], left[i].isStronger(right[j]))
		if left[i].isStronger(right[j]) {
			res = append(res, right[j])
			j++
		} else {
			res = append(res, left[i])
			i++
		}
	}

	for i < len(left) {
		res = append(res, left[i])
		i++
	}

	for j < len(right) {
		res = append(res, right[j])
		j++
	}

	return res
}

func (hand hand) evaluateStrength(part2 bool) int {
	var maxDuplicates int
	countMap := make(map[rune]int)
	for _, card := range hand.hand {
		if _, ok := countMap[card]; !ok {
			countMap[card] = 0
		}
		countMap[card] += 1
	}

	for card, count := range countMap {
		if maxDuplicates < count {
			if part2 && card == 'J' {
				continue
			}
			maxDuplicates = count
		}
	}

	if part2 && countMap['J'] != 0 {
		maxDuplicates += countMap['J']
		delete(countMap, 'J')
	}

	var strength int

	switch maxDuplicates {
	case 1:
		// High Card
		strength = 0
	case 2:
		if len(countMap) == 3 {
			// Two pair
			strength = 2
			break
		}
		// One pair
		strength = 1
	case 3:
		if len(countMap) == 2 {
			// full house
			strength = 4
			break
		}
		strength = 3
	case 4:
		strength = 5
	case 5:
		strength = 6
	default:
		return -1
	}

	return strength

}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	inputs := strings.Split(input, "\n")

	inputs = inputs[:len(inputs)-1]

	if part2 {
		cardMap['J'] = 1
	}

	var hands []hand
	for _, str := range inputs {
		handInfo := strings.Fields(str)
		bid, err := strconv.Atoi(handInfo[1])
		if err != nil {
			return err
		}
		hands = append(hands, hand{handInfo[0], bid})
	}

	hands = mergeSort(hands)

	var sorted [7][]hand

	for _, hand := range hands {
		index := hand.evaluateStrength(part2)
		sorted[index] = append(sorted[index], hand)
	}

	for _, power := range sorted {
		fmt.Println(power)
	}

	var rank, result int

	for i := 0; i < 7; i++ {
		if len(sorted[i]) == 0 {
			continue
		}
		for _, hand := range sorted[i] {
			rank++
			result += hand.bid * rank
		}
	}

	// solve part 1 here
	return result
}
