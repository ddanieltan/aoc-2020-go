package main

import (
	"aoc-2020-go/utils"
	"fmt"
	"strings"
)

func part1(path string) int {
	lines, err := utils.ReadDoubleLines(path)
	if err != nil {
		return 0
	}

	var count int
	for _, group := range lines {
		set := make(map[rune]int)
		for _, person := range strings.Split(group, "\n") {
			for _, rune := range person {
				set[rune] += 1
			}
		}
		count += len(set)
	}
	return count

}

func part2(path string) int {
	lines, err := utils.ReadDoubleLines(path)
	if err != nil {
		return 0
	}

	var count int
	for _, group := range lines {
		set := make(map[rune]int)
		people := strings.Split(group, "\n")
		for _, person := range people {
			for _, rune := range person {
				set[rune] += 1
			}
		}

		var allYes int
		for _, yes := range set {
			if yes == len(people) {
				allYes += 1
			}
		}

		count += allYes
	}
	return count

}

func main() {
	part1 := part1("input.txt")
	part2 := part2("input.txt")
	fmt.Println(part1)
	fmt.Println(part2)
}
