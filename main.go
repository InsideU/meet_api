package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err.Error())
	}
	ctx, tle := context.WithTimeout(context.Background(), 10*time.Second)
	defer tle()

	http.HandleFunc("/meetings", meetHandler)
	http.HandleFunc("/meeting/", idHandler)

	http.ListenAndServe(":3000", nil)
}
