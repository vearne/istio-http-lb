package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		client := http.Client{
			Timeout: time.Second * 3,
		}
		log.Println("new request", r.RemoteAddr, r.Header)

		url := "http://ip.vearne.cc/"
		req, _ := http.NewRequest(http.MethodGet, url, nil)

		resp, err := client.Do(req)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		w.Write([]byte(body))
	})
	fmt.Println("starting, port 3001")
	http.ListenAndServe(":3001", nil)
}
