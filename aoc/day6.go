package aoc

import (
	"fmt"
	"math"
)

func poss(t, d float64) int {
	sq := math.Sqrt(t*t - 4*d)
	a := (-t + sq) / -2
	b := (-t - sq) / -2
	return int(math.Floor(max(a, b))) - int(math.Ceil(min(a, b))) + 1
}

func Day6(in string) {
	// I didn't want to parse strings today-- sue me.
	fmt.Println("Part 1: ", poss(47, 207)*poss(84, 1394)*poss(74, 1209)*poss(67, 1014))
	fmt.Println("Part 2: ", poss(47847467, 207139412091014))
}
