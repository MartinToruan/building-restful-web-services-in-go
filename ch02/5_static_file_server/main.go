package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main(){
	router := httprouter.New()
	router.ServeFiles("/static/*filepath", http.Dir("C:\\Users\\someone\\Documents\\go_project\\src\\github.com\\MartinToruan\\building-restful-web-services-in-go\\ch02\\5_static_file_server\\static"))

	log.Fatal(http.ListenAndServe(":8000", router))
}
