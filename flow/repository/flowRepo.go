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
		BaseURL: "http://localhost:3000/v1",
		APIKEY:  "ud044Ju8oDxNnni/BGplsR/3LKmqKSMEuyLdvb3y6aE=",
	}
}

func (r FlowRepo) ExposeNewContent(content entities.Content) error {
	body, _ := json.Marshal(content)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/content", r.BaseURL), bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer"+r.APIKEY)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(req)

	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		return fmt.Errorf("failed to expose content: %v", response.Status)
	}
	return nil
}
