package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {

	r, _ := regexp.Compile("mul\\((\\d{1,3}),(\\d{1,3})\\)")
	rdo, _ := regexp.Compile("do\\(\\)")
	rdont, _ := regexp.Compile("don't\\(\\)")

	log.Print("Advent of Code 2024 03")

	file, err := os.Open("./example-03")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	mulSum := 0
	conditionalMulSum := 0
	do := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		for {
			log.Print(text)
			var nextDoSwitch []int
			if do {
				nextDoSwitch = rdont.FindStringIndex(text)
			} else {
				nextDoSwitch = rdo.FindStringIndex(text)
			}

			log.Print("do: ", do, ", next switch: ", nextDoSwitch)

			match := r.FindStringSubmatchIndex(text)
			log.Print("mul: ", match)

			if nextDoSwitch != nil && (match == nil || nextDoSwitch[0] < match[0]) {
				do = !do
				text = text[nextDoSwitch[1]:]
				continue
			}

			if match == nil {
				break
			}

			log.Print("do: ", do, ", match: ", text[match[0]:match[1]])
			x, _ := strconv.Atoi(text[match[2]:match[3]])
			y, _ := strconv.Atoi(text[match[4]:match[5]])
			
			mulSum += x * y
			
			if do {
				conditionalMulSum += x * y
			}

			text = text[match[1]:]
		}
	}

	log.Print("sum of mul instructions: ", mulSum)
	log.Print("sum of mul instructions in do: ", conditionalMulSum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
