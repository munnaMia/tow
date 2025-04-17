package tow_test

import (
	"testing"

	"github.com/munnaMia/tow"
)

func TestFromCharCode(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected string
	}{
		{"Single char", []int{65}, "A"},
		{"Multiple chars", []int{72, 101, 108, 108, 111}, "Hello"},
		{"Empty input", []int{}, ""},
		{"Unicode values", []int{9731, 9733, 9842}, "☃★♲"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tow.New("")
			result := s.FromCharCode(tt.input...).String()
			if result != tt.expected {
				t.Errorf("FromCharCode(%v) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCharAt(t *testing.T) {
	testString := "Let's have a home" // String to test on it
	tests := []struct {
		name     string
		idx      int
		expected string
	}{
		{"Positive idx", 4, "s"},
		{"Negative idx", -3, "o"},
		{"Out of range", 100, ""},
		{"Out of range", -100, ""},
		{"Positive length", len(testString), ""},
		{"Negative length", -len(testString), "L"},
		{"Negative out of range", -len(testString) - 1, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tow.New(testString)

			result := s.CharAt(tt.idx).String()

			if result != tt.expected {
				t.Errorf("CharAt(%v) = %q; want %q", tt.idx, result, tt.expected)
			}

		})
	}

}
