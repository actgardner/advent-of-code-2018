package main

import (
	"fmt"
)

func main() {
	sums := []int{0}
	sum := 0
	for {
		for _, i := range input {
			sum += i
			for _, j := range sums {
				if j == sum {
					fmt.Printf("repeat - %v", j)
					return
				}
			}
			sums = append(sums, sum)
		}
	}
	fmt.Printf("%v", sum)
}
