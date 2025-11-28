package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	result := part1("example.txt")
	expected := 820

	if result != expected {
		t.Errorf("Part 1 Example: result=%d, expected=%d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	result := part2("example.txt")
	expected := 0

	if result != expected {
		t.Errorf("Part 2 Example: result=%d, expected=%d", result, expected)
	}
}
