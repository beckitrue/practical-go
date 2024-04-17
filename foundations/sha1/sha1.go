package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	sig, err := sha1sum("/home/becki/Downloads/nvim-linux64.tar.gz")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Println(sig)
}

// function to open a zip file and
// calculate the sha1 hash of the file
// and return the hash
func sha1sum(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	r, err := gzip.NewReader(file)
	if err != nil {
		return "", err
	}

	w := sha1.New()

	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}

	// get signature
	signature := w.Sum(nil)

	// return the signature as a string
	return fmt.Sprintf("%x", signature), nil
}
