package main

import "fmt"

func main() {
	ints := [5]int{1, 2, 3, 4, 5}
	reverse(&ints)
	fmt.Println(ints)
}

func reverse(ptr *[5]int) {
	for i, j := 0, len(ptr)-1; i < j; i, j = i+1, j-1 {
		ptr[i], ptr[j] = ptr[j], ptr[i]
	}
}
