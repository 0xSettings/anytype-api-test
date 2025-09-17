package service

import (
	"anytype-flow-crud/flow/entities"
	"anytype-flow-crud/flow/repository"
	"errors"
)

type ContentService struct {
	repo  *repository.FlowRepo
	cache map[string]*entities.Content
}

func NewContentService(repo *repository.FlowRepo) *ContentService {
	return &ContentService{
		repo:  repo,
		cache: make(map[string]*entities.Content),
	}
}

func (s *ContentService) Create(content *entities.Content) error {
	s.cache[content.ID] = content
	return s.repo.CreateContent(*content)
}

func (s *ContentService) Update(id string, update *entities.Content) error {
	if _, ok := s.cache[id]; !ok {
		return errors.New("content not found")
	}
	s.cache[id] = update
	return nil
}

func (s *ContentService) Delete(id string) error {
	if _, ok := s.cache[id]; !ok {
		return errors.New("content not found")
	}
	delete(s.cache, id)
	return nil
}
