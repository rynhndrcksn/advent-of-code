package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Advent of Code - Day 01")
	safeCount := day02("day02_input.txt")
	fmt.Println("Part 1 Result:", safeCount)
}

func day02(filename string) int {
	// Open the file and prepare for iterating over it.
	file, err := os.Open(filename)
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
	rowIsSafe := make([]bool, 0, 10)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := make([]int, 0, 10)
		for _, field := range strings.Fields(scanner.Text()) {
			var num int // Declare this here so `err` isn't a shadowed variable.
			num, err = strconv.Atoi(field)
			if err != nil {
				log.Fatalf("error converting string to int: %s", err)
			}
			parts = append(parts, num)
		}
		partsLen := len(parts)

		// If numbers are ascending
		if parts[0] < parts[1] {
			for i := 0; i < partsLen-1; i++ {
				// If order isn't right, fail.
				if parts[i] > parts[i+1] {
					rowIsSafe = append(rowIsSafe, false)
					break
				}
				// If the numbers are the same, fail.
				if parts[i] == parts[i+1] {
					rowIsSafe = append(rowIsSafe, false)
					break
				}
				// Numbers can only have a difference of 1, 2, or 3.
				if abs(parts[i]-parts[i+1]) > 3 {
					rowIsSafe = append(rowIsSafe, false)
					break
				}
				// If we made it this far, it's a safe row.
				if i == partsLen-2 {
					rowIsSafe = append(rowIsSafe, true)
				}
			}
		} else if parts[0] > parts[1] { // If numbers are descending
			for i := 0; i < partsLen-1; i++ {
				// If order isn't right, fail.
				if parts[i] < parts[i+1] {
					rowIsSafe = append(rowIsSafe, false)
					break
				}
				// If the numbers are the same, fail.
				if parts[i] == parts[i+1] {
					rowIsSafe = append(rowIsSafe, false)
					break
				}
				// Numbers can only have a difference of 1, 2, or 3.
				if abs(parts[i]-parts[i+1]) > 3 {
					rowIsSafe = append(rowIsSafe, false)
					break
				}
				// If we made it this far, it's a safe row.
				if i == partsLen-2 {
					rowIsSafe = append(rowIsSafe, true)
				}
			}
		} else {
			// Numbers are the same
			rowIsSafe = append(rowIsSafe, false)
		}
	}

	safeReports := 0
	for _, v := range rowIsSafe {
		if v == true {
			safeReports += 1
		}
	}

	return safeReports
}

// abs returns the absolute version of an int.
// For example, -4 => 4 and 5 => 5.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
