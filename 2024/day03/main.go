package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/rynhndrcksn/advent-of-code/internal"
)

var part01Regex = regexp.MustCompile("mul\\(\\d+,\\d+\\)")
var part02Regex = regexp.MustCompile("(?:don't|do|mul)\\(\\d*,?\\d*\\)")

func main() {
	fmt.Println("Advent of Code - Day 03")
	result := part01("input.txt")
	fmt.Println("Day 03 Part 1 Result:", result)
	result = part02("input.txt")
	fmt.Println("Day 03 Part 2 Result:", result)
}

func part01(filename string) int {
	// Read file contents.
	lines := internal.ReadFileLinesAsLines(filename)

	// Grab valid regex matches.
	matches := make([]string, 0, 16)
	for _, line := range lines {
		matches = append(matches, part01Regex.FindAllString(line, -1)...)
	}
	nums := extractNumbers(matches)

	// Perform the multiplications and add everything together.
	total := 0
	for _, ints := range nums {
		total += ints[0] * ints[1]
	}

	return total
}

func part02(filename string) int {
	// Read file contents.
	lines := internal.ReadFileLinesAsLines(filename)

	// Grab valid regex matches (including don't() and do()).
	matches := make([]string, 0, 16)
	for _, line := range lines {
		matches = append(matches, part02Regex.FindAllString(line, -1)...)
	}

	// Only grab the valid matches, determined by the don't() and do() keywords.
	// Note: we start in an enabled state, and only disable when/if we encounter "don't()".
	enabledMatches := make([]string, 0, 16)
	enabled := true
	for _, match := range matches {
		switch match {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if enabled {
				enabledMatches = append(enabledMatches, match)
			}
		}
	}
	nums := extractNumbers(enabledMatches)

	// Perform the multiplications and add everything together.
	total := 0
	for _, ints := range nums {
		total += ints[0] * ints[1]
	}

	return total
}

// extractNumbers iterates over a slice of strings that match the regex /mul\(\d+,\d+\)/ and
// returns a multidimensional slice of ints containing the numbers within the ()s.
func extractNumbers(expressions []string) [][]int {
	re := regexp.MustCompile(`\d+`)
	result := make([][]int, len(expressions))

	for i, expr := range expressions {
		matches := re.FindAllString(expr, -1)
		nums := make([]int, len(matches))
		for j, match := range matches {
			nums[j], _ = strconv.Atoi(match)
		}
		result[i] = nums
	}
	return result
}
