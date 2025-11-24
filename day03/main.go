package main

import (
	"aoc-2020-go/utils"
	"fmt"
)

func countTrees(lines []string, right int, down int) int {
	x := 0
	trees := 0
	for i, line := range lines {
		// Skip the first line
		if i == 0 {
			continue
		}
		if i%down == 0 {
			x += right
			x = x % len(line)
			if string(line[x]) == "#" {
				trees += 1
			}
		}
	}

	return trees
}

func part1(path string) int {
	lines, err := utils.ReadLines(path)
	if err != nil {
		return 0
	}
	trees := countTrees(lines, 3, 1)

	return trees
}

type Slope struct {
	right int
	down  int
}

func part2(path string) int {
	lines, err := utils.ReadLines(path)
	if err != nil {
		return 0
	}

	slopes := []Slope{
		{right: 1, down: 1},
		{right: 3, down: 1},
		{right: 5, down: 1},
		{right: 7, down: 1},
		{right: 1, down: 2},
	}

	var product int
	for i, slope := range slopes {
		if i == 0 {
			product = countTrees(lines, slope.right, slope.down)
		} else {
			product *= countTrees(lines, slope.right, slope.down)
		}
	}

	return product
}

func main() {
	part1 := part1("input.txt")
	part2 := part2("input.txt")
	fmt.Println(part1)
	fmt.Println(part2)
}
