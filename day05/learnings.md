# Learnings from Day 05

This solution is straightforward and correct, but there are a few opportunities to write more idiomatic and efficient Go code.

## 1. Simplify `parseSeat` with Binary Conversion

The current `parseSeat` function correctly implements a binary search algorithm. However, the boarding pass encoding (`FBFBBFFRLR`) is a direct binary representation of the seat ID.

- `F` (Front) and `L` (Left) can be treated as `0`.
- `B` (Back) and `R` (Right) can be treated as `1`.

So, a boarding pass like `BFFFBBFRRR` is equivalent to the binary number `1000110111`. The first 7 characters represent the row, and the last 3 represent the column. The final seat ID is `row * 8 + col`.

The entire 10-character string can be treated as a single 10-bit binary number that directly gives you the seat ID.

We can leverage this insight to create a much simpler and more efficient function using Go's standard library.

### Suggested Refactor

```go
import (
	"strconv"
	"strings"
)

// parseSeatBinary converts a boarding pass string into a seat ID.
// It returns the ID and an error if the string is invalid.
func parseSeatBinary(pass string) (int, error) {
	// Build a binary string by replacing letters with 0s and 1s.
	binaryStr := strings.NewReplacer(
		"F", "0",
		"B", "1",
		"L", "0",
		"R", "1",
	).Replace(pass)

	// Parse the binary string into an integer.
	id, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		// This handles cases where the input string is not valid.
		return 0, err
	}
	return int(id), nil
}
```

This version is more concise and declarative. It tells the story of *what* the transformation is (a binary conversion) rather than *how* to calculate it step-by-step.

## 2. Improve the Logic in `part2`

The implementation for `part2` correctly finds the missing seat, but has two minor issues:

1.  **A subtle bug:** If there are no gaps in the seat list, the function incorrectly returns the very last seat ID. A function should have a predictable return value for a "not found" case (e.g., `0`, `-1`, or an error).
2.  **Inefficient iteration:** The manual update `seat += 1` is slightly less clear than a standard indexed-based `for` loop for comparing adjacent elements in a slice.

### Correcting the Loop

A more standard way to iterate and find a gap in a sorted slice is:

```go
func part2(path string) int {
	// ... (reading lines and populating seats)
	sort.Ints(seats)

	for i := 0; i < len(seats)-1; i++ {
		// If the next seat ID is not one greater than the current one...
		if seats[i+1] != seats[i]+1 {
			// ...we've found our missing seat.
			return seats[i] + 1
		}
	}

	// Return 0 or -1 to indicate the seat was not found.
	return 0
}
```

### Alternative Algorithm: Using a Map

Sorting the entire list of seats (`O(n log n)`) isn't the most efficient approach if memory isn't a concern. A more performant way (`O(n)`) is to use a map to store the existing seat IDs.

This involves two passes:
1.  Iterate through all boarding passes, calculate their IDs, and add them to a map (`map[int]bool`). Also, keep track of the minimum and maximum IDs found.
2.  Iterate from `minID` to `maxID` and check for the first ID that is *not* in the map.

```go
func part2WithMap(path string) int {
	lines, err := utils.ReadLines(path)
	if err != nil {
		return 0
	}

	seatSet := make(map[int]bool)
	minID, maxID := 1023, 0 // Max possible ID is 127*8+7 = 1023

	for _, line := range lines {
		id := parseSeat(line) // Or the improved parseSeatBinary
		seatSet[id] = true
		if id < minID {
			minID = id
		}
		if id > maxID {
			maxID = id
		}
	}

	for id := minID; id <= maxID; id++ {
		if !seatSet[id] {
			return id
		}
	}

	return 0 // Not found
}
```

This approach is often faster for large inputs and demonstrates a common trade-off between CPU time and memory usage.
