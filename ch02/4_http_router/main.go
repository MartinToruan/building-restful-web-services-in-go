package main

import (
	"bytes"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os/exec"
)

func getCommandOutput(command string, params ...string) string{
	cmd := exec.Command(command, params...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Start
	if err := cmd.Start(); err != nil{
		log.Fatalf("Can't run the command: %v", err)
	}

	// Wait
	if err := cmd.Wait(); err != nil{
		log.Fatalf("Error while waiting the command running: %v", err)
	}

	// End
	return stdout.String()
}

func goVersion(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	_, _ = fmt.Fprintf(w, getCommandOutput("go", "version"))
}

func getFileContent(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	_, _ = fmt.Fprintf(w, getCommandOutput("cat", params.ByName("name")))
}

func main(){
	router := httprouter.New()

	router.GET("/api/v1/go-version", goVersion)
	router.GET("/api/v1/show-file/:name", getFileContent)

	log.Fatal(http.ListenAndServe(":8000", router))

}
