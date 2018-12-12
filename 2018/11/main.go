package main

import (
	"fmt"
)

func main() {
	cells := make([][]int, 0)
	for i := 0; i < 300; i++ {
		cells = append(cells, make([]int, 300))
	}
	serial := 8979
	for x := 0; x < 300; x++ {
		for y := 0; y < 300; y += 1 {
			cells[x][y] = ((((((x + 11) * (y + 1)) + serial) * (x + 11)) / 100) % 10) - 5
		}
	}
	maxX := -1
	maxY := -1
	maxSize := -1
	maxP := 0
	for bs := 1; bs <= 300; bs++ {
		for x := 0; x < 300-(bs-1); x++ {
			for y := 0; y < 300-(bs-1); y += 1 {
				sum := 0
				for sx := x; sx < x+bs; sx++ {
					for sy := y; sy < y+bs; sy++ {
						sum += cells[sx][sy]
					}
				}
				if sum > maxP {
					maxX = x + 1
					maxY = y + 1
					maxSize = bs
					maxP = sum
					fmt.Printf("%v, %v, %v\n", maxX, maxY, maxSize)
				}
			}
		}
	}
}

var input = `8979`
