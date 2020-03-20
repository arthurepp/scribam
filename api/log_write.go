package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/arthurepp/goutil"
	"github.com/gin-gonic/gin"
)

func CreateLog(c *gin.Context) {
	var messageLog MessageLog
	if c.Bind(&messageLog) == nil {
		b, err := json.Marshal(&messageLog)
		goutil.FailOnError(err, "Failed to serialize message to put on queue")

		fmt.Println("get queue connection")
		ch, err := goutil.GetChannel(os.Getenv("SCRIBAM_QUEUE"))
		goutil.FailOnError(err, "Failed to create connection to queue")

		fmt.Println("Put message on queue")
		err = goutil.Publish(ch, os.Getenv("SCRIBAM_QUEUE_NAME"), b)
		goutil.FailOnError(err, "Failed to put message on queue log_message")
	}
	c.Data(201, gin.MIMEHTML, nil)
}
