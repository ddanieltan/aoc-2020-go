# Day 01 - Report Repair - Code Review Learnings

## Original Solution Overview

The provided solution correctly solves both Part 1 and Part 2 of Advent of Code 2020 Day 1.

### `checkComplement` function:
This function efficiently finds two numbers in a given list that sum up to a target using a hash map. It iterates through the list twice: once to populate a map with complements, and a second time to find a match. This results in an O(N) time complexity.

### `part1` function:
Reads integers from a file, then calls `checkComplement` to find two numbers that sum to 2020, returning their product.

### `part2` function:
Reads integers from a file. It then iterates through each number `n`. For each `n`, it creates a `sublist` by appending slices before and after `n` (excluding `n`). It then calls `checkComplement` on this `sublist` to find two numbers that sum to `2020 - n`, returning the product of the three numbers.

### `main` function:
Calls `part1` and `part2` with "input.txt" and prints the results.

### `main_test.go`:
Contains basic tests for `part1` and `part2` using "example.txt".

## Suggested Improvements

While the solution is correct, there are opportunities to improve efficiency and make the Go code more idiomatic.

### 1. More Efficient `checkComplement` (Single Pass)

The original `checkComplement` function iterates through the list of numbers twice. This can be optimized to a single pass for better performance.

**Original (two-pass):**
```go
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
```

**Improved (single-pass):**
This version iterates through the numbers once. For each number, it checks if the required complement has been seen before. If it has, the pair is found. If not, it adds the current number to a `seen` map.

```go
func checkComplement(nums []int, target int) (int, int, error) {
	seen := make(map[int]bool)
	for _, n := range nums {
		complement := target - n
		if seen[complement] {
			return n, complement, nil
		}
		seen[n] = true
	}
	return 0, 0, fmt.Errorf("No complement found")
}
```

### 2. More Efficient `part2` (Avoid Costly Slice Operations)

In the original `part2` function, the line `sublist := append(nums[:i], nums[i+1:]...)` is executed in every iteration of the outer loop. Creating a new slice by appending other slices is a computationally expensive operation as it involves memory allocation and copying data.

This can be avoided by making `checkComplement` calls on a sub-slice reference or by adapting the logic.

**Original (inefficient slice creation):**
```go
func part2(path string) int {
	nums, err := utils.ReadInts(path)
	if err != nil {
		return 0
	}

	for i, n := range nums {
		sublist := append(nums[:i], nums[i+1:]...) // This line is inefficient
		target := 2020 - n
		n1, n2, err := checkComplement(sublist, target)
		if err != nil {
			continue
		}
		return n1 * n2 * n

	}
	return 0
}
```

**Improved (avoiding costly slice creation):**
This version avoids the costly `append` operation by passing a slice reference to the numbers *after* the current number. This maintains the O(N^2) complexity but significantly reduces the constant factor related to memory operations.

```go
func part2(path string) int {
	nums, err := utils.ReadInts(path)
	if err != nil {
		return 0
	}

	for i, n1 := range nums {
		target := 2020 - n1
		// Find two numbers in the rest of the list that sum to target
		sublist := nums[i+1:] // Efficiently creates a sub-slice view
		n2, n3, err := checkComplement(sublist, target)
		if err == nil {
			return n1 * n2 * n3
		}
	}
	return 0
}
```