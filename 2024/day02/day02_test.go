package main

import "testing"

var filename = "day02_input_test.txt"

func TestDay02(t *testing.T) {
	got := day02(filename)
	want := 2
	if got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}
