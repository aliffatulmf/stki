package word

import (
	"encoding/json"
	"io"
	"os"
	"regexp"
	"strings"

	"aliffatulmf/stki/model"
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
	re := regexp.MustCompile(`[^a-zA-Z0-9]`)

	for _, row := range data {
		title := re.ReplaceAllString(strings.ToLower(row.Title), " ")
		body := re.ReplaceAllString(strings.ToLower(row.Body), " ")

		res = append(res, model.Corpus{
			ID:       row.ID,
			Title:    CleanWhiteSpace(title),
			Body:     CleanWhiteSpace(body),
			Document: row.Document,
		})
	}

	return res
}

func CleanWhiteSpace(s string) string {
	sep := strings.Fields(s)
	return strings.TrimSpace(strings.Join(sep, " "))
}


func FindByDocument(s string, m []model.Corpus) model.Corpus {
	for idx, val := range m {
		if val.Document == s {
			return m[idx]
		}
	}
	return model.Corpus{}
}

func Unique(sliceList []string) []string {
	keys := make(map[string]bool)
	var list []string


	for _, entry := range sliceList {
		if _, value := keys[CleanWhiteSpace(entry)]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}
