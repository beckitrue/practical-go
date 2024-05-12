package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("goroutine")
	fmt.Println("main")

	for i := 0; i < 3; i++ { // all go routines use "i" for the loop, so we need to pass it as a parameter
		go func(n int) {
			fmt.Println(n)
		}(i) // go evaluates the parameter before calling the function
	}

	time.Sleep(10 * time.Millisecond)

}
