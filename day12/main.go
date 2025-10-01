package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const INPUT = "input.txt"

var jsonStr string

func getSumObject(jsonObj map[string]any) int {
	sum := 0
	for _, val := range jsonObj {
		switch val := val.(type) {
		case float64:
			sum += int(val)
		case string:
			str := val
			if str == "red" {
				return 0
			}
		case []any:
			sum += getSumArray(val)
		case map[string]any:
			sum += getSumObject(val)
		default:
			log.Fatalf("unknown object: %v", val)
		}
	}
	return sum
}

func getSumArray(jsonObj []any) int {
	sum := 0
	for _, val := range jsonObj {
		switch val := val.(type) {
		case float64:
			sum += int(val)
		case string:
			// do nothing
		case []any:
			sum += getSumArray(val)
		case map[string]any:
			sum += getSumObject(val)
		default:
			log.Fatalf("unknown object: %v", val)
		}
	}
	return sum
}

func getSum() int {
	pattern := regexp.MustCompile(`-?\d+`)
	matches := pattern.FindAllString(jsonStr, -1)

	sum := 0
	for _, match := range matches {
		val, err := strconv.Atoi(match)
		if err != nil {
			log.Fatalf("Failed to parse int: %v", match)
		}
		sum += val
	}

	return sum
}

func setup() {
	content, err := os.ReadFile(INPUT)
	if err != nil {
		log.Fatal(err)
	}

	jsonStr = string(content)
}

func part1() {
	sum := getSum()
	fmt.Println(sum)
}

func part2() {
	var jsonObj map[string]any

	err := json.Unmarshal([]byte(jsonStr), &jsonObj)

	if err != nil {
		log.Fatal("failed to parse json")
	}

	sum := getSumObject(jsonObj)
	fmt.Println(sum)

}

func main() {
	setup()
	part1()
	part2()
}
