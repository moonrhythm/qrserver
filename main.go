package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/moonrhythm/parapet"
	"github.com/skip2/go-qrcode"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := parapet.NewBackend()
	srv.Addr = ":" + port
	srv.Handler = http.HandlerFunc(generator)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

const defaultSize = 256

func generator(w http.ResponseWriter, r *http.Request) {
	c := r.FormValue("c")
	if c == "" {
		http.Error(w, "empty content", http.StatusBadRequest)
		return
	}

	l := qrcode.Medium
	switch r.FormValue("l") {
	case "0":
		l = qrcode.Low
	case "2":
		l = qrcode.High
	case "3":
		l = qrcode.Highest
	}

	s, _ := strconv.Atoi(r.FormValue("s"))
	switch {
	case s == 0:
		s = defaultSize
	case s < -10:
		s = -10
	case s > 1000:
		s = 1000
	}

	qr, err := qrcode.New(c, l)
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	_ = qr.Write(s, w)
}
