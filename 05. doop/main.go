package main

import (
	"os"

	student ".."
	"github.com/01-edu/z01"
)

func Div(a, b int) int {
	return a / b
}
func Mult(a, b int) int {
	return a * b
}
func Mod(a, b int) int {
	return a % b
}
func Plus(a, b int) int {
	return a + b
}
func Minus(a, b int) int {
	return a - b
}

type Calculator struct {
	symbol  rune
	arrFunc func(int, int) int
}

func main() {
	arg := os.Args
	arguments := []string(arg[1:])
	lenArg := 0
	arrCalculator := []Calculator{
		{'/', Div},
		{'*', Mult},
		{'%', Mod},
		{'+', Plus},
		{'-', Minus},
	}
	for range arguments {
		lenArg++
	}
	result := 0
	errSym := -1
	if lenArg == 3 {
		if IsNumericD(arguments[0]) && IsNumericD(arguments[2]) {
			for i, fun := range arrCalculator {
				if arguments[1] == string(fun.symbol) {
					errSym++
					if student.Atoi(arguments[2]) == 0 && i == 0 {
						Print("No division by 0\n")
						break
					} else if student.Atoi(arguments[2]) == 0 && i == 2 {
						Print("No Modulo by 0\n")
						break
					} else {
						if student.Atoi(arguments[0]) == 0 && LenStr(arguments[0]) > 1 ||
							student.Atoi(arguments[2]) == 0 && LenStr(arguments[2]) > 1 {
							result = 0
							break
						}
						result = fun.arrFunc(student.Atoi(arguments[0]), student.Atoi(arguments[2]))

						// check result
						var position int
						if fun.symbol == '+' {
							if student.Atoi(arguments[0]) > 0 && student.Atoi(arguments[2]) > 0 && result < 0 {
								result = 0
								break
							}
							if student.Atoi(arguments[0]) < 0 && student.Atoi(arguments[2]) < 0 && result > 0 {
								result = 0
								break
							}
						} else if fun.symbol == '-' {
							if student.Atoi(arguments[0]) > 0 && student.Atoi(arguments[2]) < 0 && result < 0 {
								result = 0
								break
							}
							if student.Atoi(arguments[0]) < 0 && student.Atoi(arguments[2]) > 0 && result > 0 {
								result = 0
								break
							}
						} else if fun.symbol == '*' {
							position = 0
							if student.Atoi(arguments[0]) != arrCalculator[position].arrFunc(result, student.Atoi(arguments[2])) {
								result = 0
								break
							}
						} else if fun.symbol == '/' {
							position = 1
							if student.Atoi(arguments[0]) != arrCalculator[position].arrFunc(result, student.Atoi(arguments[2])) {
								result = 0
								break
							}
						}
						break
					}
				}
			}
			if errSym == -1 {
				Print("0\n")
				return
			}
		} else {
			Print("0\n")
			return
		}
		PrintInt(result)
		z01.PrintRune('\n')
	}
}

func Print(s string) {
	for _, l := range []rune(s) {
		z01.PrintRune(l)
	}
}

func PrintInt(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	minus := ""
	if n < 0 {
		minus += "-"
		n *= -1
	}
	result := ""
	for n != 0 {
		num := n % 10
		counter := 0
		elem := '0'
		for i := '0'; i <= '9'; i++ {
			if counter == num {
				elem = i
				break
			}
			counter++
		}
		result = string(elem) + result
		n /= 10
	}
	if minus == "-" {
		result = minus + result
	}
	Print(result)
}

func IsNumericD(str string) bool {
	sAsRune := []rune(str)
	counter := 0
	strLen := LenStr(str)
	for _, letter := range sAsRune {
		if letter >= '0' && letter <= '9' {
			counter++
		}
	}
	if sAsRune[0] == '-' {
		counter++
	}
	if counter == strLen {
		return true
	}
	return false
}

func LenStr(s string) int {
	strLen := 0
	for range []rune(s) {
		strLen++
	}
	return strLen
}
