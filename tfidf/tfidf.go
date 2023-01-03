package tfidf

import (
	"errors"
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

	// hapus spesial karakter
	for _, doc := range corpus {
		stemmed := factory.Stemmer(doc.Body, stopword)
		list = append(list, stemmed...)
	}

	// hapus kata duplikat
	token = append(token, word.Unique(list)...)

	return &TFIDF{
		Documents: word.CleanSpecialChar(corpus),
		Stopword:  stopword,
		Token:     token,
	}
}

// func (t *TFIDF) AsyncTermFrequency() []TField {
// 	var tfield []TField
//
// 	ch := make(chan map[string]int)
//
// 	for _, doc := range t.Documents {
// 		go factory.AsyncStringFrequncy(doc.Body, t.Stopword, ch)
//
// 		tfield = append(tfield, TField{
// 			Document: doc.Document,
// 			TF:       <-ch,
// 		})
// 	}
//
// 	return tfield
// }

func (t *TFIDF) TermFrequency() []TField {
	var tfield []TField

	for _, doc := range t.Documents {
		freq := factory.StringFrequency(doc.Body, t.Stopword)
		tfield = append(tfield, TField{
			Document: doc.Document,
			TF:       freq,
		})
	}

	kk := strings.Join(t.Keyword, " ")
	freq := factory.StringFrequency(kk, t.Stopword)
	tfield = append(tfield, TField{
		Document: "kk",
		TF:       freq,
	})

	return tfield
}

// manyDocs menghitung jumlah dokumen yang mengandung kata (df)
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

// SetKeywords menetapkan kata kunci
func (t *TFIDF) SetKeywords(keyword ...string) error {
	if len(keyword) < 1 {
		return errors.New("SetKeywords: no input")
	}

	for _, kk := range keyword {
		t.Keyword = append(t.Keyword, sastrawi.Tokenize(kk)...)
	}
	return nil
}

// Weight menghitung bobot dari tiap kata per dokumen
func (t *TFIDF) Weight(tf []TField, idf map[string]float64) []TWeight {
	var we []TWeight

	// TF[index] * IDF
	for _, f := range tf {
		w := make(map[string]float64)

		for key, val := range f.TF {
			w[key] = idf[key] * float64(val)
		}

		we = append(we, TWeight{
			Document: f.Document,
			Weight:   w,
		})
	}

	return we
}

type Rank struct {
	Document string
	Weight   float64
}

// Rank mencari peringkat berdasarkan jumlah bobot dari kata kunci di tiap dokumen
func (t *TFIDF) Rank(w []TWeight) []Rank {
	list := make(map[string]float64)
	var rank []Rank

	// menjumlahkan bobot berdasarkan kata kunci "kk"
	sumWeight := func(doc TWeight) {
		for _, key := range t.Keyword {
			list[doc.Document] += doc.Weight[key]
		}
	}

	for _, doc := range w {
		// dokumen yang memiliki key kata kunci "kk" akan dilewati
		// "kk" tidak diperlukan untuk menampilkan informasi
		if doc.Document != "kk" {
			sumWeight(doc)
		}
	}

	for key, weight := range list {
		rank = append(rank, Rank{
			Document: key,
			Weight:   weight,
		})
	}

	SortRank(rank)
	return rank
}

// func (t *TFIDF) Search(keyword ...string) ([]Documents, error) {
// 	if err := t.SetKeywords(keyword...); err != nil {
// 		return []Documents{}, err
// 	}
//
// 	tf := t.TermFrequency()
// 	idf := t.InverseDocumentFrequency()
// 	weight := t.Weight(tf, idf)
//
// 	return FindDocument(t.Documents, weight), nil
// }
