package main

import "testing"

var ContainsTests = []struct {
	name  string
	slice []string
	check string
	out   bool
}{
	{"happy", []string{"a", "b"}, "a", true},
	{"negative", []string{"a", "b"}, "c", false},
	{"words happy", []string{"cat", "bat", "goat"}, "goat", true},
}

func TestContains(t *testing.T) {
	for _, tt := range ContainsTests {
		t.Run(tt.name, func(t *testing.T) {
			result := contains(tt.slice, tt.check)
			if result != tt.out {
				t.Errorf("got %t, want %t", result, tt.out)
			}
		})
	}
}
