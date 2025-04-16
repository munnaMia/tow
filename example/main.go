package main

import (
	"fmt"

	"github.com/munnaMia/tow"
)

func main() {
	text := tow.New("   test").FromCharCode(66).String()
	

	fmt.Println(text)
}
