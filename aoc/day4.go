package aoc

import (
	"fmt"
	"strings"
)

func Day4(in string) {
	cards := strings.Split(in, "\n")

	total := 0
	copies := []int{}
	for i := 0; i < len(cards); i++ {
		copies = append(copies, 1)
	}

	for i, c := range cards {
		points := 0
		wins := 0
		card := strings.Split(strings.Split(c, ": ")[1], " | ")
		w := tointarray(strings.Fields(card[0]))
		a := tointarray(strings.Fields(card[1]))

		win := map[int]bool{}
		for _, x := range w {
			win[x] = true
		}

		for _, y := range a {
			if win[y] {
				wins++
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}

		total += points

		for p := 1; p <= wins; p++ {
			copies[i+p] += copies[i]
		}
	}

	fmt.Println("Part 1: ", total)
	fmt.Println("Part 2: ", sum(copies))
}
