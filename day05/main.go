package main

import (
	"aoc-2020-go/utils"
	"fmt"
	"sort"
)

func parseSeat(letters string) (id int) {
	minRow := 0
	maxRow := 127
	minCol := 0
	maxCol := 7
	for _, letter := range letters {
		switch letter {
		case 'F':
			maxRow -= (maxRow - minRow + 1) / 2
		case 'B':
			minRow += (maxRow - minRow + 1) / 2
		case 'L':
			maxCol -= (maxCol - minCol + 1) / 2
		case 'R':
			minCol += (maxCol - minCol + 1) / 2
		}
	}
	id = minRow*8 + minCol
	return id
}

func part1(path string) int {
	lines, err := utils.ReadLines(path)
	if err != nil {
		return 0
	}

	var highest int
	for _, line := range lines {
		id := parseSeat(line)
		if id > highest {
			highest = id
		}
	}

	return highest
}

func part2(path string) int {
	lines, err := utils.ReadLines(path)
	if err != nil {
		return 0
	}
	var seats []int
	for _, line := range lines {
		id := parseSeat(line)
		seats = append(seats, id)
	}

	sort.Ints(seats)

	seat := seats[0]

	for _, next := range seats[1:] {
		if seat+1 != next {
			return seat + 1
		} else {
			seat += 1
		}
	}

	return seat
}

func main() {
	part1 := part1("input.txt")
	part2 := part2("input.txt")
	fmt.Println(part1)
	fmt.Println(part2)
}
