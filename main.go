package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	size           = flag.Int("s", 8, "Password size")
	numbers        = flag.String("N", "1234567890", "List of numbers to be used")
	symbols        = flag.String("Y", "!#$%&'()*+,.:;<=>?@[]^{|}~", "List of symbols to be used")
	lower          = flag.String("L", "abcdefghijklmnopqrstuvwxyz", "List of lower case letters to be used")
	upper          = flag.String("U", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "List of upper case letters to be used")
	includeSymbols = flag.Bool("y", true, "Include symbols")
	includeLower   = flag.Bool("l", true, "Include lower letters")
	includeUpper   = flag.Bool("u", true, "Include capital letters")
	includeNumbers = flag.Bool("n", true, "Include numbers")
)

func convertToList(s string) []string {
	return strings.Split(s, "")
}

func main() {
	flag.Parse()

	var allChars []string

	if *includeSymbols {
		allChars = append(allChars, convertToList(*symbols)...)
	}
	if *includeNumbers {
		allChars = append(allChars, convertToList(*numbers)...)
	}
	if *includeLower {
		allChars = append(allChars, convertToList(*lower)...)
	}
	if *includeUpper {
		allChars = append(allChars, convertToList(*upper)...)
	}

	rand.Seed(time.Now().UnixNano())
	var password []string
	for char := 0; char < *size; char++ {
		password = append(password, allChars[rand.Intn(len(allChars))])
	}

	fmt.Println(strings.Join(password, ""))
}
