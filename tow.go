package tow

import "unicode"

type Str struct {
	value string
}

// constructor
func New(s string) *Str {
	return &Str{
		value: s,
	}
}

// Converted raw string
func (s *Str) String() string {
	return s.value
}

// return a string from a sequence of Unicode values (character codes).
func (s *Str) FromCharCode(charCode ...int) *Str {
	s.value = ""
	for _, char := range charCode {
		s.value += string(rune(char))
	}
	return s
}

// Take an Index and return a single string character
func (s *Str) CharAt(index int) *Str {
	length := len(s.value)

	// Work on negative indexing
	if index < 0 {
		index = length + index // Allows negative indexing (like Python)
	}

	// If index is out of range retrun "" empty string
	if length <= index || index < 0 {
		s.value = ""
		return s
	}

	s.value = string(s.value[index])
	return s
}

// return Unicode value of a charater
func (s *Str) CharCodeAt(index int) int {
	runes := []rune(s.value)
	if index >= len(runes) || index < 0 {
		return -1
	}
	return int(runes[index])
}

// Concat multiple strings
func (s *Str) Concat(strs ...string) *Str {
	for _, str := range strs {
		s.value += string(str)
	}

	return s
}

// Check if the given Sub-string exist or not on the string and return boolean
func (s *Str) Includes(str string) bool {
	strLength := len(str)

	// If length is 0 return true
	if strLength == 0 {
		return true
	}

	// Checking the string exist of not
	for i := 0; i+strLength < len(s.value); i++ {
		if s.value[i:i+strLength] == str {
			return true
		}
	}

	return false // return false if match not found
}

// Checks if string ends with given substring
func (s *Str) EndsWith(str string) bool {
	sLength := len(s.value) // given main string length
	strLength := len(str)   // Substring length
	if strLength == 0 {
		return false
	}
	return s.value[sLength-strLength:sLength] == str // checking the last string
}

// This method searches on string and returns the index of the first occurrence of the specified substring.
func (s *Str) IndexOf(str string) int {
	strLength := len(str)

	// If substring length is 0 return -1
	if strLength == 0 {
		return -1
	}

	// Search for the string
	for i := 0; i+strLength < len(s.value); i++ {
		if str == s.value[i:i+strLength] {
			return i
		}
	}

	return -1
}

// This method take a string and removes whitespace from both ends of the string
func (s *Str) Trim() *Str {
	startIdx := 0
	endIdx := len(s.value) - 1

	// start spaces
	for startIdx <= endIdx {
		if !unicode.IsSpace(rune(s.value[startIdx])) {
			break
		}
		startIdx++
	}

	// end spaces
	for startIdx <= endIdx {
		if !unicode.IsSpace(rune(s.value[endIdx])) {
			break
		}
		endIdx--
	}

	s.value = s.value[startIdx : endIdx+1]
	return s
}

// This method take a string and removes whitespace from ends of the string
func (s *Str) TrimEnd() *Str {
	startIdx := 0
	endIdx := len(s.value) - 1
	// end spaces
	for startIdx <= endIdx {
		if !unicode.IsSpace(rune(s.value[endIdx])) {
			break
		}
		endIdx--
	}

	s.value = s.value[startIdx : endIdx+1]
	return s
}

// This method take a string and removes whitespace from ends of the string
func (s *Str) TrimStart() *Str {
	startIdx := 0
	endIdx := len(s.value) - 1
	// start spaces
	for startIdx <= endIdx {
		if !unicode.IsSpace(rune(s.value[startIdx])) {
			break
		}
		startIdx++
	}

	s.value = s.value[startIdx : endIdx+1]
	return s
}
