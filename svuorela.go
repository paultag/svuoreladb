package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func outputData(w http.ResponseWriter, code int, value *string) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(code)
	if value != nil {
		w.Write([]byte(*value))
	}
}

func main() {
	db, err := LoadDatabase("svuorela.db")
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Printf("%s\n", req.Method)
		switch req.Method {
		case "GET":
			outputData(w, 200, db.Get(req.URL.Path))
		case "POST":
			data, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Printf("%s\n", err)
				outputData(w, 500, nil)
			}
			db.Write(req.URL.Path, string(data))
		default:
			return
		}
	})
	http.ListenAndServe(":8000", mux)
}
