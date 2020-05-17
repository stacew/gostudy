package main

import (
	"log"
	"net/http"
	"time"

	"github.com/stacew/gostudy/tuckersweb/web4/decoHandler"
	"github.com/stacew/gostudy/tuckersweb/web4/myapp"
)

func logger1(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Println("[LOGGER1] Started")
	h.ServeHTTP(w, r)
	log.Println("[LOGGER1] Completed:", time.Since(start).Milliseconds())
}

func logger2(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Println("[LOGGER2] Started")
	h.ServeHTTP(w, r)
	log.Println("[LOGGER2] Completed:", time.Since(start).Milliseconds())
}

func NewHandler() http.Handler {
	h := myapp.NewHandler()
	h = decoHandler.NewDecoHandler(h, logger1)
	h = decoHandler.NewDecoHandler(h, logger2)
	return h
}

func main() {
	mux := NewHandler()
	http.ListenAndServe(":8080", mux)
}
