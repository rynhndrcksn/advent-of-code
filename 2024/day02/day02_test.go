package main

import "testing"

var filename = "input_test.txt"

func TestDay02(t *testing.T) {
	got := day02(filename)
	want := 4
	if got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}

func BenchmarkDay02(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day02(filename)
	}
}
