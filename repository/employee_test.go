package repository

import (
	"context"
	"log"
	"mongodb/model"
	"testing"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoClient() *mongo.Client {
	mongoTestClient, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI("mongodb+srv://admin:admin123@cluster0.rs5nrtm.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"),
	)

	if err != nil {
		log.Fatal("Not Able to connect to MongoDB", err)
	}
	log.Println("Able to connect to Mongo DB")

	err = mongoTestClient.Ping(
		context.Background(),
		readpref.Primary(),
	)

	if err != nil {
		log.Fatal("Not Able to Ping the Primary Cluster", err)
	}

	log.Println("Able to Ping the Primary Cluster")

	return mongoTestClient
}

func TestMongoOperations(t *testing.T) {
	mongoTestClient := NewMongoClient()
	defer mongoTestClient.Disconnect(context.Background())

	// dumy data
	emp1 := uuid.NewString()
	emp2 := uuid.NewString()

	// create collection
	// mongodb will create new connection if it does no exist
	coll := mongoTestClient.Database("companydb").Collection("employee_test")

	// Initiate EmployeeRepo
	// Employee Repo accepts mongo collection instance
	empRepo := EmployeeRepo{
		MongoCollection: coll,
	}

	// Creating First Employee

	t.Run("Create First Employee", func(t *testing.T) {
		emp := model.Employee{
			Name:       "Rohit Sharma",
			Department: "Physics",
			EmployeeID: emp1,
		}
		result, err := empRepo.InsertEmployee(emp)

		if err != nil {
			t.Fatal("Not able to create Employee 1\n", err)
		}

		t.Log("Employee 1 Create Succesfully\n", result)
	})

	// Add Second Employee

	t.Run("Create Second Employee", func(t *testing.T) {
		emp := model.Employee{
			Name:       "Virat Kholi",
			Department: "Chemistry",
			EmployeeID: emp2,
		}
		result, err := empRepo.InsertEmployee(emp)

		if err != nil {
			t.Fatal("Not able to create Employee 2\n", err)
		}

		t.Log("Employee 2 Create Succesfully\n", result)
	})

	// Find Emplpoyee

	t.Run("Fetch Employee One", func(t *testing.T) {
		result, err := empRepo.FindEmployeeById(emp1)

		if err != nil {
			t.Fatal("Employee One not Found\n", err)
		}

		t.Log("Employee One\n", result.Name)
	})

	// List All Emplpoyees

	t.Run("Print All Employees", func(t *testing.T) {
		result, err := empRepo.FindAllEmployee()

		if err != nil {
			t.Fatal("No Employees Found\n", err)
		}

		t.Log("Listing All Employees\n", result)
	})

	// Update First Employee

	t.Run("Update Employee Name", func(t *testing.T) {
		emp := model.Employee{
			Name:       "Rohit Gurunat Sharma",
			Department: "Chemistry",
			EmployeeID: emp1,
		}
		result, err := empRepo.UpdateEmployeeByID(emp1, &emp)

		if err != nil {
			t.Fatal("Update Opreation Failed\n", err)
		}

		t.Log("Update Count\n", result)

	})

	// List All Emplpoyees After Update

	t.Run("List All Emplpoyees After Update", func(t *testing.T) {
		result, err := empRepo.FindAllEmployee()

		if err != nil {
			t.Fatal("No Employees Found\n", err)
		}

		t.Log("Listing All Employees\n", result)
	})

	// Delete Emplpoyee

	t.Run("Delete Emplpoyee ", func(t *testing.T) {
		result, err := empRepo.DeleteEmployeeByID(emp1)

		if err != nil {
			t.Fatal("Employee One not Found\n", err)
		}

		t.Log("Delete Count \n", result)
	})

	// List All Emplpoyees After Deletion

	t.Run("After Deletion Print All Employees", func(t *testing.T) {
		result, err := empRepo.FindAllEmployee()

		if err != nil {
			t.Fatal("No Employees Found\n", err)
		}

		t.Log("Listing All Employees\n", result)
	})

	// Delete All Employees for Cleanup
	t.Run("Delete All Employees for Cleanup", func(t *testing.T) {
		result, err := empRepo.DeleteAllEmployees()

		if err != nil {
			t.Fatal("Delete Oreation Failed\n", err)
		}

		t.Log("Log Count\n", result)
	})
}
