package aoc

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type Range struct {
	start int
	end   int
	add   int
}

var seeds []int
var seedRanges []Range
var almanac [][]Range

func Day5(in string) {
	parseInput(in)

	fmt.Println("Part 1: ", part1())
	fmt.Println("Part 2: ", part2())
}

func part1() int {
	low := int(math.MaxUint >> 1)
	for _, seed := range seeds {
		low = min(low, mapThroughAlmanac(Range{start: seed, end: seed})[0].start)
	}
	return low
}

func part2() int {
	low := int(math.MaxUint >> 1)
	for _, r := range seedRanges {
		res := mapThroughAlmanac(r)
		low = min(low, res[0].start)
	}
	return low
}

func mapThroughAlmanac(from Range) []Range {
	nextRange := []Range{from}
	var r []Range
	for i := 0; i < 7; i++ {
		r = nextRange
		nextRange = []Range{}
		var result []Range
		for j := 0; j < len(r); j++ {
			result = mapRange(r[j], almanac[i])
			nextRange = append(nextRange, result...)
		}
		sort.Slice(nextRange, func(a, b int) bool {
			return nextRange[a].start < nextRange[b].start
		})
	}

	return nextRange
}

func mapRange(from Range, into []Range) []Range {
	split := []Range{}

	f1 := from.start
	f2 := from.end
	next := f1

	if into[0].start > f1 {
		split = []Range{
			{
				start: f1,
				end:   min(into[0].start-1, f2),
				add:   0,
			},
		}
		next = min(into[0].start-1, f2) + 1
	}

	for next <= f2 {
		i := sort.Search(len(into), func(i int) bool {
			return next <= into[i].end
		})

		if i != len(into) && into[i].start <= next {
			end := min(into[i].end, f2)
			split = append([]Range{
				{
					start: next + into[i].add,
					end:   end + into[i].add,
					add:   0,
				}}, split...)
			next = end + 1
		} else {
			split = append(split, []Range{
				{
					start: next,
					end:   f2,
					add:   0,
				}}...)
			next = f2 + 1
		}
	}

	sort.Slice(split, func(i, j int) bool {
		return split[i].start < split[j].start
	})

	return split
}

func parseInput(in string) {
	lines := strings.Split(in, "\n")

	seeds = tointarray(strings.Fields(lines[0])[1:])
	for i := 0; i < len(seeds)-1; i += 2 {
		seedRanges = append(seedRanges, Range{start: seeds[i], end: seeds[i] + seeds[i+1] - 1, add: 0})
	}

	sort.Slice(seedRanges, func(i, j int) bool {
		return seedRanges[i].start < seedRanges[j].start
	})

	for _, line := range lines[1:] {
		if len(line) > 0 && !isNumber(rune(line[0])) {
			almanac = append(almanac, []Range{})
		} else if len(line) > 0 {
			vals := tointarray(strings.Fields(line))
			s := vals[1]
			e := s + vals[2] - 1
			a := vals[0] - s
			almanac[len(almanac)-1] = append(almanac[len(almanac)-1], Range{start: s, end: e, add: a})
		}
	}

	for a := range almanac {
		sort.Slice(almanac[a], func(i, j int) bool {
			return almanac[a][i].start < almanac[a][j].start
		})
	}
}
