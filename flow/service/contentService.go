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

func (s *ContentService) Create(content *entities.Content) (*entities.Content, error) {
	created, err := s.repo.CreateContent(*content)
	if err != nil {
		return nil, err
	}
	s.cache[created.ID] = created
	return created, nil
}

func (s *ContentService) Update(id string, update *entities.Content) (*entities.Content, error) {
	if _, ok := s.cache[id]; !ok {
		return nil, errors.New("content not found")
	}
	update.ID = id
	updated, err := s.repo.UpdateContent(*update)
	if err != nil {
		return nil, err
	}
	s.cache[id] = updated
	return updated, nil
}

func (s *ContentService) Delete(id string) error {
	if _, ok := s.cache[id]; !ok {
		return errors.New("content not found")
	}
	if err := s.repo.DeleteContent(id); err != nil {
		return err
	}
	delete(s.cache, id)
	return nil
}
