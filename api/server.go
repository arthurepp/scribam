package main

import (
	"github.com/gin-gonic/gin"
)

type MessageLog struct {
	Source   string `form:"source" json:"source" binding:"required"`
	Message  string `form:"message" json:"message" binding:"required"`
	DateTime string `form:"dateTime" json:"datetime"`
	Type     string `form:"type" json:"type" binding:"required"`
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
