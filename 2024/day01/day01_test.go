package main

import "testing"

func TestDay01(t *testing.T) {
	got := day01("input_test.txt")
	want := 11
	if got != want {
		t.Errorf("got: %d, want %d", got, want)
	}
}
