package internal

import (
	"math/rand/v2"
	"testing"
)

func TestAbsInt(t *testing.T) {
	tests := map[string]struct {
		input int
		want  int
	}{
		"negative int": {
			input: -5,
			want:  5,
		},
		"positive int": {
			input: 5,
			want:  5,
		},
		"zero": {
			input: 0,
			want:  0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := AbsInt(test.input)
			if got != test.want {
				t.Errorf("got: %d, want: %d", got, test.want)
			}
		})
	}
}

func BenchmarkAbsInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AbsInt(rand.IntN(100))
	}
}

func TestReadFileLinesAsInts(t *testing.T) {
	slice := ReadFileLinesAsInts("read_file_lines_as_ints.txt")
	expectedLen := 10
	t.Run("has correct length", func(t *testing.T) {
		if len(slice) != expectedLen {
			t.Errorf("got: %d, expectedLen: %d", len(slice), expectedLen)
		}
	})
	t.Run("has correct elements", func(t *testing.T) {
		for i := 0; i < len(slice); i++ {
			if i != slice[i] {
				t.Errorf("got: %d, want: %d", slice[i], i)
			}
		}
	})
}

func BenchmarkReadFileLinesAsInts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadFileLinesAsInts("read_file_lines_as_ints.txt")
	}
}

func TestReadFileLinesAsWords(t *testing.T) {
	slice := ReadFileLinesAsWords("read_file_lines_as_words.txt")
	expectedLen := 8
	t.Run("has correct length", func(t *testing.T) {
		if len(slice) != expectedLen {
			t.Errorf("got: %d, want: %d", len(slice), expectedLen)
		}
	})

	// File contents: This is a file that contains eight words.
	expectedElems := []string{"This", "is", "a", "file", "that", "contains", "eight", "words."}
	t.Run("has correct elements", func(t *testing.T) {
		for i := 0; i < len(slice); i++ {
			if slice[i] != expectedElems[i] {
				t.Errorf("got: %s, want: %s", slice[i], expectedElems[i])
			}
		}
	})
}

func BenchmarkReadFileLinesAsWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadFileLinesAsWords("read_file_lines_as_words.txt")
	}
}
