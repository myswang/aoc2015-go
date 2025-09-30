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

	strs = make([]string, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		strs = append(strs, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func part1() {
	count := 0
	bannedWords := regexp.MustCompile(`ab|cd|pq|xy`)
	vowels := regexp.MustCompile(`a|e|i|o|u`)
	for _, str := range strs {
		match := bannedWords.FindString(str)
		if match != "" {
			continue
		}
		matchCount := len(vowels.FindAllString(str, -1))
		if matchCount < 3 {
			continue
		}

		for i := 0; i < len(str); i++ {
			if i > 0 && str[i] == str[i-1] {
				count++
				break
			}
		}
	}
	fmt.Println(count)
}

func part2() {
	count := 0
	for _, str := range strs {
		found := false
		for i := 0; !found && i < len(str)-2; i++ {
			p1 := str[i : i+2]
			for j := i + 2; j < len(str)-1; j++ {
				p2 := str[j : j+2]
				if p1 == p2 {
					found = true
					break
				}
			}
		}
		if found {
			for i := 0; i < len(str)-2; i++ {
				if str[i] == str[i+2] {
					count++
					break
				}
			}
		}
	}
	fmt.Println(count)
}

func main() {
	setup()
	part1()
	part2()
}
