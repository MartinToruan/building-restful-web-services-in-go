package main

import (
	"github.com/gorilla/mux"
	"gopkg.in/gin-gonic/gin.v1/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"time"
)

// Movie holds a movie data
type Movie struct {
	ID			bson.ObjectId	`json:"id" bson:"_id,omitempty"`
	Name 		string 			`json:"name" bson:"name"`
	Year 		string 			`json:"year" bson:"year"`
	Directors 	[]string 		`json:"directors" bson:"directors"`
	Writers 	[]string 		`json:"writers" bson:"writers"`
	BoxOffice 	BoxOffice 		`json:"boxOffice" bson:"boxOffice"`
}

type BoxOffice struct {
	Budget uint64 `bson:"budget"`
	Gross uint64 `bson:"gross"`
}

type DB struct {
	session *mgo.Session
	collection *mgo.Collection
}

func (db *DB) GetMovie(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	// Get data From database
	var movie Movie
	if err := db.collection.Find(bson.M{"_id": bson.ObjectIdHex(vars["id"])}).One(&movie); err != nil{
		w.Header().Set("Content-Type", "text")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("Data Not Found"))
	} else{
		// Marshal Data
		response, _ := json.Marshal(movie)

		// Return response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(response)
	}
}

func (db *DB) PostMovie(w http.ResponseWriter, r *http.Request){
	// Prepare Request Body and save to movie
	var movie Movie
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&movie); err != nil{
		log.Fatalf("error while decode request body: %v\n", err)
	}

	// set movie ID
	movie.ID = bson.NewObjectId()

	// Save the data to Database
	if err := db.collection.Insert(movie); err != nil{
		log.Fatalf("error while insert data to database: %v", err)
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	} else{
		// Unmarshal response
		response, _ := json.Marshal(&movie)
		w.Header().Set("Content-Type", "json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write(response)
	}
}

func (db *DB) UpdateMovie(w http.ResponseWriter, r *http.Request){
	var movie Movie

	// Get Path parameter
	vars := mux.Vars(r)

	// Get Updated Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil{
		log.Fatalf("error while decode data: %v\n", err)
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	}

	// Update data in Database
	if err := db.collection.Update(bson.M{"_id": bson.ObjectIdHex(vars["id"])}, bson.M{"$set": movie}); err != nil{
		log.Fatalf("error while update data in database: %v\n", err)
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	}
}

func (db *DB) DeleteMovie(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	// Delete Data
	if err := db.collection.Remove(bson.M{"_id": bson.ObjectIdHex(vars["id"])}); err != nil{
		log.Fatalf("error while delete data: %v\n", err)
		w.Header().Set("Content-Type", "plain/text")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Data Not found!"))
	} else{
		w.Header().Set("Content-Type", "text")
		_, _ = w.Write([]byte("Delete Succsfull!"))
	}
}

func main(){
	// Prepare Connection to MongoDB
	s, err := mgo.Dial("127.0.0.1")
	if err != nil{
		log.Fatalf("error while connect to MongoDB Server: %v\n", err)
	}
	defer s.Close()
	c := s.DB("appdb").C("movies")

	// Create DB Object
	db := &DB{s, c}

	// Route
	r := mux.NewRouter()

	r.HandleFunc("/v1/movies/{id:[a-zA-Z0-9]*}", db.GetMovie).Methods("GET")
	r.HandleFunc("/v1/movies/", db.PostMovie).Methods("POST")
	r.HandleFunc("/v1/movies/{id:[a-zA-Z0-9]*}", db.UpdateMovie).Methods("PUT")
	r.HandleFunc("/v1/movies/{id:[a-zA-Z0-9]*}", db.DeleteMovie).Methods("DELETE")

	// Prepare Server
	srv := http.Server{
		Addr: "127.0.0.1:8000",
		Handler: r,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start the server
	log.Fatal(srv.ListenAndServe())
}
