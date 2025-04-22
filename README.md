# TOW

<img align="right" width="160px" src="./assets/test.png">

### 📦 Tow – A Simple Go String Utility Library with Method Chaining

#### tow is a lightweight Go package that enables method chaining for string operations, inspired by JavaScript, and implemented in pure Go. This library is designed for easy manipulation of strings while offering a fluent interface that allows multiple operations to be chained together.

---

- 🧠 Built by a beginner, for beginners.
- 🌱 First step into Go library development.
- 🔗 JS-style chaining for strings in pure Go.

## ✨ Features
- ✅ Method chaining for string operations (inspired by JavaScript).

- ✅ A collection of useful string manipulation methods.

- ✅ Written from scratch without using go string package &  tested as cleanly as possible, with a focus on learning

- ✅ Written entirely in Go without relying on external dependencies.

- ✅ Built with simplicity and readability in mind for Go developers.

## ⚠ Limitations
- ❌ No regex support (yet)

- 🧪 Limited methods as of now

- 🚧 Only basic string manipulation functions are included

- 🔒 Currently does not support complex string patterns or transformations
 
- 🛠 No advanced error handling for edge cases (e.g., invalid input types)
 
- ⚡ Performance not yet optimized for large-scale string operations


## 📦 Installation
You can install the tow package via go get:
```bash
go get github.com/munnaMia/tow
```

## 🔍 Example Usage

### Javascript
```js
const text = "   Hello World!   ".trim().replace("World","JS").toUpperCase();

console.log(text); // Output: HELLO JS!
```


### GO
```go
package main

import (
    "fmt"
    "github.com/munnaMia/tow"
)

func main() {
    text := tow.New("   hello world   ").Trim().Replace("World","JS").ToUpperCase().ToString()
                
    fmt.Println(text) // Output: HELLO JS!
}

```


## 🛠 Methods Available
### Constructor:
- **New(str string)** – Creates a new Str object with the given string.

### String Manipulation:
- **ToString()** - Returns the raw string stored in the Str object.
- **Length()** - Returns the length of the string.
- **FromCharCode(charCode ...int)** - Converts a sequence of Unicode values (character codes) into a string.
- **CharAt(index int)** - Returns the character at a specific index.
- **CharCodeAt(index int)** - Returns the Unicode value of the character at a specific index.
- **Concat(strs ...string)** - Concatenates multiple strings together.
- **Includes(str string)** - Checks if the string contains the specified substring.
- **EndsWith(str string)** - Checks if the string ends with the specified substring.
- **StartsWith(str string)** - Checks if the string starts with the specified text.
- **IndexOf(str string)** - Returns the index of the first occurrence of a specified substring.
- **Trim()** - Removes whitespace from both ends of the string.
- **TrimStart()** - Removes whitespace from the beginning of the string.
- **TrimEnd()** - Removes whitespace from the end of the string.
- **ToUpperCase()** - Converts the string to uppercase.
- **ToLowerCase()** - Converts the string to lowercase.
- **PadStart(padLength int, padChar rune)** - Pads the start of the string with the specified character.
- **PadEnd(padLength int, padChar rune)** - Pads the end of the string with the specified character.
- **Repeat(count int)** - Repeats the string count times.
- **Replace(search, replace string)** - Replaces the first occurrence of a substring with another string.
- **ReplaceAll(search, replace string)** - Replaces all occurrences of a substring with another string.
- **Match(pattern string)** - Returns a slice of strings that match the specified regular expression pattern.
- **Slice(start, end int)** - Returns a substring of the original string between the specified indices.
- **Split(separator string)** - Splits the string into a slice of substrings based on the given separator.




## 🙌 Why TOW?
> “It was always my dream to build a small library of my own. As a beginner, I looked for a simple problem to solve — and I found one. Go lacks the syntactic sugar of method chaining in string manipulation, something I enjoyed in JavaScript. So, I created TOW – a minimalistic string utility library inspired by JavaScript.”