package main

import (
	"reflect"
	"testing"

	"github.com/rynhndrcksn/advent-of-code/internal"
)

func TestPart01(t *testing.T) {
	got := part01("input_test.txt")
	want := 161
	if got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}

func BenchmarkPart01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part01("input_test.txt")
	}
}

func TestPart01Regex(t *testing.T) {
	lines := internal.ReadFileLinesAsLines("input_test.txt")
	var got []string
	for _, line := range lines {
		got = part01Regex.FindAllString(line, -1)
	}
	want := 4
	if len(got) != want {
		t.Errorf("got: %d, want: %d", len(got), want)
	}
}

func TestExtractNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected [][]int
	}{
		{
			name:     "Basic multiplication expressions",
			input:    []string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"},
			expected: [][]int{{2, 4}, {5, 5}, {11, 8}, {8, 5}},
		},
		{
			name:     "Empty slice",
			input:    []string{},
			expected: [][]int{},
		},
		{
			name:     "Single expression",
			input:    []string{"mul(10,20)"},
			expected: [][]int{{10, 20}},
		},
		{
			name:     "Expression with larger numbers",
			input:    []string{"mul(100,200)", "mul(300,600)"},
			expected: [][]int{{100, 200}, {300, 600}},
		},
		{
			name:     "Expression with no numbers",
			input:    []string{"mul(,)", "mul()"},
			expected: [][]int{{}, {}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractNumbers(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("extractNumbers() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func BenchmarkExtractNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		extractNumbers([]string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"})
	}
}
