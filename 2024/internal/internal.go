package internal

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// AbsInt takes an integer and returns it's absolute value.
// For example: -4 returns 4, 5 returns 5, etc.
func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// ReadFileLinesAsLines takes a filename, parses it by newlines, and returns a slice of lines.
func ReadFileLinesAsLines(filename string) []string {
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

	parts := make([]string, 0, 16)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts = append(parts, scanner.Text())
	}

	// Make sure scanner didn't encounter any issues.
	if err = scanner.Err(); err != nil {
		log.Fatalf("err from bufio scanner: %s", err.Error())
	}

	return parts
}

// ReadFileLinesAsInts takes a filename, parses it by whitespace, and returns a slice of ints.
func ReadFileLinesAsInts(filename string) []int {
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

	parts := make([]int, 0, 16)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, field := range strings.Fields(scanner.Text()) {
			var num int // Declare this here so `err` isn't a shadowed variable.
			num, err = strconv.Atoi(field)
			if err != nil {
				log.Fatalf("error converting string to int: %s", err)
			}
			parts = append(parts, num)
		}
	}

	// Make sure scanner didn't encounter any issues.
	if err = scanner.Err(); err != nil {
		log.Fatalf("err from bufio scanner: %s", err.Error())
	}

	return parts
}

// ReadFileLinesAsWords takes a filename, parses it by whitespace, and returns a slice of the words.
func ReadFileLinesAsWords(filename string) []string {
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

	parts := make([]string, 0, 16)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, word := range strings.Fields(scanner.Text()) {
			parts = append(parts, word)
		}
	}

	// Make sure scanner didn't encounter any issues.
	if err = scanner.Err(); err != nil {
		log.Fatalf("err from bufio scanner: %s", err.Error())
	}

	return parts
}
