package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {

	log.Print("Advent of Code 2024 11")

	file, err := os.Open("./example-11")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	count25 := 0
	count75 := 0
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		count25 += blink(25, num)
		count75 += blink(75, num)
	}
	log.Print("count after 25 blinks: ", count25)
	log.Print("count after 75 blinks: ", count75)
}

func blink(n, val int) int {
	cache := make(map[Key]int)
	return memoBlink(n, val, cache)
}

func memoBlink(n, val int, cache map[Key]int) int {
	if n == 0 {
		return 1
	}

	cached := cache[Key{val, n}]
	if cached != 0 {
		return cached
	}

	var ret int
	if val == 0 {
		ret = memoBlink(n-1, 1, cache)
	} else if digits(val) % 2 == 0 {
		s := strconv.Itoa(val)
		left, _ := strconv.Atoi(s[:len(s)/2])
		right, _ := strconv.Atoi(s[len(s)/2:])
		ret = memoBlink(n-1, left, cache) + memoBlink(n-1, right, cache)
	} else {
		ret = memoBlink(n-1, val * 2024, cache)
	}

	cache[Key{val, n}] = ret
	return ret
}

type Key struct {
	Value, Iterations int
}

func digits(n int) int {
	l10 := math.Log10(float64(n+1))
	return int(math.Ceil(l10))
}
