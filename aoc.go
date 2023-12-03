package main

import (
	"aoc/aoc"
	"os"
)

func main() {
	in := getinput()
	aoc.Day3(in)
}

func getinput() string {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return string(dat)
}
