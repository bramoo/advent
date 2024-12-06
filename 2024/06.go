package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {

	log.Print("Advent of Code 2024 06")

	file, err := os.Open("./input-06")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var maap []string
	var x, y int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		maap = append(maap, text)
		loc := strings.Index(text, "^")
		if loc != -1 {
			x = loc
			y = len(maap) - 1
		}
	}

	w := len(maap[0])
	h := len(maap)
	log.Print(maap)
	// log.Print("guard at ", x, ", ", y)
	sumPatrolled := getPatrol(x, y, maap)

	sumStuck := 0
	for i := range w {
		for j := range h {
			if i == x && j == y {
				continue
			}
			
			t := maap[j]
			row := []rune(t)
			row[i] = '#'
			maap[j] = string(row)
			patrolled := getPatrol(x, y, maap)
			maap[j] = t
			// log.Print(patrolled)
			if patrolled == -1 {
				sumStuck++
			}
		}
	}

	log.Print(sumPatrolled, " locations visited")
	log.Print(sumStuck, " stuck patrols")

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getPatrol(x int, y int, maap []string) int {
	w := len(maap[0])
	h := len(maap)
	dx := 0
	dy := -1
	hx := x
	hy := y
	hdx := dx
	hdy := dy

	patrol := make([][]byte, h)
	for i, _ := range patrol {
		patrol[i] = make([]byte, w)
	}

	for {
		// log.Print("guard at ", x, ", ", y)
		patrol[y][x] = 1
		x, y, dx, dy = move(x, y, dx, dy, maap)
		if x < 0 || x >= w || y < 0 || y >= h {
			break
		}

		hx, hy, hdx, hdy = move(hx, hy, hdx, hdy, maap)
		hx, hy, hdx, hdy = move(hx, hy, hdx, hdy, maap)

		if x == hx && y == hy && dx == hdx && dy == hdy {
			// in a loop
			return -1
		}
	}

	sumPatrolled := 0
	log.Print("guard left the map")
	for _, p := range patrol {
		for _, v := range p {
			sumPatrolled += int(v)
		}
	}

	return sumPatrolled
}

func move(x int, y int, dx int, dy int, maap []string) (int, int, int, int) {
	w := len(maap[0])
	h := len(maap)
	nx := x + dx
	ny := y + dy

	if nx < 0 || nx >= w || ny < 0 || ny >= h {
		return nx, ny, dx, dy
	} else if maap[ny][nx] == '#' {
		return x, y, -dy, dx
	}

	return nx, ny, dx, dy
}
