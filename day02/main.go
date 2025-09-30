package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

const INPUT = "input.txt"

var dimensions [][]int

func setup() {
	f, err := os.Open(INPUT)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), "x")
		vals := make([]int, 3)

		for i, str := range strs {
			val, err := strconv.Atoi(str)
			if err != nil {
				log.Fatal(err)
			}
			vals[i] = val
		}
		dimensions = append(dimensions, vals)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func part1() {
	sum := 0
	for _, dim := range dimensions {
		s1 := dim[0] * dim[1]
		s2 := dim[0] * dim[2]
		s3 := dim[1] * dim[2]
		slack := min(s1, s2, s3)
		sum += 2*(s1+s2+s3) + slack
	}

	fmt.Printf("Part 1 answer: %d\n", sum)
}

func part2() {
	sum := 0
	for _, dim := range dimensions {
		slices.Sort(dim)
		ribbon := 2*dim[0] + 2*dim[1]
		volume := dim[0] * dim[1] * dim[2]
		sum += ribbon + volume
	}

	fmt.Printf("Part 2 answer: %d\n", sum)
}

func main() {
	setup()
	part1()
	part2()
}
