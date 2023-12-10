package aoc

import (
	"fmt"
	"strings"
)

func Day9(in string) {
	lines := strings.Split(in, "\n")
	sum1 := 0
	sum2 := 0

	for _, l := range lines {
		nums := tointarray(strings.Fields(l))
		hist := [][]int{nums}
		all0 := false

		for !all0 {
			all0 = true
			last := &hist[len(hist)-1]
			next := make([]int, len(*last)-1)
			for i := 0; i < len(*last)-1; i++ {
				next[i] = (*last)[i+1] - (*last)[i]
				all0 = all0 && next[i] == 0
			}
			hist = append(hist, next)
		}

		curr := len(hist) - 1
		hist[curr] = append(append([]int{0}, hist[curr]...), 0)

		for curr > 0 {
			hist[curr-1] = append([]int{hist[curr-1][0] - hist[curr][0]}, hist[curr-1]...)
			hist[curr-1] = append(hist[curr-1], hist[curr-1][len(hist[curr-1])-1]+hist[curr][len(hist[curr])-1])
			curr--
		}

		sum1 += hist[0][len(hist[0])-1]
		sum2 += hist[0][0]
	}

	fmt.Println("Part 1: ", sum1)
	fmt.Println("Part 2: ", sum2)
}
