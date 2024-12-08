package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Advent of Code - Day 01")
	distance, similarity := day01("input.txt")
	fmt.Println("Part 1 Result:", distance)
	fmt.Println("Part 2 Results:", similarity)
}

func day01(fileName string) (int, int) {
	// Open the file and prepare for iterating over it.
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("err opening filename: %s", err.Error())
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatalf("err while closing filename: %s", err.Error())
		}
	}(file)

	// Prepare some variables to keep track of things, and iterate over the scanner.
	col1Ints := make([]int, 0, 10)
	col2Ints := make([]int, 0, 10)
	col1Occurrences := make(map[int]int)
	col2Occurrences := make(map[int]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "   ")

		// Parse the first int and add it to the col1Ints slice.
		var num1 int
		num1, err = strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("err converting parts[0] to int: %s", parts[0])
		}
		col1Ints = append(col1Ints, num1)

		// Parse the second int and add it to the col2Ints slice.
		var num2 int
		num2, err = strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("err converting parts[1] to int: %s", parts[1])
		}
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

	// Make sure scanner didn't encounter any issues.
	if err = scanner.Err(); err != nil {
		log.Fatalf("err from bufio scanner: %s", err.Error())
	}

	// Sort each column.
	slices.Sort(col1Ints)
	slices.Sort(col2Ints)

	// Compute similarity score.
	var similarity int
	for k, v := range col1Occurrences {
		num := col2Occurrences[k]
		for i := 0; i < v; i++ {
			similarity += k * num
		}
	}

	// Iterate over the slices and grab the distance.
	var distance int
	for k := range col1Ints {
		distance += abs(col1Ints[k] - col2Ints[k])
	}

	return distance, similarity
}

// abs returns the absolute version of an int.
// For example, -4 => 4 and 5 => 5.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
