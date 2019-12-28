package piscine

func ActiveBits(n int) uint {
	counter := 0
	for n != 0 {
		if n%2 == 1 {
			counter++
		}
		n /= 2
	}
	return uint(counter)
}
