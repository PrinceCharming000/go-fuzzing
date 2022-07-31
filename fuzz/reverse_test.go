package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tesecases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"1234567890", "0987654321"},
	}
	for _, tc := range tesecases {
		rev := Reverse(tc.in)
		if rev != tc.want {
			t.Errorf("Reverse: %q want %q", rev, tc.want)
		}
	}
}
