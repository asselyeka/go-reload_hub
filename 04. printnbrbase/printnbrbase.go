package student

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !CheckBase(base) || !ParseInt(nbr) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	result := ""
	baseLen := 0
	for range []rune(base) {
		baseLen++
	}
	if nbr < 0 {
		z01.PrintRune('-')
		nbr *= -1 
	}
	for nbr != 0 {
		index := nbr%baseLen
		for i, l := range []rune(base) {
			if i == index {
				result = string(l) + result
			}
		}
		nbr /= baseLen 
	}
	for _, l := range []rune(result){
		z01.PrintRune(l)
	}
}

func CheckBase(s string) bool {
	strLen := 0
	sRune := []rune(s)
	for _, l := range sRune {
		if l == '+' || l == '-' {
			return false
		}
		strLen++
	}
	if strLen < 2 {
		return false
	}
	for i, l := range sRune {
		for j := i+1; j < strLen; j++ {
			if l == sRune[j] {
				return false
			}
		}
	}
	return true
}

func ParseInt(n int) bool {
	if -9223372036854775808 <= n && n <= 9223372036854775807 {
		return true
	}
	return false
}