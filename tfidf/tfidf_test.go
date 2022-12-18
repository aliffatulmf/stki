package tfidf_test

import (
	"testing"

	"aliffatulmf/stki/tfidf"
	"aliffatulmf/stki/word"
	"github.com/stretchr/testify/assert"
)


func TFIDFNew(t *testing.T) *tfidf.TFIDF {
	stopword := word.OpenStopwordFile("../stopword/stopword.tala.txt")
	corpus, err := word.ReadJSON("../data_valid.json")
	assert.NoError(t, err)

	n := tfidf.New(corpus, stopword)
	//n.SetKeyword(kk)
	return n
}

//func TestTFIDF_SetKeyword(t *testing.T) {
//	n := TFIDFNew(t)
//	assert.Equal(t, kk, n.Keyword)
//}

func TestTFIDF_Token(t *testing.T) {
	n := TFIDFNew(t)
	if assert.NotEmpty(t, n.Token) {
		t.Log(n.Token)
	}
}

func TestTFIDF_TermFrequency(t *testing.T) {
	//docs := []string{"doc1", "doc2", "doc3"}
	n := TFIDFNew(t)
	tf := n.TermFrequency()

	//for _, val := range tf {
	//	assert.Contains(t, docs, val.Document)
	//}
	for _, val := range tf {
		t.Log(val)
	}
}

func TestTFIDF_InverseDocumentFrequency(t *testing.T) {
	n := TFIDFNew(t)
	tf := n.TermFrequency()

	idf := n.InverseDocumentFrequency(tf)
	//t.Log(idf)
	for key, val := range idf {
		t.Log(key, val)
	}
}

//var kk = "ilmu"

func TestTFIDF_Weight(t *testing.T) {
	n := TFIDFNew(t)
	tf := n.TermFrequency()
	idf := n.InverseDocumentFrequency(tf)

	w := n.Weight(tf, idf)
	//t.Log(w)

	tfidf.Find(w, true, "ilmu", "logika")
}
