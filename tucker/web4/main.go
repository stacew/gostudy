package main

import (
	"gostudy/tucker/web4/decoHandler"
	"gostudy/tucker/web4/myapp"
	"log"
	"net/http"
	"time"
)

func logger1(w http.ResponseWriter, r *http.Request, h http.Handler) {
	println("STD print logger1() - Started")
	start := time.Now()
	log.Println("[LOGGER1] Started")
	h.ServeHTTP(w, r)
	log.Println("[LOGGER1] Completed:", time.Since(start).Milliseconds())
	println("STD print logger1() - Completed")
}

func logger2(w http.ResponseWriter, r *http.Request, h http.Handler) {
	println("STD print logger2() - Started")
	start := time.Now()
	log.Println("[LOGGER2] Started")
	h.ServeHTTP(w, r)
	log.Println("[LOGGER2] Completed:", time.Since(start).Milliseconds())
	println("STD print logger2() - Completed")
}

func NewHandler() http.Handler {
	h := myapp.NewHandler()
	h = decoHandler.NewDecoHandler(h, logger1)
	h = decoHandler.NewDecoHandler(h, logger2)
	return h
}

func main() {
	mux := NewHandler()
	http.ListenAndServe(":3000", mux)
}
