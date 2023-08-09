package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
)

const (
	lowerChars = "abcdefghijklmnopqrstuvwxyz"
	upperChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numChars   = "0123456789"
	symChars   = "!@#$%^&*()-_+=<>?"
)

var (
	length      int
	includeNum  bool
	includeSym  bool
	includeUpp  bool
	includeLow  bool
)

func init() {
	flag.IntVar(&length, "length", 8, "Length of the password")
	flag.BoolVar(&includeNum, "numbers", true, "Include numbers")
	flag.BoolVar(&includeSym, "symbols", true, "Include symbols")
	flag.BoolVar(&includeUpp, "uppercase", true, "Include uppercase letters")
	flag.BoolVar(&includeLow, "lowercase", true, "Include lowercase letters")
}

func generatePassword() (string, error) {
	var chars string

	if includeLow {
		chars += lowerChars
	}
	if includeUpp {
		chars += upperChars
	}
	if includeNum {
		chars += numChars
	}
	if includeSym {
		chars += symChars
	}

	password := make([]byte, length)
	max := big.NewInt(int64(len(chars)))

	for i := 0; i < length; i++ {
		random, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		password[i] = chars[random.Int64()]
	}

	return string(password), nil
}

func main() {
	flag.Parse()

	password, err := generatePassword()
	if err != nil {
		fmt.Println("Error generating password:", err)
		return
	}
	fmt.Println("Generated Password:", password)
}

