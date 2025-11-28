package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadString(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(bytes), nil
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func ReadDoubleLines(path string) ([]string, error) {
	str, err := ReadString(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(str, "\n\n")

	return lines, nil
}

func ReadInts(path string) ([]int, error) {
	lines, err := ReadLines(path)
	if err != nil {
		return nil, err
	}

	var ints []int
	for _, ss := range lines {
		n, err := strconv.Atoi(ss)
		if err != nil {
			return nil, err
		}
		ints = append(ints, n)
	}

	return ints, nil
}
