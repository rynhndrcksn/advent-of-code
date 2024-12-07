package main

import "testing"

var filename = "day01_input_test.txt"

func TestDay01(t *testing.T) {
	got := day01(filename)
	want := 11
	if got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}

func TestAbs(t *testing.T) {
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
			got := abs(test.input)
			if got != test.want {
				t.Errorf("got: %d, want: %d", got, test.want)
			}
		})
	}
}

func BenchmarkDay01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day01(filename)
	}
}
