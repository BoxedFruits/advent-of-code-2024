package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.OpenFile("part1_input.txt", os.O_RDONLY, 0644)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fmt.Println(part1(scanner))
}

func part1(scanner *bufio.Scanner) int {
	count := 0
	processUpdates := false
	rules := make(map[string][]string)

	//Page ordering rules
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			processUpdates = true
			continue
		}

		if !processUpdates {
			rule := strings.Split(line, "|")
			rules[rule[0]] = append(rules[rule[0]], rule[1])
		} else {
			count += processUpdate(rules, line)
		}
	}

	return count
}

func processUpdate(rules map[string][]string, update string) int {
	updateArr := strings.Split(update, ",")

	for _, v := range updateArr {
		pair := rules[v]

		idxFirst := containsA(updateArr, v)
		idxSec := containsB(updateArr, pair)

		if idxFirst > idxSec && idxSec != -1 {
			return 0
		}
	}
	ans, _ := strconv.Atoi(updateArr[len(updateArr)/2])
	return ans
}

func containsA(updateArr []string, target string) int {
	for k, v := range updateArr {
		if v == target {
			return k
		}
	}
	return -1
}

func containsB(updateArr []string, targets []string) int {
	for k, v := range updateArr {
		if slices.Contains(targets, v) {
			return k
		}
	}
	return -1
}
