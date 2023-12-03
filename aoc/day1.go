package aoc

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var NumberStrings = map[int]string{
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

type num struct {
	value int
	index int
}

type ByIdx []num

func (a ByIdx) Len() int           { return len(a) }
func (a ByIdx) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByIdx) Less(i, j int) bool { return a[i].index < a[j].index }

func Day1(in string) {
	arr := strings.Split(in, "\n")
	sum1 := 0
	sum2 := 0
	for _, val := range arr {
		nums := []num{}
		for i, c := range val {
			if d, err := strconv.Atoi(string(c)); err == nil {
				nums = append(nums, num{value: d, index: i})
			}
		}

		sort.Sort(ByIdx(nums))
		sum1 += 10*nums[0].value + nums[len(nums)-1].value

		for i, s := range NumberStrings {
			loc1 := strings.Index(val, s)
			if loc1 != -1 {
				nums = append(nums, num{value: i, index: loc1})
			}
			loc2 := strings.LastIndex(val, s)
			if loc2 != -1 && loc2 != loc1 {
				nums = append(nums, num{value: i, index: loc2})
			}
		}

		sort.Sort(ByIdx(nums))
		sum2 += 10*nums[0].value + nums[len(nums)-1].value
	}

	fmt.Printf("Part 1: %d\n", sum1)
	fmt.Printf("Part 2: %d\n", sum2)
}
