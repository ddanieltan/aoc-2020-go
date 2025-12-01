package main

import (
	"aoc-2020-go/utils"
	"fmt"
)

func part1(path string) int {
	lines, err := utils.ReadLines(path)
	if err != nil {
		return 0
	}

	processed := make(map[int]bool)

	for {
		:
	}


	for _, line := range lines {
		fmt.Println(line)
	}
	return 0

}

func part2(path string) int {
	lines, err := utils.ReadLines(path)
	if err != nil {
		return 0
	}

	for _, line := range lines {
		fmt.Println(line)
	}
	return 0

}

func main() {
	part1 := part1("input.txt")
	part2 := part2("input.txt")
	fmt.Println(part1)
	fmt.Println(part2)
}
