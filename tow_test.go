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

func TestCharCodeAt(t *testing.T) {
	testString := "Let's have a ☃★♲" // String to test on it
	tests := []struct {
		name     string
		idx      int
		expected int
	}{
		{"Positive idx", 4, 115},
		{"Negative idx", -3, -1},
		{"Out of range", 100, -1},
		{"Positive length", len(testString), -1},
		{"Charater Test", len([]rune(testString)) - 1, 9842},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tow.New(testString)
			result := s.CharCodeAt(tt.idx)
			if result != tt.expected {
				t.Errorf("CharCodeAt(%v) = %q; want %q", tt.idx, result, tt.expected)
			}

		})
	}
}

func TestConcat(t *testing.T) {
	testString := "hello" // String to test on it
	tests := []struct {
		name     string
		texts    []string
		expected string
	}{
		{"Simple word", []string{" ", "world"}, "hello world"},
		{"Simple texts", []string{" to the ", "world"}, "hello to the world"},
		{"Special char", []string{"☃★♲"}, "hello☃★♲"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tow.New(testString)
			result := s.Concat(tt.texts...).String()
			if result != tt.expected {
				t.Errorf("CharCodeAt(%v) = %q; want %q", tt.texts, result, tt.expected)
			}

		})
	}
}

func TestIncludes(t *testing.T){
	testString := "The quick brown fox jumps over the lazy dog." // String to test on it
	tests := []struct {
		name     string
		text    string
		expected bool
	}{
		{"Simple word", "fox", true},
		{"Invalid word", "cow", false},
		{"Empty word", "", true},
	}

	for _, tt := range tests{
		t.Run(tt.name,func(t *testing.T) {
			s := tow.New(testString)
			result := s.Includes(tt.text)
			if result != tt.expected {
				t.Errorf("Includes(%v) = %t; want %t", tt.text, result, tt.expected)
			}
		})
	}

}


func TestIndexOf(t *testing.T){
	testString := "The quick brown fox jumps over the lazy dog." // String to test on it
	tests := []struct {
		name     string
		text    string
		expected int
	}{
		{"Simple word", "fox", 16},
		{"Invalid word", "cow", -1},
		{"Empty word", "", -1},
	}

	for _, tt := range tests{
		t.Run(tt.name,func(t *testing.T) {
			s := tow.New(testString)
			result := s.IndexOf(tt.text)
			if result != tt.expected {
				t.Errorf("Includes(%v) = %q; want %q", tt.text, result, tt.expected)
			}
		})
	}

}
