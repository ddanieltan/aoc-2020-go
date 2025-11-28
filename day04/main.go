package main

import (
	"aoc-2020-go/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func hasRequiredFields(fields []string) bool {
	compulsory := map[string]int{
		"byr": 0,
		"iyr": 0,
		"eyr": 0,
		"hgt": 0,
		"hcl": 0,
		"ecl": 0,
		"pid": 0,
	}

	for _, field := range fields {
		code := string(field[:3])
		compulsory[code]++
	}

	for _, value := range compulsory {
		if value == 0 {
			return false
		}
	}

	return true
}

func hasValidFields(fields []string) bool {
	compulsory := map[string]bool{
		"byr": false,
		"iyr": false,
		"eyr": false,
		"hgt": false,
		"hcl": false,
		"ecl": false,
		"pid": false,
	}

	for _, field := range fields {
		split := strings.Split(field, ":")
		code := split[0]
		value := split[1]
		switch code {
		case "byr":
			digits, err := strconv.Atoi(value)
			if err == nil && (digits >= 1920) && (digits <= 2002) && (len(strconv.Itoa(digits)) == 4) {
				compulsory[code] = true
			}
		case "iyr":
			digits, err := strconv.Atoi(value)
			if err == nil && (digits >= 2010) && (digits <= 2020) && (len(strconv.Itoa(digits)) == 4) {
				compulsory[code] = true
			}
		case "eyr":
			digits, err := strconv.Atoi(value)
			if err == nil && (digits >= 2020) && (digits <= 2030) && (len(strconv.Itoa(digits)) == 4) {
				compulsory[code] = true
			}
		case "hgt":
			if strings.HasSuffix(value, "cm") {
				digits, err := strconv.Atoi(strings.TrimSuffix(value, "cm"))
				if err == nil && (digits >= 150) && (digits <= 193) {
					compulsory[code] = true
				}
			}
			if strings.HasSuffix(value, "in") {
				digits, err := strconv.Atoi(strings.TrimSuffix(value, "in"))
				if err == nil && (digits >= 59) && (digits <= 76) {
					compulsory[code] = true
				}
			}
		case "hcl":
			r := regexp.MustCompile(`^#[0-9a-f]+`)
			if r.MatchString(value) && len(value) == 7 {
				compulsory[code] = true
			}
		case "ecl":
			correct := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
			for _, eye := range correct {
				if value == eye {
					compulsory[code] = true
				}
			}
		case "pid":
			r := regexp.MustCompile(`[0-9]+`)
			if r.MatchString(value) && len(value) == 9 {
				compulsory[code] = true
			}
		}
	}

	for _, boolean := range compulsory {
		if boolean == false {
			return false
		}
	}

	return true
}

func part1(path string) int {
	passports, err := utils.ReadDoubleLines(path)
	if err != nil {
		return 0
	}

	var valid int
	for _, passport := range passports {
		fields := strings.Fields(passport)
		if hasRequiredFields(fields) {
			valid++
		}
	}

	return valid
}

func part2(path string) int {
	passports, err := utils.ReadDoubleLines(path)
	if err != nil {
		return 0
	}

	var valid int
	for _, passport := range passports {
		fields := strings.Fields(passport)
		if hasRequiredFields(fields) && hasValidFields(fields) {
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
