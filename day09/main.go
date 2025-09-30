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

var dist = make(map[string]map[string]int)

func dfsMin(node string, cost int, step int, visited map[string]bool, minCost int) int {
	if visited[node] {
		return minCost
	}
	if step == len(dist)-1 {
		minCost = min(minCost, cost)
		return minCost
	}
	visited[node] = true
	for k, v := range dist[node] {
		minCost = min(minCost, dfsMin(k, cost+v, step+1, visited, minCost))
	}
	visited[node] = false
	return minCost
}

func dfsMax(node string, cost int, step int, visited map[string]bool, maxCost int) int {
	if visited[node] {
		return maxCost
	}
	if step == len(dist)-1 {
		maxCost = max(maxCost, cost)
		return maxCost
	}
	visited[node] = true
	for k, v := range dist[node] {
		maxCost = max(maxCost, dfsMax(k, cost+v, step+1, visited, maxCost))
	}
	visited[node] = false
	return maxCost
}

func setup() {
	f, err := os.Open(INPUT)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		p1, p2 := parts[0], parts[2]
		d, err := strconv.Atoi(parts[4])
		if err != nil {
			log.Fatalf("Failed to parse int: %v", parts[4])
		}
		if dist[p1] == nil {
			dist[p1] = make(map[string]int)
		}

		if dist[p2] == nil {
			dist[p2] = make(map[string]int)
		}
		dist[p1][p2] = d
		dist[p2][p1] = d
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func part1() {
	minCost := math.MaxInt
	for k := range dist {
		minCost = min(minCost, dfsMin(k, 0, 0, make(map[string]bool), math.MaxInt))
	}
	fmt.Println(minCost)
}

func part2() {
	maxCost := math.MinInt
	for k := range dist {
		maxCost = max(maxCost, dfsMax(k, 0, 0, make(map[string]bool), math.MinInt))
	}
	fmt.Println(maxCost)
}

func main() {
	setup()
	part1()
	part2()
}
