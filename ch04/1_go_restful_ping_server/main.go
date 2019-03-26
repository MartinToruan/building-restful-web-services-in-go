package main

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"io"
	"log"
	"net/http"
	"time"
)

func main(){
	// Create web service
	webservice := new(restful.WebService)

	// Create route
	webservice.Route(webservice.GET("/ping").To(pingTime))

	restful.Add(webservice)

	// Start Server
	log.Fatal(http.ListenAndServe(":8000", nil))

}

func pingTime(req* restful.Request, resp *restful.Response){
	_, _ = io.WriteString(resp, fmt.Sprintf("%s", time.Now()))
}
