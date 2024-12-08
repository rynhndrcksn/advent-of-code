package main

import (
	"fmt"
	"slices"

	"github.com/rynhndrcksn/advent-of-code/internal"
)

func main() {
	fmt.Println("Advent of Code - Day 01")
	distance, similarity := day01("input.txt")
	fmt.Println("Part 1 Result:", distance)
	fmt.Println("Part 2 Results:", similarity)
}

func day01(filename string) (int, int) {
	parts := internal.ReadFileLinesAsInts(filename)
	col1Ints := make([]int, 0, 10)
	col2Ints := make([]int, 0, 10)
	col1Occurrences := make(map[int]int)
	col2Occurrences := make(map[int]int)

	for i := 0; i < len(parts); i += 2 {
		num1 := parts[i]
		num2 := parts[i+1]
		col1Ints = append(col1Ints, num1)
		col2Ints = append(col2Ints, num2)

		// Update the maps tracking number occurrences.
		if _, ok := col1Occurrences[num1]; ok {
			col1Occurrences[num1] += 1
		} else {
			col1Occurrences[num1] = 1
		}
		if _, ok := col2Occurrences[num2]; ok {
			col2Occurrences[num2] += 1
		} else {
			col2Occurrences[num2] = 1
		}
	}

	// Sort each column.
	slices.Sort(col1Ints)
	slices.Sort(col2Ints)

	// Compute distance.
	var distance int
	for k := range col1Ints {
		distance += internal.AbsInt(col1Ints[k] - col2Ints[k])
	}

	// Compute similarity score.
	var similarity int
	for k, v := range col1Occurrences {
		num := col2Occurrences[k]
		for i := 0; i < v; i++ {
			similarity += k * num
		}
	}

	return distance, similarity
}
