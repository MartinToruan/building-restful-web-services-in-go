package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

type CustomServeMux struct {
}

func (p *CustomServeMux) ServeHTTP(w http.ResponseWriter, req *http.Request){
	if req.URL.Path == "/"{
		giveRandom(w, req)
		return
	}
	http.NotFound(w, req)
	return
}

func giveRandom(w http.ResponseWriter, req *http.Request){
	_, _ = fmt.Fprintf(w, "Your random number is: %f", rand.Float64())
}

func main(){
	mux := &CustomServeMux{}
	log.Fatal(http.ListenAndServe(":8000", mux))
}
