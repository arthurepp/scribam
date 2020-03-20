package main

import (
	"github.com/gin-gonic/gin"
)

type MessageLog struct {
	Application string `form:"application" json:"application" binding:"required"`
	User        string `form:"user" json:"user" binding:"required"`
	Data        string `form:"data" json:"data" binding:"required"`
	DateTime    string `form:"dateTime" json:"datetime"`
	TrackID     string `form:"trackID" json:"trackID"`
}

func main() {
	r := gin.Default()

	v1 := r.Group("api/")
	{
		v1.POST("/create", CreateLog)
		v1.POST("/find", QueryLog)
	}
	r.Run(":8001")
}
