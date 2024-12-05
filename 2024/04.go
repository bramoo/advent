package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	log.Print("Advent of Code 2024 04")

	file, err := os.Open("./input-04")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var rules [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		log.Print("rule: ", text)
		split := strings.Split(text, "|")
		first, _ := strconv.Atoi(split[0])
		second, _ := strconv.Atoi(split[1])
		rules = append(rules, []int{first, second})
	}

	midSum := 0
	var unordered [][]int
	update:
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, ",")
		update := make([]int, 0, len(split))
		for _, e := range split {
			i, _ := strconv.Atoi(e)
			update = append(update, i)
		}
		log.Print("update: ", update)

		for _, rule := range rules {
			first := slices.Index(update, rule[0])
			second := slices.Index(update, rule[1])
			if first == -1 || second == -1 {
				continue
			}
			if first > second {
				log.Print("violates ", rule)
				unordered = append(unordered, update)
				continue update
			}
		}

		middle := update[len(update)/2]
		log.Print("correctly ordered, mid: ", middle)
		midSum += middle
	}

	orderedMidSum := 0
	for _, update := range unordered {
		log.Print("reordering ", update)
		for {
			ordered := true
			for _, rule := range rules {
				first := slices.Index(update, rule[0])
				second := slices.Index(update, rule[1])
				if first == -1 || second == -1 {
					continue
				}
				if first > second {
					ordered = false
					better := make([]int, 0, len(update))
					better = append(better, update[:second]...)
					better = append(better, update[first])
					better = append(better, update[second])
					better = append(better, update[second+1:first]...)
					better = append(better, update[first+1:]...)
					log.Print(update, " + ", rule, " = ", better)
					update = better
				}
			}
			if ordered {
				orderedMidSum += update[len(update)/2]
				break
			}
		}
	}

	log.Print("Sum of middle values: ", midSum)
	log.Print("Sum of fixed middle values: ", orderedMidSum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
