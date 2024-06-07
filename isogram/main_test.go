package main

import "testing"

func TestIsIsogram(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			input: "a",
			want:  true,
		},
		{
			input: "ab",
			want:  true,
		},
		{
			input: "abb",
			want:  false,
		},
		{
			input: "a-b-c",
			want:  true,
		},
	}
	for _, test := range tests {
		if got := IsIsogram(test.input); got != test.want {
			t.Errorf("IsIsogram(%q) = %t, want %t", test.input, got, test.want)
		}
	}
}
