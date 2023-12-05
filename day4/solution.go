package day4

import (
	_ "embed"
	"strings"
)

//go:embed input
var input string

func Part1() int {

	sum := 0

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		winning := strings.Split(line[10:40], " ")
		have := strings.Split(line[42:], " ")

		winningMap := map[string]bool{}
		for _, win := range winning {
			if win == "" {
				continue
			}
			winningMap[win] = true
		}
		partial := 0

		for _, h := range have {
			if h == "" {
				continue
			}
			if winningMap[h] {
				if partial == 0 {
					partial = 1
				} else {
					partial *= 2
				}
			}
		}

		sum += partial
	}
	return sum

}

func Part2() int {
	sum := 0

	points := map[int]int{}

	for i, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		winning := strings.Split(line[10:40], " ")
		have := strings.Split(line[42:], " ")

		winningMap := map[string]bool{}
		for _, win := range winning {
			if win == "" {
				continue
			}
			winningMap[win] = true
		}
		partial := 0

		for _, h := range have {
			if h == "" {
				continue
			}
			if winningMap[h] {
				partial += 1
			}
		}

		points[i] = partial

	}

	var dig func(val int)
	dig = func(val int) {
		sum++
		for i := 1; i <= points[val]; i++ {
			dig(val + i)
		}
	}

	for index := range points {
		dig(index)
	}

	return sum

}
