package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"strconv"
)

func main() {

	log.Print("Advent of Code 2024 07")

	file, err := os.Open("./example-07")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalCalibration := 0
	totalCalibration2 := 0

	for scanner.Scan() {
		text := scanner.Text()
		colon := strings.Index(text, ": ")
		if colon == -1 {
			continue
		}
		target, _ := strconv.Atoi(text[:colon])
		var values []int
		for _, v := range strings.Split(text[colon+2:], " ") {
			val, _ := strconv.Atoi(v)
			values = append(values, val)
		}

		log.Print(target, values)

		if sovl(target, values[0], values[1:], Part1) {
			totalCalibration += target
			totalCalibration2 += target
			log.Print("  solved (1) ", target)
		} else if sovl(target, values[0], values[1:], Part2) {
			totalCalibration2 += target
			log.Print("  solved (2) ", target)
		}
	}

	log.Print("total calibration result: ", totalCalibration)
	log.Print("total calibration with || result: ", totalCalibration2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func sovl(target, acc int, remaining []int, part Part) bool {
	if acc > target {
		return false // operators can only increase value of result
	}

	if len(remaining) == 0 {
		return acc == target
	}

	next := remaining[0]
	tail := remaining[1:]
	p1 := sovl(target, acc + next, tail, part) || sovl(target, acc * next, tail, part)

	if part == Part1 {
		return p1
	} else {
		concat, _ := strconv.Atoi(strconv.Itoa(acc) + strconv.Itoa(next))
		return p1 || sovl(target, concat, tail, part)
	}
}

type Part int
const (
	Part1 Part = 1
	Part2 Part = 2
)
