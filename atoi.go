package student

func Atoi(s string) int {
	num := 0
	if !CheckStr(s) {
		return 0
	}
	strAsRune := []rune(s)
	minus := 1
	if strAsRune[0] == '+' || strAsRune[0] == '-' {
		if strAsRune[0] == '-' {
			minus = -1
		}
		strAsRune = strAsRune[1:]
	}
	for index, n := range strAsRune {
		if n != '0' {
			strAsRune = strAsRune[index:]
			for _, number := range strAsRune {
				counter := 0
				for i := '0'; i < number; i++ {
					counter++
				}
				//min -9,223,372,036,854,775,808; max 9,223,372,036,854,775,807
				if num > 922337203685477580 {
					return 0
				} else if num == 922337203685477580 {
					if minus == 1 {
						if counter > 7 {
							return 0
						}
					} else {
						if counter > 8 {
							return 0
						}
					}
					num = 10*num + counter
				} else {
					num = 10*num + counter
				}
			}
			break
		}
	}
	return num * minus
}

func CheckStr(s string) bool {
	if s == "" {
		return false
	}
	strAsRune := []rune(s)
	if strAsRune[0] == '+' || strAsRune[0] == '-' {
		strAsRune = strAsRune[1:]
	}
	if string(strAsRune) == "" {
		return false
	}
	numberLengh := 0
	for _, n := range strAsRune {
		if n < '0' || n > '9' {
			return false
		}
		numberLengh++
	}
	//min -9,223,372,036,854,775,808; max 9,223,372,036,854,775,807
	if numberLengh > 19 {
		return false
	}
	return true
}
