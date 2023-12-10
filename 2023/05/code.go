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

func walkMap(frame []int, map_info [][]int) [][]int {

	if len(map_info) == 0 {
		return [][]int{frame}
	}

	for _, rows := range map_info {
		frames = }

	new_frames := breakdownFrame(frame)
}

func breakdownFrame(frame []int, source []int, dest []int) (result [][]int) {
	if frame[0] > source[0] {
		if frame[0] <= source[1] {
			if source[1] < frame[1] {
				// Left Overlap
				// Return 2 frames
				leftFrame := []int{
					dest[0] + (frame[0] - source[0]),
					dest[1],
				}
				rightFrame := []int{
					dest[1],
					frame[1],
				}
				return [][]int{leftFrame, rightFrame}
			} else {
				// Full Overlap
				// Return 1 frames
				fullFrame := []int{
					dest[0] + (frame[0] - source[0]),
					dest[1] - (source[1] - frame[1]),
				}
				return [][]int{fullFrame}
			}
		}
	} else if frame[1] >= source[0] {
		if frame[1] <= source[1] {
			// Right overlap
			// Return 2 frames
			leftFrame := []int{
				frame[0],
				dest[0],
			}
			rightFrame := []int{
				dest[0],
				dest[1] - (source[1] - frame[1]),
			}
			return [][]int{leftFrame, rightFrame}
		} else {
			// Contains
			// Return 3 frames
			leftFrame := []int{
				frame[0],
				dest[0],
			}
			middleFrame := []int{
				dest[0],
				dest[1],
			}
			rightFrame := []int{
				dest[1],
				frame[1],
			}
			return [][]int{leftFrame, middleFrame, rightFrame}
		}
	}
	return [][]int{{}}
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	var seeds []int
	var maps [][][]int

	inputs := strings.Split(input, "\n\n")

	seeds_str := strings.Split(strings.TrimPrefix(inputs[0], "seeds:"), " ")
	for _, seed_str := range seeds_str {
		if seed_str != "" {
			seed, err := strconv.Atoi(seed_str)
			if err != nil {
				return err
			}
			seeds = append(seeds, seed)
		}
	}

	// Remove seed information
	inputs = inputs[1:]

	var rows []string

	for _, str := range inputs {
		rows = append(rows, strings.Split(str, "\n")...)
	}

	// fmt.Println(rows)

	var map_info [][]int
	for _, row_str := range rows {
		var row_bounds []int
		if strings.Contains(row_str, "map") || len(row_str) == 0 {
			// Add the map to maps
			if len(map_info) != 0 {
				maps = append(maps, map_info)
				map_info = nil
			}
			continue
		}

		row_info_str := strings.SplitN(row_str, " ", 3)
		// fmt.Println("ROW", row_info_str)

		for _, row_info := range row_info_str {
			bounds, err := strconv.Atoi(row_info)
			if err != nil {
				return err
			}

			row_bounds = append(row_bounds, bounds)
		}

		// Swap Destination bound to Source bound
		row_bounds[0], row_bounds[1] = row_bounds[1], row_bounds[0]
		map_info = append(map_info, row_bounds)

		row_bounds = nil
	}

	// fmt.Println(seeds)
	// fmt.Println(maps)

	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		var frames [][]int
		var locations [][]int
		var frame []int
		for i, seed := range seeds {
			if i%2 == 0 {
				frame = append(frame, seed)
			} else {
				frame = append(frame, frame[0]+seed)
				frames = append(frames, frame)
				fmt.Println(frame)
				frame = nil
			}
		}

		// fmt.Println(frames)

		for _, frame := range frames {
			for _, map_info := range maps {
				for _, row := range map_info {
					frame = breakdownFrame(frame)
					// for i := 0; i < row[2]; i++ {
					// 	if row[i]+i == 0 {
					// 		heap[row[0]+i] = -1
					// 	} else {
					// 		heap[row[0]+i] = row[1] + i
					// 	}
					// }
				}
			}

		}

		// fmt.Println(heap)
		// return "not implemented"
	}

	var locations []int
	for _, seed := range seeds {
		var dest int
		dest = seed
		for _, map_info := range maps {
			for _, row := range map_info {
				if row[0] > dest {
					continue
				}

				if row[2]+row[0] > dest {
					dest = row[1] + (dest - row[0])
					break
				}
			}
		}
		locations = append(locations, dest)
	}

	// fmt.Println(locations)
	min := locations[0]

	for _, val := range locations {
		if val < min {
			min = val
		}
	}
	// solve part 1 here
	return min
}
