package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const INPUT = "input.txt"

type CommandType int

const (
	TurnOff CommandType = iota
	TurnOn
	Toggle
)

var commandTypeNames = map[CommandType]string{
	TurnOff: "turn off",
	TurnOn:  "turn on",
	Toggle:  "toggle",
}

func (ct CommandType) String() string {
	return commandTypeNames[ct]
}

type Point struct {
	X int
	Y int
}

type Command struct {
	Name     CommandType
	StartPos Point
	EndPos   Point
}

var commands []Command

func parseInt(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("error when parsing int: %v", str)
	}
	return val
}

func parsePoint(str string) Point {
	pointVals := strings.Split(str, ",")
	x := parseInt(pointVals[0])
	y := parseInt(pointVals[1])

	return Point{x, y}
}

func setup() {
	f, err := os.Open(INPUT)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	cmdPattern := regexp.MustCompile(`turn (?:off|on)|toggle`)
	pointPattern := regexp.MustCompile(`\d+,\d+`)
	for scanner.Scan() {
		str := scanner.Text()
		cmd := Command{}
		switch cmdPattern.FindString(str) {
		case "turn off":
			cmd.Name = TurnOff
		case "turn on":
			cmd.Name = TurnOn
		case "toggle":
			cmd.Name = Toggle
		default:
			log.Fatal("no valid command name found")
		}

		pointStrs := pointPattern.FindAllString(str, -1)
		cmd.StartPos = parsePoint(pointStrs[0])
		cmd.EndPos = parsePoint(pointStrs[1])

		commands = append(commands, cmd)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func part1() {
	var lights [1000][1000]bool

	for _, cmd := range commands {
		start, end := cmd.StartPos, cmd.EndPos
		for x := start.X; x <= end.X; x++ {
			for y := start.Y; y <= end.Y; y++ {
				switch cmd.Name {
				case TurnOff:
					lights[x][y] = false
				case TurnOn:
					lights[x][y] = true
				case Toggle:
					lights[x][y] = !lights[x][y]
				}
			}
		}
	}

	count := 0
	for _, row := range lights {
		for _, light := range row {
			if light {
				count++
			}
		}
	}
	fmt.Println(count)
}

func part2() {
	var lights [1000][1000]int

	for _, cmd := range commands {
		start, end := cmd.StartPos, cmd.EndPos
		for x := start.X; x <= end.X; x++ {
			for y := start.Y; y <= end.Y; y++ {
				switch cmd.Name {
				case TurnOff:
					lights[x][y] = max(lights[x][y]-1, 0)
				case TurnOn:
					lights[x][y] += 1
				case Toggle:
					lights[x][y] += 2
				}
			}
		}
	}

	sum := 0
	for _, row := range lights {
		for _, light := range row {
			sum += light
		}
	}
	fmt.Println(sum)
}

func main() {
	setup()
	part1()
	part2()
}
