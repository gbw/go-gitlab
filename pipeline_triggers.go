//
// Copyright 2021, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package gitlab

import (
	"fmt"
	"net/http"
	"time"
)

type (
	PipelineTriggersServiceInterface interface {
		ListPipelineTriggers(pid any, opt *ListPipelineTriggersOptions, options ...RequestOptionFunc) ([]*PipelineTrigger, *Response, error)
		GetPipelineTrigger(pid any, trigger int, options ...RequestOptionFunc) (*PipelineTrigger, *Response, error)
		AddPipelineTrigger(pid any, opt *AddPipelineTriggerOptions, options ...RequestOptionFunc) (*PipelineTrigger, *Response, error)
		EditPipelineTrigger(pid any, trigger int, opt *EditPipelineTriggerOptions, options ...RequestOptionFunc) (*PipelineTrigger, *Response, error)
		DeletePipelineTrigger(pid any, trigger int, options ...RequestOptionFunc) (*Response, error)
		RunPipelineTrigger(pid any, opt *RunPipelineTriggerOptions, options ...RequestOptionFunc) (*Pipeline, *Response, error)
	}

	// PipelineTriggersService handles Project pipeline triggers.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/api/pipeline_triggers/
	PipelineTriggersService struct {
		client *Client
	}
)

var _ PipelineTriggersServiceInterface = (*PipelineTriggersService)(nil)

// PipelineTrigger represents a project pipeline trigger.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pipeline_triggers/
type PipelineTrigger struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	CreatedAt   *time.Time `json:"created_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	LastUsed    *time.Time `json:"last_used"`
	Token       string     `json:"token"`
	UpdatedAt   *time.Time `json:"updated_at"`
	Owner       *User      `json:"owner"`
}

// ListPipelineTriggersOptions represents the available ListPipelineTriggers() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pipeline_triggers/#list-project-trigger-tokens
type ListPipelineTriggersOptions ListOptions

// ListPipelineTriggers gets a list of project triggers.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pipeline_triggers/#list-project-trigger-tokens
func (s *PipelineTriggersService) ListPipelineTriggers(pid any, opt *ListPipelineTriggersOptions, options ...RequestOptionFunc) ([]*PipelineTrigger, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/triggers", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var pt []*PipelineTrigger
	resp, err := s.client.Do(req, &pt)
	if err != nil {
		return nil, resp, err
	}

	return pt, resp, nil
}

// GetPipelineTrigger gets a specific pipeline trigger for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pipeline_triggers/#get-trigger-token-details
func (s *PipelineTriggersService) GetPipelineTrigger(pid any, trigger int, options ...RequestOptionFunc) (*PipelineTrigger, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/triggers/%d", PathEscape(project), trigger)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	pt := new(PipelineTrigger)
	resp, err := s.client.Do(req, pt)
	if err != nil {
		return nil, resp, err
	}

	return pt, resp, nil
}

// AddPipelineTriggerOptions represents the available AddPipelineTrigger() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pipeline_triggers/#create-a-trigger-token
type AddPipelineTriggerOptions struct {
	Description *string `url:"description,omitempty" json:"description,omitempty"`
}

// AddPipelineTrigger adds a pipeline trigger to a specified project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pipeline_triggers/#create-a-trigger-token
func (s *PipelineTriggersService) AddPipelineTrigger(pid any, opt *AddPipelineTriggerOptions, options ...RequestOptionFunc) (*PipelineTrigger, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/triggers", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	pt := new(PipelineTrigger)
	resp, err := s.client.Do(req, pt)
	if err != nil {
		return nil, resp, err
	}

	return pt, resp, nil
}

// EditPipelineTriggerOptions represents the available EditPipelineTrigger() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pipeline_triggers/#update-a-pipeline-trigger-token
type EditPipelineTriggerOptions struct {
	Description *string `url:"description,omitempty" json:"description,omitempty"`
}

// EditPipelineTrigger edits a trigger for a specified project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pipeline_triggers/#update-a-pipeline-trigger-token
func (s *PipelineTriggersService) EditPipelineTrigger(pid any, trigger int, opt *EditPipelineTriggerOptions, options ...RequestOptionFunc) (*PipelineTrigger, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/triggers/%d", PathEscape(project), trigger)

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	pt := new(PipelineTrigger)
	resp, err := s.client.Do(req, pt)
	if err != nil {
		return nil, resp, err
	}

	return pt, resp, nil
}

// DeletePipelineTrigger removes a trigger from a project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pipeline_triggers/#remove-a-pipeline-trigger-token
func (s *PipelineTriggersService) DeletePipelineTrigger(pid any, trigger int, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/triggers/%d", PathEscape(project), trigger)

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// RunPipelineTriggerOptions represents the available RunPipelineTrigger() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pipeline_triggers/#trigger-a-pipeline-with-a-token
type RunPipelineTriggerOptions struct {
	Ref       *string           `url:"ref" json:"ref"`
	Token     *string           `url:"token" json:"token"`
	Variables map[string]string `url:"variables,omitempty" json:"variables,omitempty"`
}

// RunPipelineTrigger starts a trigger from a project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pipeline_triggers/#trigger-a-pipeline-with-a-token
func (s *PipelineTriggersService) RunPipelineTrigger(pid any, opt *RunPipelineTriggerOptions, options ...RequestOptionFunc) (*Pipeline, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/trigger/pipeline", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	pt := new(Pipeline)
	resp, err := s.client.Do(req, pt)
	if err != nil {
		return nil, resp, err
	}

	return pt, resp, nil
}
