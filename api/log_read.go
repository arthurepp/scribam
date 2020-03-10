package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MessageLogQuery struct {
	Source          string `form:"source" json:"source"`
	InitialDateTime string `form:"initialdatetime" json:"initialdatetime" binding:"required"`
	FinalDateTime   string `form:"finaldatetime" json:"finaldatetime" binding:"required"`
}

func QueryLog(c *gin.Context) {
	var messageLogQuery MessageLogQuery
	if c.Bind(&messageLogQuery) == nil {
		fmt.Println(messageLogQuery)
		session, err := mgo.Dial("mongodb_scribam")
		if err != nil {
			fmt.Println(err)
		}
		defer session.Close()
		collection := session.DB("scribamlog").C("log")
		var results []MessageLog
		conditions := bson.M{
			"$and": []bson.M{
				bson.M{"datetime": bson.M{"$gt": messageLogQuery.InitialDateTime}},
				bson.M{"datetime": bson.M{"$lt": messageLogQuery.FinalDateTime}},
				bson.M{"source": messageLogQuery.Source},
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
