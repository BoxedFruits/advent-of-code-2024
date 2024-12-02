package main

import (
	// "errors"
	//"fmt"
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type IntHeap []float64

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(float64))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	// Check if the input file exists
	file, err := os.OpenFile("input.txt", os.O_RDONLY, 0644)

	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// part1(scanner)
	part2(scanner)
}

func part1(scanner *bufio.Scanner) {
	leftCol := &IntHeap{}
	rightCol := &IntHeap{}

	heap.Init(leftCol)
	heap.Init(rightCol)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "   ")
		if line[0] == "" {
			break
		}

		left, _ := strconv.ParseFloat(line[0], 64)
		right, _ := strconv.ParseFloat(line[1], 64)
		heap.Push(leftCol, left)
		heap.Push(rightCol, right)
	}

	runningSum := 0.0

	for l := leftCol.Len(); l > 0; l-- {
		leftVal := heap.Pop(leftCol).(float64)
		rightVal := heap.Pop(rightCol).(float64)
		diff := math.Abs(leftVal - rightVal)

		runningSum += diff
	}

	fmt.Println("Part 1 Answer is: ", runningSum)

}

func part2(scanner *bufio.Scanner) {
	hm_left := make(map[string]int)
	hm_right := make(map[string]int)
	ans := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "   ")
		if line[0] == "" {
			break
		}

		hm_left[line[0]] += 1
		hm_right[line[1]] += 1
	}

	for key, val := range hm_left {
		fmt.Println(hm_left[key], hm_right[key], key, ans)
		k_s, _ := strconv.Atoi(key)
		ans += k_s * val * hm_right[key]
	}

	fmt.Println("Part 2 Answer is: ", ans)
}
