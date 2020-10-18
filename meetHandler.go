package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Meeting struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title        string             `json:"title,omitempty" bson:"title,omitempty"`
	Participants []participant      `json:"participants,omitempty" bson:"participants,omitempty"`
	Starttime    string             `json:"starttime,omitempty" bson:"starttime,omitempty"`
	Endtime      string             `json:"endtime,omitempty" bson:"endtime,omitempty"`
	Creationtime string             `json:"creationtime,omitempty" bson:"creationtime,omitempty"`
}

type participant struct {
	Name  string `json:"name,omitempty" bson:"name,omitempty"`
	Email string `json:"email,omitempty" bson:"email,omitempty"`
	Rsvp  string `json:"rsvp,omitempty" bson:"rsvp,omitempty"`
}

func meetHandler(res http.ResponseWriter, req *http.Request) {

	if req.Method == "POST" {
		createMeeting(res, req)
	}
	if req.Method == "GET" {
		GetTimesMeeting(res, req)
	}
}

func BusyUser(users Meeting) error {

	collection := client.Database("zoom").Collection("meets")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var meet Meeting

	for _, user := range users.Participants {
		if user.Rsvp == "YES" {
			filter := bson.M{
				"endtime": bson.M{"$match": string(time.Now().Format(time.RFC3339))},
			}

			cursor, _ := collection.Find(ctx, filter)
			for cursor.Next(ctx) {
				cursor.Decode(&meet)
				if (users.Starttime >= meet.Starttime && users.Starttime <= meet.Endtime) || (users.Endtime >= meet.Starttime && users.Endtime <= meet.Endtime) {
					return errors.New("Clased Meeting")
				}
			}
		}
	}
	return nil
}

func createMeeting(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("contet-type", "application/json")
	var meet Meeting
	json.NewDecoder(req.Body).Decode(&meet)

	if meet.Starttime < meet.Creationtime {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(`{"error" : "true"},{"message" : "Invalid Starttime"}`))
		return
	}

	if meet.Starttime > meet.Endtime {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(`{"error": "true"},{"message" : "Meeting Cannot start after end time"}`))
		return
	}
	busy := BusyUser(meet)
	if busy != nil {
		res.WriteHeader(http.StatusConflict)
		res.Write([]byte(`"msg" : "Conflicting error"`))
		return
	}
	collection := client.Database("zoom").Collection("meets")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	result, _ := collection.InsertOne(ctx, meet)
	meet.ID = result.InsertedID.(primitive.ObjectID)
	json.NewEncoder(res).Encode(meet)

}

func CheckTime(starttime string, endtime string) []Meeting {
	collection := client.Database("zoom").collection("meetings")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	filter := bson.D{
		{Key: "starttime", Value: bson.M{"$match1": starttime}},
		{key: "endtime", Value: bson.M{"$match2": endtime}},
	}
	cursor, _ := collection.Find(ctx, filter)

	var return_type []Meeting

	var meet Meeting
	for cursor.Next(ctx) {
		cursor.Decode(&meet)
		return_type = append(return_type, meet)
	}
	return return_type

}

func GetTimesMeeting(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("contet-type", "application/json")

	starttime := request.URL.Query()["start"][0]
	endtime := request.URL.Query()["end"][0]

	timing := CheckTime(starttime, endtime)
	json.NewEncoder(res).Encode(timing)

}
