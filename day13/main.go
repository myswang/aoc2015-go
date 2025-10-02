package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const INPUT = "input.txt"

var n = 8
var dist [][]int

func setup() {
	f, err := os.Open(INPUT)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	dist = make([][]int, n)
	for i := range n {
		dist[i] = make([]int, n)
	}

	people := make(map[string]int)
	scanner := bufio.NewScanner(f)
	nextId := 0
	for scanner.Scan() {
		line := scanner.Text()
		line = line[:len(line)-1]
		parts := strings.Fields(line)
		p1, p2, gl := parts[0], parts[10], parts[2]
		val, err := strconv.Atoi(parts[3])
		if err != nil {
			log.Fatal(err)
		}
		if _, ok := people[p1]; !ok {
			people[p1] = nextId
			nextId++
		}
		if _, ok := people[p2]; !ok {
			people[p2] = nextId
			nextId++
		}
		if gl == "gain" {
			dist[people[p1]][people[p2]] += val
			dist[people[p2]][people[p1]] += val
		} else {
			dist[people[p1]][people[p2]] -= val
			dist[people[p2]][people[p1]] -= val
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// Held-Karp algorithm:
// https://en.wikipedia.org/wiki/Held%E2%80%93Karp_algorithm
func heldKarp() {

	// dp[mask][j] = max cost to reach subset "mask" ending at person j
	dp := make([][]int, 1<<n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MinInt32
		}
	}
	// Base case: start at person 0
	dp[1][0] = 0

	// Iterate over all subsets of persons
	for mask := 1; mask < 1<<n; mask++ {
		for j := range n {
			if mask&(1<<j) == 0 { // person j not in mask
				continue
			}
			prevMask := mask ^ (1 << j)
			if prevMask == 0 {
				continue
			}
			for i := range n {
				if prevMask&(1<<i) != 0 {
					cost := dp[prevMask][i] + dist[i][j]
					if cost > dp[mask][j] {
						dp[mask][j] = cost
					}
				}
			}
		}
	}

	// Find best tour cost that returns to 0
	fullMask := (1 << n) - 1
	maxCost := math.MinInt32
	for j := 1; j < n; j++ {
		cost := dp[fullMask][j] + dist[j][0]
		if cost > maxCost {
			maxCost = cost
		}
	}

	fmt.Println(maxCost)
}

func part1() {
	heldKarp()
}

func part2() {
	for i := range n {
		dist[i] = append(dist[i], 0)
	}
	n++
	dist = append(dist, make([]int, n))

	heldKarp()
}

func main() {
	setup()
	part1()
	part2()
}
