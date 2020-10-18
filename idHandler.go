package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"path"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func checkmeet(id primitive.ObjectID) (Meeting, error) {
	var meet Meeting
	collection := client.Database("zoom").Collection("meets")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, Meeting{ID: id}).Decode(&meet)
	if meet.ID != id {
		err = errors.New("ID not found")
	}
	return meet, err
}

func idHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		res.Header().Set("content-type", "application/json")
		id, _ := primitive.ObjectIDFromHex(path.Base(req.URL.Path))
		meetingwithID, err := checkmeet(id)
		if err != nil {
			log.Fatal(err.Error())
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		json.NewEncoder(res).Encode(meetingwithID)

	}
	res.WriteHeader(http.StatusBadRequest)
	res.Write([]byte(`{"message":"Bad Request"}`))

}
