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
		BaseURL: "http://localhost:31009/v1",
		APIKEY:  "ud044Ju8oDxNnni/BGplsR/3LKmqKSMEuyLdvb3y6aE=",
	}
}

func (r FlowRepo) post(endpoint string, payload interface{}) error {
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
	return nil
}

func (r FlowRepo) CreateSpace(space entities.Space) error {
	return r.post("spaces", space)
}

func (r FlowRepo) CreatePage(page entities.Page) error {
	return r.post("pages", page)
}

func (r FlowRepo) CreateContent(content entities.Content) error {
	return r.post("content", content)
}
