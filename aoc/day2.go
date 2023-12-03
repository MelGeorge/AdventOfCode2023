package aoc

import (
	"fmt"
	"strconv"
	"strings"
)

func Day2(in string) {
	list := strings.Split(in, "\n")

	part1 := 0
	part2 := 0

	for num, str := range list {
		games := strings.Split(strings.Split(str, ":")[1], ";")

		possible := true
		r, g, b := 0, 0, 0
		for _, game := range games {
			m := getMarbles(game)

			if m["red"] > 12 || m["green"] > 13 || m["blue"] > 14 {
				possible = false
			}

			r, g, b = max(m["red"], r), max(m["green"], g), max(m["blue"], b)
		}

		if possible {
			part1 += (num + 1)
		}

		part2 += r * g * b
	}

	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", part2)
}

func getMarbles(game string) map[string]int {
	colors := strings.Split(game, ",")
	m := make(map[string]int)
	for _, color := range colors {
		c := strings.Fields(color)
		num, _ := strconv.Atoi(string(c[0]))
		m[string(c[1])] = num
	}
	return m
}
