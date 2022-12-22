package word

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"strings"

	"aliffatulmf/stki/model"
	"github.com/RadhiFadlillah/go-sastrawi"
)

type OpenType int

const (
	Raw OpenType = iota
	Cleaned
)

func ReadJSON(f string, t OpenType) ([]model.Corpus, error) {
	var data []model.Corpus

	readFile, err := os.Open(f)
	if err != nil {
		return data, err
	}
	defer readFile.Close()

	byteValue, _ := io.ReadAll(readFile)
	if err := json.Unmarshal(byteValue, &data); err != nil {
		return data, err
	}

	switch t {
	case Raw:
		return data, nil
	case Cleaned:
		return CleanSpecialChar(data), nil
	default:
		return data, errors.New("unrecognized type")
	}
}

func CleanSpecialChar(data []model.Corpus) []model.Corpus {
	var res []model.Corpus

	for _, row := range data {
		title := sastrawi.Tokenize(row.Title)
		body := sastrawi.Tokenize(row.Body)

		res = append(res, model.Corpus{
			ID:       row.ID,
			Title:    strings.Join(title, " "),
			Body:     strings.Join(body, " "),
			Document: row.Document,
		})
	}

	return res
}

func Unique(sliceList []string) []string {
	keys := make(map[string]bool)
	var list []string


	for _, entry := range sliceList {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}
