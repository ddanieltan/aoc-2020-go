package main

import (
	"aoc-2020-go/utils"
	"fmt"
)

// Given a list of nums and a target sum, find if any 2 nums can add up to that target
func checkComplement(nums []int, target int) (int, int, error) {
	complement := make(map[int]int)
	for _, n := range nums {
		complement[target-n] = n
	}
	for _, n := range nums {
		value, ok := complement[n]
		if ok {
			return value, n, nil
		}
	}
	return 0, 0, fmt.Errorf("No complement found")
}

func part1(path string) int {
	nums, err := utils.ReadInts(path)
	if err != nil {
		return 0
	}

	n1, n2, err := checkComplement(nums, 2020)
	if err != nil {
		return 0
	}

	return n1 * n2

}

func part2(path string) int {
	nums, err := utils.ReadInts(path)
	if err != nil {
		return 0
	}

	for i, n := range nums {
		sublist := append(nums[:i], nums[i+1:]...)
		target := 2020 - n
		n1, n2, err := checkComplement(sublist, target)
		if err != nil {
			continue
		}
		return n1 * n2 * n

	}
	return 0

}

func main() {
	part1 := part1("input.txt")
	part2 := part2("input.txt")
	fmt.Println(part1)
	fmt.Println(part2)
}
