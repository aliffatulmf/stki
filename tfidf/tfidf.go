package tfidf

import (
	"fmt"
	"math"
	"sort"
	"strings"

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

type TWeight struct {
	Document string
	Weight   map[string]float64
}

func New(corpus []model.Corpus, stopword sastrawi.Dictionary) *TFIDF {
	var list []string
	var token []string

	// remove special character
	for _, doc := range corpus {
		stemmed := word.Stemmer(doc.Body, stopword)
		for _, val := range strings.Fields(stemmed) {
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

func (t *TFIDF) SetKeyword(kk string) *TFIDF {
	t.Keyword = kk
	return t
}

// count the same word using the tokenized word
func (t *TFIDF) keywordInDocument(s string) map[string]int {
	list := make(map[string]int)

	for _, keyword := range t.Token {
		list[keyword] = strings.Count(s, keyword)
	}
	return list
}

// count the number of the same words
func (t *TFIDF) SumFloat64(tf map[string]int) float64 {
	var total float64

	for _, val := range tf {
		total += float64(val)
	}
	return total
}

func (t *TFIDF) TermFrequency() []TField {
	var tfield []TField

	for _, doc := range t.Documents {
		kid := t.keywordInDocument(doc.Body)
		tfield = append(tfield, TField{
			Document: doc.Document,
			TF:       kid,
		})
	}

	return tfield
}

func (t *TFIDF) countSimilar(word string, tf []TField) float64 {
	var total float64

	for _, entry := range tf {
		for key, val := range entry.TF {
			if strings.EqualFold(key, word) {
				total += float64(val)
			}
		}
	}
	return total
}

// count the number of documents containing the word (df)
func (t *TFIDF) notZero(word string, tf []TField) float64 {
	var total float64

	for _, doc := range t.Documents {
		if strings.Contains(doc.Body, word) {
			total += 1
		}
	}

	return total
}

func (t *TFIDF) InverseDocumentFrequency(tf []TField) map[string]float64 {
	list := make(map[string]float64)

	// count the number of words in the document
	for _, key := range t.Token {
		list[key] = t.notZero(key, tf)
	}

	idf := make(map[string]float64)
	docLength := float64(len(t.Documents))

	for key, val := range list {
		for _, kk := range t.Token {
			if strings.EqualFold(key, kk) {
				// idf = log(document length (dl) / the number of documents containing the word (list))
				idf[key] = math.Log10(docLength / val)
			}
		}
	}

	return idf
}

func (t *TFIDF) Weight(tf []TField, idf map[string]float64) []TWeight {
	var list []TWeight

	for _, f := range tf {
		w := make(map[string]float64)

		for key, val := range f.TF {
			w[key] = idf[key] * float64(val)
		}

		list = append(list, TWeight{
			Document: f.Document,
			Weight:   w,
		})
	}
	return list
}

func Find(w []TWeight, showNumber bool, keyword ...string) {
	list := make(map[string]float64)

	for _, v := range w {
		for _, key := range keyword {
			list[v.Document] += v.Weight[key]
		}
	}

	var rankNumber []float64
	for _, val := range list {
		rankNumber = append(rankNumber, val)
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(rankNumber)))

	var rankString []string
	for _, rank := range rankNumber {
		for key, val := range list {
			if rank == val {
				rankString = append(rankString, key)
			}
		}
	}

	if showNumber {
		for key, val := range list {
			fmt.Printf("Dokumen %s dengan skor %f\n", key, val)
		}
	} else {
		rankString = word.Unique(rankString)
		for idx, val := range rankString {
			fmt.Printf("Ranking %d Dokumen %s\n", idx + 1, val)
		}
	}

}
