package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("<h1>Ola</h1>"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
