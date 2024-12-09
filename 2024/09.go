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

	log.Print("Advent of Code 2024 09")

	file, err := os.Open("./input-09")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		disk := strings.Split(text, "")
		log.Print("disk: ", disk)
		var files []int

		for i, e := range disk {
			val, _ := strconv.Atoi(e)
			for range val {
				if i % 2 == 0 {
					files = append(files, i/2)
				} else {
					files = append(files, -1)
				}
			}
		}
		log.Print("expanded file table ", files)

		defrag := make([]int, len(files))
		copy(defrag, files)

		// rearrange files
		free := 0
		last := len(files) -1
		for free < last {
			for free < last && files[free] != -1 {
				free++
			}

			for free < last && files[last] == -1 {
				last--
			}

			files[free] = files[last]
			files[last] = -1
		}

		log.Print("rearranged files ", files)

		// defrag
		for id := len(disk) / 2; id >= 0; id-- {
			first := slices.Index(defrag, id)
			count := 1

			// count file block
			for first + count < len(defrag) && defrag[first + count] == id {
				count++
			}

			// find big enough gap
			gap := 0
			space := 0
			for gap < first {
				if defrag[gap] == -1 {
					space++
					if space == count {
						break
					}
				} else {
					space = 0
				}
				gap++
			}

			// move file block
			if space == count {
				for i := range count {
					defrag[gap - i] = id
					defrag[first + i] = -1
				}
			}
		}

		log.Print("defragged ", defrag)

		// checksum
		checksum := 0
		for i, e := range files {
			if e == -1 { continue }
			checksum += i * e
		}

		checksumDefrag := 0
		for i, e := range defrag {
			if e == -1 { continue }
			checksumDefrag += i * e
		}

		log.Print("checksum ", checksum)
		log.Print("checksum defrag ", checksumDefrag)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
