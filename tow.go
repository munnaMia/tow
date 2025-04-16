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
	if index < 0 {
		index = len(s.value) + index // Allows negative indexing (like Python)
	}
	s.value = string(s.value[index])
	return s
}

// // return Unicode value of a charater
// func (s *Str) CharCodeAt(index int) int {

// }
