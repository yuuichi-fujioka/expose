package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
)

type Handler struct{}

func (h *Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodGet {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	log.Println(req.RequestURI)

	filepath := req.RequestURI[1:]

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	rw.WriteHeader(http.StatusOK)

	file, err := os.Open(filepath)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()
	io.Copy(rw, file)
}

func NewHandler() *Handler {
	return &Handler{}
}

func main() {

	addr := flag.String("bind", ":8080", "Bind Address")
	flag.Parse()

	s := &http.Server{
		Addr:    *addr,
		Handler: NewHandler(),
	}

	log.Printf("listen %s\n", *addr)

	s.ListenAndServe()

}
