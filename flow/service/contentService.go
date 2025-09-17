package service

import (
	"anytype-flow-crud/flow/entities"
	"anytype-flow-crud/flow/repository"
)

type ContentService struct {
	rep *repository.FlowRepo

	//Since API isnt supporting Update via body params
	cache map[string]*entities.Content
}

func ExposeNewContentService(rep *repository.FlowRepo) *ContentService {
	return &ContentService{
		rep:   rep,
		cache: make(map[string]*entities.Content),
	}
}
