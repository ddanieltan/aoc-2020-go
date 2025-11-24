# Learnings from Day 4

Your solution is a great start—it's clear, passes the tests, and correctly solves the puzzle. The following suggestions are aimed at refining your Go code to be more idiomatic, efficient, and maintainable, which are key goals in professional software development.

## 1. Parse First, Then Validate

A more robust and cleaner approach is to parse the raw passport strings into a more useful data structure *before* you start validating. A `map[string]string` is a perfect fit for a passport, mapping field keys (like `"byr"`) to their values.

This simplifies validation logic immensely. Instead of iterating through slices of `"key:value"` strings, you get direct, O(1) access to any field.

**Example: Refactored Passport Parsing**

```go
// In your part1/part2 functions, you could transform the input like this:
passportsData := []map[string]string{}
for _, passportStr := range strings.Split(input, "\n\n") {
    passport := make(map[string]string)
    fields := strings.Fields(passportStr)
    for _, field := range fields {
        parts := strings.SplitN(field, ":", 2)
        if len(parts) == 2 {
            passport[parts[0]] = parts[1]
        }
    }
    passportsData = append(passportsData, passport)
}

// Now you can process the structured data
var valid int
for _, passport := range passportsData {
    if hasRequiredFields(passport) && hasValidFields(passport) {
        valid++
    }
}
```

## 2. Reduce Redundancy with a Single Source of Truth

Both `hasRequiredFields` and `hasValidFields` define their own list of compulsory fields. This duplication can lead to bugs if one list is updated and the other is not.

Define required fields once as a package-level variable. A `map` is great for quick lookups.

```go
var requiredFields = map[string]struct{}{
    "byr": {}, "iyr": {}, "eyr": {}, "hgt": {}, "hcl": {}, "ecl": {}, "pid": {},
}

// The struct{}{} is a zero-memory value, making it a memory-efficient choice for set-like behavior.
```

Your `hasRequiredFields` function then becomes much simpler:

```go
// (Assumes passport is a map[string]string)
func hasRequiredFields(passport map[string]string) bool {
    for key := range requiredFields {
        if _, found := passport[key]; !found {
            return false
        }
    }
    return true
}
```

## 3. Improve Efficiency: Compile Regex Once

In your `hasValidFields` function, `regexp.MustCompile()` is called inside a loop. This is inefficient because the Go runtime has to re-compile the same regular expression for every single passport.

Regex compilation is an expensive operation. You should compile them once when your program starts and reuse the compiled object. A common pattern is to define them as package-level variables.

**Before:**
```go
// Inside a loop
r := regexp.MustCompile(`^#[0-9a-f]+`)
if r.MatchString(value) && len(value) == 7 {
    // ...
}
```

**After:**
```go
// At the top level of your package
var hclRegex = regexp.MustCompile(`^#[0-9a-f]{6}$`)
var pidRegex = regexp.MustCompile(`^[0-9]{9}$`)

// Inside your validation logic
if hclRegex.MatchString(value) {
    // ...
}
```

Note that the regex patterns have also been made more specific (`{6}` and `{9}`) to perform the length check at the same time, simplifying the code.

## 4. Enhance Readability with Validator Functions

The `hasValidFields` function is doing too much. Its large `switch` statement makes it hard to read and test. You can break down the validation logic for each field into its own dedicated function. This follows the Single Responsibility Principle.

This approach makes the code cleaner, easier to debug, and allows for unit testing of individual validation rules.

**Example: A `byr` validator**
```go
func validateBYR(value string) bool {
    year, err := strconv.Atoi(value)
    if err != nil {
        return false
    }
    return year >= 1920 && year <= 2002
}
```

You can then create a "dispatch table" using a map of functions to make the main validation logic elegant and extensible.

**Example: Using a Validator Map**
```go
var fieldValidators = map[string]func(string) bool{
    "byr": func(s string) bool { year, _ := strconv.Atoi(s); return year >= 1920 && year <= 2002 },
    "iyr": func(s string) bool { year, _ := strconv.Atoi(s); return year >= 2010 && year <= 2020 },
    "eyr": func(s string) bool { year, _ := strconv.Atoi(s); return year >= 2020 && year <= 2030 },
    "hgt": validateHGT, // Can still point to a more complex function
    "hcl": hclRegex.MatchString,
    "ecl": func(s string) bool {
        _, ok := map[string]struct{}{"amb":{},"blu":{},"brn":{},"gry":{},"grn":{},"hzl":{},"oth":{}}[s]
        return ok
    },
    "pid": pidRegex.MatchString,
    "cid": func(s string) bool { return true }, // 'cid' is always valid if present
}

// (Assumes passport is a map[string]string)
func hasValidFields(passport map[string]string) bool {
    for key, validator := range fieldValidators {
        value, ok := passport[key]
        // If a required field is missing, it's not valid.
        if !ok && key != "cid" {
            return false
        }
        // If the field is present, validate it.
        if ok && !validator(value) {
            return false
        }
    }
    return true
}
```

This revised `hasValidFields` is much cleaner. It checks for both presence and validity in one pass. With this, you can simplify your `part2` function to call only this one function.

By applying these patterns, your Go code will be more aligned with the practices used in production systems, making you a more effective Go programmer. Keep up the great work!