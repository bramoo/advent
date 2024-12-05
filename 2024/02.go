package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"strconv"
)

func main() {

	log.Print("Advent of Code 2024 02")

	file, err := os.Open("./example-02")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var countSafe, countSafeDampened int

	line:
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		report := make([]int, len(split))
		for i, e := range split {
			num, err := strconv.Atoi(e)
			if err != nil {
				log.Fatal(err)
			}
			report[i] = num
		}

		log.Print(report)
		
		if isReportSafe(report) {
			countSafe++
			countSafeDampened++
			log.Print("report ", report, " safe")
			continue
		}

		for i := range len(report) {
			sub := make([]int, 0, len(report) - 1)
			sub = append(sub, report[:i]...)
			sub = append(sub, report[i+1:]...)

			if isReportSafe(sub) {
				log.Print("report ", report, " safe with ", i, " removed: ", sub)
				countSafeDampened++
				continue line
			}
		}

		log.Print("report ", report, " unsafe")
	}

	log.Print(countSafe, " safe reports")
	log.Print(countSafeDampened, " safe reports with dampening")
}

func isReportSafe(report []int) bool {
	if len(report) < 2 {
		return true
	}

	first := report[0]
	second := report[1]
	delta := second - first

	if delta < -3 || -1 < delta && delta < 1 || 3 < delta {
		log.Print("unsafe: delta out of range")
		return false
	}

	for _, e := range report[2:] {
		first = second
		second = e
		d := second - first

		if d < -3 || -1 < d && d < 1 || 3 < d {
			log.Print("unsafe: delta out of range")
			return false
		}

		if delta * d < 1 {
			log.Print("unsafe: non monotonic")
			return false
		}
	}

	return true
}
