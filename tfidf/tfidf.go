package tfidf

import (
	"math"
	"strings"

	"aliffatulmf/stki/factory"
	"aliffatulmf/stki/model"
	"aliffatulmf/stki/word"
	"github.com/RadhiFadlillah/go-sastrawi"
)

type TFIDF struct {
	Keyword   string
	Token     []string
	Documents []model.Corpus
	Stopword  sastrawi.Dictionary
}

type TField struct {
	Document string
	TF       map[string]int
}

func New(corpus []model.Corpus, stopword sastrawi.Dictionary) *TFIDF {
	var list []string
	var token []string

	// remove special character
	for _, doc := range corpus {
		stemmed := factory.Stemmer(doc.Body, stopword)
		for _, val := range stemmed {
			list = append(list, val)
		}
	}

	// remove duplicate value
	for _, val := range word.Unique(list) {
		token = append(token, val)
	}

	return &TFIDF{
		Documents: corpus,
		Stopword:  stopword,
		Token:     token,
	}
}

func (t *TFIDF) TermFrequency() []TField {
	var tfield []TField

	for _, doc := range t.Documents {
		freq := factory.StringFrequency(doc.Body, t.Stopword)
		tfield = append(tfield, TField{
			Document: doc.Document,
			TF:       freq,
		})
	}

	return tfield
}

// count the number of documents containing the word (df)
func (t *TFIDF) manyDocs(word string) float64 {
	var total float64

	for _, doc := range t.Documents {
		if strings.Contains(doc.Body, word) {
			total += 1
		}
	}

	return total
}

func (t *TFIDF) InverseDocumentFrequency() map[string]float64 {
	idf := make(map[string]float64)
	docLength := float64(len(t.Documents))

	for _, key := range t.Token {
		docs := t.manyDocs(key)
		// idf = log(document length (dl) / the number of documents containing the word (list))
		idf[key] = math.Log10(docLength / docs)
	}

	return idf
}
