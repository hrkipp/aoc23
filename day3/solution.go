package day3

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input
var input string

func init() {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for lineNum := range lines {
		lines[lineNum] = "." + lines[lineNum] + "."
	}
	paddingLine := strings.Repeat(".", len(lines[0]))
	lines = append([]string{paddingLine}, lines...)
	lines = append(lines, paddingLine)
	input = strings.Join(lines, "\n")
}

func Part1() int {

	sum := 0
	lines := strings.Split(input, "\n")

	for lineNum, line := range lines {
		matches := regexp.MustCompile("[0-9]+").FindAllStringIndex(line, -1)

		for _, match := range matches {
			val, err := strconv.Atoi(line[match[0]:match[1]])
			if err != nil {
				panic(err)
			}
			set := lines[lineNum-1][match[0]-1:match[1]+1] +
				string(line[match[0]-1]) +
				string(line[match[1]]) +
				lines[lineNum+1][match[0]-1:match[1]+1]
			cleaned := strings.ReplaceAll(set, ".", "")
			if len(cleaned) != 0 {
				sum += val
			}
		}
	}

	return sum
}

func isNum(c byte) bool {
	return c > 47 && c < 58
}

func adjacentNumbers(line string, index int) []int {
	var out []int

	if isNum(line[index]) {
		/*
					......456.......
					.....345........
					.......567......
			           ^
		*/
		if isNum(line[index-1]) {
			if isNum(line[index+1]) {
				num, err := strconv.Atoi(line[index-1 : index+2])
				if err != nil {
					panic(err)
				}
				out = append(out, num)
			} else {
				match := regexp.MustCompile("[0-9]+$").FindString(line[:index+1])
				num, err := strconv.Atoi(match)
				if err != nil {
					panic(err)
				}
				out = append(out, num)
			}
		} else {
			match := regexp.MustCompile("^[0-9]+").FindString(line[index:])
			num, err := strconv.Atoi(match)
			if err != nil {
				panic(err)
			}
			out = append(out, num)
		}

	} else {
		match := regexp.MustCompile("[0-9]+$").FindString(line[:index])
		if match != "" {
			num, err := strconv.Atoi(match)
			if err != nil {
				panic(err)
			}
			out = append(out, num)
		}
		match = regexp.MustCompile("^[0-9]+").FindString(line[index+1:])
		if match != "" {
			num, err := strconv.Atoi(match)
			if err != nil {
				panic(err)
			}
			out = append(out, num)
		}

	}
	return out
}

func Part2() int {

	sum := 0
	lines := strings.Split(input, "\n")
	for lineNum, line := range lines {
		for charNum, char := range line {
			if char != '*' {
				continue
			}

			numbers := []int{}

			numbers = append(numbers, adjacentNumbers(lines[lineNum-1], charNum)...)
			match := regexp.MustCompile("[0-9]+$").FindString(line[:charNum])
			if match != "" {
				num, err := strconv.Atoi(match)
				if err != nil {
					panic(err)
				}
				numbers = append(numbers, num)
			}
			match = regexp.MustCompile("^[0-9]+").FindString(line[charNum+1:])
			if match != "" {
				num, err := strconv.Atoi(match)
				if err != nil {
					panic(err)
				}
				numbers = append(numbers, num)
			}

			numbers = append(numbers, adjacentNumbers(lines[lineNum+1], charNum)...)

			if len(numbers) == 2 {
				sum += numbers[0] * numbers[1]
			} else {
				fmt.Println(lines[lineNum-1][charNum-4 : charNum+4])
				fmt.Println(line[charNum-4 : charNum+4])
				fmt.Println(lines[lineNum+1][charNum-4 : charNum+4])
				fmt.Println(numbers)
				fmt.Println()

			}
		}
	}

	return sum
}
