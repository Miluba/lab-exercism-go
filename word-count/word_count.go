package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	f := make(map[string]int)
	re := regexp.MustCompile("\\d+|\\w+[']?\\w+")
	words := re.FindAllString(phrase, -1)
	for _, w := range words {
		w = strings.ToLower(w)
		i := f[w]
		i++
		f[w] = i
	}
	return f
}
