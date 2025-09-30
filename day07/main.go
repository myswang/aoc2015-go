package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const INPUT = "input.txt"

type Result struct {
}

var env = make(map[string][]string)
var cache = make(map[string]uint16)

func eval(id string) uint16 {
	if val, ok := cache[id]; ok {
		return val
	}

	if v, err := strconv.Atoi(id); err == nil {
		val := uint16(v)
		cache[id] = val
		return val
	}

	op := env[id]
	if op == nil {
		log.Fatalf("Unknown wire or value: %v", id)
	}

	var val uint16
	switch len(op) {
	case 3:
		switch op[1] {
		case "AND":
			val = eval(op[0]) & eval(op[2])
		case "OR":
			val = eval(op[0]) | eval(op[2])
		case "LSHIFT":
			val = eval(op[0]) << eval(op[2])
		case "RSHIFT":
			val = eval(op[0]) >> eval(op[2])
		default:
			log.Fatalf("Invalid op: %v", op[1])
		}
	case 2:
		if op[0] == "NOT" {
			val = ^eval(op[1])
		} else {
			log.Fatalf("Invalid op: %v", op[0])
		}
	case 1:
		val = eval(op[0])
	default:
		log.Fatalf("Unexpected op length: %v (%v)", len(op), op)
	}

	cache[id] = val
	return val
}

func setup() {
	f, err := os.Open(INPUT)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " -> ")
		env[parts[1]] = strings.Fields(parts[0])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func part1() {
	res := eval("a")
	fmt.Println(res)
}

func part2() {
	env["b"] = []string{strconv.Itoa(int(cache["a"]))}
	cache = make(map[string]uint16)
	res := eval("a")
	fmt.Println(res)
}

func main() {
	setup()
	part1()
	part2()
}
