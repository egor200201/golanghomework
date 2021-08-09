package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var Tokens map[int64]Token = make(map[int64]Token)

type Token struct {
	Token     string
	CreatedAt string
	ExpiredAt string
}
type Handlers struct {
}

func NewHandlers() *Handlers {
	return &Handlers{}
}

func (h *Handlers) Configure() *http.ServeMux {
	userMux := http.NewServeMux()
	userMux.HandleFunc("/", mainHandler)
	userMux.HandleFunc("/list", listHandler)
	return userMux
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case http.MethodGet:
		body, _ := ioutil.ReadFile("page.html")
		fmt.Fprint(w, string(body))
	case http.MethodPost:
		cookie := http.Cookie{
			Name:  "token",
			Value: r.PostFormValue("name") + ":" + r.PostFormValue("address"),
		}
		http.SetCookie(w, &cookie)
		newToken := Token{
			Token:     cookie.Value,
			CreatedAt: time.Now().Local().String(),
			ExpiredAt: time.Now().Local().Add(time.Hour * time.Duration(240)).String(),
		}
		Tokens[int64(len(Tokens))] = newToken
		fmt.Fprintf(w, "saved!")
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/list" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		return
	} else {
		for i, val := range Tokens {
			fmt.Fprintf(w, "%d \t %s \t %s \t %s \n", i, val.Token, val.CreatedAt, val.ExpiredAt)
		}
	}
}

type Server struct {
	httpServer *http.Server
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}
	fmt.Printf("listening at %s", port)
	return s.httpServer.ListenAndServe()

}

func main() {
	server := NewServer()
	handlers := NewHandlers()

	server.Run("8080", handlers.Configure())
}
