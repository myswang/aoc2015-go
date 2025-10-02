package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Reindeer struct {
	Speed    int
	Duration int
	Cooldown int
}

type ReindeerState struct {
	Distance int
	Points   int
	Timer    int
	IsMoving bool
}

const INPUT = "input.txt"

var reindeer []Reindeer

func setup() {
	f, err := os.Open(INPUT)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	re := regexp.MustCompile(`\d+`)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		matches := re.FindAllString(scanner.Text(), -1)
		vals := []int{}
		for _, match := range matches {
			val, err := strconv.Atoi(match)
			if err != nil {
				log.Fatal(err)
			}
			vals = append(vals, val)
		}
		reindeer = append(reindeer, Reindeer{
			vals[0], vals[1], vals[2],
		})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func part1() {
	maxDist := 0
	totalTime := 2503

	for _, rd := range reindeer {
		curTime := totalTime
		curDist := 0
		for curTime > 0 {
			curDist += rd.Speed * min(rd.Duration, curTime)
			curTime = max(curTime-rd.Duration-rd.Cooldown, 0)
		}
		if curDist > maxDist {
			maxDist = curDist
		}
	}
	fmt.Println(maxDist)
}

func part2() {
	totalTime := 2503
	rdStates := []ReindeerState{}

	for i := range reindeer {
		rdStates = append(rdStates, ReindeerState{
			0, 0, reindeer[i].Duration, true,
		})
	}

	for range totalTime {
		for i := range rdStates {
			if rdStates[i].IsMoving {
				rdStates[i].Distance += reindeer[i].Speed
			}
			rdStates[i].Timer--
			if rdStates[i].Timer == 0 {
				if rdStates[i].IsMoving {
					rdStates[i].Timer = reindeer[i].Cooldown
				} else {
					rdStates[i].Timer = reindeer[i].Duration
				}
				rdStates[i].IsMoving = !rdStates[i].IsMoving
			}
		}
		leadIdx := []int{}
		leadDist := 0

		for i := range rdStates {
			if rdStates[i].Distance > leadDist {
				leadIdx = []int{i}
				leadDist = rdStates[i].Distance
			} else if rdStates[i].Distance == leadDist {
				leadIdx = append(leadIdx, i)
			}
		}
		for _, idx := range leadIdx {
			rdStates[idx].Points++
		}
	}

	maxPoints := 0
	for i := range rdStates {
		if rdStates[i].Points > maxPoints {
			maxPoints = rdStates[i].Points
		}
	}

	fmt.Println(maxPoints)
}

func main() {
	setup()
	part1()
	part2()
}
