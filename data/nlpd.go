package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/353solutions/nlp"
)

func main() {
	http.HandleFunc("/healthz", healthHandler)
	http.HandleFunc("/tokenize", tokenizeHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
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
