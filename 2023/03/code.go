package main

import (
	// "fmt"
	"fmt"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func isNumber(in byte) bool {
	s := rune(in)
	if '0' <= s && s <= '9' {
		return true
	}
	return false
}

func isPeriod(in byte) bool {
	s := rune(in)
	if s == '.' {
		return true
	}
	return false
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	var num string
	var valid_num bool
	var nearby_gears [][]int
	var engine []string
	var result int

	boundary := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) != 0 {
			engine = append(engine, line)
		}
	}

	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		gear_map := make(map[string][]int)

		for i, line := range engine {
			for j := 0; j < len(line); j++ {
				if isNumber(line[j]) {
					num += string(line[j])
					for _, pos := range boundary {
						if 0 <= i+pos[0] && i+pos[0] < len(line) && 0 <= j+pos[1] && j+pos[1] < len(engine) {
							if engine[i+pos[0]][j+pos[1]] == '*' && !valid_num {
								valid_num = true
								gear_pos := []int{i + pos[0], j + pos[1]}
								nearby_gears = append(nearby_gears, gear_pos)
							}
						}
					}
				} else {
					if len(num) != 0 && valid_num {
						fmt.Println("Valid", num)
						part, err := strconv.Atoi(num)
						if err != nil {
							return err
						}

						for _, gear_pos := range nearby_gears {
							gear_pos_str := strconv.Itoa(gear_pos[0]) + "," + strconv.Itoa(gear_pos[1])
							fmt.Println(gear_pos, "to", gear_pos_str)
							if _, ok := gear_map[gear_pos_str]; !ok {
								gear_map[gear_pos_str] = []int{part}
							} else {
								gear_map[gear_pos_str] = append(gear_map[gear_pos_str], part)
							}
						}
					}
					num = ""
					nearby_gears = nil
					valid_num = false
				}

				if j == len(line)-1 {
					if len(num) != 0 && valid_num {
						part, err := strconv.Atoi(num)
						if err != nil {
							return err
						}
						// result += part

						for _, gear_pos := range nearby_gears {
							gear_pos_str := strconv.Itoa(gear_pos[0]) + "," + strconv.Itoa(gear_pos[1])
							if _, ok := gear_map[gear_pos_str]; !ok {
								gear_map[gear_pos_str] = []int{part}
							} else {
								gear_map[gear_pos_str] = append(gear_map[gear_pos_str], part)
							}
						}
					}
					num = ""
					nearby_gears = nil
					valid_num = false
				}
			}
		}

		fmt.Print(gear_map)

		for _, values := range gear_map {
			if len(values) == 2 {
				result += values[0] * values[1]
			}
		}
		return result
	}
	// solve part 1 here

	for i, line := range engine {
		for j := 0; j < len(line); j++ {
			if isNumber(line[j]) {
				num += string(line[j])
				for _, pos := range boundary {
					if 0 <= i+pos[0] && i+pos[0] < len(line) && 0 <= j+pos[1] && j+pos[1] < len(engine) {
						if !isNumber(engine[i+pos[0]][j+pos[1]]) && !isPeriod(engine[i+pos[0]][j+pos[1]]) {
							valid_num = true
						}
					}
				}
			} else {
				if len(num) != 0 && valid_num {
					fmt.Println("Valid", num)
					part, err := strconv.Atoi(num)
					if err != nil {
						return err
					}

					result += part
				}
				num = ""
				valid_num = false
			}

			if j == len(line)-1 {
				if len(num) != 0 && valid_num {
					part, err := strconv.Atoi(num)
					if err != nil {
						return err
					}
					result += part
				}
				num = ""
				valid_num = false
			}
		}
	}
	return result
}
