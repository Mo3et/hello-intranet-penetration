package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mo3et/hello-intranet-penetration/define"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		q := request.URL.Query()
		byte, err := json.Marshal(q)
		if err != nil {
			log.Printf("Marshal Error: %v", err)
		}
		if _, err := writer.Write(byte); err != nil {
			panic(err)
			// log.Printf("Writer fail", err)
		}
	})
	log.Println("Local Server is Running", define.LocalServerAddr)

	if err := http.ListenAndServe(define.LocalServerAddr, nil); err != nil {
		panic(err)
	}
}
