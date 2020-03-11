package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	Name     string
	Abbrev   string
	Chapters [][]string
}

func handler(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)
	data, _ := ioutil.ReadFile("assets/nvi.json")
	//Correção de erro: https://stackoverflow.com/questions/31398044/got-error-invalid-character-%C3%AF-looking-for-beginning-of-value-from-json-unmar
	data = bytes.TrimPrefix(data, []byte("\xef\xbb\xbf"))

	newsList := make([]Book, 0)
	err := json.NewDecoder(strings.NewReader(string(data))).Decode(&newsList)
	var sb strings.Builder
	if err == nil {
		sb.WriteString("[")
		for i, book := range newsList {
			sb.WriteString("\"")
			sb.WriteString(book.Name)
			sb.WriteString("\"")
			if i < len(newsList)-1 {
				sb.WriteString(",")
			}
		}
		sb.WriteString("]")
		fmt.Fprintf(w, "%v", sb.String())
	} else {
		fmt.Fprintf(w, "Erro inesperado. Panico. %v", err)
	}

}

func handlerVersiculo(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)
	id := r.URL.Query().Get("id")
	cap := r.URL.Query().Get("cap")
	vs := r.URL.Query().Get("vs")
	fmt.Printf("Query %v, %v, %v", id, cap, vs)

	data, _ := ioutil.ReadFile("assets/nvi.json")
	//Correção de erro: https://stackoverflow.com/questions/31398044/got-error-invalid-character-%C3%AF-looking-for-beginning-of-value-from-json-unmar
	data = bytes.TrimPrefix(data, []byte("\xef\xbb\xbf"))

	newsList := make([]Book, 0)
	err := json.NewDecoder(strings.NewReader(string(data))).Decode(&newsList)
	var sb strings.Builder
	if err == nil {
		for _, book := range newsList {
			if book.Name == id {
				capInt, _ := strconv.ParseInt(cap, 10, 32)
				vsInt, _ := strconv.ParseInt(vs, 10, 32)

				sb.WriteString(book.Chapters[capInt-1][vsInt-1])
			}
		}
		fmt.Fprintf(w, "%v", sb.String())
	} else {
		fmt.Fprintf(w, "Erro inesperado. Panico. %v", err)
	}

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	http.HandleFunc("/api/v1/books", handler)
	http.HandleFunc("/api/v1/nvi", handlerVersiculo)
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
