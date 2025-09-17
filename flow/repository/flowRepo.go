package repository

import "anytype-flow-crud/flow/entities"

type FlowRepo struct {
	BaseURL string
	APIKEY  string
}

func NewFlowRepo() *FlowRepo {
	return &FlowRepo{
		BaseURL: "http://localhost:3000/v1",
		APIKEY:  "ud044Ju8oDxNnni/BGplsR/3LKmqKSMEuyLdvb3y6aE=",
	}
}

func (r FlowRepo) ExposeNewContent(content entities.Content) error {

}
