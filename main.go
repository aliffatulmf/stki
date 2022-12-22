package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"aliffatulmf/stki/stopword"
	"aliffatulmf/stki/tfidf"
	"aliffatulmf/stki/word"
)

func main() {
	sw := flag.String("stopword", "stopword/stopword.tala.txt", "set stopword location")
	cp := flag.String("corpus", "data_short.json", "set corpus location")
	kw := flag.String("keyword", "", "input keyword")
	flag.Parse()

	dict, err := stopword.OpenStopwordFile(*sw)
	dict.Remove("tahu")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	corpus, err := word.ReadJSON(*cp, word.Cleaned)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	s := tfidf.New(corpus, dict)
	result, err := s.Search(strings.Fields(*kw)...)
	if err != nil {
		log.Fatal(err.Error())
	}

	var top float64
	var docKey string
	for _, val := range result {
		if top < val.Score {
			docKey = val.Document
			top = val.Score
		}
	}

	documents, err := word.ReadJSON(*cp, word.Raw)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, val := range documents {
		if val.Document == docKey {
			fmt.Printf("[Title:\t%s]\n", val.Title)
			fmt.Printf("Content:\t%s\n", val.Body)
		}
	}
}
