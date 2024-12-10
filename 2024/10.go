package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"strconv"
)

func main() {

	log.Print("Advent of Code 2024 10")

	file, err := os.Open("./example-10")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var topo [][]int
	var heads [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		var row []int
		for x, e := range strings.Split(text, "") {
			height, _ := strconv.Atoi(e)
			row = append(row, height)
			if height == 0 {
				heads = append(heads, []int{x, len(topo)})
			}
		}
		topo = append(topo, row)
	}

	log.Print(topo)
	log.Print(heads)

	totalScore := 0
	totalScore2 := 0
	for _, head := range heads {
		totalScore += score(head, topo, make(map[int]bool))
		totalScore2 += score2(head, topo)
	}

	log.Print(totalScore)
	log.Print(totalScore2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func score(head []int, topo [][]int, ends map[int]bool) int {
	x := head[0]
	y := head[1]
	width := len(topo[0])
	idx := y * width + x
	val := topo[y][x]
	if val == 9 {
		if ends[idx] {
			return 0
		} else {
			ends[idx] = true
			return 1
		}
	}

	next := val + 1
	sum := 0
	if x > 0 && topo[y][x-1] == next {
		sum += score([]int{x-1, y}, topo, ends)
	}
	if y > 0 && topo[y-1][x] == next {
		sum += score([]int{x, y-1}, topo, ends)
	}
	if x < width - 1 && topo[y][x+1] == next {
		sum += score([]int{x+1, y}, topo, ends)
	}
	if y < len(topo) - 1 && topo[y+1][x] == next {
		sum += score([]int{x, y+1}, topo, ends)
	}
	return sum
}

func score2(head []int, topo [][]int) int {
	x := head[0]
	y := head[1]
	width := len(topo[0])
	val := topo[y][x]
	if val == 9 {
		return 1
	}

	next := val + 1
	sum := 0
	if x > 0 && topo[y][x-1] == next {
		sum += score2([]int{x-1, y}, topo)
	}
	if y > 0 && topo[y-1][x] == next {
		sum += score2([]int{x, y-1}, topo)
	}
	if x < width - 1 && topo[y][x+1] == next {
		sum += score2([]int{x+1, y}, topo)
	}
	if y < len(topo) - 1 && topo[y+1][x] == next {
		sum += score2([]int{x, y+1}, topo)
	}
	return sum
}
