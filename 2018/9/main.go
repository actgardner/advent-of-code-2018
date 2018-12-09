package main

import (
	"fmt"
)

type marble struct {
	ccw   *marble
	cw    *marble
	value int
}

func main() {
	players := 411
	last := 7205900
	scores := make([]int, players)
	zero := &marble{nil, nil, 0}
	one := &marble{nil, nil, 1}
	two := &marble{nil, nil, 2}

	zero.ccw = one
	zero.cw = two
	two.ccw = zero
	two.cw = one
	one.ccw = two
	one.cw = zero

	current := two
	player := 3
	for value := 3; value <= last; value++ {
		if value%23 != 0 {
			nextCcw := current.cw
			nextCw := current.cw.cw
			new := &marble{nextCcw, nextCw, value}
			nextCcw.cw = new
			nextCw.ccw = new
			current = new
		} else {
			scores[player] += value
			remove := current.ccw.ccw.ccw.ccw.ccw.ccw.ccw
			remove.ccw.cw = remove.cw
			remove.cw.ccw = remove.ccw
			scores[player] += remove.value
			current = remove.cw
		}

		player = ((player + 1) % players)
	}
	top := 0
	for i, score := range scores {
		if score > top {
			top = score
			fmt.Printf("%v, %v\n", i, score)
		}
	}
}

var input = `411 players; last marble is worth 72059 points`
