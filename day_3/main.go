package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.OpenFile("part1_input.txt", os.O_RDONLY, 0644)

	if err != nil {
		log.Fatalf("Could not open the file")
	}

	defer file.Close()

	var validOp = regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	scanner := bufio.NewScanner(file)

	fmt.Println(test(scanner, validOp))
}

func test(scanner *bufio.Scanner, re *regexp.Regexp) int64 {
	var total int64

	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)

		for _, match := range matches {
			num1, _ := strconv.ParseInt(match[1], 10, 64)
			num2, _ := strconv.ParseInt(match[2], 10, 64)
			fmt.Println(num1, num2)
			total += (num1 * num2)
		}
	}

	return total
}
