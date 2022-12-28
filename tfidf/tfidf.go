package tfidf

import (
	"fmt"
	"math"
	"strings"

	"aliffatulmf/stki/factory"
	"aliffatulmf/stki/model"
	"aliffatulmf/stki/word"

	"github.com/RadhiFadlillah/go-sastrawi"
)

type TFIDF struct {
	Keyword   []string
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
		stemmed := factory.Stemmer(doc.Body, stopword)
		list = append(list, stemmed...)
	}

	// remove duplicate value
	token = append(token, word.Unique(list)...)

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

func (t *TFIDF) SetKeywords(keyword ...string) error {
	if len(keyword) > 0 {
		t.Keyword = keyword
		return nil
	}
	return fmt.Errorf("SetKeywords: %w", ErrNoInput)
}

func (t *TFIDF) Weight(tf []TField, idf map[string]float64) map[string]float64 {
	var list []TWeight

	// TF[index] * IDF
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

	weight := make(map[string]float64)

	for _, v := range list {
		for _, key := range t.Keyword {
			weight[v.Document] += v.Weight[key]
		}
	}

	return weight
}

func (t *TFIDF) Search(keyword ...string) ([]Documents, error) {
	if err := t.SetKeywords(keyword...); err != nil {
		return []Documents{}, err
	}

	tf := t.TermFrequency()
	idf := t.InverseDocumentFrequency()
	weight := t.Weight(tf, idf)

	return FindDocument(t.Documents, weight), nil
}
