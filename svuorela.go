package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/syndtr/goleveldb/leveldb"
)

func outputData(w http.ResponseWriter, code int, value []byte) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(code)
	if value != nil {
		w.Write(value)
	}
}

func main() {
	db, err := leveldb.OpenFile("svuorela.db", nil)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Printf("%s\n", req.Method)
		switch req.Method {
		case "GET":
			data, _ := db.Get([]byte(req.URL.Path), nil)
			// err is set when the key's not found
			// if err != nil {
			// 	log.Printf("%s\n", err)
			// 	outputData(w, 500, nil)
			// 	return
			// }
			outputData(w, 200, data)
		case "POST":
			data, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Printf("%s\n", err)
				outputData(w, 500, nil)
				return
			}
			db.Put([]byte(req.URL.Path), []byte(data), nil)
		default:
			return
		}
	})
	http.ListenAndServe(":8000", mux)
}
