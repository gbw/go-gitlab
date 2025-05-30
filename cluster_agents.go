//
// Copyright 2022, Timo Furrer <tuxtimo@gmail.com>
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
	ClusterAgentsServiceInterface interface {
		ListAgents(pid any, opt *ListAgentsOptions, options ...RequestOptionFunc) ([]*Agent, *Response, error)
		GetAgent(pid any, id int, options ...RequestOptionFunc) (*Agent, *Response, error)
		RegisterAgent(pid any, opt *RegisterAgentOptions, options ...RequestOptionFunc) (*Agent, *Response, error)
		DeleteAgent(pid any, id int, options ...RequestOptionFunc) (*Response, error)
		ListAgentTokens(pid any, aid int, opt *ListAgentTokensOptions, options ...RequestOptionFunc) ([]*AgentToken, *Response, error)
		GetAgentToken(pid any, aid int, id int, options ...RequestOptionFunc) (*AgentToken, *Response, error)
		CreateAgentToken(pid any, aid int, opt *CreateAgentTokenOptions, options ...RequestOptionFunc) (*AgentToken, *Response, error)
		RevokeAgentToken(pid any, aid int, id int, options ...RequestOptionFunc) (*Response, error)
	}

	// ClusterAgentsService handles communication with the cluster agents related
	// methods of the GitLab API.
	//
	// GitLab API docs: https://docs.gitlab.com/api/cluster_agents/
	ClusterAgentsService struct {
		client *Client
	}
)

var _ ClusterAgentsServiceInterface = (*ClusterAgentsService)(nil)

// Agent represents a GitLab agent for Kubernetes.
//
// GitLab API docs: https://docs.gitlab.com/api/cluster_agents/
type Agent struct {
	ID              int           `json:"id"`
	Name            string        `json:"name"`
	CreatedAt       *time.Time    `json:"created_at"`
	CreatedByUserID int           `json:"created_by_user_id"`
	ConfigProject   ConfigProject `json:"config_project"`
}

type ConfigProject struct {
	ID                int        `json:"id"`
	Description       string     `json:"description"`
	Name              string     `json:"name"`
	NameWithNamespace string     `json:"name_with_namespace"`
	Path              string     `json:"path"`
	PathWithNamespace string     `json:"path_with_namespace"`
	CreatedAt         *time.Time `json:"created_at"`
}

func (a Agent) String() string {
	return Stringify(a)
}

// AgentToken represents a GitLab agent token.
//
// GitLab API docs:
// https://docs.gitlab.com/api/cluster_agents/#list-tokens-for-an-agent
type AgentToken struct {
	ID              int        `json:"id"`
	Name            string     `json:"name"`
	Description     string     `json:"description"`
	AgentID         int        `json:"agent_id"`
	Status          string     `json:"status"`
	CreatedAt       *time.Time `json:"created_at"`
	CreatedByUserID int        `json:"created_by_user_id"`
	LastUsedAt      *time.Time `json:"last_used_at"`
	Token           string     `json:"token"`
}

func (a AgentToken) String() string {
	return Stringify(a)
}

// ListAgentsOptions represents the available ListAgents() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/cluster_agents/#list-the-agents-for-a-project
type ListAgentsOptions ListOptions

// ListAgents returns a list of agents registered for the project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/cluster_agents/#list-the-agents-for-a-project
func (s *ClusterAgentsService) ListAgents(pid any, opt *ListAgentsOptions, options ...RequestOptionFunc) ([]*Agent, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	uri := fmt.Sprintf("projects/%s/cluster_agents", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, uri, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var as []*Agent
	resp, err := s.client.Do(req, &as)
	if err != nil {
		return nil, resp, err
	}

	return as, resp, nil
}

// GetAgent gets a single agent details.
//
// GitLab API docs:
// https://docs.gitlab.com/api/cluster_agents/#get-details-about-an-agent
func (s *ClusterAgentsService) GetAgent(pid any, id int, options ...RequestOptionFunc) (*Agent, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	uri := fmt.Sprintf("projects/%s/cluster_agents/%d", PathEscape(project), id)

	req, err := s.client.NewRequest(http.MethodGet, uri, nil, options)
	if err != nil {
		return nil, nil, err
	}

	a := new(Agent)
	resp, err := s.client.Do(req, a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, nil
}

// RegisterAgentOptions represents the available RegisterAgent()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/cluster_agents/#register-an-agent-with-a-project
type RegisterAgentOptions struct {
	Name *string `url:"name,omitempty" json:"name,omitempty"`
}

// RegisterAgent registers an agent to the project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/cluster_agents/#register-an-agent-with-a-project
func (s *ClusterAgentsService) RegisterAgent(pid any, opt *RegisterAgentOptions, options ...RequestOptionFunc) (*Agent, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	uri := fmt.Sprintf("projects/%s/cluster_agents", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPost, uri, opt, options)
	if err != nil {
		return nil, nil, err
	}

	a := new(Agent)
	resp, err := s.client.Do(req, a)
	if err != nil {
		return nil, resp, err
	}

	return a, resp, nil
}

// DeleteAgent deletes an existing agent registration.
//
// GitLab API docs:
// https://docs.gitlab.com/api/cluster_agents/#delete-a-registered-agent
func (s *ClusterAgentsService) DeleteAgent(pid any, id int, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("projects/%s/cluster_agents/%d", PathEscape(project), id)

	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ListAgentTokensOptions represents the available ListAgentTokens() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/cluster_agents/#list-tokens-for-an-agent
type ListAgentTokensOptions ListOptions

// ListAgentTokens returns a list of tokens for an agent.
//
// GitLab API docs:
// https://docs.gitlab.com/api/cluster_agents/#list-tokens-for-an-agent
func (s *ClusterAgentsService) ListAgentTokens(pid any, aid int, opt *ListAgentTokensOptions, options ...RequestOptionFunc) ([]*AgentToken, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	uri := fmt.Sprintf("projects/%s/cluster_agents/%d/tokens", PathEscape(project), aid)

	req, err := s.client.NewRequest(http.MethodGet, uri, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var ats []*AgentToken
	resp, err := s.client.Do(req, &ats)
	if err != nil {
		return nil, resp, err
	}

	return ats, resp, nil
}

// GetAgentToken gets a single agent token.
//
// GitLab API docs:
// https://docs.gitlab.com/api/cluster_agents/#get-a-single-agent-token
func (s *ClusterAgentsService) GetAgentToken(pid any, aid int, id int, options ...RequestOptionFunc) (*AgentToken, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	uri := fmt.Sprintf("projects/%s/cluster_agents/%d/tokens/%d", PathEscape(project), aid, id)

	req, err := s.client.NewRequest(http.MethodGet, uri, nil, options)
	if err != nil {
		return nil, nil, err
	}

	at := new(AgentToken)
	resp, err := s.client.Do(req, at)
	if err != nil {
		return nil, resp, err
	}

	return at, resp, nil
}

// CreateAgentTokenOptions represents the available CreateAgentToken() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/cluster_agents/#create-an-agent-token
type CreateAgentTokenOptions struct {
	Name        *string `url:"name,omitempty" json:"name,omitempty"`
	Description *string `url:"description,omitempty" json:"description,omitempty"`
}

// CreateAgentToken creates a new token for an agent.
//
// GitLab API docs:
// https://docs.gitlab.com/api/cluster_agents/#create-an-agent-token
func (s *ClusterAgentsService) CreateAgentToken(pid any, aid int, opt *CreateAgentTokenOptions, options ...RequestOptionFunc) (*AgentToken, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	uri := fmt.Sprintf("projects/%s/cluster_agents/%d/tokens", PathEscape(project), aid)

	req, err := s.client.NewRequest(http.MethodPost, uri, opt, options)
	if err != nil {
		return nil, nil, err
	}

	at := new(AgentToken)
	resp, err := s.client.Do(req, at)
	if err != nil {
		return nil, resp, err
	}

	return at, resp, nil
}

// RevokeAgentToken revokes an agent token.
//
// GitLab API docs:
// https://docs.gitlab.com/api/cluster_agents/#revoke-an-agent-token
func (s *ClusterAgentsService) RevokeAgentToken(pid any, aid int, id int, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("projects/%s/cluster_agents/%d/tokens/%d", PathEscape(project), aid, id)

	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
