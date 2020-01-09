package main

import (
	"fmt"

	student ".."
)

func main() {
	fmt.Println(student.RecursivePower((-7), (-2)))
	fmt.Println(student.RecursivePower((-8), (-7)))
	fmt.Println(student.RecursivePower((4), (8)))
	fmt.Println(student.RecursivePower((1), (3)))
	fmt.Println(student.RecursivePower((-1), (1)))
	fmt.Println(student.RecursivePower((-6), (5)))
	fmt.Println(student.RecursivePower((9223372036854775807), (2)))
}
