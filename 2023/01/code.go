package main

import (
	// "fmt"
	"strconv"

	"github.com/jpillora/puzzler/harness/aoc"

	"strings"
)

func main() {
	aoc.Harness(run)
}

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			node.children[ch] = NewTrieNode()
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			return false
		}
		node = node.children[ch]
	}
	return node != nil
}

func reverse(word string) (result string) {
	for _, str := range word {
		result = string(str) + result
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
	// when you're ready to do part 2, remove this "not implemented" block

	var lines []string
	var result int

	lines = strings.Split(input, "\n")

	if part2 {

		var lstack, rstack string
		trie := NewTrie()
		reverse_trie := NewTrie()

		word_map := map[string]string{
			"one":   "1",
			"two":   "2",
			"three": "3",
			"four":  "4",
			"five":  "5",
			"six":   "6",
			"seven": "7",
			"eight": "8",
			"nine":  "9",
		}

		for word := range word_map {
			trie.Insert(word)
			reverse_trie.Insert(reverse(word))
		}

		for _, line := range lines {
			n := len(line)

			var val1, val2 string
			var err error

			if n == 0 {
				break
			}

			for i := 0; i < n; i++ {
				lstack += string(line[i])
				if '0' <= line[i] && line[i] <= '9' {
					val1 = string(line[i])
					lstack = ""
					break
				} else if trie.Search(lstack) == false {
					for {
						if len(lstack) == 0 || trie.Search(lstack) {
							break
						}
						lstack = lstack[1:]
					}
				} else {
					//check if lstack exists in word
					if value, ok := word_map[lstack]; ok {
						val1 = value
						lstack = ""
						break
					}
				}
			}

			for i := n - 1; i >= 0; i-- {
				rstack += string(line[i])
				if '0' <= line[i] && line[i] <= '9' {
					val2 = string(line[i])
					rstack = ""
					break
				} else if reverse_trie.Search(rstack) == false {
					for {
						if len(rstack) == 0 || reverse_trie.Search(rstack) {
							break
						}
						rstack = rstack[1:]
					}
				} else {
					//check if lstack exists in word
					if value, ok := word_map[reverse(rstack)]; ok {
						val2 = value
						rstack = ""
						break
					}
				}
			}

			val := val1 + val2
			intVal, err := strconv.Atoi(val)
			if err != nil {
				return err
			}
			result += intVal
		}
		return result
	}

	// solve part 1 here
	for _, line := range lines {
		n := len(line)

		if n == 0 {
			break
		}

		var val1, val2 string
		var err error

		for i := 0; i < n; i++ {
			if '0' <= line[i] && line[i] <= '9' {
				val1 = string(line[i])
				break
			}
		}

		for i := n - 1; i >= 0; i-- {
			if '0' <= line[i] && line[i] <= '9' {
				val2 = string(line[i])
				break
			}
		}

		val := val1 + val2
		intVal, err := strconv.Atoi(val)
		if err != nil {
			return err
		}
		result += intVal
	}

	return result
}
