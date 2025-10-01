package main

import (
	"fmt"
	"strconv"
)

type Run struct {
	Value int
	Count int
}

func convertToRun(input int) []Run {
	inputStr := strconv.Itoa(input)
	output := []Run{}
	prev := inputStr[0]
	count := 0
	for i := 0; i < len(inputStr); i++ {
		if inputStr[i] == prev {
			count++
		} else {
			val := int(prev - '0')
			output = append(output, Run{val, count})
			count = 1
			prev = inputStr[i]
		}
	}
	val := int(prev - '0')
	output = append(output, Run{val, count})
	return output
}

func transform(input []Run) []Run {
	output := make([]Run, 0, len(input)*2)
	for _, run := range input {
		nextRuns := []Run{{run.Count, 1}, {run.Value, 1}}
		for _, r := range nextRuns {
			if len(output) > 0 && output[len(output)-1].Value == r.Value {
				output[len(output)-1].Count += r.Count
			} else {
				output = append(output, r)
			}
		}
	}
	return output
}

func getLength(input []Run) int {
	output := 0
	for _, run := range input {
		output += run.Count
	}
	return output
}

func part1() {
	input := convertToRun(1113122113)
	for range 40 {
		input = transform(input)
	}
	fmt.Println(getLength(input))
}

func part2() {
	input := convertToRun(1113122113)
	for range 50 {
		input = transform(input)
	}
	fmt.Println(getLength(input))
}

func main() {
	part1()
	part2()
}
