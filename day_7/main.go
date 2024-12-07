package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Numbers cannot be rearranged
// Start by mulitplying everything first
// if you go under the target, not possible
// else if you go over the target, backtrack to the last number and add it instead

func main() {
	fmt.Println("Hello, World!")

	file, _ := os.OpenFile("input1.txt", os.O_RDONLY, 0644)
	defer file.Close()
	input := make(map[int][]int)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		leftRight := strings.Split(scanner.Text(), ":")
		left, _ := strconv.Atoi(leftRight[0])

		right := strings.Split(leftRight[1], " ")[1:]

		for _, v := range right {
			if v != " " || v != "" {
				num, _ := strconv.Atoi(v)
				input[left] = append(input[left], num)
			}
		}
	}
	fmt.Println(part1(input))
}

func part1(input map[int][]int) int {
	ans := 0

	for k, v := range input {
		if isPossible(k, v) {
			fmt.Println("---------------------------------- ", k)
			ans += k
		}
	}

	return ans
}

func isPossible(value int, ops []int) bool {
	var results []int
	permuteOperations(ops, &results)

	for _, b := range results {
		if b == value {
			return true
		}
	}
	return false
}

func permuteOperations(nums []int, results *[]int) {
	var permute func(current int, index int)

	permute = func(current int, index int) {
		if index == len(nums)-1 {
			*results = append(*results, current)
			return
		}

		num := nums[index+1]
		permute(current+num, index+1)
		permute(current*num, index+1)
	}

	permute(nums[0], 0)
}
