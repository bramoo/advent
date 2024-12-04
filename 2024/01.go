package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {

	r, err := regexp.Compile("(\\d+)\\s+(\\d+)")
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Advent of Code 2024 01")
	var a, b []int
	counts := make(map[int]int)

	file, err := os.Open("./input-01")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		match := r.FindStringSubmatch(scanner.Text())

		ai, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal(err)
		}

		bi, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatal(err)
		}	

		a = append(a, ai)
		b = append(b, bi)
		counts[bi] = counts[bi] + 1
	}

	log.Print(len(a), " as and ", len(b), " bs")
	log.Print("first 8 as and bs")
	log.Print(a[:8])
	log.Print(b[:8])

	slices.Sort(a)
	slices.Sort(b)

	log.Print("first 8 sorted as and bs")
	log.Print(a[:8])
	log.Print(b[:8])

	var sum, match_sum int
	for i, el := range a {
		sum += abs(el - b[i])
		match_sum += el * counts[el]
	}

	log.Print("Sum of absolute differences: ", sum)
	log.Print("Sum of product of a with count of matches in b: ", match_sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
