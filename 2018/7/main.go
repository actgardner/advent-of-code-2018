package main

import (
	"fmt"
	"strings"
)

func main() {
	delim := "\n"
	parts := strings.Split(input, delim)
	adj := make([][]int, 26)
	for i := range adj {
		adj[i] = make([]int, 26)
	}
	for _, p := range parts {
		lParts := strings.Split(p, " ")
		dep := int(rune(lParts[1][0]) - 'A')
		ant := int(rune(lParts[7][0]) - 'A')
		adj[ant][dep] = 1
	}

	used := make([]int, 26)
	completed := make([]int, 26)
	tick := 1
	workerTime := make([]int, 5)
	workerTasks := make([]int, 5)
	for {
		avail := make([]int, 0)
		for i := 0; i < len(adj); i++ {
			valid := true
			for j := 0; j < len(adj[i]); j++ {
				if adj[i][j] != 0 {
					valid = false
				}
			}
			if valid && len(avail) < 5 && used[i] == 0 {
				avail = append(avail, i)
			}
		}
		for i := 0; i < len(workerTime); i++ {
			if workerTime[i] == 0 && len(avail) > 0 {
				workerTasks[i] = avail[0]
				workerTime[i] = 60 + avail[0] + 1
				used[avail[0]] = 1
				avail = avail[1:]
			}
			if workerTime[i] > 0 {
				workerTime[i] -= 1
				if workerTime[i] == 0 {
					completed[workerTasks[i]] = 1
					for j := 0; j < len(adj); j++ {
						adj[j][workerTasks[i]] = 0
					}
				}
			}
		}
		fmt.Printf("%v %v %v\n", tick, workerTime, workerTasks)
		complete := true
		for _, c := range completed {
			if c == 0 {
				complete = false
			}
		}
		if complete {
			return
		}
		tick += 1
	}
	fmt.Printf("%v", adj)
}

var input = `Step F must be finished before step R can begin.
Step I must be finished before step P can begin.
Step C must be finished before step O can begin.
Step H must be finished before step K can begin.
Step Y must be finished before step N can begin.
Step M must be finished before step J can begin.
Step D must be finished before step W can begin.
Step B must be finished before step N can begin.
Step T must be finished before step A can begin.
Step R must be finished before step L can begin.
Step P must be finished before step S can begin.
Step O must be finished before step J can begin.
Step X must be finished before step N can begin.
Step A must be finished before step K can begin.
Step N must be finished before step G can begin.
Step W must be finished before step U can begin.
Step Q must be finished before step U can begin.
Step V must be finished before step U can begin.
Step J must be finished before step G can begin.
Step G must be finished before step S can begin.
Step Z must be finished before step U can begin.
Step U must be finished before step S can begin.
Step E must be finished before step L can begin.
Step K must be finished before step L can begin.
Step L must be finished before step S can begin.
Step M must be finished before step N can begin.
Step T must be finished before step E can begin.
Step J must be finished before step U can begin.
Step G must be finished before step L can begin.
Step D must be finished before step P can begin.
Step T must be finished before step Z can begin.
Step U must be finished before step L can begin.
Step Z must be finished before step K can begin.
Step Q must be finished before step V can begin.
Step G must be finished before step K can begin.
Step Z must be finished before step E can begin.
Step Q must be finished before step Z can begin.
Step J must be finished before step S can begin.
Step G must be finished before step U can begin.
Step I must be finished before step M can begin.
Step W must be finished before step K can begin.
Step Y must be finished before step V can begin.
Step B must be finished before step Q can begin.
Step Y must be finished before step D can begin.
Step I must be finished before step G can begin.
Step A must be finished before step S can begin.
Step X must be finished before step S can begin.
Step O must be finished before step N can begin.
Step M must be finished before step X can begin.
Step V must be finished before step Z can begin.
Step W must be finished before step Z can begin.
Step C must be finished before step L can begin.
Step Q must be finished before step G can begin.
Step A must be finished before step U can begin.
Step G must be finished before step Z can begin.
Step P must be finished before step Q can begin.
Step C must be finished before step Z can begin.
Step U must be finished before step K can begin.
Step Q must be finished before step L can begin.
Step X must be finished before step U can begin.
Step A must be finished before step N can begin.
Step N must be finished before step S can begin.
Step Z must be finished before step L can begin.
Step F must be finished before step D can begin.
Step D must be finished before step A can begin.
Step J must be finished before step K can begin.
Step W must be finished before step Q can begin.
Step T must be finished before step J can begin.
Step T must be finished before step W can begin.
Step E must be finished before step K can begin.
Step P must be finished before step U can begin.
Step O must be finished before step Z can begin.
Step D must be finished before step B can begin.
Step R must be finished before step J can begin.
Step O must be finished before step A can begin.
Step N must be finished before step E can begin.
Step D must be finished before step G can begin.
Step M must be finished before step Q can begin.
Step F must be finished before step W can begin.
Step T must be finished before step L can begin.
Step U must be finished before step E can begin.
Step X must be finished before step L can begin.
Step M must be finished before step G can begin.
Step Z must be finished before step S can begin.
Step F must be finished before step Y can begin.
Step N must be finished before step Z can begin.
Step T must be finished before step U can begin.
Step D must be finished before step O can begin.
Step H must be finished before step X can begin.
Step V must be finished before step E can begin.
Step M must be finished before step T can begin.
Step Y must be finished before step O can begin.
Step P must be finished before step E can begin.
Step C must be finished before step E can begin.
Step P must be finished before step L can begin.
Step M must be finished before step A can begin.
Step F must be finished before step T can begin.
Step I must be finished before step C can begin.
Step X must be finished before step Z can begin.
Step Y must be finished before step U can begin.
Step B must be finished before step E can begin.`
