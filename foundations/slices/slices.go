package main

import (
	"fmt"
)

func main() {
	var s []int
	fmt.Println("len", len(s))
	if s == nil {
		fmt.Println("nil slice")
	}

	s2 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("s2 = %#v\n", s2)
	fmt.Printf("s2: len=%d, cap=%d\n", len(s2), cap(s2))

	s3 := s2[1:4]
	fmt.Printf("s3 = %#v\n", s3)
	fmt.Printf("s3: len=%d, cap=%d\n", len(s3), cap(s3))

	s3 = append(s3, 100)
	fmt.Printf("s3(append) = %#v\n", s3)
	fmt.Printf("s2(append) = %#v\n", s2) // s2 is changed
	fmt.Printf("s2: len=%d, cap=%d\n", len(s2), cap(s2))
	fmt.Printf("s3: len=%d, cap=%d\n", len(s3), cap(s3))

}
