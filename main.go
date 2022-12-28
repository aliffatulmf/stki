package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"aliffatulmf/stki/stopword"
	"aliffatulmf/stki/tfidf"
	"aliffatulmf/stki/word"
)

var (
	cmdStopword string
	cmdCorpus   string
	cmdKeyword  string
)

func init() {
	flag.StringVar(&cmdStopword, "stopword", "stopword.tala.txt", "set stopword location")
	flag.StringVar(&cmdCorpus, "corpus", "data_short.json", "set corpus location")
	flag.StringVar(&cmdKeyword, "keyword", "", "input keyword")
	flag.Parse()
}

func callAll() ([]tfidf.Documents, error) {
	dict, err := stopword.OpenStopwordFile(cmdStopword)
	dict.Remove("tahu")
	if err != nil {
		return []tfidf.Documents{}, err
	}

	corpus, err := word.ReadJSON(cmdCorpus, word.Cleaned)
	if err != nil {
		return []tfidf.Documents{}, err
	}

	s := tfidf.New(corpus, dict)
	result, err := s.Search(strings.Fields(cmdKeyword)...)
	if err != nil {
		return []tfidf.Documents{}, err
	}

	return result, nil
}

func main() {
	docs, err := callAll()
	if err != nil {
		fmt.Println(errors.Unwrap(err))
		os.Exit(1)
	}

	var top float64
	var docKey string
	for _, val := range docs {
		if top < val.Score {
			docKey = val.Document
			top = val.Score
		}
	}

	documents, err := word.ReadJSON(cmdCorpus, word.Raw)
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
