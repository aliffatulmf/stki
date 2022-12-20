package factory

import (
	"strings"

	"github.com/RadhiFadlillah/go-sastrawi"
)

func StringFrequency(str string, dict sastrawi.Dictionary) map[string]int {
	list := make(map[string]int)
	token := Stemmer(str, dict)

	for _, word := range token {
		list[word] = strings.Count(str, word)
	}

	return list
}

func Stemmer(sliceList string, dict sastrawi.Dictionary) []string {
	var list []string

	stemmer := sastrawi.NewStemmer(dict)
	for _, word := range sastrawi.Tokenize(sliceList) {
		if dict.Contains(word) {
			continue
		}
		list = append(list, stemmer.Stem(word))
	}

	return list
}

