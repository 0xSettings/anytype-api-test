package service

import (
	"anytype-flow-crud/flow/entities"
	"anytype-flow-crud/flow/repository"
)

type SpaceService struct {
	repo  *repository.FlowRepo
	cache map[string]*entities.Space
}

func NewSpaceService(repo *repository.FlowRepo) *SpaceService {
	return &SpaceService{
		repo:  repo,
		cache: make(map[string]*entities.Space),
	}
}

func (s *SpaceService) Create(space *entities.Space) error {
	s.cache[space.ID] = space
	return s.repo.CreateSpace(*space)
}
