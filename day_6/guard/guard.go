package guard

import (
	"fmt"
)

type Direction int // Enum

type Guard struct {
	Direction Direction
}

const (
	Up Direction = iota + 1
	Down
	Left
	Right
)

func (d Direction) String() string {
	return [...]string{"^", "V", "<", ">"}[d-1]
}

func GetDirection(c string) Direction {
	switch c {
	case "^":
		return Up
	case "V":
		return Down
	case "<":
		return Left
	case ">":
		return Right
	}
	fmt.Println("Invalid direction")
	return 0
}

func (g *Guard) ChangeDirection(tile string) {
	fmt.Println("Moving guard from", g.Direction)
	switch g.Direction {
	case Up:
		g.Direction = Right
	case Down:
		g.Direction = Left
	case Left:
		g.Direction = Up
	case Right:
		g.Direction = Down
	}

	fmt.Println("Moving guard to", g.Direction)
}

func GetDirectionCoordinates(d Direction) [2]int {
	switch d {
	case Up:
		return [2]int{0, -1}
	case Down:
		return [2]int{0, 1}
	case Left:
		return [2]int{-1, 0}
	case Right:
		return [2]int{1, 0}
	}
	return [2]int{0, 0}
}
