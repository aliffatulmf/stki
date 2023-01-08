package factory

import (
	"github.com/RadhiFadlillah/go-sastrawi"
)

// func AsyncStringFrequncy(str string, dict sastrawi.Dictionary, ch chan map[string]int) {
// 	list := make(map[string]int)
// 	token := Stemmer(str, dict)
//
// 	for _, word := range token {
// 		list[word] = strings.Count(str, word)
// 	}
//
// 	ch <- list
// }

func findInArray(str string, token []string) int {
	var count int

	for _, word := range token {
		if str == word {
			count += 1
		}
	}

	return count
}

func StringFrequency(str string, dict sastrawi.Dictionary) map[string]int {
	list := make(map[string]int)
	token := Stemmer(str, dict)

	for _, word := range token {
		list[word] = findInArray(word, token)
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
