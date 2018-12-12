package main

import (
	"fmt"
	"strings"
)

func main() {
	delim := "\n"
	parts := strings.Split(input, delim)
	initialState := strings.Split(parts[0], ": ")[1]
	padding := 100
	state := make([]int, len(parts)+padding*2)
	for i := 0; i < len(initialState); i++ {
		if initialState[i] == '#' {
			state[i+padding] = 1
		}
	}

	parts = parts[2:]
	rules := make([][]int, 0)
	for _, p := range parts {
		inOut := strings.Split(p, " => ")
		if inOut[1] == "#" {
			rule := make([]int, 0)
			for i := 0; i < len(inOut[0]); i++ {
				if inOut[0][i] == '#' {
					rule = append(rule, 1)
				} else {
					rule = append(rule, 0)
				}
			}
			rules = append(rules, rule)
		}
	}

	for t := 0; t < 20; t++ {
		newState := []int{0, 0}
		for p := 2; p < len(state)-2; p++ {
			input := state[p-2 : p+3]
			result := 0
			for _, rule := range rules {
				valid := true
				for i := 0; i < 5; i++ {
					if rule[i] != input[i] {
						valid = false
					}
				}
				if valid {
					result = 1
					break
				}
			}
			newState = append(newState, result)
		}
		newState = append(newState, 0, 0)
		state = newState
		sum := 0
		for p := 2; p < len(state)-2; p++ {
			sum += (p - padding) * state[p]
		}
		fmt.Printf("%v = %v\n", t, sum)
	}
	sum := 0
	for p := 2; p < len(state)-2; p++ {
		sum += (p - padding) * state[p]
	}
	fmt.Printf("sum: %v", sum)
}

var input = `initial state: #...#..###.#.###.####.####.#..#.##..#..##..#.....#.#.#.##.#...###.#..##..#.##..###..#..##.#..##...

...#. => #
#..## => #
..... => .
##.## => .
.##.. => #
.##.# => .
####. => #
.#.#. => .
..#.# => .
.#.## => .
.#..# => .
##... => #
#...# => #
##### => .
#.### => #
..### => #
###.. => .
#.#.# => #
##..# => #
..#.. => #
.#### => .
#.##. => .
....# => .
...## => .
#.... => .
#..#. => .
..##. => .
###.# => #
#.#.. => #
##.#. => #
.###. => .
.#... => .`
