package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	return NbrBase(AtoiBase(nbr, baseFrom), baseTo)
}

func NbrBase(nbr int, base string) string {
	l := LenLen(base)
	if CheckBase(base) == false || l < 2 {
		return "0"
	} else if nbr > 2000000000 || nbr < -2000000000 {
		return string(nbr)
	} else {
		str := ""
		for nbr >= l {
			str += string(base[nbr%l])
			nbr = nbr / l
		}
		str += string(base[nbr])
		return RevStr(str)
	}
}
