package word

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/RadhiFadlillah/go-sastrawi"
)

type Stemmed []string

func OpenStopwordFile(s string) sastrawi.Dictionary {
	var words []string

	readFile, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		words = append(words, fileScanner.Text())
	}

	return sastrawi.NewDictionary(words...)
}

func Stemmer(sliceList string, stopword sastrawi.Dictionary) string {
	var list []string

	dict := sastrawi.DefaultDictionary()
	stemmer := sastrawi.NewStemmer(dict)
	dict.Remove("tahu")

	for _, word := range sastrawi.Tokenize(sliceList) {
		if stopword.Contains(word) {
			continue
		}
		list = append(list, stemmer.Stem(word))
	}

	return strings.Join(list, " ")
}
