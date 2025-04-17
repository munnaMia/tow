package tow

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

// Check if the given Sub-string exist or not on the string
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
