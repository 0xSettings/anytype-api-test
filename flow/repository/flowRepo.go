package repository

import (
	"anytype-flow-crud/flow/entities"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type FlowRepo struct {
	BaseURL string
	APIKEY  string
}

func NewFlowRepo() *FlowRepo {
	return &FlowRepo{
		BaseURL: "http://localhost:8080/api/v1", // matches main.go server
		APIKEY:  "ud044Ju8oDxNnni/BGplsR/3LKmqKSMEuyLdvb3y6aE=",
	}
}

func (r FlowRepo) post(endpoint string, payload interface{}, out interface{}) error {
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/%s", r.BaseURL, endpoint), bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+r.APIKEY)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return fmt.Errorf("failed request %s: %v", endpoint, res.Status)
	}
	if out != nil {
		return json.NewDecoder(res.Body).Decode(out)
	}
	return nil
}

func (r FlowRepo) put(endpoint string, payload interface{}, out interface{}) error {
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("PUT", fmt.Sprintf("%s/%s", r.BaseURL, endpoint), bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+r.APIKEY)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return fmt.Errorf("failed request %s: %v", endpoint, res.Status)
	}
	if out != nil {
		return json.NewDecoder(res.Body).Decode(out)
	}
	return nil
}

func (r FlowRepo) delete(endpoint string) error {
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/%s", r.BaseURL, endpoint), nil)
	req.Header.Set("Authorization", "Bearer "+r.APIKEY)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return fmt.Errorf("failed delete %s: %v", endpoint, res.Status)
	}
	return nil
}

// Space
func (r FlowRepo) CreateSpace(space entities.Space) (*entities.Space, error) {
	var created entities.Space
	err := r.post("spaces", space, &created)
	return &created, err
}

// Page
func (r FlowRepo) CreatePage(page entities.Page) (*entities.Page, error) {
	var created entities.Page
	err := r.post("objects", page, &created)
	return &created, err
}

// Content
func (r FlowRepo) CreateContent(content entities.Content) (*entities.Content, error) {
	var created entities.Content
	err := r.post("content", content, &created)
	return &created, err
}

func (r FlowRepo) UpdateContent(content entities.Content) (*entities.Content, error) {
	var updated entities.Content
	err := r.put(fmt.Sprintf("content/%s", content.ID), content, &updated)
	return &updated, err
}

func (r FlowRepo) DeleteContent(id string) error {
	return r.delete(fmt.Sprintf("content/%s", id))
}
