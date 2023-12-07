package day6

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/hrkipp/aoc23/util"
)

//go:embed input
var input string

func Part1() int {

	lines := util.Lines(input)

	times := util.SplitInts(lines[0][9:], " ")
	distances := util.SplitInts(lines[1][9:], " ")
	product := 1

	for i := 0; i < len(times); i++ {
		sum := 0
		record := distances[i]

		for time := 0; time < times[i]; time++ {
			distance := (times[i] - time) * time
			if distance > record {
				sum++
			}
		}
		product *= sum
	}

	return product
}

func Part2() int {

	lines := util.Lines(input)

	times := util.SplitInts(strings.ReplaceAll(lines[0][9:], " ", ""), " ")
	distances := util.SplitInts(strings.ReplaceAll(lines[1][9:], " ", ""), " ")
	fmt.Println(times, distances)
	product := 1

	for i := 0; i < len(times); i++ {
		sum := 0
		record := distances[i]

		for time := 0; time < times[i]; time++ {
			distance := (times[i] - time) * time
			if distance > record {
				sum++
			}
		}
		product *= sum
	}

	return product
}
