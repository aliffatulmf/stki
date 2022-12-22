package tfidf

import "aliffatulmf/stki/model"

type Documents struct {
	model.Corpus
	Score float64
}

func FindDocument(corpus []model.Corpus, score map[string]float64) []Documents {
	var docs []Documents

	for _, val := range corpus {
		docs = append(docs, Documents{
			Corpus: val,
			Score:  score[val.Document],
		})
	}
	return docs
}
