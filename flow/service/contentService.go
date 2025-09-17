package service

import (
	"anytype-flow-crud/flow/entities"
	"anytype-flow-crud/flow/repository"
	"errors"
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

func (s ContentService) Create(content *entities.Content) error {
	s.cache[content.ID] = content
	return s.rep.ExposeNewContent(*content)
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
