package main

import (
	"fmt"
	"log"
	"net/http"
)

func middleware(hander http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Println("Executing middleware before request phase!")
		// Pass control back to the handler
		hander.ServeHTTP(w, r)
		fmt.Println("Executing middleware after response phase!")
	})
}

func mainLogic(w http.ResponseWriter, r *http.Request){
	// Business Logic goes here
	fmt.Println("Executing mainHandler...")
	_, _ = w.Write([]byte("OK"))
}

func main(){
	mainLogicHandler := http.HandlerFunc(mainLogic)
	http.Handle("/", middleware(mainLogicHandler))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
