package repository

import (
	"RESTFul_MongoDB/model"
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"testing"
)

func newMongoClient() *mongo.Client {
	mongoTestClient, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI("mongodb+srv://admin:nCSofvLztjfukDVD@cluster0.szakrdx.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))
	if err != nil {
		log.Fatal("error while connecting mongo", err)
	}

	log.Println("mongodb successfully connected.")
	err = mongoTestClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("ping failed", err)
	}

	log.Println("ping successfully")
	return mongoTestClient
}

func TestMongoOperations(t *testing.T) {
	mongoTestClient := newMongoClient()
	defer mongoTestClient.Disconnect(context.Background())

	// dummy data
	emp1 := uuid.New().String()
	//emp2 := uuid.New().String()

	// connect to collection
	conn := mongoTestClient.Database("companydb").Collection("employee_test")

	empRepo := EmployeeRepo{MongoCollection: conn}

	// insert Employee 1 data
	t.Run("Insert Employee 1", func(t *testing.T) {
		emp := model.Employee{
			Name:       "employee1",
			Department: "Department",
			EmployeeID: emp1,
		}

		result, err := empRepo.InsertEmployee(&emp)
		if err != nil {
			t.Fatal("insert employee failed", err)
		}

		t.Log("insert 1 successful", result)
	})

	// Get Employee
	t.Run("Get Employee", func(t *testing.T) {
		result, err := empRepo.FindEmployeeByID(emp1)
		if err != nil {
			t.Fatal("get employee failed", err)
		}

		t.Log("emp1", result.Name)
	})

	//Get All
	t.Run("Get All Employee", func(t *testing.T) {
		results, err := empRepo.FindAllEmployee()
		if err != nil {
			t.Fatal("get all employee failed", err)
		}

		t.Log(results)
	})

	// Update Employee 1 data
	t.Run("Get All Employee", func(t *testing.T) {
		empU := model.Employee{
			Name:       "Update",
			Department: "Update",
			EmployeeID: emp1,
		}

		result, err := empRepo.UpdateEmployeeID(emp1, &empU)
		if err != nil {
			t.Fatal("update employee failed", err)
		}

		t.Log("employee", result)
	})

	t.Run("Delete Employee", func(t *testing.T) {
		result, err := empRepo.DeleteEmployeeByID(emp1)
		if err != nil {
			t.Fatal("Delete employee failed", err)
		}

		t.Log("Delete count", result)
	})

	t.Run("Get all Employee ", func(t *testing.T) {
		results, err := empRepo.FindAllEmployee()
		if err != nil {
			t.Fatal("get all employee failed", err)
		}

		t.Log("employee", results)
	})

	t.Run("Delete all Employee", func(t *testing.T) {
		result, err := empRepo.DeleteAllEmployee()
		if err != nil {
			t.Fatal("Delete all employee failed", err)
		}

		t.Log("Delete count", result)
	})
}
