package main

import (
	"context"
	"log"
	"mongodb/usecase"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to Load Env file", err)
	}
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err = mongo.Connect(context.Background(), opts)

	if err != nil {
		log.Fatal("Unable to Connect to MongoDB", err)
	}
	log.Println("Connection String is Correct")

	// Ping the Mongo DB

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Unable to Ping MongoDB", err)
	}
	log.Println("Mongo DB ping successfull")

}

func main() {

	//Create Collection

	coll := client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))

	// create employe Service

	empService := usecase.EmployeeService{MongoCollection: coll}

	// creating a router

	router := mux.NewRouter()

	// create a health handler for /health
	router.HandleFunc("/health", healthHandler).Methods(http.MethodGet)
	router.HandleFunc("/employee", empService.CreateEmployee).Methods(http.MethodPost)
	router.HandleFunc("/employee/{id}", empService.GetEmployeeByID).Methods(http.MethodGet)
	router.HandleFunc("/employee", empService.GetAllEmployee).Methods(http.MethodGet)
	router.HandleFunc("/employee/{id}", empService.UpdateEmployeeByIDr).Methods(http.MethodPut)
	router.HandleFunc("/employee/{id}", empService.DeleteEmployeeByID).Methods(http.MethodDelete)
	router.HandleFunc("/employee", empService.DeleteAllEmployee).Methods(http.MethodDelete)

	// create listen and server for router
	log.Println("Server is Running on port 4444")
	http.ListenAndServe(":4444", router)

	// closing the mongo connection
	defer client.Disconnect(context.Background())
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("running..."))
}
