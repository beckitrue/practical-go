package main

import (
	"fmt"
	"sort"
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

	s4 := appendInt(s3, 200)
	fmt.Printf("s4 = %#v\n", s4)

	s2 = []int{10, 20, 30, 40, 50, 60, 70}

	s5 := appendInt(s2, 1000)
	fmt.Printf("s5 = %#v\n", s5)
	fmt.Printf("s5: len=%d, cap=%d\n", len(s5), cap(s5))

	fmt.Println(concat([]string{"A", "B"}, []string{"C", "D", "E"})) // [A B C D E]

	vs := []float64{2, 1, 3}
	fmt.Println(median(vs))

	vs = []float64{2, 1, 4, 3}
	fmt.Println(median(vs))

	fmt.Println(median(nil))

}

func median(values []float64) (float64, error) {
	// calculate the median of a slice of float64 values
	// values is the slice of float64 values
	// returns the median value
	// check for empty slice
	if len(values) == 0 {
		return 0, fmt.Errorf("empty slice")
	}
	s := make([]float64, len(values))
	copy(s, values)
	sort.Float64s(s)
	mid := len(s) / 2
	if mid%2 == 1 {
		return s[mid], nil
	}
	v := (s[mid-1] + s[mid]) / 2
	return v, nil
}

func concat(s1 []string, s2 []string) []string {
	// concatenate two slices
	// s1 and s2 are the slices to concatenate
	s := make([]string, len(s1)+len(s2))
	copy(s, s1)
	copy(s[len(s1):], s2)
	return s
}

func appendInt(s []int, v int) []int {
	// append an integer to the end of a slice
	// s is the slice and v is the value to append
	i := len(s)

	if len(s) < cap(s) {
		s = s[:len(s)+1]
	} else {
		fmt.Println("slice is full, need to allocate a new slice")
		newSlice := make([]int, 2*len(s)+1)
		copy(newSlice, s)
		s = newSlice
	}

	s[i] = v
	return s
}
