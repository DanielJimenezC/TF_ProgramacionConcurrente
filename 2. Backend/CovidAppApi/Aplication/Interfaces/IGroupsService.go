package services

import (
	entity "../../domain/entities"
)

// IGroupsService Interface
type IGroupsService interface {
	GetGroups() (entity.ClustersString, error)
}
