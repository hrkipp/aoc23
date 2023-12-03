package main

import (
	"fmt"
	"log"

	"github.com/hrkipp/aoc23/day1"
	"github.com/hrkipp/aoc23/day2"
	"github.com/hrkipp/aoc23/day3"
)

func main() {
	sol, err := day1.Part1()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Day 1 Part 1: ", sol)

	sol, err = day1.Part2()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Day 1 Part 2: ", sol)

	sol, err = day2.Part1()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Day 2 Part 1: ", sol)

	sol, err = day2.Part2()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Day 2 Part 2: ", sol)

	fmt.Println("Day 3 Part 1: ", day3.Part1())
	fmt.Println("Day 3 Part 2: ", day3.Part2())
}
