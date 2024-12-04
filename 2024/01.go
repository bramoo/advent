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

	log.Print("Advent of Code 2024 01a")
	var a, b []int

	file, err := os.Open("./input-01")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		match := r.FindStringSubmatch(scanner.Text())
		// log.Print("match: ", match)
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
	}

	log.Print("a: ", len(a), "(", cap(a), ") b: ", len(b), "(", cap(b), ")")
	log.Print(a[:8])
	log.Print(b[:8])

	slices.Sort(a)
	slices.Sort(b)

	log.Print(a[:8])
	log.Print(b[:8])

	var sum int
	for i, _ := range a {
		sum += abs(a[i] - b[i])
	}

	b_index, b_count := advance(0, b)

	var match_sum int64
	match_loop:
	for i, left := range a {
		log.Print("at ", i, " with value ", left)

		for {
			if left <= b[b_index] {
				break
			}
			b_index, b_count = advance(b_index + 1, b)
			if b_index == len(b) {
				break match_loop
			}
		}

		if left == b[b_index] {
			log.Print("adding ", left * b_count)
			match_sum += int64(left * b_count)
		}
	}

	log.Print("Sum of absolute differences: ", sum)
	log.Print("Sum of product of a with count of matches in b: ", match_sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func advance(start_index int, values []int) (int, int) {
	if start_index >= len(values) {
		log.Print("Advance: start index past end of values")
		return start_index, 0
	}

	index := start_index
	count := 1
	
	for {
		if index + 1 >= len(values) || values[index] != values[index + 1] {
			break
		}
		index++
		count++
	}
	log.Print("Advanced to '", values[index], "' at: ", index, ", count: ", count)
	return index, count
}
