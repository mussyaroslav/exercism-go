package main

import "testing"

func TestScore(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "Scrabble",
			want:  14,
		},
		{
			input: "Happy",
			want:  15,
		},
		{
			input: "Sad",
			want:  4,
		},
	}
	for _, test := range tests {
		if got := Score(test.input); got != test.want {
			t.Errorf("Score(%q) = %d, want %d", test.input, got, test.want)
		}
	}
}
