package main

import (
	"flag"
	"log"

	"aliffatulmf/stki/tfidf"
	"aliffatulmf/stki/word"
)

type TfIdf interface {
	TermFrequency() []tfidf.TField
	InverseDocumentFrequency(tf []tfidf.TField) map[string]float64
	Weight(tf []tfidf.TField, idf map[string]float64) []tfidf.TWeight
}

func main() {
	sw := flag.String("stopword", "stopword/stopword.tala.txt", "set stopword location")
	cp := flag.String("corpus", "data_short.json", "set corpus location")
	kw := flag.String("keyword", "", "input keyword")
	swg := flag.Bool("showScore", false, "to display the score")
	flag.Parse()

	stopword := word.OpenStopwordFile(*sw)
	corpus, err := word.ReadJSON(*cp)
	if err != nil {
		log.Fatal(err)
	}

	var n TfIdf = tfidf.New(corpus,stopword)
	tf := n.TermFrequency()
	idf := n.InverseDocumentFrequency(tf)
	w := n.Weight(tf, idf)

	if len(*kw) > 0 {
		tfidf.Find(w, *swg, *kw)
	}
}
