package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/gorilla/mux"

	"nlp"
)

var (
	// Version is the software version
	Version    = "unknown"
	okResponse = []byte("OK\n")
	usage      = `usage: %s

Environment variables:
NLPD_ADDR - address to listen on (default :8080)
`
)

type nlpFunc func(string) []string

func checkHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(okResponse)
}

func nlpHandler(w http.ResponseWriter, r *http.Request, fn nlpFunc) {
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("can't read request body - %s", err)
		http.Error(w, "can't read", http.StatusInternalServerError)
		return
	}

	tokens := fn(string(data))
	out, err := json.Marshal(tokens)
	if err != nil {
		log.Printf("can't run %v - %s", fn, err)
		http.Error(w, "can't tokenize", http.StatusInternalServerError)
		return
	}
	w.Write(out)
}

func makeHandler(fn nlpFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		nlpHandler(w, r, fn)
	}
}

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "version", false, "show version & exit")
	flag.Usage = func() {
		name := path.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, usage, name)
	}
	flag.Parse()

	if showVersion {
		fmt.Printf("nlp version %s\n", Version)
		os.Exit(0)
	}

	if flag.NArg() > 0 {
		fmt.Fprintf(os.Stderr, "error: wrong number of arguments\n")
		os.Exit(1)
	}
	addr := os.Getenv("NLPD_ADDR")
	if len(addr) == 0 {
		addr = ":8080"
	}

	r := mux.NewRouter()
	r.HandleFunc("/_/check", checkHandler).Methods("GET")
	r.HandleFunc("/tokenize", makeHandler(nlp.Tokenize)).Methods("POST")
	r.HandleFunc("/sentencize", makeHandler(nlp.Sentencize)).Methods("POST")

	log.Printf("serving NLP on %s\n", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
