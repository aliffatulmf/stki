package tfidf

import (
	"fmt"
	"sort"

	"aliffatulmf/stki/word"
)

type TWeight struct {
	Document string
	Weight   map[string]float64
}

func Weight(tf []TField, idf map[string]float64) []TWeight {
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
			fmt.Printf("Ranking %d Dokumen %s\n", idx+1, val)
		}
	}

}
