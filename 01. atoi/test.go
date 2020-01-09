package main

import (
	"fmt"

	student ".."
)

func main() {
	s := ""
	s2 := "-"
	s3 := "--123"
	s4 := "1"
	s5 := "-3"
	s6 := "8292"
	s7 := "9223372036854775807"
	s8 := "-9223372036854775808"

	n := student.Atoi(s)
	n2 := student.Atoi(s2)
	n3 := student.Atoi(s3)
	n4 := student.Atoi(s4)
	n5 := student.Atoi(s5)
	n6 := student.Atoi(s6)
	n7 := student.Atoi(s7)
	n8 := student.Atoi(s8)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Println(n5)
	fmt.Println(n6)
	fmt.Println(n7)
	fmt.Println(n8)
}
