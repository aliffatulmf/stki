package stopword

import (
	"bufio"
	"os"

	"github.com/RadhiFadlillah/go-sastrawi"
)

func OpenStopwordFile(s string) (sastrawi.Dictionary, error) {
	var words []string

	readFile, err := os.Open(s)
	if err != nil {
		return sastrawi.Dictionary{}, err
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		words = append(words, fileScanner.Text())
	}

	return sastrawi.NewDictionary(words...), nil
}
