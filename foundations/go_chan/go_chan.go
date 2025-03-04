package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("goroutine")
	fmt.Println("main")

	for i := 0; i < 3; i++ { // all go routines use "i" for the loop, so we need to pass it as a parameter
		i := i // shadowing the variable
		go func() {
			fmt.Println(i)
		}()

		/*=
		go func(n int) {
			fmt.Println(n)
		}(i) // go evaluates the parameter before calling the function
		*/

	}

	time.Sleep(10 * time.Millisecond)

	shadowExample()

	ch := make(chan string)
	go func() {
		ch <- "Hello" // send to channel
	}()
	msg := <-ch // receive from channel
	fmt.Println(msg)

	go func() {
		for i := 0; i < 3; i++ {
			msg := fmt.Sprintf("msg %d", i)
			ch <- msg
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println("got:", msg)
	}

	msg = <-ch // channel is closed, so it will return the zero value of the type
	fmt.Printf("closed: %#v\n", msg)

	msg, ok := <-ch // ch is closed
	fmt.Printf("closed: %#v, (ok: %v)\n", msg, ok)

	values := []int{15, 8, 42, 16, 4, 23}
	fmt.Println(sleepSort(values))

}

/* For every value "n" in values, spin a goroutine that will
- sleep for "n" milliseconds
- Send "n" to the channel

In the function body, collect values from the channel to a slice and return it
*/

func sleepSort(values []int) []int {
	ch := make(chan int)
	var sorted []int

	for _, n := range values {
		go func(n int) {
			time.Sleep(time.Duration(n) * time.Millisecond)
			ch <- n
		}(n)
	}

	for i := 0; i < len(values); i++ {
		sorted = append(sorted, <-ch)
	}

	return sorted
}

func shadowExample() {
	n := 7
	{
		n := 2 // from here to } this is "n
		fmt.Println("inner", n)
	}
	fmt.Println("outer", n)
}
