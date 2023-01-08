package tfidf

import "sort"

type By func(doc1, doc2 *Rank) bool

type documentSorter struct {
	documents []Rank
	by        func(doc1, doc2 *Rank) bool
}

func (ds *documentSorter) Len() int {
	return len(ds.documents)
}
func (ds *documentSorter) Swap(i, j int) {
	ds.documents[i], ds.documents[j] = ds.documents[j], ds.documents[i]
}
func (ds *documentSorter) Less(i, j int) bool {
	return ds.by(&ds.documents[i], &ds.documents[j])
}

func (by By) Sort(documents []Rank) {
	ps := &documentSorter{
		documents: documents,
		by:        by,
	}
	sort.Sort(sort.Reverse(ps))
}

func SortRank(documents []Rank) {
	document := func(doc1, doc2 *Rank) bool {
		return doc1.Weight < doc2.Weight
	}

	By(document).Sort(documents)
}
