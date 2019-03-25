package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func ArticleHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "Category is: %v\n", vars["category"])
	_, _ = fmt.Fprintf(w, "ID is: %v\n", vars["id"])
}

func main(){
	router := mux.NewRouter()

	router.HandleFunc("/article/{category}/{id:[0-9]+}", ArticleHandler)

	srv := http.Server{
		Handler: router,
		Addr: "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
