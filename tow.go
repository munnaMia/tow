package tow

import (
	"regexp"
	"unicode"
)

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
func (s *Str) ToString() string {
	return s.value
}

// Return length of the string
func (s *Str) Length() int {
	return len(s.value)
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

// This method  returns string converted to uppercase.
func (s *Str) ToUpperCase() *Str {
	runes := []rune(s.value)
	for idx, word := range runes {
		if word >= 'a' && word <= 'z' {
			runes[idx] = word - 32
		}
	}
	s.value = string(runes)
	return s
}

// This method  returns string converted to lowercase.
func (s *Str) ToLowerCase() *Str {
	runes := []rune(s.value)
	for idx, word := range runes {
		if word >= 'A' && word <= 'Z' {
			runes[idx] = word + 32
		}
	}
	s.value = string(runes)
	return s
}

// Add padding start of the string
func (s *Str) PadStart(padLength int, padChar rune) *Str {
	// If padding is not needed, return the original string
	if padLength < len(s.value) {
		return s // no pad will be added
	}

	padCount := padLength - len([]rune(s.value)) // get the length for pad charater length

	result := make([]rune, padCount)

	// Create a slice with the padding characters
	for i := 0; i < padCount; i++ {
		result[i] = padChar
	}

	result = append(result, []rune(s.value)...) // Combine padding and original string
	s.value = string(result)
	return s
}

// Add padding end of the string
func (s *Str) PadEnd(padLength int, padChar rune) *Str {
	// If padding is not needed, return the original string
	if padLength < len(s.value) {
		return s // no pad will be added
	}

	padCount := padLength - len([]rune(s.value)) // get the length for pad charater length

	padding := make([]rune, padCount)

	// Create a slice with the padding characters
	for i := 0; i < padCount; i++ {
		padding[i] = padChar
	}

	s.value = string(append([]rune(s.value), padding...)) // Combine input string with padding and convert them to a string

	return s
}

// Check the string is start with the given string and retrun boolean.
func (s *Str) StartsWith(text string) bool {
	return len(text) <= len(s.value) && s.value[:len(text)] == text
}

func (s *Str) Repeat(count int) *Str {
	word := s.value
	for i := 1; i < count; i++ {
		s.value += word
	}
	return s
}

// replace the string with the given one and return a string.
// will not modify the orignal string.
func (s *Str) Replace(search, replace string) *Str {
	if len(search) > len(s.value) || search == "" {
		return s
	}

	for i := 0; i <= len(s.value)-len(search); i++ {
		if s.value[i:i+len(search)] == search {
			s.value = s.value[:i] + replace + s.value[i+len(search):]
			return s
		}
	}
	return s
}

// replace all the string with the given one and return a string.
// will not modify the orignal string.
func (s *Str) ReplaceAll(search, replace string) *Str {
	if len(search) > len(s.value) || search == "" {
		return s
	}
	result := ""
	i := 0
	for i <= len(s.value)-len(search) {
		if s.value[i:i+len(search)] == search {
			result += replace
			i += len(search)
		} else {
			result += string(s.value[i])
			i++
		}
	}
	result += s.value[i:]
	s.value = result
	return s
}

// Match returns the first match of the pattern.
// If the pattern includes groups, returns only the matched string.
func (s *Str) Match(pattern string) []string {
	re := regexp.MustCompile(pattern)
	return re.FindAllString(s.value, -1)
}

// Return a slice of a string.
func (s *Str) Slice(start, end int) string {
	runes := []rune(s.value)
	length := len(runes)
	// handling negative numbers
	if start < 0 {
		start += length
	}
	if end < 0 {
		end += length
	}
	// Clamp to valid bounds
	if start < 0 {
		start = 0
	}
	if end < 0 {
		end = 0
	}
	if end > length {
		end = length
	}
	if start > end {
		return ""
	}

	return string(runes[start:end])
}

// Return a string array.
func (s *Str) Split(separator string) []string {
	if separator == "" {
		return []string{s.value} //for empty separator return array with whole string value
	}
	var result []string
	text := s.value
	separatorLen := len([]rune(separator))
	start := 0
	for i := 0; i <= len(text)-separatorLen; {
		if text[i:i+separatorLen] == separator {
			result = append(result, text[start:i])
			i += separatorLen
			start = i
		} else {
			i++
		}
	}
	result = append(result, text[start:])
	return result
}
