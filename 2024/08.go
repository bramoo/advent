package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	log.Print("Advent of Code 2024 08")

	file, err := os.Open("./example-08")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	width := 0
	height := 0
	antennae := make(map[rune][][]int)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		for i, e := range scanner.Text() {
			width = max(width, i+1)
			if e != '.' {
				antennae[e] = append(antennae[e], []int{i, height})
			}
		}
		height++
	}

	log.Print("width ", width, " height ", height)
	log.Print(antennae)

	antinodes := make([][]int, height)
	for i := range len(antinodes) {
		antinodes[i] = make([]int, width)
	}

	antinodes2 := make([][]int, height)
	for i := range len(antinodes2) {
		antinodes2[i] = make([]int, width)
	}

	for key, locs := range antennae {
		log.Print("calculating antinodes for ", key)
		for i, a := range locs {
			for _, b := range locs[i+1:] {
				log.Print("  antinodes for ", a, " and ", b)
				dx := b[0] - a[0]
				dy := b[1] - a[1]
				// antinode 1
				x := a[0] - dx
				y := a[1] - dy
				if 0 <= x && x < width && 0 <= y && y < height {
					antinodes[y][x] = 1
				}
				
				x = b[0] + dx
				y = b[1] + dy
				if 0 <= x && x < width && 0 <= y && y < height {
					antinodes[y][x] = 1
				}

				// part 2
				dx, dy = reduceStep(dx, dy)
				x = a[0]
				y = a[1]
				for {
					if x < 0 || width <= x || y < 0 || height <= y {
						break
					}
					antinodes2[y][x] = 1
					x += dx
					y += dy
				}

				x = a[0]
				y = a[1]
				for {
					if x < 0 || width <= x || y < 0 || height <= y {
						break
					}
					antinodes2[y][x] = 1
					x -= dx
					y -= dy
				}
			}
		}
	}

	antinodeCount := 0
	for _, row := range antinodes {
		log.Print(row)
		for _, e := range row {
			antinodeCount += e
		}
	}
	
	antinodeCount2 := 0
	for _, row := range antinodes2 {
		log.Print(row)
		for _, e := range row {
			antinodeCount2 += e
		}
	}

	log.Print("antinode count ", antinodeCount)
	log.Print("antinode2 count ", antinodeCount2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func reduceStep(dx, dy int) (int, int) {
	if dx == 0 || dy == 0 {
		return dx, dy
	}
	div := 1
	if dx < 0 && dy < 0 {
		div = gcd(-dx, -dy)
	} else if dx < 0 {
		div = gcd(-dx, dy)
	} else if dy < 0 {
		div = gcd(dx, -dy)
	} else {
		div = gcd(dx, dy)
	}
	return dx / div, dy / div
}

func gcd(a, b int) int {
	log.Print("gcd of ", a, " and ", b)
	if a == b { return a }
	if a < b { return gcd(b, a) }
	return gcd (a - b, b)
}
