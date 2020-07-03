package communication

import (
	model "../../domain/entities"
)

type GroupsResponse struct {
	Clusters model.ClustersString `json:"clusters,omitempty"`
}
