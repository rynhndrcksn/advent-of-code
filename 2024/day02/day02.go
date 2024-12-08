package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/rynhndrcksn/advent-of-code/internal"
)

func main() {
	fmt.Println("Advent of Code - Day 01")
	safeCount := day02("input.txt")
	fmt.Println("Day 02 Result:", safeCount)
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
	scanner := bufio.NewScanner(file)
	safeNum := 0
	for scanner.Scan() {
		parts := make([]int, 0, 16)
		for _, field := range strings.Fields(scanner.Text()) {
			var num int // Declare this here so `err` isn't a shadowed variable.
			num, err = strconv.Atoi(field)
			if err != nil {
				log.Fatalf("error converting string to int: %s", err)
			}
			parts = append(parts, num)
		}

		safe := checkPart(parts)

		// If it's safe, no need to run the second loop below.
		if safe {
			safeNum++
			continue
		}

		for i := range len(parts) {
			newParts := make([]int, 0, len(parts)-1)
			newParts = append(newParts, parts[:i]...)
			newParts = append(newParts, parts[i+1:]...)

			if checkPart(newParts) {
				safeNum++
				break
			}
		}
	}

	return safeNum
}

func checkPart(parts []int) bool {
	prevNum := 0
	isIncreasing := false
	for i, num := range parts {
		if i == 0 {
			prevNum = num
			continue
		}

		if prevNum == num {
			return false
		}

		if internal.AbsInt(prevNum-num) > 3 {
			return false
		}

		if i == 1 {
			isIncreasing = num > prevNum

			prevNum = num
			continue
		}

		if (num > prevNum) != isIncreasing {
			return false
		}

		prevNum = num
	}
	return true
}
