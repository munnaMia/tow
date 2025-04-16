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

// Converted raw strinig
func (s *Str) String() string {
	return s.value
}

// return a string from a sequence of Unicode values (character codes).
func (s *Str) FromCharCode(charCode ...int) *Str {
	for _, char := range charCode {
		s.value += string(rune(char))
	}
	return s
} 
