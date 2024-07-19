package model

type Employee struct {
	EmployeeID string `json:"employee_id,omitempty" bson:"employe_id"`
	Name       string `json:"name,omitempty" bson:"name"`
	Department string `json:"department,omitempty" bson:"department"`
}
