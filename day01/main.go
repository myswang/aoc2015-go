package main

import (
	"fmt"
	"log"
	"os"
)

const INPUT = "input.txt"

var directions string

func setup() {
	content, err := os.ReadFile(INPUT)
	if err != nil {
		log.Fatal(err)
	}
	directions = string(content)
}

func part1() {
	floor := 0
	for _, char := range directions {
		if char == '(' {
			floor++
		} else {
			floor--
		}
	}
	fmt.Printf("Part 1 answer: %d\n", floor)
}

func part2() {
	floor := 0
	for pos, char := range directions {
		if char == '(' {
			floor++
		} else {
			floor--
		}

		if floor == -1 {
			fmt.Printf("Part 2 answer: %d\n", pos+1)
			return
		}
	}
	fmt.Println("did not enter the basement")
}

func main() {
	setup()
	part1()
	part2()
}
