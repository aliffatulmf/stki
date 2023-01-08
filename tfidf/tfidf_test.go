package tfidf_test

import (
	"aliffatulmf/stki/stopword"
	"aliffatulmf/stki/tfidf"
	"aliffatulmf/stki/word"
	"testing"

	"github.com/stretchr/testify/assert"
)

var keyword = "pengetahuan logistik"

func newTfidf(t *testing.T) *tfidf.TFIDF {
	dict, err := stopword.OpenStopwordFile("../stopword/stopword.tala.txt")
	dict.Remove("tahu")
	assert.NoError(t, err)

	corpus, err := word.ReadJSON("../data_short.json", word.Cleaned)
	assert.NoError(t, err)

	n := tfidf.New(corpus, dict)
	n.SetKeywords(keyword)
	return n
}

func TestTermFrequency(t *testing.T) {
	n := newTfidf(t)

	// berdasarkan dokumen
	// https://drive.google.com/file/d/0B_FXUwD6KqsmeFFvYXRhRXZyZVk/view
	freq := map[string]int{
		"d1": 3,
		"d2": 2,
		"d3": 4,
	}

	tf := n.TermFrequency()
	for _, docs := range tf {
		switch docs.Document {
		case "d1":
			assert.Equal(t, freq[docs.Document], len(docs.TF))
		case "d2":
			assert.Equal(t, freq[docs.Document], len(docs.TF))
		case "d3":
			assert.Equal(t, freq[docs.Document], len(docs.TF))
		}
	}
}
