package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Info struct {
	Host       string      `json:"host"`
	UserAgent  string      `json:"user_agent"`
	RequestURI string      `json:"request_uri"`
	Header     http.Header `json:"headers"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	info := Info{
		Host:       r.Host,
		UserAgent:  r.UserAgent(),
		RequestURI: r.RequestURI,
		Header:     r.Header,
	}
	njson, _ := json.Marshal(info)
	//	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, "%s\n", njson)
}

const port = 8080

func main() {

	fmt.Print("port number:", port)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
