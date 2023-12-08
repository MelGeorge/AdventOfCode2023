package aoc

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type HandType int

const (
	fiveOAK   HandType = 0
	fourOAK   HandType = 1
	fullHouse HandType = 2
	threeOAK  HandType = 3
	twoPair   HandType = 4
	onePair   HandType = 5
	highCard  HandType = 6
)

var order = map[byte]int{'A': 0, 'K': 1, 'Q': 2, 'J': 3, 'T': 4, '9': 5, '8': 6, '7': 7, '6': 8, '5': 9, '4': 10, '3': 11, '2': 12}

var jokerOrder = map[byte]int{'A': 0, 'K': 1, 'Q': 2, 'T': 3, '9': 4, '8': 5, '7': 6, '6': 7, '5': 8, '4': 9, '3': 10, '2': 11, 'J': 12}

type hand struct {
	c string
	b int
}

func Day7(in string) {
	lines := strings.Split(in, "\n")

	var hands []hand
	for _, line := range lines {
		f := strings.Fields(line)
		bid, _ := strconv.Atoi(f[1])
		hands = append(hands, hand{c: f[0], b: bid})
	}

	withJokers := []bool{false, true}
	for part, jokers := range withJokers {
		sort.Slice(hands, func(i, j int) bool {
			it := handType(hands[i], jokers)
			jt := handType(hands[j], jokers)

			if it == jt {
				for n := range hands[i].c {
					if hands[i].c[n] != hands[j].c[n] {
						o := order
						if jokers {
							o = jokerOrder
						}
						return o[hands[i].c[n]] > o[hands[j].c[n]]
					}
				}
			}

			return handType(hands[i], jokers) > handType(hands[j], jokers)
		})

		winnings := 0
		for i := range hands {
			winnings += hands[i].b * (i + 1)
		}

		fmt.Println("Part", part, ": ", winnings)
	}
}

func handType(h hand, jokers bool) HandType {
	cards := map[byte]int{}
	var hiFreq byte
	freq := 0
	for i := 0; i < 5; i++ {
		cards[h.c[i]]++
		if cards[h.c[i]] > freq && h.c[i] != 'J' {
			freq = cards[h.c[i]]
			hiFreq = h.c[i]
		}
	}

	if jokers && cards['J'] > 0 {
		cards[hiFreq] += cards['J']
		cards['J'] = 0
	}

	fives, quads, trips, pairs := 0, 0, 0, 0
	for _, count := range cards {
		switch count {
		case 5:
			fives++
		case 4:
			quads++
		case 3:
			trips++
		case 2:
			pairs++
		}
	}

	if fives > 0 {
		return fiveOAK
	} else if quads > 0 {
		return fourOAK
	} else if trips > 0 && pairs > 0 {
		return fullHouse
	} else if trips > 0 {
		return threeOAK
	} else if pairs > 1 {
		return twoPair
	} else if pairs > 0 {
		return onePair
	} else {
		return highCard
	}
}
