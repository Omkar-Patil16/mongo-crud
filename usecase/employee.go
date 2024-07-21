package usecase

import (
	"encoding/json"
	"log"
	"mongodb/model"
	"mongodb/repository"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeService struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func (svc EmployeeService) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	// writing the header
	w.Header().Add("Content Type", "applicatiom/json")

	//creating the response
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	//Post Request body
	var emp model.Employee

	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid body", err)
		res.Error = err.Error()
		return
	}

	// create empID
	emp.EmployeeID = uuid.NewString()

	// Inser Employe

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	insertID, err := repo.InsertEmployee(emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("insert error", err)
		res.Error = err.Error()
		return
	}

	res.Data = emp.EmployeeID
	w.WriteHeader(http.StatusOK)
	log.Println("Employee with ID inserted", insertID, emp)

}

func (svc EmployeeService) GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content Type", "application/Json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	empID := mux.Vars(r)["id"]
	log.Println("ID of the employee", empID)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	emp, err := repo.FindEmployeeById(empID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error", err)
		res.Error = err.Error()
		return
	}
	res.Data = emp
	w.WriteHeader(http.StatusOK)

}

func (svc EmployeeService) GetAllEmployee(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content Type", "application/Json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	emp, err := repo.FindAllEmployee()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error", err)
		res.Error = err.Error()
		return
	}
	res.Data = emp
	w.WriteHeader(http.StatusOK)
}

func (svc EmployeeService) UpdateEmployeeByIDr(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content Type", "application/Json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	empID := mux.Vars(r)["id"]
	log.Println("ID of the employee", empID)

	if empID == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid EmployeeID")
		res.Error = "Invalid EmployeeID"
		return
	}

	var emp model.Employee

	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid body", err)
		res.Error = err.Error()
		return
	}
	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}
	emp.EmployeeID = empID

	count, err := repo.UpdateEmployeeByID(empID, &emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error", err)
		res.Error = err.Error()
		return
	}
	res.Data = count
	w.WriteHeader(http.StatusOK)
}
func (svc EmployeeService) DeleteEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content Type", "application/Json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	empID := mux.Vars(r)["id"]
	log.Println("ID of the employee", empID)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	count, err := repo.DeleteEmployeeByID(empID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error", err)
		res.Error = err.Error()
		return
	}
	res.Data = count
	w.WriteHeader(http.StatusOK)

}

func (svc EmployeeService) DeleteAllEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content Type", "application/Json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	count, err := repo.DeleteAllEmployees()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error", err)
		res.Error = err.Error()
		return
	}

	res.Data = count
	w.WriteHeader(http.StatusOK)
}
