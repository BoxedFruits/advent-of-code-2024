package main

import (
	"fmt"
	"net/http"
)

func retrieve_input() {
	fmt.Println("Hello world")

	resp, err := http.Get("https://adventofcode.com/2024/day/1/input")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}
