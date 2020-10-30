package services

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"

	entities "github.com/heroku/deploy/entities"
)

func getData() []byte {
	data, _ := ioutil.ReadFile("assets/nvi.json")
	//Correção de erro: https://stackoverflow.com/questions/31398044/got-error-invalid-character-%C3%AF-looking-for-beginning-of-value-from-json-unmar
	return bytes.TrimPrefix(data, []byte("\xef\xbb\xbf"))
}

func UnMarshallNVI() ([]entities.Book, error) {
	newsList := make([]entities.Book, 0)
	data := getData()
	err := json.NewDecoder(strings.NewReader(string(data))).Decode(&newsList)
	return newsList, err
}

func BuildJSONCapitulosVersiculos(chapters [][]string) string {
	var sb strings.Builder
	sb.WriteString("[")
	for i := range chapters {
		sb.WriteString("{\"chapter\":\"")
		sb.WriteString(strconv.Itoa(i + 1))
		sb.WriteString("\", \"qtdVerse\":\"")
		sb.WriteString(strconv.Itoa(len(chapters[i])))
		sb.WriteString("\"}")
		if i < len(chapters[i])-1 {
			sb.WriteString(",")
		}
	}
	sb.WriteString("]")
	return sb.String()
}

func BuildContext(cap int64, vs int64, chapter []string) string {
	var sb strings.Builder
	j := vs - 5
	if j < 0 {
		j = 0
	}

	for j < vs {
		sb.WriteString(strconv.Itoa(int(j)))
		sb.WriteString(". ")
		sb.WriteString(chapter[j])
		sb.WriteString("\n")
		j++

	}

	for p := vs; (p < (vs + 5)) && (p < int64(len(chapter))); p++ {
		sb.WriteString(strconv.Itoa(int(p)))
		sb.WriteString(". ")
		sb.WriteString(chapter[p])
		sb.WriteString("\n")
	}

	return sb.String()
}
