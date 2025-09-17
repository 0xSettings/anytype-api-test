package repository

import (
	"anytype-flow-crud/flow/entities"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// FlowRepo represents the Anytype repository adapter
type FlowRepo struct {
	BaseURL string
	APIKEY  string
}

// NewFlowRepo initializes repo with Anytype MCP base URL + key
func NewFlowRepo() *FlowRepo {
	return &FlowRepo{
		BaseURL: "http://localhost:31009/v1",                    // Anytype MCP JSON-RPC endpoint
		APIKEY:  "ud044Ju8oDxNnni/BGplsR/3LKmqKSMEuyLdvb3y6aE=", // replace if needed
	}
}

// generic RPC caller
func (r FlowRepo) callRPC(method string, params interface{}) (map[string]interface{}, error) {
	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      "1",
		"method":  method,
		"params":  params,
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", r.BaseURL, bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+r.APIKEY)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("failed request %s: %v", method, res.Status)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

// ------------------ SPACE ------------------

// CreateSpace -> calls space.create RPC
func (r FlowRepo) CreateSpace(space entities.Space) (string, error) {
	resp, err := r.callRPC("space.create", map[string]interface{}{
		"name": space.Name,
	})
	if err != nil {
		return "", err
	}
	if result, ok := resp["result"].(map[string]interface{}); ok {
		if id, ok := result["spaceId"].(string); ok {
			return id, nil
		}
	}
	return "", fmt.Errorf("spaceId not found in response")
}

// ------------------ OBJECT / PAGE ------------------

// CreatePage -> calls object.create inside a space
func (r FlowRepo) CreatePage(page entities.Page) (string, error) {
	resp, err := r.callRPC("object.create", map[string]interface{}{
		"spaceId": page.SpaceID,
		"title":   page.Title,
		"type":    "page",
	})
	if err != nil {
		return "", err
	}
	if result, ok := resp["result"].(map[string]interface{}); ok {
		if id, ok := result["objectId"].(string); ok {
			return id, nil
		}
	}
	return "", fmt.Errorf("objectId not found in response")
}

// ------------------ CONTENT ------------------

// CreateContent -> calls object.create with body content
func (r FlowRepo) CreateContent(content entities.Content, spaceID string) (string, error) {
	resp, err := r.callRPC("object.create", map[string]interface{}{
		"spaceId": spaceID,
		"title":   content.Title,
		"body":    content.Body,
		"type":    "content",
	})
	if err != nil {
		return "", err
	}
	if result, ok := resp["result"].(map[string]interface{}); ok {
		if id, ok := result["objectId"].(string); ok {
			return id, nil
		}
	}
	return "", fmt.Errorf("objectId not found in response")
}

// UpdateContent -> calls object.update
func (r FlowRepo) UpdateContent(content entities.Content) error {
	_, err := r.callRPC("object.update", map[string]interface{}{
		"objectId": content.ID,
		"title":    content.Title,
		"body":     content.Body,
	})
	return err
}

// DeleteContent -> calls object.delete
func (r FlowRepo) DeleteContent(id string) error {
	_, err := r.callRPC("object.delete", map[string]interface{}{
		"objectId": id,
	})
	return err
}

// PublishContent -> calls object.publish
func (r FlowRepo) PublishContent(id string) error {
	_, err := r.callRPC("object.publish", map[string]interface{}{
		"objectId": id,
	})
	return err
}
