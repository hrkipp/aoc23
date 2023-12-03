package day1

import (
	_ "embed"
	"strings"
)

//go:embed input
var input string

func Part1() int {

	sum := 0

	options := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		first := FirstNum(line, options)
		last := LastNum(line, options)

		sum += first*10 + last

	}
	return sum

}

func Part2() int {

	sum := 0

	options := map[string]int{
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		first := FirstNum(line, options)
		last := LastNum(line, options)

		sum += first*10 + last

	}
	return sum

}

func FirstNum(str string, options map[string]int) int {
	for i := 0; i < len(str); i++ {
		for opt, val := range options {
			if strings.HasPrefix(str[i:], opt) {
				return val
			}
		}
	}
	return 0
}

func LastNum(str string, options map[string]int) int {
	for i := 0; i < len(str); i++ {
		for opt, val := range options {
			if strings.HasSuffix(str[:len(str)-i], opt) {
				return val
			}
		}
	}
	return 0
}
