package day2

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1 Day 2: ", part1("test_input.txt"))
}

func part1(filename string) int {
	// scanner := GetScanner(filename )
	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)

	if err != nil {
		log.Fatalf("Could not open file")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	goodLevels := 0

	fmt.Println("HERE???")
	for scanner.Scan() {
		fmt.Println("starting??")

		if scanner.Text() == "" {
			break
		}

		levels := convertToInt(strings.Split(scanner.Text(), " "))
		shouldIncrease := false

		if levels[0]-levels[1] > 0 {
			shouldIncrease = true
		}

		if isSafeReport(levels, shouldIncrease) {
			goodLevels += 1
		}
	}

	return goodLevels
}

func isSafeReport(levels []int, shouldIncrease bool) bool {
	lenLevels := len(levels)

	for i := 1; i < lenLevels; i++ {
		curr := float64(levels[i])
		prev := float64(levels[i-1])

		if prev-curr == 0 || shouldIncrease && prev-curr < 0 || !shouldIncrease && prev-curr > 0 {
			return false
		}

		if math.Abs(curr-prev) > 3 {
			return false
		}
	}
	return true
}

func convertToInt(s []string) []int {
	arr := []int{}

	for _, v := range s {
		intVal, err := strconv.Atoi(v)

		if err != nil {
			log.Fatalf("Error converting input to int array")
		}

		arr = append(arr, intVal)
	}

	return arr
}

func GetScanner(filename string) *bufio.Scanner {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)

	if err != nil {
		log.Fatalf("Could not open file")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	return scanner
}
