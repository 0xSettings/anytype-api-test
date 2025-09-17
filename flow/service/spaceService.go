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

func (s *SpaceService) Create(space *entities.Space) (*entities.Space, error) {
	created, err := s.repo.CreateSpace(*space)
	if err != nil {
		return nil, err
	}
	s.cache[created.ID] = created
	return created, nil
}
