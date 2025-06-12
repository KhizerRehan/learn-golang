package structs

import (
	"errors"
	"strings"
)

type User struct {
	// Explicitly set to lowercase to make it unexported
	name  string
	age   int
	email string
}

type Admin struct {
	User
	role string
}


func New(name string, age int, email string) (*User, error) {
	errorsList := []string{}

	if(name == "") {
		errorsList = append(errorsList, "name is required")
	}
	if(age <= 0) {
		errorsList = append(errorsList, "age is required and should be greater than 0")
	}
	if(email == "") {
		errorsList = append(errorsList, "email is required")
	}	

	if(len(errorsList) > 0) {
		return nil, errors.New(strings.Join(errorsList, ", "))
	}

	return &User{
		name: name,
		age: age,
		email: email,
	}, nil
}

func NewAdmin(name string, age int, email string, role string) (*Admin, error) {
	user, err := New(name, age, email)

	if err != nil {
		return nil, err
	}

	if(role == "") {
		return nil, errors.New("role is required")
	}

	return &Admin{
		User: *user,
		role: role,
	}, nil
}

func (a *Admin) GetRole() string {
	return a.role
}

func (a *Admin) GetName() string {
	return a.User.name
}
	
func (a *Admin) GetAge() int {
	return a.User.age
}

func (a *Admin) GetEmail() string {
	return a.User.email
}

func NewAdminStatic() *Admin {
	return &Admin{
		User: User{
			name: "Static",
			age: 20,
			email: "static@example.com",
		},
		role: "STATIC",
	}
}
