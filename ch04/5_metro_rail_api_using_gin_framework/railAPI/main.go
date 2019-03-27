package main

import (
	"database/sql"
	"github.com/MartinToruan/building-restful-web-services-in-go/ch04/5_metro_rail_api_using_gin_framework/dbutils"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/gin-gonic/gin.v1/json"
	"log"
	"net/http"
)

var DB *sql.DB
var err error

func main(){
	DB, err = sql.Open("sqlite3", "./railapi.db")
	if err != nil{
		log.Fatalf("error while setting up driver: %v", err)
	}

	// Create Tables
	dbutils.Initialize(DB)

	// Gin Default Server
	r := gin.Default()

	// Router
	r.GET("/v1/stations/:station_id", GetStation)
	r.POST("/v1/stations/", CreateStation)
	r.DELETE("/v1/stations/:station_id", DeleteStation)

	// Run Server
	r.Run(":8000")


}

// StationResource holds information about locations
type StationPersistence struct {
	ID int `json:"id"`
	Name string `json:"name"`
	OpeningTime string `json:"opening_time"`
	ClosingTime string `json:"closing_time"`
}

// GetStation returns the station detail
func GetStation(c *gin.Context) {
	var station StationPersistence
	id := c.Param("station_id")
	err := DB.QueryRow("select ID, NAME, CAST(OPENING_TIME as CHAR), CAST(CLOSING_TIME as CHAR) from station where id=?", id).Scan(&station.ID, &station.Name, &station.OpeningTime, &station.ClosingTime)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"result": station,
		})
	}
}

func CreateStation(c *gin.Context){
	var station StationPersistence
	decoder := json.NewDecoder(c.Request.Body)
	if err := decoder.Decode(&station); err != nil{
		log.Fatalf("error while decode: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	statement, err := DB.Prepare("Insert into station(NAME, opening_time, closing_time) VALUES(?, ?, ?)")
	if err != nil{
		log.Fatalf("error while prepare statement: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	res, err := statement.Exec(station.Name, station.OpeningTime, station.ClosingTime)
	if err != nil{
		log.Fatalf("error while insert data: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}else{
		newID, _ := res.LastInsertId()
		station.ID = int(newID)
		c.JSON(http.StatusCreated, gin.H{
			"result": station,
		})
	}
}

func DeleteStation(c *gin.Context){
	id := c.Param("station_id")

	// Prepare Delete Statement
	statement, err := DB.Prepare("delete from station where id=?")
	if err != nil{
		log.Fatalf("error while prepare delete statement: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	_, err = statement.Exec(id)
	if err != nil{
		log.Fatalf("error while update data in database: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else{
		c.JSON(http.StatusOK, "")
	}

}