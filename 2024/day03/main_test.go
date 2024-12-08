package main

import "testing"

func TestPart01(t *testing.T) {
	got := part01("input_test.txt")
	want := 161
	if got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}
