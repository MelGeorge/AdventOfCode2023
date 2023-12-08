package aoc

import (
	"strconv"
	"strings"
)

func isNumber(r rune) bool {
	_, err := strconv.Atoi(string(r))
	return err == nil
}

func sum(arr []int) int {
	s := 0
	for _, v := range arr {
		s += v
	}
	return s
}

func tointarray(s []string) []int {
	i := []int{}
	for _, val := range s {
		n, _ := strconv.Atoi(val)
		i = append(i, n)
	}

	return i
}

/*
	arr := strings.Split(in, "\n")
	e := [][]int{}
	e = append(e, []int{})
	i := 0
	for j, val := range arr {
		if val != "" {
			n, _ := strconv.Atoi(val)
			e[i] = append(e[i], n)
		}

		if val == "" && j < len(arr)-1 {
			i++
			e = append(e, []int{})
		}
	}
*/

func separate(file string, sep string) func() string {
	offset := 0

	return func() string {
		if offset+len(sep) < len(file) {
			i := strings.Index(file[offset:], sep)
			if i == -1 {
				return ""
			}
			temp := file[offset : offset+i]
			offset += i + len(sep)
			return temp
		}
		return ""
	}
}

/*
iter := separate(in, "\n")
for i, row := iter(), 0; len(i) != 0; row++ {
	fmt.Println(i)
	i = iter()
}
*/
