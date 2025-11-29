# Learnings from day06/main.go

Here are some suggestions for improving the Go solution for Day 6 of Advent of Code 2020.

## 1. Reduce Code Duplication

Both `part1` and `part2` functions share a significant amount of boilerplate code: reading the input file, iterating through groups, and counting character frequencies within each group. This duplicated logic can be extracted into a shared helper function.

This improves maintainability. If the input processing logic needed to change, you would only have to update it in one place.

For example, you could create a function `parseGroups` that processes the input and returns a data structure that both `part1` and `part2` can use.

### Original Code Snippet (Duplication)

```go
// In part1
lines, err := utils.ReadDoubleLines(path)
if err != nil {
    return 0
}
// ...
for _, group := range lines {
    set := make(map[rune]int)
    for _, person := range strings.Split(group, "\n") {
        for _, rune := range person {
            set[rune] += 1
        }
    }
    // ...
}

// In part2
lines, err := utils.ReadDoubleLines(path)
if err != nil {
    return 0
}
// ...
for _, group := range lines {
    set := make(map[rune]int)
    people := strings.Split(group, "\n")
    for _, person := range people {
        for _, rune := range person {
            set[rune] += 1
        }
    }
    // ...
}
```

### Refactored Approach

A helper function could return the processed data for each group. A good data structure for this would be a struct that holds the frequency map and the number of people in the group.

```go
type GroupAnswers struct {
    QuestionCounts map[rune]int
    NumPeople      int
}

func parseGroups(path string) ([]GroupAnswers, error) {
    groups, err := utils.ReadDoubleLines(path)
    if err != nil {
        return nil, err
    }

    var result []GroupAnswers
    for _, group := range groups {
        questionCounts := make(map[rune]int)
        people := strings.Split(group, "\n")
        for _, person := range people {
            for _, r := range person {
                questionCounts[r]++
            }
        }
        result = append(result, GroupAnswers{
            QuestionCounts: questionCounts,
            NumPeople:      len(people),
        })
    }
    return result, nil
}

func part1(path string) int {
    groups, err := parseGroups(path)
    if err != nil {
        return 0
    }

    var totalCount int
    for _, group := range groups {
        totalCount += len(group.QuestionCounts)
    }
    return totalCount
}

func part2(path string) int {
    groups, err := parseGroups(path)
    if err != nil {
        return 0
    }

    var totalCount int
    for _, group := range groups {
        var groupCount int
        for _, count := range group.QuestionCounts {
            if count == group.NumPeople {
                groupCount++
            }
        }
        totalCount += groupCount
    }
    return totalCount
}
```

## 2. Use More Idiomatic Data Structures for Sets

In `part1`, you are interested in the number of unique questions anyone answered "yes" to in a group. The `map[rune]int` works, but its purpose is to store a *set* of unique `rune`s.

In Go, a more idiomatic and memory-efficient way to represent a set is `map[T]struct{}` where `T` is the type of the elements in the set. The `struct{}` type has a width of zero, so it doesn't consume extra memory for the map values.

### Original `part1` Map

```go
set := make(map[rune]int)
for _, person := range strings.Split(group, "\n") {
    for _, rune := range person {
        set[rune] += 1
    }
}
count += len(set)
```

### Improved Version with a "Set"

```go
// Inside the loop for part1
questionSet := make(map[rune]struct{})
for _, person := range strings.Split(group, "\n") {
    for _, r := range person {
        questionSet[r] = struct{}{}
    }
}
count += len(questionSet)
```
This change makes the intent of the code clearer: you are collecting a unique set of questions, not counting their frequencies.

## 3. Use Descriptive Variable Names

The variable `set` in `part2` is used as a frequency map to count how many people in a group answered "yes" to a specific question. Naming it `set` can be misleading, as "set" usually implies a collection of unique items with no associated values.

A more descriptive name like `questionCounts` or `freqMap` would make the code's purpose easier to understand at a glance.

### Original Code

```go
func part2(path string) int {
    // ...
    for _, group := range lines {
		set := make(map[rune]int)
        // ...
		for _, person := range people {
			for _, rune := range person {
				set[rune] += 1
			}
		}

		var allYes int
		for _, yes := range set {
            // ...
		}
		// ...
	}
    // ...
}
```

### With Clearer Variable Names

```go
func part2(path string) int {
    // ...
    for _, group := range lines {
		questionCounts := make(map[rune]int)
        // ...
		for _, person := range people {
			for _, r := range person {
				questionCounts[r]++
			}
		}

		var unanimousYesCount int
		for _, count := range questionCounts {
			if count == len(people) {
				unanimousYesCount++
			}
		}
		totalCount += unanimousYesCount
	}
    // ...
}
```
Using clearer names like `questionCounts`, `r` (for rune), and `unanimousYesCount` improves readability and makes the code self-documenting.

By applying these suggestions, the code becomes more readable, maintainable, and idiomatic to the Go language.
