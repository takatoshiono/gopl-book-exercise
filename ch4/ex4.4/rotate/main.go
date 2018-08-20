package main

import "fmt"

func main() {
	ints := []int{1, 2, 3, 4, 5}
	ints = rotate(ints, 2)
	fmt.Println(ints) // expect [3 4 5 1 2]
}

// Rotate s left by n positions.
func rotate(s []int, n int) []int {
	s = append(s, s[0:n]...)
	return s[n:]
}
