package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {

	r, err := regexp.Compile("mul\\((\\d{1,3}),(\\d{1,3})\\)")
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Advent of Code 2024 03")

	file, err := os.Open("./input-03")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	mulSum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches := r.FindAllStringSubmatch(scanner.Text(), -1)

		for _, match := range matches {
			log.Print(match)
			ai, err := strconv.Atoi(match[1])
			if err != nil {
				log.Fatal(err)
			}

			bi, err := strconv.Atoi(match[2])
			if err != nil {
				log.Fatal(err)
			}	

			log.Print("match with values ", ai, " and ", bi)
			mulSum += ai * bi
		}
	}

	log.Print("sum of mul instructions ", mulSum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
