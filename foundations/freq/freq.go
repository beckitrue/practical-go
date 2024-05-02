package main

// find the word frequency in a text file and print the result

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var filename string = "sherlock.txt"
	// open the file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	// create a map to store the word frequency
	freq := make(map[string]int)

	// create a buffer to store the word
	word := make([]byte, 0)

	// read the file byte by byte
	for {
		b := make([]byte, 1)
		// read the byte
		_, err := file.Read(b)
		if err != nil {
			break
		}
		//fmt.Printf("byte: %v\n", b)
		// skip the byte if it is not a letter
		if !isLetter(b[0]) {
			// append word to frequency map and increment the count
			if len(word) > 0 {
				updateFreq(freq, word)
			}
			// reset the word buffer
			word = make([]byte, 0)
		} else {
			// append the byte to the word buffer
			word = append(word, b[0])
		}
	}
	fmt.Printf("word frequency: %v\n", freq)
}

func isLetter(b byte) bool {
	return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z'
}

func updateFreq(freq map[string]int, word []byte) {
	// convert the word to lowercase
	w := strings.ToLower(string(word))
	// increment the count
	freq[w]++
}
