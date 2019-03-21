package main

import (
	"fmt"
	"github.com/MartinToruan/building-restful-web-services-in-go/ch01/romanNumerals"
	"html"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	// http package has methods for dealing with requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPathElements := strings.Split(r.URL.Path, "/")

		if urlPathElements[1] == "roman_number" {
			number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
			if number == 0 || number > 10 {
				w.WriteHeader(http.StatusNotFound)
				_, _ = w.Write([]byte("404 - Not Found"))
			} else {
				_, _ = fmt.Fprintf(w, "%q", html.EscapeString(romanNumerals.Numerals[number]))
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("400 - Bad request"))
		}
	})

	s := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("Can't run the server: %v", err)
	}
}
