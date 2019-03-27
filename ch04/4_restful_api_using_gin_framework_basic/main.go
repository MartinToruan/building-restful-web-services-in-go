package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main(){
	r := gin.Default()

	r.GET("/pingTime", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"serverTime": time.Now().UTC(),
		})
	})

	r.Run(":8000")
}
