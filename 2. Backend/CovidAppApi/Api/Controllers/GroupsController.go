package controllers

import (
	"encoding/json"
	"net/http"

	responses "../../aplication/communication"
	interfaces "../../aplication/interfaces"
	service "../../aplication/services"
	header "../../config"
	servError "../errors"
)

type groupscontroller struct{}

var (
	groupServ interfaces.IGroupsService = service.GroupsService()
)

type IGroupsController interface {
	GetGroups(response http.ResponseWriter, request *http.Request)
}

func GroupsController() IGroupsController {
	return &groupscontroller{}
}

func (*groupscontroller) GetGroups(response http.ResponseWriter, request *http.Request) {
	header.AddHeaders(&response)
	if (*request).Method == "OPTIONS" {
		response.WriteHeader(http.StatusOK)
	}

	result, err := groupServ.GetGroups()
	resultGroupsResponse := responses.GroupsResponse{Clusters: result}
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(servError.ServiceError{Message: err.Error()})
	} else {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(resultGroupsResponse)
	}
}
