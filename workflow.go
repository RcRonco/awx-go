package awx

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// JobTemplateService implements awx job template apis.
type WorkflowService struct {
	client *Client
}

// ListWorkflowsResponse represents `ListWorkflow` endpoint response.
type ListWorkflowsResponse struct {
	Pagination
	Results []*WorkflowJT `json:"results"`
}

// ListWorkflow shows a list of workflow templates.
func (jt *WorkflowService) ListWorkflow(params map[string]string) ([]*WorkflowJT, *ListWorkflowsResponse, error) {
	result := new(ListWorkflowsResponse)
	endpoint := "/api/v2/workflow_job_templates/"
	resp, err := jt.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

// Launch lauchs a job with the workflow job template.
func (jt *WorkflowService) Launch(id int, data map[string]interface{}, params map[string]string) (*JobLaunch, error) {
	result := new(JobLaunch)
	endpoint := fmt.Sprintf("/api/v2/workflow_job_templates/%d/launch/", id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := jt.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	// in case invalid job id return
	if result.Job == 0 {
		return nil, errors.New("invalid job id 0")
	}

	return result, nil
}

// CreateWorkflow creates a workflow job template
func (jt *WorkflowService) CreateWorkflow(data map[string]interface{}, params map[string]string) (*WorkflowJT, error) {
	result := new(WorkflowJT)
	mandatoryFields = []string{"name", "job_type", "inventory", "project"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}
	endpoint := "/api/v2/workflow_job_templates/"
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := jt.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}
	if err := CheckResponse(resp); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateWorkflow updates a workflow job template
func (jt *WorkflowService) UpdateWorkflow(id int, data map[string]interface{}, params map[string]string) (*WorkflowJT, error) {
	result := new(WorkflowJT)
	endpoint := fmt.Sprintf("/api/v2/workflow_job_templates/%d", id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := jt.client.Requester.PatchJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}
	if err := CheckResponse(resp); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteWorkflow deletes a workflow job template
func (jt *WorkflowService) DeleteWorkflow(id int) (*WorkflowJT, error) {
	result := new(WorkflowJT)
	endpoint := fmt.Sprintf("/api/v2/workflow_job_templates/%d", id)

	resp, err := jt.client.Requester.Delete(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

