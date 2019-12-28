package piscine

func SplitWhiteSpaces(str string) []string {
	strAsRune := []rune(str)
	len := 0
	size := 0
	for i := range strAsRune {
		len = i + 1
	}
	if len > 0 && IsSeparator(strAsRune[0]) == false {
		size++
		for i := 0; i < len; i++ {
			if IsSeparator(strAsRune[i]) == true {
				if i+1 < len {
					if IsSeparator(strAsRune[i+1]) == false {
						size++
					}
				}
			}
		}
	}
	answer := make([]string, size)
	word := ""
	index := 0
	if len > 0 && IsSeparator(strAsRune[0]) == false {
		for i := 0; i < len; i++ {
			if IsSeparator(strAsRune[i]) == true {
				if i+1 < len {
					if IsSeparator(strAsRune[i+1]) == false {
						answer[index] = word
						index++
						word = ""
					}
				}
			} else {
				word += string(strAsRune[i])
			}
		}
		answer[size-1] = word
	}
	return answer
}

func IsSeparator(r rune) bool {
	if r == ' ' || r == '\n' || r == '\t' {
		return true
	}
	return false
}
