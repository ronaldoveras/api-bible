package main

import (
	"encoding/json"
	"fmt"
	 "io/ioutil"
	"log"
	"net/http"
	"strings"
	"bytes"
	"os"
)

type Bible struct{
	Name string
	Abbrev string
	//Chapters []string
}

func handler(w http.ResponseWriter, r *http.Request) {

	data, _ := ioutil.ReadFile("assets/nvi.json")
	//Correção de erro: https://stackoverflow.com/questions/31398044/got-error-invalid-character-%C3%AF-looking-for-beginning-of-value-from-json-unmar
	data = bytes.TrimPrefix(data, []byte("\xef\xbb\xbf"))

	newsList := make([]Bible,0)
	err := json.NewDecoder(strings.NewReader(string(data))).Decode(&newsList)
	var sb strings.Builder
	if err == nil {
		sb.WriteString("[")
		for i, book := range newsList {
			sb.WriteString("\"")
			sb.WriteString(book.Name)
			sb.WriteString("\"")
			if(i < len(newsList) -1){
				sb.WriteString(",")
			}
		}
		sb.WriteString("]")
		fmt.Fprintf(w,"%v",sb.String())
	} 	else {
	 	fmt.Fprintf(w, "Erro inesperado. Panico. %v", err)
	} 

}

func main() {
	http.HandleFunc("/api/v1/books", handler)
        port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(port, nil))
}
