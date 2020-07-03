package services

import (
	entity "../../domain/entities"
)

// IPredictionService Interface
type IGroupsService interface {
	GetGroups() (entity.ClustersString, error)
}
