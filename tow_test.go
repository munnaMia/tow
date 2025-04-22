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
			result := s.FromCharCode(tt.input...).ToString()
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
		{"Positive out of range", 100, ""},
		{"Negative out of range", -100, ""},
		{"Positive length", len(testString), ""},
		{"Negative length", -len(testString), "L"},
		{"Negative out of range", -len(testString) - 1, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tow.New(testString)

			result := s.CharAt(tt.idx).ToString()

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
				t.Errorf("CharCodeAt(%v) = %d; want %d", tt.idx, result, tt.expected)
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
			result := s.Concat(tt.texts...).ToString()
			if result != tt.expected {
				t.Errorf("Concat(%v) = %q; want %q", tt.texts, result, tt.expected)
			}

		})
	}
}
func TestIncludes(t *testing.T) {
	testString := "The quick brown fox jumps over the lazy dog." // String to test on it
	tests := []struct {
		name     string
		text     string
		expected bool
	}{
		{"Simple word", "fox", true},
		{"Invalid word", "cow", false},
		{"Empty word", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tow.New(testString)
			result := s.Includes(tt.text)
			if result != tt.expected {
				t.Errorf("Includes(%v) = %t; want %t", tt.text, result, tt.expected)
			}
		})
	}

}
func TestIndexOf(t *testing.T) {
	testString := "The quick brown fox jumps over the lazy dog." // String to test on it
	tests := []struct {
		name     string
		text     string
		expected int
	}{
		{"Simple word", "fox", 16},
		{"Invalid word", "cow", -1},
		{"Empty word", "", -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tow.New(testString)
			result := s.IndexOf(tt.text)
			if result != tt.expected {
				t.Errorf("IndexOf(%v) = %d; want %d", tt.text, result, tt.expected)
			}
		})
	}

}
func TestEndsWith(t *testing.T) {
	testString := "The fox is fast." // String to test on it
	tests := []struct {
		name     string
		text     string
		expected bool
	}{
		{"Valid word", "fast.", true},
		{"Invalid word", "cow", false},
		{"Empty word", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tow.New(testString)
			result := s.EndsWith(tt.text)
			if result != tt.expected {
				t.Errorf("EndsWith(%v) = %t; want %t", tt.text, result, tt.expected)
			}
		})
	}
}
func TestLength(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected int
	}{
		{"Small text", "Test", 4},
		{"Midium text", "This is a cow", 13},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tow.New(tt.text)
			result := s.Length()
			if result != tt.expected {
				t.Errorf("Length() = %d; want %d", result, tt.expected)
			}
		})
	}

}
func TestTrim(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected string
	}{
		{"Space both side", "   Hello   ", "Hello"},
		{"Space at the start", "   cow", "cow"},
		{"Space at the end", "test   ", "test"},
		{"More space at the end", "test     ", "test"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tow.New(tt.text)
			result := s.Trim().ToString()
			if result != tt.expected {
				t.Errorf("Trim() = %s; want %s", result, tt.expected)
			}
		})
	}
}
func TestTrimStart(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected string
	}{
		{"Space both side", "   Hello   ", "Hello   "},
		{"Space at the start", "   cow", "cow"},
		{"Space at the end", "tryout   ", "tryout   "},
		{"More space at the end", "test     ", "test     "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tow.New(tt.text)
			result := s.TrimStart().ToString()
			if result != tt.expected {
				t.Errorf("TrimStart() = %s; want %s", result, tt.expected)
			}
		})
	}

}
func TestTrimEnd(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected string
	}{
		{"Space both side", "   Hello   ", "   Hello"},
		{"Space at the start", "   cow", "   cow"},
		{"Space at the end", "try   ", "try"},
		{"More space at the end", "test     ", "test"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tow.New(tt.text)
			result := s.TrimEnd().ToString()
			if result != tt.expected {
				t.Errorf("TrimEnd() = %s; want %s", result, tt.expected)
			}
		})
	}

}

func TestToUpperCase(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected string
	}{
		{"Simple string", "this is a simple string", "THIS IS A SIMPLE STRING"},
		{"Mixed string", "#this is Mixsed string", "#THIS IS MIXSED STRING"},
		{"Uppercase string", "THIS IS A UPPERCASE STRING", "THIS IS A UPPERCASE STRING"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tow.New(tt.text)
			result := s.ToUpperCase().ToString()
			if result != tt.expected {
				t.Errorf("ToUpperCase() = %s; want %s", result, tt.expected)
			}
		})
	}
}

func TestToLowerCase(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected string
	}{
		{"Simple lowercase string", "this is a simple string", "this is a simple string"},
		{"Mixed string", "#tHis is Mixsed stRing", "#this is mixsed string"},
		{"Uppercase string", "THIS IS A UPPERCASE STRING", "this is a uppercase string"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tow.New(tt.text)
			result := s.ToLowerCase().ToString()
			if result != tt.expected {
				t.Errorf("ToUpperCase() = %s; want %s", result, tt.expected)
			}
		})
	}
}

func TestPadStart(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		pLen     int
		pCh      rune
		expected string
	}{
		{"No pad added", "text text", 4, '*', "text text"},
		{"Adding Simple padding", "232", 8, '*', "*****232"},
		{"With end space padding", "232 ", 8, '*', "****232 "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tow.New(tt.text)
			result := s.PadStart(tt.pLen, tt.pCh).ToString()
			if result != tt.expected {
				t.Errorf("PadStart() = %s; want %s", result, tt.expected)
			}
		})
	}
}

func TestEnd(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		pLen     int
		pCh      rune
		expected string
	}{
		{"No pad added", "text text", 4, '.', "text text"},
		{"Adding Simple padding", "232", 8, '.', "232....."},
		{"With start space padding", " 232", 8, '.', " 232...."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tow.New(tt.text)
			result := s.PadEnd(tt.pLen, tt.pCh).ToString()
			if result != tt.expected {
				t.Errorf("PadStart() = %s; want %s", result, tt.expected)
			}
		})
	}
}

func TestStartsWith(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		input    string
		expected bool
	}{
		{"Basic match", "hello world", "hello", true},
		{"Not at the beginning", "hello world", "world", false},
		{"Empty string and prefix", "", "", true},
		{"Empty string, non-empty prefix", "", "hello", false},
		{"Non-empty string, empty prefix", "hello world", "", true},
		{"Prefix longer than string", "hello", "hello world", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tow.New(tt.text)
			result := s.StartsWith(tt.input)
			if result != tt.expected {
				t.Errorf("StartsWith(%s) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestRepeat(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		count    int
		expected string
	}{
		{"Basic text", "hello", 3, "hellohellohello"},
		{"With space at the end", "hello ", 3, "hello hello hello "},
		{"Empty string", "", 3, ""},
		{"Only space", " ", 3, "   "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tow.New(tt.input)
			result := s.Repeat(tt.count).ToString()
			if result != tt.expected {
				t.Errorf("Repeat(%v) = %s; want %s", tt.count, result, tt.expected)
			}
		})
	}
}
