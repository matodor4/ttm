package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)


var port JsPort

type response struct {
	Quote string `json:"Quote"`
}

func StartServer(syncCh chan struct{}, quoteCh chan []byte) {
	err := GetPort(JsonPath, &port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server start at :%d port", port.Port)
	err = http.ListenAndServe(":"+strconv.Itoa(port.Port), getHandler(syncCh, quoteCh))

	if err != nil {
		log.Fatal(err)
	}

}

func getHandler(syncCh chan struct{}, quoteCh chan []byte) http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:

			//sync two goroutine
			syncCh <- struct{}{}

			//get actual quote
			quote := <- quoteCh

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			resp := &response{Quote: string(quote)}
			jsonResp, err := json.Marshal(resp)
			if err != nil {
				log.Fatal(err)
			}
			_, err = w.Write(jsonResp)
			if err != nil {
				log.Fatal(err)
			}
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	return r
}
