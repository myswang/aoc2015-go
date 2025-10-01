package main

import (
	"fmt"
	"strings"
)

const LETTERS = "abcdefghijklmnopqrstvwxyz"

func hasPair(input []int) bool {
	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			for j := i + 2; j < len(input)-1; j++ {
				if input[j] == input[j+1] && input[j] != input[i] {
					return true
				}
			}
		}
	}
	return false
}

func hasRun(input []int) bool {
	for i := 0; i < len(input)-2; i++ {
		if input[i+1]-input[i] == 1 && input[i+2]-input[i+1] == 1 {
			return true
		}
	}
	return false
}

func skipBad(input []int) bool {
	ok := false
	for i := range input {
		if input[i] == 8 || input[i] == 11 || input[i] == 14 {
			increment(input, i)
			for j := i + 1; j < len(input); j++ {
				input[j] = 0
			}
			ok = true
			break
		}
	}
	return ok
}

func increment(input []int, idx int) {
	input[idx]++
	if input[idx] >= len(LETTERS) {
		input[idx] = 0
		increment(input, idx-1)
	}
}

func printPassword(input []int) {
	output := ""
	for _, val := range input {
		output += string(LETTERS[val])
	}

	fmt.Println(output)
}

func part1() {
	inputStr := "vzbxkghb"
	input := []int{}

	for _, c := range inputStr {
		val := strings.IndexRune(LETTERS, c)
		input = append(input, val)
	}

	for {
		ok := skipBad(input)
		if hasPair(input) && hasRun(input) {
			break
		}
		if !ok {
			increment(input, len(input)-1)
		}
	}

	printPassword(input)
}

func part2() {
	inputStr := "vzbxxzaa"
	input := []int{}

	for _, c := range inputStr {
		val := strings.IndexRune(LETTERS, c)
		input = append(input, val)
	}

	for {
		ok := skipBad(input)
		if hasPair(input) && hasRun(input) {
			break
		}
		if !ok {
			increment(input, len(input)-1)
		}
	}

	printPassword(input)
}

func main() {
	part1()
	part2()
}
