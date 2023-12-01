package main

import (
	"os"
)

func main() {
	in := getinput()
	day1(in)
}

func getinput() string {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return string(dat)
}
