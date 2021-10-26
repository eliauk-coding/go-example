package main

import "fmt"

// As in all languages in the C family, everything in Go is passed by value
func main() {
	s1 := make([]int, 0, 10)
	var appendFunc = func(s []int) {
		s = append(s, 10, 20, 30)
		fmt.Println(s)
	}
	fmt.Println(s1)
	appendFunc(s1)
	fmt.Println(s1)
	fmt.Println(s1[:10])
	fmt.Println(s1[:])
}
