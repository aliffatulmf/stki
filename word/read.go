package word

import (
	"encoding/json"
	"io"
	"os"
	"strings"

	"aliffatulmf/stki/model"
	"github.com/RadhiFadlillah/go-sastrawi"
)

func ReadJSON(file string) ([]model.Corpus, error) {
	var data []model.Corpus

	readFile, err := os.Open(file)
	if err != nil {
		return data, err
	}
	defer readFile.Close()

	byteValue, _ := io.ReadAll(readFile)
	if err := json.Unmarshal(byteValue, &data); err != nil {
		return data, err
	}

	return CleanSpecialChar(data), nil
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
