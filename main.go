package main

import (
	"RESTFul_MongoDB/usecase"
	"context"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
	"os"
)

var mongoClient *mongo.Client

func init() {
	//load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("env load error", err)
	}

	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	if err != nil {
		log.Fatal("connection error", err)
	}

	err = mongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("ping failed", err)
	}

	log.Println("mongo connected")
}

func main() {
	defer mongoClient.Disconnect(context.Background())

	conn := mongoClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))

	//create employee service
	empService := usecase.EmployeeService{MongoCollection: conn}

	r := mux.NewRouter()

	r.HandleFunc("/create_employee", empService.CreateEmployee).Methods("POST")
	r.HandleFunc("/employee/{id}", empService.GetEmployeeByID).Methods("GET")
	r.HandleFunc("/employees", empService.GetAllEmployee).Methods("GET")
	r.HandleFunc("/update_employees/{id}", empService.UpdateEmployeeByID).Methods("PUT")
	r.HandleFunc("/delete_employees/{id}", empService.GetEmployeeByID).Methods("DELETE")

	log.Println("server is running on 8080")
	http.ListenAndServe(":8080", r)
}
