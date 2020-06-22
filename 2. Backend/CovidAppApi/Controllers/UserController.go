package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "../entities"
	servError "../errors"
	service "../services"
	"github.com/gorilla/mux"
)

type controller struct{}

var (
	userServ service.IUserService = service.UserService()
)

// IUserController interface
type IUserController interface {
	GetAll(response http.ResponseWriter, request *http.Request)
	Create(response http.ResponseWriter, request *http.Request)
	GetByID(response http.ResponseWriter, request *http.Request)
	Update(response http.ResponseWriter, request *http.Request)
	Delete(response http.ResponseWriter, request *http.Request)
}

// UserController Implementation
func UserController() IUserController {
	return &controller{}
}

// GetAll Users
func (*controller) GetAll(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	users, err := userServ.GetAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(servError.ServiceError{Message: "Error getting the users"})
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(users)
}

// Create User
func (*controller) Create(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user model.User

	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(servError.ServiceError{Message: "Error decode Json users"})
	}

	err2 := userServ.Validate(&user)
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(servError.ServiceError{Message: err2.Error()})
	} else {
		err1 := userServ.Create(user)
		if err1 != nil {
			response.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(response).Encode(servError.ServiceError{Message: "Error creating users"})
		}
		response.WriteHeader(http.StatusCreated)
		json.NewEncoder(response).Encode(user)
	}
}

// GetbByID User
func (*controller) GetByID(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(servError.ServiceError{Message: err.Error()})
	}
	user, err1 := userServ.GetByID(id)
	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(servError.ServiceError{Message: err1.Error()})
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(user)
}

// Delete user
func (*controller) Delete(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(servError.ServiceError{Message: err.Error()})
	}
	err1 := userServ.Delete(id)
	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(servError.ServiceError{Message: err1.Error()})
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(id)
}

// Update user
func (*controller) Update(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user model.User

	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(servError.ServiceError{Message: "Error decode Json users"})
	}

	vars := mux.Vars(request)
	id, err1 := strconv.Atoi(vars["id"])
	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(servError.ServiceError{Message: err1.Error()})
	}

	err2 := userServ.Update(id, user)
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(servError.ServiceError{Message: err2.Error()})
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(user)
}
