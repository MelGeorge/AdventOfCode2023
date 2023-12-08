package aoc

import (
	"fmt"
	"strings"
)

type node struct {
	left  string
	right string
}

var net = map[string]node{}
var dir string

func Day8(in string) {
	lines := strings.Split(in, "\n")
	dir = lines[0]
	a_paths := []string{}

	for i := 2; i < len(lines); i++ {
		src := lines[i][0:3]
		l := lines[i][7:10]
		r := lines[i][12:15]
		net[src] = node{left: l, right: r}
		if src[2] == 'A' {
			a_paths = append(a_paths, src)
		}
	}

	// Part 1: Start at AAA, end at ZZZ
	fmt.Println("Part 1: ", numSteps("AAA", func(loc string) bool {
		return loc == "ZZZ"
	}))

	// Part 2: Start at XXA, end at XXZ
	steps := []int{}
	for _, path := range a_paths {
		steps = append(steps, numSteps(path, func(loc string) bool {
			return loc[2] == 'Z'
		}))
	}

	fmt.Println("Part 2: ", LCM(steps[0], steps[1], steps[2:]...))
}

func numSteps(start string, done func(string) bool) int {
	curr, steps := 0, 0
	loc := start

	for !done(loc) {
		switch dir[curr] {
		case 'L':
			loc = net[loc].left
		case 'R':
			loc = net[loc].right
		}

		curr, steps = (curr+1)%len(dir), steps+1
	}

	return steps
}

// GCD & LCM functions STOLEN from:
// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
