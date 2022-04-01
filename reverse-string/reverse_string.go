package reverse

func Reverse(input string) string {
	ra := []rune(input)
	j := len(ra)
	for i := 0; i < len(ra)/2; i++ {
		j--
		ra[i], ra[j] = ra[j], ra[i]
	}
	return string(ra)
}
