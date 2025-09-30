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

var env = make(map[string][]string)
var cache = make(map[string]uint16)

func eval(id string) uint16 {
	if val, ok := cache[id]; ok {
		return val
	}

	if val, err := strconv.Atoi(id); err == nil {
		cache[id] = uint16(val)
		return cache[id]
	}

	op := env[id]
	if op == nil {
		log.Fatalf("Unknown wire or value: %v", id)
	}

	switch len(op) {
	case 3:
		switch op[1] {
		case "AND":
			cache[id] = eval(op[0]) & eval(op[2])
		case "OR":
			cache[id] = eval(op[0]) | eval(op[2])
		case "LSHIFT":
			cache[id] = eval(op[0]) << eval(op[2])
		case "RSHIFT":
			cache[id] = eval(op[0]) >> eval(op[2])
		default:
			log.Fatalf("Invalid op: %v", op[1])
		}
	case 2:
		if op[0] == "NOT" {
			cache[id] = ^eval(op[1])
		} else {
			log.Fatalf("Invalid op: %v", op[0])
		}
	case 1:
		cache[id] = eval(op[0])
	default:
		log.Fatalf("Unexpected op length: %v (%v)", len(op), op)
	}

	return cache[id]
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
