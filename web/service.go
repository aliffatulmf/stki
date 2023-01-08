package web

import (
	"aliffatulmf/stki/model"
	"aliffatulmf/stki/stopword"
	"aliffatulmf/stki/tfidf"
	"errors"
	"fmt"
	"strings"

	"github.com/RadhiFadlillah/go-sastrawi"
)

type Service struct {
	Repository *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo}
}

func (s *Service) Create(model *model.Corpus) error {
	if err := s.Repository.Create(model); err != nil {
		return errors.Unwrap(err)
	}
	return nil
}

func (s *Service) Finds(keyword ...string) ([]model.Corpus, error) {
	corpus, err := s.Repository.Finds()
	if err != nil {
		return corpus, errors.Unwrap(err)
	}

	kk := strings.Join(keyword, " ")
	if len(kk) < 1 {
		return corpus, nil
	} else {
		var articles []model.Corpus
		rank, err := UseTFIDF(corpus, kk)
		if err != nil {
			return articles, fmt.Errorf("ServiceRank: %w", err)
		}

		for _, doc := range rank {
			record, err := s.Repository.Find("document = ?", doc.Document)
			if err != nil {
				return corpus, fmt.Errorf("ServiceFindArticle: %w", err)
			}

			articles = append(articles, record)
		}

		return articles, nil
	}
}

func UseTFIDF(corpus []model.Corpus, keyword string) ([]tfidf.Rank, error) {
	var rank []tfidf.Rank

	dict, err := stopword.OpenStopwordFile("stopword/stopword.tala.txt")
	if err != nil {
		return rank, err
	}
	dict.Remove("tahu")

	m := tfidf.New(corpus, dict)
	m.SetKeywords(sastrawi.Tokenize(keyword)...)

	tf := m.TermFrequency()
	idf := m.InverseDocumentFrequency()
	weight := m.Weight(tf, idf)
	rank = m.Rank(weight)

	return rank, nil
}
