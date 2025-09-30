package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
)

const INPUT = "input.txt"

var secret string

func setup() {
	content, err := os.ReadFile(INPUT)
	if err != nil {
		log.Fatal(err)
	}
	secret = string(content)
}

func part1() {
	for n := 0; ; n++ {
		str := secret + strconv.Itoa(n)
		hash := md5.Sum([]byte(str))
		hashStr := hex.EncodeToString(hash[:])

		if hashStr[:5] == "00000" {
			fmt.Println(n)
			return
		}
	}
}

func part2() {
	for n := 0; ; n++ {
		str := secret + strconv.Itoa(n)
		hash := md5.Sum([]byte(str))
		hashStr := hex.EncodeToString(hash[:])

		if hashStr[:6] == "000000" {
			fmt.Println(n)
			return
		}
	}
}

func main() {
	setup()
	part1()
	part2()
}
