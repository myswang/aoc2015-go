package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

const INPUT = "input.txt"

var strs []string

func setup() {
	f, err := os.Open(INPUT)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		strs = append(strs, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func part1() {
	pattern := regexp.MustCompile(`\\(?:\\|"|x[0-9a-f]{2})|[a-z]`)
	sum := 0
	for _, str := range strs {
		matches := pattern.FindAllString(str, -1)
		sum += len(str) - len(matches)
	}

	fmt.Println(sum)
}

func part2() {
	sum := 0
	for _, str := range strs {
		newStr := ""
		for _, char := range str {
			switch char {
			case '\\':
				newStr += "\\\\"
			case '"':
				newStr += "\\\""
			default:
				newStr += string(char)
			}
		}
		sum += len(newStr) - len(str) + 2
	}
	fmt.Println(sum)
}

func main() {
	setup()
	part1()
	part2()
}
