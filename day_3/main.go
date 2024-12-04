package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	// "text/scanner"
)

func main() {
	file, err := os.OpenFile("part1_input.txt", os.O_RDONLY, 0644)

	if err != nil {
		log.Fatalf("Could not open the file")
	}

	defer file.Close()

	// var validOp = regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	scanner := bufio.NewScanner(file)

	//fmt.Println(test(scanner, validOp)) //also part 1

	var part2Op = regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)|don\'t\(\)|do\(\)`)

	fmt.Println(part2(scanner, part2Op))
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

func part2(scanner *bufio.Scanner, re *regexp.Regexp) int64 {
	var total int64
	doIt := true

	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		fmt.Println(matches)

		for _, match := range matches {
			if match[0] == `do()` {
				doIt = true
				continue
			} else if match[0] == `don't()` {
				doIt = false
				continue
			} else if doIt == true {
				num1, _ := strconv.ParseInt(match[1], 10, 64)
				num2, _ := strconv.ParseInt(match[2], 10, 64)
				fmt.Println(num1, num2)
				total += (num1 * num2)

			}
		}
	}
	return total
}
