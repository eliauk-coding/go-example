package main

import (
	"fmt"
	"reflect"
)

func testSlice(sLen, sCap int, sAppend []string) {
	s := make([]string, sLen, sCap)
	fmt.Printf("slice len = %d cap = %d\n", len(s), cap(s))
	s = append(s, sAppend...)
	fmt.Printf("slice grow len = %d cap = %d\n\n", len(s), cap(s))
}

func main() {
	// 2 + 2 <= 4 No need to expand
	//slice len = 2 cap = 4
	//slice len = 4 cap = 4
	testSlice(2, 4, []string{"1", "2"})
	// 2 + 2 > 3 && 3 * 2 > 2 + 2 Double the capacity
	//slice len = 2 cap = 3
	//slice len = 4 cap = 6
	testSlice(2, 3, []string{"1", "2"})
	// 2 + 5 > 3 * 2 Expand to cap
	//slice len = 2 cap = 3
	//slice len = 7 cap = 7
	testSlice(2, 3, []string{"1", "2", "4", "5", "6"})
	// Capacity is greater than 1024，Expanded to 1.15 times
	// slice len = 1024 cap = 1024
	// slice len = 1026 cap = 1280
	testSlice(1024, 1024, []string{"1", "2"})
	//
	// slice len = 1023 cap = 1024
	// slice len = 1026 cap = 1360
	testSlice(1024, 1025, []string{"1", "2", "3"})
	var str string
	fmt.Println(reflect.TypeOf(str).Size())
	var num int
	fmt.Println(reflect.TypeOf(num).Size())
}
