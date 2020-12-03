package service

import (
	"Go-000/Week02/model"
)

type Service interface {
	GetAllEmployee() ([]model.Employee, error)

	GetEmployeeById(id int) (*model.Employee, error)

	SaveEmployee(employee model.Employee) error

	UpdateEmployee(employee model.Employee) error

	DeleteEmployeeById(id int) error
}
