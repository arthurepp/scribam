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
	Source   string `form:"source" json:"source" binding:"required"`
	Message  string `form:"message" json:"message" binding:"required"`
	DateTime string `form:"dateTime" json:"datetime"`
	Type     string `form:"type" json:"type" binding:"required"`
}

func main() {
	fmt.Println("init application")
	fmt.Println("get queue connection")
	ch, err := goutil.GetChannel(os.Getenv("SCRIBAM_QUEUE"))
	goutil.FailOnError(err, "Failed to create connection to queue")

	fmt.Println("get consumer queue")
	msgs, err := goutil.GetConsummer(ch, "log_message")
	goutil.FailOnError(err, "Failed to create connection to queue")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Println(fmt.Sprintf("Received a message: %s", d.Body))

			var messageLog MessageLog
			err := json.Unmarshal(d.Body, &messageLog)
			goutil.FailOnError(err, "Failed to deserialize queue message")

			//valida mensagem de alguma forma

			session, err := mgo.Dial("mongodb_scribam")
			if err != nil {
				goutil.FailOnError(err, "Failed to connect to database")
			}
			defer session.Close()
			messageLog.DateTime = time.Now().String()
			c := session.DB("scribamlog").C("log")
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
