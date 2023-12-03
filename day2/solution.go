package day2

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var input string

func Part1() (int, error) {

	max := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	sum := 0

gameLoop:
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		id, games, _ := strings.Cut(line, ": ")
		for _, game := range strings.Split(games, "; ") {
			for _, pair := range strings.Split(game, ", ") {
				number, color, _ := strings.Cut(pair, " ")

				num, err := strconv.Atoi(number)
				if err != nil {
					return 0, fmt.Errorf("number: %v", err)
				}

				if max[color] < num {
					idNum, err := strconv.Atoi(id[5:])
					if err != nil {
						return 0, fmt.Errorf("id: %v", err)
					}
					sum += idNum
					continue gameLoop
				}
			}
		}
	}

	return 5050 - sum, nil
}

func Part2() (int, error) {

	sum := 0

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		gameMaxes := map[string]int{}

		_, games, _ := strings.Cut(line, ": ")
		for _, game := range strings.Split(games, "; ") {
			for _, pair := range strings.Split(game, ", ") {
				number, color, _ := strings.Cut(pair, " ")

				num, err := strconv.Atoi(number)
				if err != nil {
					return 0, fmt.Errorf("number: %v", err)
				}

				if gameMaxes[color] < num {
					gameMaxes[color] = num
				}
			}
		}

		power := 1
		for _, val := range gameMaxes {
			power *= val
		}
		sum += power

	}

	return sum, nil
}
