package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func QueryHandler(w http.ResponseWriter, r *http.Request){
	params := r.URL.Query()
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "ID: %v\n", params["id"])
	_, _ = fmt.Fprintf(w, "ID: %v\n", params["category"])
}

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/articles/", QueryHandler)
	r.Queries("id", "category")

	log.Fatal(http.ListenAndServe(":8000", r))
}
