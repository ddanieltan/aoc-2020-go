# Learnings from Advent of Code 2020 Day 3 Go Solution Review

This document provides a critique and suggestions for improving the Go solution for Advent of Code 2020, Day 3. The goal is to highlight idiomatic Go practices and opportunities for cleaner, more efficient code.

## 1. `countTrees` Function Analysis

### Looping Logic
**Observation:**
The function uses a `for...range` loop over all lines and includes checks to skip the first line (`if i == 0`) and process lines based on the `down` slope (`if i % down == 0`).

```go
// Original loop structure
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
```

**Suggestion:**
A standard `for` loop that directly implements the traversal logic can be cleaner and more efficient by avoiding conditional checks on every iteration. This makes the intent of the loop clearer at a glance.

**Improved `for` loop:**
```go
func countTrees(lines []string, right int, down int) int {
	x := 0
	trees := 0
	// Start at the first step 'down' and increment by 'down' each time.
	for i := down; i < len(lines); i += down {
		x += right
		x %= len(lines[i]) // Use len of the current line
		if lines[i][x] == '#' {
			trees++
		}
	}
	return trees
}
```
This revised loop directly models the movement down the slope, eliminating the need for the `if i == 0` and `if i % down == 0` checks.

### Character Comparison
**Observation:**
A character from the line is compared to `"#"` by first converting it to a `string`.
`if string(line[x]) == "#"`

**Suggestion:**
Indexing a string in Go (e.g., `line[x]`) accesses the `byte` at that position. It is more efficient to compare bytes directly, as this avoids the overhead of allocating a new string for the comparison.

**Improved comparison:**
```go
// Compare bytes directly. '#' is a byte literal (type rune).
if line[x] == '#' {
    trees++
}
```

## 2. `part2` Function Analysis

### Product Calculation
**Observation:**
The calculation of the final product in `part2` uses a conditional to handle the first element differently.

```go
// Original product calculation
var product int
for i, slope := range slopes {
    if i == 0 {
        product = countTrees(lines, slope.right, slope.down)
    } else {
        product *= countTrees(lines, slope.right, slope.down)
    }
}
```

**Suggestion:**
You can simplify this by initializing the `product` variable to `1` before the loop. This removes the need for the conditional `if i == 0` check, resulting in cleaner code.

**Improved product calculation:**
```go
// Initialize product to 1
product := 1
for _, slope := range slopes {
    product *= countTrees(lines, slope.right, slope.down)
}
```
