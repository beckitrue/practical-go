package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	banner("Go", 25)
	banner("G\u2318", 25)

	s := "G\u2318"
	fmt.Println("len:", len(s))
	// code point = rune ~= unicode character
	for i, r := range s {
		fmt.Println(i, r)

		if i == 1 {
			fmt.Printf("%c of type %T\n", r, r)
		}
	}

	// byte (uint8)
	// rune (int32)

	b := s[0]
	fmt.Printf("%c of type %T\n", b, b)

	// define a list of strings
	texts := []string{
		"bob",
		"taco",
		"racecar",
		"taco cat",
		"level",
		"a",
	}

	// loop through the list of strings

	for _, text := range texts {
		fmt.Printf("text: %#v, reverse: %#v, palindrome: %t\n", text, reverse(text), isPalindrome(text))
	}

}

func banner(text string, width int) {
	padding := (width - utf8.RuneCountInString(text)) / 3
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

// func that reverses a text string
func reverse(text string) string {
	// get length of text
	len := utf8.RuneCountInString(text)
	// check if text is empty
	if len == 0 {
		return ""
	}
	// convert text string to a slice of runes
	runes := make([]rune, len)
	for i, r := range text {
		runes[len-i-1] = r
	}
	// convert slice of runes to a string
	return string(runes)
}

// func that checks if a text string is a palindrome
func isPalindrome(text string) bool {
	// get length of text
	len := utf8.RuneCountInString(text)
	rs := []rune(text)
	for i := 0; i < len/2; i++ {
		if rs[i] != rs[len-i-1] {
			return false
		}
	}
	return true
}
