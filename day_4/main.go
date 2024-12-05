package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.OpenFile("part1_input.txt", os.O_RDONLY, 0644)

	if err != nil {
		log.Fatalf("Could not open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	fmt.Println(part1(scanner))
}

func part1(scanner *bufio.Scanner) int {
	count := 0
	arr := [][]string{}

	// Format input to 2d array
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		arr = append(arr, row)
	}
	m := len(arr)
	n := len(arr[0])
	// DFS for the next letters
	for i, row := range arr {
		for j := range row {
			if arr[i][j] == "X" {
				count += Dfs(arr, i, j, m, n)
			}
		}
	}

	return count
}

func Dfs(arr [][]string, i int, j int, m int, n int) int {
	//[[1,0],[1,0],[1,0]] right
	//[[0,1],[0,1],[0,1]] down
	//[[-1,0],[-1,0],[-1,0]] left
	//[[0,1],[0,1],[0,1]] up
	matches := 0
	letters := []string{"M", "A", "S"}

	for _, pair := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}, {1, 1}, {1, -1}, {-1, -1}, {-1, 1}} {
		x := i
		y := j
		pointer := 0
		for pointer <= 2 { // loop three times
			x += pair[0]
			y += pair[1]
			fmt.Println(x, y)
			// fmt.Println(arr[x][y])
			if x < 0 || x >= m || y < 0 || y >= n || arr[x][y] != letters[pointer] {
				break
			}
			pointer += 1
		}

		if pointer == 3 {
			matches += 1
		}
	}

	return matches
}
