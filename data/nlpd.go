package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/353solutions/nlp"
)

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/tokenize", tokenizeHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, err := ioutil.ReadAll(io.LimitReader(r.Body, 1<<20))
	if err != nil {
		log.Printf("can't read request - %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tokens := nlp.Tokenize(string(data))
	out, err := json.Marshal(tokens)
	if err != nil {
		log.Printf("can't marshal output - %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(out)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Sanity checks
	fmt.Fprintf(w, "OK\n")
}
