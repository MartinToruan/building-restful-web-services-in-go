package main

import (
	"database/sql"
	"encoding/json"
	"github.com/MartinToruan/building-restful-web-services-in-go/ch04/3_metro_rail_api_using_go_restful/dbutils"
	"github.com/emicklei/go-restful"
	_ "github.com/mattn/go-sqlite3"
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

	// Create Continer
	wsContiner := restful.NewContainer()
	wsContiner.Router(restful.CurlyRouter{})

	// Add service to Container
	t := TrainPersistence{}
	t.Register(wsContiner)

	// Start Server
	server := &http.Server{
		Addr: ":8000",
		Handler: wsContiner,
	}
	log.Fatal(server.ListenAndServe())
}

type TrainPersistence struct {
	ID int
	DriverName string
	OperatingStatus bool
}

func (t *TrainPersistence) Register(container *restful.Container){
	ws := new(restful.WebService)
	ws.
		Path("/v1/trains").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	// You can Specify this per route as well
	ws.Route(ws.GET("/{train-id}").To(t.getTrain))
	ws.Route(ws.POST("").To(t.createTrain))
	ws.Route(ws.DELETE("/{train-id}").To(t.removeTrain))

	// Add service to Container
	container.Add(ws)
}

func (t TrainPersistence) getTrain(req *restful.Request, resp *restful.Response){
	// Get ID from request
	id := req.PathParameter("train-id")

	// Get data from database
	if err := DB.QueryRow("select id, driver_name, operating_status from train where id=?", id).Scan(&t.ID, &t.DriverName, &t.OperatingStatus); err != nil{
		log.Println(err)
		resp.AddHeader("Content-Type", "text/plain")
		resp.WriteErrorString(http.StatusNotFound, "Train could not be found.")
	} else {
		resp.WriteEntity(t)
	}
}

func (t *TrainPersistence) createTrain(req *restful.Request, resp *restful.Response){
	log.Println(req.Request.Body)
	var b TrainPersistence

	decoder := json.NewDecoder(req.Request.Body)
	if err := decoder.Decode(&b); err != nil{
		log.Fatalf("error while decode request body: %v", err)
	}
	defer func(){
		_ = req.Request.Body.Close()
	}()

	statement, err := DB.Prepare("insert into train(driver_name, operating_status) values(?, ?)")
	if err != nil{
		log.Fatalf("error while prepare instert statement: %v\n", err)
	}

	res, err := statement.Exec(b.DriverName, b.OperatingStatus)
	if err != nil{
		log.Fatalf("error while exec insert query: %v\n", err)
		resp.AddHeader("Content-Type", "text/plain")
		_ =  resp.WriteErrorString(http.StatusInternalServerError, err.Error())
	} else{
		newID, _ := res.LastInsertId()
		b.ID = int(newID)
		_ = resp.WriteHeaderAndEntity(http.StatusCreated, b)
	}
}

func (t *TrainPersistence) removeTrain(req *restful.Request, resp * restful.Response){
	id := req.PathParameter("train-id")

	statement, err := DB.Prepare("delete from train where id=?")
	if err != nil{
		log.Fatalf("error while prepare statement: %v\n", err)
		resp.AddHeader("Content-Type", "text/plain")
		_ = resp.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
	_, err = statement.Exec(id)
	if err != nil{
		log.Fatal("error while delete data in database: %v", err)
		resp.AddHeader("Content-Type", "plain/text")
		_ = resp.WriteErrorString(http.StatusInternalServerError, err.Error())
	} else{
		resp.WriteHeader(http.StatusOK)
	}

}