package main

import (
	"flag"
	"fmt"

	"math/rand"
)

var (
	size           = flag.Int("s", 24, "Length of the random string")
	includeSpecial = flag.Bool("x", false, "Include special characters in the random string")
)

func main() {
	flag.Parse()
	fmt.Print(generateRandomString(*size, *includeSpecial))
}

func generateRandomString(size int, includeSpecial bool) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	specials := "!@#$%^&*()-_=+[]{}<>?/|"

	charset := letters + digits
	if includeSpecial {
		charset += specials
	}

	res := make([]byte, size)
	for i := 0; i < size; i++ {
		res[i] = charset[rand.Intn(len(charset))]
	}
	return string(res)
}
