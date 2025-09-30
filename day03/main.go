package main

import (
	"fmt"
	"log"
	"maps"
	"os"
)

type Point struct {
	X int
	Y int
}

const INPUT = "input.txt"

var directions string

func walk(point *Point, visited map[Point]bool, dir rune) {
	switch dir {
	case '<':
		point.X--
	case '>':
		point.X++
	case '^':
		point.Y--
	case 'v':
		point.Y++
	default:
		log.Fatalf("Invalid direction: %c", dir)
		return
	}
	visited[*point] = true
}

func setup() {
	content, err := os.ReadFile(INPUT)
	if err != nil {
		log.Fatal(err)
	}
	directions = string(content)
}

func part1() {
	visited := make(map[Point]bool)
	point := Point{0, 0}
	visited[point] = true

	for _, dir := range directions {
		walk(&point, visited, dir)
	}

	fmt.Println(len(visited))
}

func part2() {
	v1 := make(map[Point]bool)
	v2 := make(map[Point]bool)

	p1 := Point{0, 0}
	p2 := Point{0, 0}
	v1[p1] = true
	v2[p2] = true

	for idx, dir := range directions {
		if idx%2 == 0 {
			walk(&p1, v1, dir)
		} else {
			walk(&p2, v2, dir)
		}
	}

	maps.Copy(v2, v1)
	fmt.Println(len(v2))
}

func main() {
	setup()
	part1()
	part2()
}
