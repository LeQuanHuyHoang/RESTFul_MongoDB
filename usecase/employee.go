package usecase

import (
	"RESTFul_MongoDB/model"
	"RESTFul_MongoDB/repository"
	"encoding/json"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

type EmployeeService struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"errors,omitempty"`
}

func (svc *EmployeeService) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)
	var emp model.Employee

	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error decoding")
		res.Error = err.Error()
		return
	}

	emp.EmployeeID = uuid.New().String()

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}
	insertID, err := repo.InsertEmployee(&emp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println()
		res.Error = err.Error()
		return
	}

	res.Data = emp.EmployeeID
	w.WriteHeader(http.StatusOK)

	log.Println("employee created successfully", insertID, emp)
}
func (svc *EmployeeService) GetEmployeeByID(w http.ResponseWriter, r *http.Request)    {}
func (svc *EmployeeService) GetAllEmployee(w http.ResponseWriter, r *http.Request)     {}
func (svc *EmployeeService) UpdateEmployeeByID(w http.ResponseWriter, r *http.Request) {}
func (svc *EmployeeService) DeleteEmployeeByID(w http.ResponseWriter, r *http.Request) {}
func (svc *EmployeeService) DeleteAllEmployee(w http.ResponseWriter, r *http.Request)  {}
