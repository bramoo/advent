package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	log.Print("Advent of Code 2024 04")

	file, err := os.Open("./example-04")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	xmasCount := 0
	x_masCount:= 0

	for i := range len(lines) {
		for j := range len(lines[i]) {
			if j < len(lines[i]) -3 {
				// horizontal
				sub := lines[i][j:j+4]
				if sub == "XMAS" {
					log.Print("XMAS (hf) at ", i, ", ", j)
					xmasCount++
				} else if sub == "SAMX" {
					log.Print("XMAS (hb) at ", i, ", ", j+3)
					xmasCount++
				}
			}

			if i < len(lines) - 3 {
				// vertical
				sub := btos(lines[i][j], lines[i+1][j], lines[i+2][j], lines[i+3][j])
				if sub == "XMAS" {
					log.Print("XMAS (vf) at ", i, ", ", j)
					xmasCount++
				} else if sub == "SAMX" {
					log.Print("XMAS (vb) at ", i+3, ", ", j)
					xmasCount++
				}
			}

			if i < len(lines) - 3 && j < len(lines[i]) - 3 {
				// diag
				sub := btos(lines[i][j], lines[i+1][j+1], lines[i+2][j+2], lines[i+3][j+3])
				if sub == "XMAS" {
					log.Print("XMAS (dr) at ", i, ", ", j)
					xmasCount++
				} else if sub == "SAMX" {
					log.Print("XMAS (ul) at ", i+3, ", ", j+3)
					xmasCount++
				}

				sub = btos(lines[i+3][j], lines[i+2][j+1], lines[i+1][j+2], lines[i][j+3])
				if sub == "XMAS" {
					log.Print("XMAS (ur) at ", i+3, j)
					xmasCount++
				} else if sub == "SAMX" {
					log.Print("XMAS (dl) at ", i, ", ", j+3)
					xmasCount++
				}
			}
			
			// x-mas
			if 0 < i && i < len(lines) - 1 && 0 < j &&j < len(lines[i]) - 1 {
				if lines[i][j] == 'A' {
					d1 := lines[i-1][j-1]
					d2 := lines[i+1][j+1]
					d3 := lines[i-1][j+1]
					d4 := lines[i+1][j-1]

					if d1 == 'M' && d2 == 'S' || d1 == 'S' && d2 == 'M' {
						if d3 == 'M' && d4 == 'S' || d3 == 'S' && d4 == 'M' {
							x_masCount++
						}
					}
				}
			}
		}
	}

	log.Print("found ", xmasCount, " XMAS and ", x_masCount, " x-mas")
}

func btos(b ...byte) string {
	return string(b)
}
