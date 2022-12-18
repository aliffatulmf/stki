package word_test

import (
	"testing"

	"aliffatulmf/stki/model"
	"aliffatulmf/stki/word"

	"github.com/stretchr/testify/assert"
)

func TestTFIDF(t *testing.T) {
	stopword := word.OpenStopwordFile("../stopword/stopword.tala.txt")
	corpus, err := word.ReadJSON("../data_medium.json")
	assert.NoError(t, err)

	var corpus1 []model.Corpus


	for _, entry := range corpus {
		 corpus1 = append(corpus1, model.Corpus{
			ID:       entry.ID,
			Title:    word.Stemmer(entry.Title, stopword),
			Body:     word.Stemmer(entry.Body, stopword),
			Document: entry.Document,
		})
	}
}
