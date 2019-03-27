package controllers

import (
	"github.com/revel/revel"
	"log"
	"net/http"
	"strconv"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

// TrainResource is the model for holding rail information
type TrainResource struct {
	ID 				int		`json:"id"`
	DriverName		string 	`json:"driver_name"`
	OperatingStatus	bool	`json:"operating_status"`
}

// GetTrain handles GET on train resource
func (c App) GetTrain() revel.Result{
	var train TrainResource
	id := c.Params.Route.Get("train-id")
	train.ID, _ = strconv.Atoi(id)
	train.DriverName = "Yusuf"
	train.OperatingStatus = true

	// Response
	c.Response.Status = http.StatusOK
	return c.RenderJSON(train)
}

func (c App) CreateTrain() revel.Result{
	var train TrainResource
	c.Params.BindJSON(&train)
	train.ID = 2
	c.Response.Status = http.StatusCreated
	return c.RenderJSON(train)
}

func (c App) RemoveTrain() revel.Result{
	id := c.Params.Route.Get("train-id")
	log.Println("successfully deleted the resource: ", id)
	c.Response.Status = http.StatusOK
	return c.RenderText("")
}