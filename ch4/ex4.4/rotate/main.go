package main

import (
	"fmt"
)

func main() {
	ints := []int{1, 2, 3, 4, 5}
	ints = rotate(ints, 2)
	fmt.Println(ints) // expect [3 4 5 1 2]
	rotate2(ints, 2)
	fmt.Println(ints) // expect [5 1 2 3 4]
}

// Rotate s left by n positions.
func rotate(s []int, n int) []int {
	s = append(s, s[0:n]...)
	return s[n:]
}

func rotate2(s []int, n int) {
	n = n % len(s)
	for start, count := 0, 0; count < len(s); start++ {
		prev := s[start]
		current := start
		for {
			next := (current - n) % len(s)
			if next < 0 {
				next = len(s) + next
			}
			temp := s[next]
			s[next] = prev
			prev = temp
			current = next
			count++
			if current == start {
				break
			}
		}
	}
}
