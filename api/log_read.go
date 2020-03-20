package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MessageLogQuery struct {
	Application     string `form:"application" json:"application"`
	InitialDateTime string `form:"initialdatetime" json:"initialdatetime" binding:"required"`
	FinalDateTime   string `form:"finaldatetime" json:"finaldatetime" binding:"required"`
}

func QueryLog(c *gin.Context) {
	var messageLogQuery MessageLogQuery
	if c.Bind(&messageLogQuery) == nil {
		fmt.Println(messageLogQuery)
		session, err := mgo.Dial(os.Getenv("SCRIBAM_DB_HOST"))
		if err != nil {
			fmt.Println(err)
		}
		defer session.Close()
		collection := session.DB(os.Getenv("SCRIBAM_DB_HOST")).C("log")
		var results []MessageLog
		conditions := bson.M{
			"$and": []bson.M{
				bson.M{"datetime": bson.M{"$gt": messageLogQuery.InitialDateTime}},
				bson.M{"datetime": bson.M{"$lt": messageLogQuery.FinalDateTime}},
				bson.M{"application": messageLogQuery.Application},
			},
		}
		err = collection.Find(conditions).Sort("-datetime").All(&results)
		fmt.Println(results)
		if err != nil {
			c.JSON(404, gin.H{"error": "Fail to reload configurations"})
		} else {
			c.JSON(200, results)
		}
	}
}
