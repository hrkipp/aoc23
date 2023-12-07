package day5

import (
	_ "embed"
	"fmt"
	"math"
	"strings"

	"github.com/hrkipp/aoc23/util"
)

//go:embed input
var input string

func Part1() int {

	type mapRange struct {
		dStart int
		sStart int
		length int
	}

	maps := []func(int) int{}
	lines := strings.Split(strings.TrimSpace(input), "\n")

	seeds := util.SplitInts(lines[0][7:], " ")
	fmt.Println(seeds)

	var currentLine []mapRange
	for _, line := range lines[2:] {
		switch {
		case line == "":
			lineSnapshot := currentLine
			maps = append(maps, func(i int) int {
				for _, line := range lineSnapshot {
					if line.sStart < i && line.sStart+line.length >= i {
						return line.dStart + (i - line.sStart)
					}
				}
				return i
			})
		case strings.HasSuffix(line, "map:"):
			currentLine = nil
		default:
			parts := util.SplitInts(line, " ")
			currentLine = append(currentLine, mapRange{
				parts[0],
				parts[1],
				parts[2],
			})
		}

	}

	min := math.MaxInt64
	for _, seed := range seeds {
		s := seed
		for _, m := range maps {
			s = m(s)
		}
		if s < min {
			min = s
		}
	}
	return min
}

func Part2() int {

	type mapRange struct {
		dStart int
		sStart int
		length int
	}

	maps := []func(int) int{}
	lines := strings.Split(strings.TrimSpace(input), "\n")

	seeds := util.SplitInts(lines[0][7:], " ")

	var currentLine []mapRange
	for _, line := range lines[2:] {
		switch {
		case line == "":
			lineSnapshot := currentLine
			maps = append(maps, func(i int) int {
				for _, line := range lineSnapshot {
					if line.sStart <= i && line.sStart+line.length > i {
						return line.dStart + (i - line.sStart)
					}
				}
				return i
			})
		case strings.HasSuffix(line, "map:"):
			currentLine = nil
		default:
			parts := util.SplitInts(line, " ")
			currentLine = append(currentLine, mapRange{
				parts[0],
				parts[1],
				parts[2],
			})
		}

	}

	min := math.MaxInt64
	for i := 0; i < len(seeds); i += 2 {
		for s := seeds[i]; s < seeds[i]+seeds[i+1]; s++ {
			start := s
			for _, m := range maps {
				start = m(start)
			}
			if start < min {
				min = start
			}
		}
	}
	return min
}
