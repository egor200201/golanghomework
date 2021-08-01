package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Info struct {
	Host       string `json:"host"`
	UserAgent  string `json:"user_agent"`
	RequestURI string `json:"request_uri"`
	Header     Header `json:"headers"`
}

type Header struct {
	Accept    []string `json:"Accept"`
	UserAgent []string `json:"User-Agent"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	info := Info{
		Host:       r.Host,
		UserAgent:  r.UserAgent(),
		RequestURI: r.RequestURI,
		Header: Header{
			Accept:    r.Header["Accept"],
			UserAgent: r.Header["User-Agent"],
		},
	}
	njson, _ := json.Marshal(info)
	fmt.Fprintf(w, "%s\n", njson)
}

func main() {
	const port = 8080

	fmt.Print("port number:", port)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
