package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/arthurepp/goutil"
	"gopkg.in/mgo.v2"
)

type MessageLog struct {
	Application string `form:"application" json:"application" binding:"required"`
	User        string `form:"user" json:"user" binding:"required"`
	Data        string `form:"data" json:"data" binding:"required"`
	DateTime    string `form:"dateTime" json:"datetime"`
	TrackID     string `form:"trackID" json:"trackID"`
}

func main() {
	fmt.Println("init application")
	fmt.Println("get queue connection")
	ch, err := goutil.GetChannel(os.Getenv("SCRIBAM_QUEUE"))
	goutil.FailOnError(err, "Failed to create connection to queue")

	fmt.Println("get consumer queue")
	msgs, err := goutil.GetConsummer(ch, os.Getenv("SCRIBAM_QUEUE_NAME"))
	goutil.FailOnError(err, "Failed to create connection to queue")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Println(fmt.Sprintf("Received a message: %s", d.Body))

			var messageLog MessageLog
			err := json.Unmarshal(d.Body, &messageLog)
			goutil.FailOnError(err, "Failed to deserialize queue message")

			//TODO validate menssage

			session, err := mgo.Dial(os.Getenv("SCRIBAM_DB_HOST"))
			if err != nil {
				goutil.FailOnError(err, "Failed to connect to database")
			}
			defer session.Close()
			messageLog.DateTime = time.Now().String()
			c := session.DB(os.Getenv("SCRIBAM_DB_HOST")).C("log")
			err = c.Insert(messageLog)
			if err != nil {
				goutil.FailOnError(err, "Failed to insert message in database")
			}

			err = d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
