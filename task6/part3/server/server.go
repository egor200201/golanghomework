package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// send Not found in such case
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		body, _ := ioutil.ReadFile("C:/Users/Егор/Desktop/golanghomework/task6/part3/page/page.html")
		fmt.Fprint(w, string(body))
	case http.MethodPost:
		cookie := http.Cookie{
			Name:  "token",
			Value: r.PostFormValue("name") + ":" + r.PostFormValue("address"),
		}
		http.SetCookie(w, &cookie)
		body, _ := ioutil.ReadFile("C:/Users/Егор/Desktop/golanghomework/task6/part3/page/page.html")
		fmt.Fprint(w, string(body))

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

const port = 8080

func main() {
	// server port number

	fmt.Printf("Launching server on port: %d \n\n", port)

	// set handler for route '/'
	http.HandleFunc("/", handler)
	// start server without ending
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
