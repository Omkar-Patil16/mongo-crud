package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("env load err", err)
	}

	log.Println("env file loaded")

	//create mongo client

	mongoClient, err = mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(os.Getenv("MONGO_URI")),
	)

	if err != nil {
		log.Fatal("Connection Error\n", err)
	}

	log.Println("env file loaded")

	err = mongoClient.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal("Ping Failed", err)
	}

	log.Println("MongoDB connected")
}

func main() {
	defer mongoClient.Disconnect(context.Background())

	r := mux.NewRouter()

	r.HandleFunc("/health", heathHandler).Methods(http.MethodGet)
	http.ListenAndServe(":4444", r)

}

func heathHandler(w http.ResponseWriter, r *http.Response) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("running"))
}
