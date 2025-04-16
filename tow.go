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
