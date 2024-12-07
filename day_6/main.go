package main

import (
	"bufio"
	"fmt"
	"guard/guard"
	"os"
	"strings"
)

func main() {
	input := [][]string{}

	file, _ := os.Open("test1_input.txt")
	defer file.Close()

	var nightGuard guard.Guard
	var startX int
	var startY int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		lineArr := make([]string, len(line))

		for i, c := range line {
			if c != "." && c != "#" { //Initiailze direction
				fmt.Println("Initiailze direction", guard.GetDirection(c))
				nightGuard.Direction = guard.GetDirection(c)
				startX = i
				startY = len(input)
			}
			lineArr[i] = c
		}

		input = append(input, lineArr)
	}

	fmt.Println(input)
	fmt.Println(part1(input, startX, startY, nightGuard))
}

func part1(input [][]string, startX int, startY int, nightGuard guard.Guard) int {
	currPosX := startX
	currPosY := startY
	seen := make(map[[2]int]struct{})
	seen[[2]int{currPosY, currPosX}] = struct{}{}

	leftBound := 0
	rightBound := len(input[0])
	topBound := 0
	bottomBound := len(input)

	directionCords := guard.GetDirectionCoordinates(nightGuard.Direction)

	for true {
		currPosX += directionCords[0]
		currPosY += directionCords[1]

		if currPosX < leftBound || currPosX >= rightBound || currPosY < topBound || currPosY >= bottomBound {
			return len(seen)
		}

		fmt.Println("currPosX", currPosX, "currPosY", currPosY, input[currPosY][currPosX])
		// Change direction is tile is "#"
		if input[currPosY][currPosX] == "#" {
			currPosX -= directionCords[0]
			currPosY -= directionCords[1]

			nightGuard.ChangeDirection(input[currPosY][currPosX])
			directionCords = guard.GetDirectionCoordinates(nightGuard.Direction)

			currPosX += directionCords[0]
			currPosY += directionCords[1]
		}

		if _, ok := seen[[2]int{currPosY, currPosX}]; !ok {
			// set.Insert([]int{currPosX, currPosY})
			seen[[2]int{currPosY, currPosX}] = struct{}{}
		}

	}

	return 0
}
