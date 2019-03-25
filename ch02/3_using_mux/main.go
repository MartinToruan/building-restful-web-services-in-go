package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("/randomFloat", func(w http.ResponseWriter, req *http.Request){
		_, _ = fmt.Fprintf(w, "Your Float Random Number: %f", rand.Float64())
	})

	mux.HandleFunc("/randomInt", func(w http.ResponseWriter, req *http.Request){
		_, _ = fmt.Fprintf(w, "Your Random Int Number: %d", rand.Int())
	})

	log.Fatal(http.ListenAndServe(":8000", mux))
}
