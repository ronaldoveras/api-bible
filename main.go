package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	services "github.com/heroku/deploy/services"
)

func handler(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	dados, err := services.UnMarshallNVI()
	var sb strings.Builder
	if err == nil {
		sb.WriteString("[")
		for i, book := range dados {
			sb.WriteString("{\"name\":\"")
			sb.WriteString(book.Name)
			sb.WriteString("\", ")

			sb.WriteString("\"qtdBooks\":\"")
			sb.WriteString(strconv.Itoa(len(book.Chapters)))
			// fmt.Printf("%s contém %d capítulos \n", book.Name, len(book.Chapters))
			sb.WriteString("\", \"chaptersResume\" : ")
			sb.WriteString(services.BuildJSONCapitulosVersiculos(book.Chapters))
			sb.WriteString("\"}")
			if i < len(dados)-1 {
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

	dados, err := services.UnMarshallNVI()
	var sb strings.Builder
	if err == nil {
		for _, book := range dados {
			if book.Name == id {
				capInt, _ := strconv.ParseInt(cap, 10, 32)
				vsInt, _ := strconv.ParseInt(vs, 10, 32)

				defer func() {
					if err := recover(); err != nil {
						log.Println("Ocorreu um erro nos indices de capítulo e versículo.")
						w.WriteHeader(http.StatusBadRequest)
						fmt.Fprintf(w, "Erro inesperado. :; Sorry! Bad Request.")
					}
				}()
				versicle := book.Chapters[capInt-1][vsInt-1]
				sb.WriteString("{")
				sb.WriteString("\"verse\": \"")
				sb.WriteString(versicle)
				sb.WriteString("\"")

				sb.WriteString(",")

				sb.WriteString("\"context\": \"")
				sb.WriteString(services.BuildContext(capInt, vsInt, book.Chapters[capInt-1]))
				sb.WriteString("\"")
				sb.WriteString("}")

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
	http.HandleFunc("/api/nvi/v1/books", handler)
	http.HandleFunc("/api/nvi/v1/verse", handlerVersiculo)
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
