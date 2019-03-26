package main

import (
	jsonparse "encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Args struct {
	Id string
}

type Book struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
}

type JSONServer struct {

}

func (s *JSONServer) GiveBookDetail(r *http.Request, args *Args, reply *Book) error{
	// Read JSON file
	var books []Book
	raw, readerr := ioutil.ReadFile("./books.json")
	if readerr != nil{
		log.Fatalf("error while read file: %v\n", readerr)
		os.Exit(1)
	}

	// Unmarshall JSON raw data into books array
	if marshalerr := jsonparse.Unmarshal(raw, &books); marshalerr != nil{
		log.Fatalf("error while unmarhsaling data: %v\n", marshalerr)
		os.Exit(1)
	}

	// Search Book
	for _, book := range books{
		if book.Id == args.Id{
			*reply = book
			break
		}
	}

	return nil
}

func main(){
	// Create a new RPC server
	s := rpc.NewServer()

	// Register Codec
	s.RegisterCodec(json.NewCodec(), "application/json")

	// Register Service
	if errStart := s.RegisterService(new(JSONServer), ""); errStart != nil{
		log.Fatalf("error while starting server: %v\n", errStart)
	}

	// Create route
	route := mux.NewRouter()
	route.Handle("/rpc", s)

	// Start server
	log.Fatal(http.ListenAndServe(":1234", route))
}
