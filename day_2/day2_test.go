package day2

import (
	"fmt"
	"testing"
)

func TestDayOne(t *testing.T) {
	fmt.Println("RUNNING ETST")
	ans := part1("input_part1.txt")
	if ans != 2 {
		fmt.Println("Got: ", ans)
		t.Fatalf("Got the wrong answer should be 2")
	}
}
