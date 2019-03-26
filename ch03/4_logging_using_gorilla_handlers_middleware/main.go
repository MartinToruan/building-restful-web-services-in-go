package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func mainLogic(w http.ResponseWriter, r *http.Request){
	log.Println("Processing request!")
	_, _ = w.Write([]byte("OK"))
	log.Println("Finish Processing request!")
}

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/", mainLogic)
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	log.Fatal(http.ListenAndServe(":8000", loggedRouter))
}
