package main

import (
	"io"
	"log"
	"net/http"
)

func handleProxy(w http.ResponseWriter, r *http.Request) {
	url := r.RequestURI[1:]
	if url == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Usage: /http://example.com or /https://example.com"))
		return
	}
	log.Printf("Proxying request to: %s", url)
	resp, err := http.DefaultClient.Do(&http.Request{
		Method: r.Method,
		URL:    r.URL,
