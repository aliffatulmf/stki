package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"aliffatulmf/stki/stopword"
	"aliffatulmf/stki/tfidf"
	"aliffatulmf/stki/word"
)

type TfIdf interface {
	TermFrequency() []tfidf.TField
	InverseDocumentFrequency() map[string]float64
	//Weight(tf []tfidf.TField, idf map[string]float64) []tfidf.TWeight
}

func main() {
	sw := flag.String("stopword", "stopword/stopword.tala.txt", "set stopword location")
	cp := flag.String("corpus", "data_short.json", "set corpus location")
	kw := flag.String("keyword", "", "input keyword")
	swg := flag.Bool("showScore", false, "to display the score")
	flag.Parse()

	dict, err := stopword.OpenStopwordFile(*sw)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	corpus, err := word.ReadJSON(*cp)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var n TfIdf = tfidf.New(corpus, dict)
	tf := n.TermFrequency()
	idf := n.InverseDocumentFrequency()

	w := tfidf.Weight(tf, idf)

	if len(*kw) > 0 {
		keyword := strings.Fields(*kw)
		tfidf.Find(w, *swg, keyword...)
	}
}
