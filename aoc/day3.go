package aoc

import (
	"fmt"
	"strconv"
	"strings"
)

type coord struct {
	r int
	c int
}

func isSymbol(r rune) bool {
	return !isNumber(r) && r != '.'
}

func idxIsSymbol(arr [][]rune, row, col int) bool {
	if row < 0 || col < 0 || row >= len(arr) || col >= len(arr[row]) {
		return false
	}

	return isSymbol(arr[row][col])
}

func idxIsGear(arr [][]rune, row, col int) bool {
	return idxIsSymbol(arr, row, col) && arr[row][col] == '*'
}

func Day3(in string) {
	lines := strings.Split(in, "\n")
	arr := [][]rune{}
	for _, l := range lines {
		arr = append(arr, []rune(l))
	}

	part_sum := 0
	is_part := false
	part_no := 0
	adj_gears := map[coord]bool{}

	gears := map[coord][]int{}

	for row, line := range arr {
		for col, n := range line {
			num, err := strconv.Atoi(string(n))
			if err != nil {
				if is_part {
					part_sum += part_no

					for g := range adj_gears {
						gears[g] = append(gears[g], part_no)
					}
				}
				part_no = 0
				is_part = false
				adj_gears = map[coord]bool{}
				continue
			}

			part_no *= 10
			part_no += num

			s := []int{-1, 0, 1}
			for _, r := range s {
				for _, c := range s {
					if idxIsSymbol(arr, row+r, col+c) {
						is_part = true
					}

					if idxIsGear(arr, row+r, col+c) {
						adj_gears[coord{row + r, col + c}] = true
					}
				}
			}
		}

		if is_part {
			part_sum += part_no
			for g := range adj_gears {
				gears[g] = append(gears[g], part_no)
			}
		}
		part_no = 0
		is_part = false
		adj_gears = map[coord]bool{}
	}

	ratio_sum := 0
	for _, parts := range gears {
		if len(parts) == 2 {
			ratio_sum += parts[0] * parts[1]
		}
	}

	fmt.Println("Part1: ", part_sum)
	fmt.Println("Part2: ", ratio_sum)
}
