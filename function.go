package main

import (
	"fmt"
	// "errors"
)
func main()  {
	x := 5
	y := 7
	s := 0
	Sum(x, y, &s)
	// 调用闭包函数
	s2 := Sum5(x, y)
	fmt.Print(s, s2())
}

// Pointer
func Sum(x int, y int, z *int) {
	*z = x + y
}
// return a value
func Sum2(x int, y int) int {
	return x + y
}
// return mutilple value
func Sum3(x int, y int) (int, int) {
	return x+y, x-y
}

// return a function 
func Sum4() func(x int, y int) int {
	return func(x int, y int) int {
			return x + y
	}
}
// closure
func Sum5(x int, y int) func() int {
	return func() int {
			return x + y
	}
}