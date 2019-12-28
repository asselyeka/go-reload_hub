package piscine

func AtoiBase(s string, base string) int {
	result := 0
	if CheckBaseAtoi(base) == false || LenAtoi(base) < 2 {
		result = 0
	} else {
		for indS, symS := range s {
			index := 0
			for indB, symB := range base {
				if symS == symB {
					index = indB
				}
			}
			result += index * RecursivePower(LenAtoi(base), LenAtoi(s)-1-indS)
		}
	}
	return result
}

func CheckBaseAtoi(base string) bool {
	baseAsRune := []rune(base)
	for index, letter := range baseAsRune {
		if letter == '-' || letter == '+' {
			return false
		}
		for i := index + 1; i < LenAtoi(base); i++ {
			if letter == baseAsRune[i] {
				return false
			}
		}
	}
	return true
}

func LenAtoi(str string) int {
	count := 0
	a := []rune(str)
	for index := range a {
		count = index + 1
	}
	return count
}
