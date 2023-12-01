package main

import (
	"strconv"
)

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
