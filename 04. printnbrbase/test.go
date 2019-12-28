package main

import (
	"github.com/01-edu/z01"
	student ".."
)

func main() {
	student.PrintNbrBase(125152, "0123456789")
	z01.PrintRune('\n')
	student.PrintNbrBase(-125418461554641354354, "01")
	z01.PrintRune('\n')
	student.PrintNbrBase(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	student.PrintNbrBase(-125, "choumi")
	z01.PrintRune('\n')
	student.PrintNbrBase(125, "aa")
	z01.PrintRune('\n')
}