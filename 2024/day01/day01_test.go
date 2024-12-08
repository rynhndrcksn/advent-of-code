package main

import (
	"testing"
)

var filename = "input_test.txt"

func TestDay01(t *testing.T) {
	distance, similarity := day01(filename)
	want := 11
	if distance != want {
		t.Errorf("got: %d, want: %d", distance, want)
	}
	want = 31
	if similarity != want {
		t.Errorf("got: %d, want: %d", similarity, want)
	}
}

func BenchmarkDay01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day01(filename)
	}
}
