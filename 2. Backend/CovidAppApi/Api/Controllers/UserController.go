package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	interfaces "../../aplication/interfaces"
	service "../../aplication/services"
	header "../../config"
	model "../../domain/entities"
	servError "../errors"
	"github.com/gorilla/mux"
)

type controller struct{}

var (
	userServ  interfaces.IUserService       = service.UserService()
	blockServ interfaces.IBlockChainService = service.BlockChainService()
)

// IUserController interface
type IUserController interface {
	GetAll(response http.ResponseWriter, request *http.Request)
	SignUp(response http.ResponseWriter, request *http.Request)
	GetByID(response http.ResponseWriter, request *http.Request)
	Update(response http.ResponseWriter, request *http.Request)
	Delete(response http.ResponseWriter, request *http.Request)
	Login(response http.ResponseWriter, request *http.Request)
	AddBlock(response http.ResponseWriter, request *http.Request)
}

// UserController Implementation
func UserController() IUserController {
	return &controller{}
}

// GetAll Users
func (*controller) GetAll(response http.ResponseWriter, request *http.Request) {
	header.AddHeaders(&response)
	//response.Header().Set("Content-Type", "application/json")
	//response.Header().Set("Access-Control-Allow-Origin", "*")
	users, err := userServ.GetAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(servError.ServiceError{Message: err.Error()})
	} else {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(users)
	}
}

// SignUp User
func (*controller) SignUp(response http.ResponseWriter, request *http.Request) {
	header.AddHeaders(&response)
	if (*request).Method == "OPTIONS" {
		response.WriteHeader(http.StatusCreated)
	}
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
			json.NewEncoder(response).Encode(servError.ServiceError{Message: fmt.Sprintf("Error creating users - %s", err1.Error())})
		} else {
			response.WriteHeader(http.StatusCreated)
			json.NewEncoder(response).Encode(user)
		}
	}
}

// GetbByID User
func (*controller) GetByID(response http.ResponseWriter, request *http.Request) {
	header.AddHeaders(&response)
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
	} else {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(user)
	}
}

// Delete user
func (*controller) Delete(response http.ResponseWriter, request *http.Request) {
	header.AddHeaders(&response)
	if (*request).Method == "OPTIONS" {
		response.WriteHeader(http.StatusOK)
	}
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
	} else {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(id)
	}
}

// Update user
func (*controller) Update(response http.ResponseWriter, request *http.Request) {
	header.AddHeaders(&response)
	if (*request).Method == "OPTIONS" {
		response.WriteHeader(http.StatusOK)
	}
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
	} else {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(user)
	}
}

// Login user
func (*controller) Login(response http.ResponseWriter, request *http.Request) {
	header.AddHeaders(&response)
	if (*request).Method == "OPTIONS" {
		response.WriteHeader(http.StatusOK)
	}
	var user model.User

	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(servError.ServiceError{Message: "Error decode Json users"})
	}

	login, _ := userServ.Login(user)
	if !login {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(servError.ServiceError{Message: "Invalid Username or Password"})
	} else {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(servError.ServiceError{Message: "Succesfull Login"})
	}
}

func (*controller) AddBlock(response http.ResponseWriter, request *http.Request) {
	header.AddHeaders(&response)
	if (*request).Method == "OPTIONS" {
		response.WriteHeader(http.StatusOK)
	}
	var block model.Block

	err := json.NewDecoder(request.Body).Decode(&block)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(servError.ServiceError{Message: "Error decode Json users"})
	}
	done, err := blockServ.BlockChainGenerated(block)
	if !done {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(servError.ServiceError{Message: "Error al insertar Bloque"})
	} else {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(servError.ServiceError{Message: "Succesfull"})
	}
}
