package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

// Movie holds a movie data
type Movie struct {
	Name string `bson:"name"`
	Year string `bson:"year"`
	Directors []string `bson:"directors"`
	Writers []string `bson:"writers"`
	BoxOffice `bson:"boxOffice"`
}

type BoxOffice struct {
	Budget uint64 `bson:"budget"`
	Gross uint64 `bson:"gross"`
}

func main(){
	session, err := mgo.Dial("localhost")
	if err != nil{
		log.Fatalf("error while connect to MongoDB: %v\n", err)
	}
	defer session.Close()

	// Connect to Database and Collection
	c := session.DB("appdb").C("movies")

	// Create a movie
	darkNigh := &Movie{
		Name: "The Dark Knight2",
		Year: "2008",
		Directors: []string{"Christoper Nolan"},
		Writers: []string{"Jonathan Nolan", "Christoper Nolan"},
		BoxOffice: BoxOffice{
			Budget: 185000000,
			Gross: 533316061,
		},
	}

	if err := c.Insert(darkNigh); err != nil{
		log.Fatalf("error while insert data to database: %v\n", err)
	}

	result:= Movie{}
	//if err := c.Find(bson.M{"boxOffice.budget": bson.M{"$gt": 1850000}}).One(&result); err != nil{
	//	log.Fatalf("error while read data from database: %v\n", err)
	//}
	if err := c.Find(bson.M{"name": "The Dark Knight2"}).One(&result); err != nil{
		log.Fatalf("error while read data from database: %v\n", err)
	}
	log.Println("===")
	log.Println(result.Name)
	log.Println(result.Year)
	log.Println(result.Directors)
	log.Println("===")
	log.Println(result)
}
