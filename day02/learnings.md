# Learnings from Advent of Code 2020 Day 2 Go Solution Review

This document summarizes potential improvements and common Go idioms identified during the review of the Advent of Code 2020 Day 2 solution.

## 1. Parsing and Data Structure

**Observation:**
The original solution duplicated parsing logic in both `part1` and `part2`.

**Learning/Improvement:**
It is good practice to separate the concern of parsing input from the core logic of solving the puzzle. This makes the code cleaner, more modular, and reduces repetition. Define `struct` types to represent the parsed input data (e.g., `policy` and `passwordEntry`) and implement a single parsing function (`parseInput`) that returns a slice of these structs. The puzzle-solving functions (`part1`, `part2`) then operate directly on this structured data.

**Example `struct` definitions:**

```go
type policy struct {
    char rune
    min  int // used for 'atLeast' in part 1, 'pos1' in part 2
    max  int // used for 'atMost' in part 1, 'pos2' in part 2
}

type passwordEntry struct {
    policy   policy
    password string
}
```

**Example `parseInput` function:**

```go
// parseInput reads the specified file and converts each line into a passwordEntry struct.
func parseInput(path string) ([]passwordEntry, error) {
    // Assumes a utility function that reads a file into a slice of strings.
    lines, err := utils.ReadLines(path)
    if err != nil {
        return nil, fmt.Errorf("could not read file: %w", err)
    }

    // `lineRegex` should be a package-level variable using `regexp.MustCompile`
    // to avoid recompilation and error handling at runtime for a fixed pattern.
    // var lineRegex = regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)
    
    entries := make([]passwordEntry, 0, len(lines))

    for _, line := range lines {
        match := lineRegex.FindStringSubmatch(line)
        if match == nil {
            return nil, fmt.Errorf("line did not match expected format: %q", line)
        }

        // match[1] and match[2] are the numbers, match[3] is the char, match[4] is the password.
        min, err := strconv.Atoi(match[1])
        if err != nil {
            return nil, fmt.Errorf("could not parse min value from line: %q", line)
        }

        max, err := strconv.Atoi(match[2])
        if err != nil {
            return nil, fmt.Errorf("could not parse max value from line: %q", line)
        }

        entry := passwordEntry{
            policy: policy{
                char: []rune(match[3])[0], // Use a rune for the character
                min:  min,
                max:  max,
            },
            password: match[4],
        }
        entries = append(entries, entry)
    }

    return entries, nil
}
```

## 2. Regular Expression Handling

**Observation:**
The `regexp.Compile` function was called repeatedly within both `part1` and `part2` functions.

**Learning/Improvement:**
For regular expressions that are constant and known at compile-time, it is more idiomatic and efficient to use `regexp.MustCompile` at the package level. This ensures the regex is compiled only once when the program starts, avoiding redundant compilation and runtime error checking for the regex pattern itself.

**Example:**

```go
// At the top of your file (package scope)
var lineRegex = regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)

// You can then use `lineRegex` directly within your functions.
```

## 3. Part 2 Validation Logic

**Observation:**
The logic for Part 2 correctly implemented the "exactly one of two conditions must be true" rule using two `if` statements.

**Learning/Improvement:**
The boolean XOR (`^`) operator provides a more concise and often more readable way to express the condition "exactly one of two things must be true."

Additionally, comparing `byte` values directly is more efficient than converting single bytes to strings for comparison. When indexing a string (e.g., `word[i-1]`), you get a `byte`. If the policy character (`letter`) is also represented as a single `byte` or `rune` (e.g., `letter[0]`), direct byte comparison is preferred.

**Example:**

```go
// Assuming 'letter' is a string like "a", then 'letter[0]' is its byte representation.
// 'word' is a string, and 'word[i-1]' is the byte at that index.
pos1Match := word[i-1] == letter[0]
pos2Match := word[j-1] == letter[0]

if pos1Match ^ pos2Match { // The ^ operator performs a boolean XOR
    valid++
}
```
