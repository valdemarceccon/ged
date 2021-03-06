package main

import (
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", test)
	r.HandleFunc("/upload", upload)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func test(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello world"))

	if err != nil {
		log.Println(err.Error())
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		_, err := w.Write([]byte("only post allowed here"))

		if err != nil {
			log.Println(err.Error())
		}
		return
	}

	defer r.Body.Close()
	f, err := os.OpenFile("a.png", os.O_CREATE | os.O_RDWR, 0755)

	if err != nil {
		log.Println(err.Error())
	}

	r.Body = http.MaxBytesReader(w, r.Body, 2000000000000)

	if err = r.ParseMultipartForm(2000000000000); err != nil {
		log.Println(err)
		return
	}

	file, _, err := r.FormFile("file")
	if  err != nil {
		log.Println(err)
		return
	}
	io.Copy(f, file)

	if err != nil {
		log.Println(err.Error())
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("upload ok"))
}
