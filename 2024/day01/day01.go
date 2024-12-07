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
	fmt.Println("Part 1 Result:", day01("day01_input.txt"))
}

func day01(fileName string) int {
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

	// Create a scanner and add each column to their own slices.
	col1 := make([]int, 0, 10)
	col2 := make([]int, 0, 10)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "   ")

		var num1 int
		num1, err = strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("err converting parts[0] to int: %s", parts[0])
		}
		col1 = append(col1, num1)

		var num2 int
		num2, err = strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("err converting parts[1] to int: %s", parts[1])
		}
		col2 = append(col2, num2)
	}

	// Make sure scanner didn't encounter any issues.
	if err = scanner.Err(); err != nil {
		log.Fatalf("err from bufio scanner: %s", err.Error())
	}

	// Sort each column
	slices.Sort(col1)
	slices.Sort(col2)

	// Iterate over the slices and grab the total.
	var total int
	for k := range col1 {
		total += abs(col1[k] - col2[k])
	}

	return total
}

// abs returns the absolute version of an int.
// For example, -4 => 4 and 5 => 5.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
