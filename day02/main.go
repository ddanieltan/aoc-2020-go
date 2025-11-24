package main

import (
	"aoc-2020-go/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func part1(path string) int {
	lines, err := utils.ReadLines(path)
	if err != nil {
		return 0
	}

	// 1-3 a: abcde
	r, err := regexp.Compile(`(\d+)-(\d+) (\w): (\w+)`)
	if err != nil {
		return 0
	}

	var valid int
	for _, line := range lines {
		match := r.FindStringSubmatch(line)
		atLeast, _ := strconv.Atoi(match[1])
		atMost, _ := strconv.Atoi(match[2])
		letter := match[3]
		word := match[4]
		cnt := strings.Count(word, letter)

		if (cnt >= atLeast) && (cnt <= atMost) {
			valid++
		}
	}

	return valid

}

func part2(path string) int {
	lines, err := utils.ReadLines(path)
	if err != nil {
		return 0
	}

	// 1-3 a: abcde
	r, err := regexp.Compile(`(\d+)-(\d+) (\w): (\w+)`)
	if err != nil {
		return 0
	}

	var valid int
	for _, line := range lines {
		match := r.FindStringSubmatch(line)
		i, _ := strconv.Atoi(match[1])
		j, _ := strconv.Atoi(match[2])
		letter := match[3]
		word := match[4]
		if (string(word[i-1]) == letter) && (string(word[j-1]) == letter) {
			continue
		}
		if (string(word[i-1]) == letter) || (string(word[j-1]) == letter) {
			valid++
		}
	}

	return valid

}

func main() {
	part1 := part1("input.txt")
	part2 := part2("input.txt")
	fmt.Println(part1)
	fmt.Println(part2)
}
