package prime

func Factors(n int64) []int64 {
	pf := []int64{}
	var f int64 = 2
	for {
		if n == 1 {
			break
		}
		if n%f == 0 {
			pf = append(pf, f)
			n /= f
		} else {
			f++
		}
	}
	return pf
}
