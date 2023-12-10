package main

import (
	// "fmt"
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

	cube_map := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	// solve part 1 here

	var result int
	games := strings.Split(input, "\n")
	for _, game := range games {
		if len(game) == 0 {
			break
		}

		inputs := strings.Split(game, ":")

		var id int
		var err error
		valid := true
		id_str := strings.Split(inputs[0], " ")[1]
		id, err = strconv.Atoi(id_str)
		if err != nil {
			return err
		}

		sets := strings.Split(inputs[1], ";")

		// when you're ready to do part 2, remove this "not implemented" block
		if part2 {
			var min_red, min_green, min_blue int
			for _, set := range sets {
				// fmt.Println(set)
				cubes_info := strings.Split(set, ",")
				for _, info := range cubes_info {
					cube_info := strings.Split(info, " ")

					count, err := strconv.Atoi(cube_info[1])
					if err != nil {
						return err
					}

					// fmt.Println(cube_info[2])
					switch cube_info[2] {
					case "red":
						if count > min_red {
							min_red = count
						}
					case "blue":
						if count > min_blue {
							min_blue = count
						}
					case "green":
						if count > min_green {
							min_green = count
						}
					default:
						continue
					}
				}
			}
			result += min_blue * min_red * min_green
		} else {
			for _, set := range sets {
				// fmt.Println(set)
				cubes_info := strings.Split(set, ",")
				for _, info := range cubes_info {
					cube_info := strings.Split(info, " ")

					// fmt.Println(cube_info[2])
					if limit, ok := cube_map[cube_info[2]]; ok {
						count, err := strconv.Atoi(cube_info[1])
						if err != nil {
							return err
						}

						// fmt.Println(count, limit)
						if count > limit {
							valid = false
							break
						}
					}
				}
				if !valid {
					break
				}
			}

			if valid {
				result += id
			}
		}
	}
	return result
}
