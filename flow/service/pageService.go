package service

import (
	"anytype-flow-crud/flow/entities"
	"anytype-flow-crud/flow/repository"
)

type PageService struct {
	repo  *repository.FlowRepo
	cache map[string]*entities.Page
}

func NewPageService(repo *repository.FlowRepo) *PageService {
	return &PageService{
		repo:  repo,
		cache: make(map[string]*entities.Page),
	}
}

func (s *PageService) Create(page *entities.Page) error {
	s.cache[page.ID] = page
	return s.repo.CreatePage(*page)
}
