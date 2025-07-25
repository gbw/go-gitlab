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
)

type (
	ProtectedEnvironmentsServiceInterface interface {
		ListProtectedEnvironments(pid any, opt *ListProtectedEnvironmentsOptions, options ...RequestOptionFunc) ([]*ProtectedEnvironment, *Response, error)
		GetProtectedEnvironment(pid any, environment string, options ...RequestOptionFunc) (*ProtectedEnvironment, *Response, error)
		ProtectRepositoryEnvironments(pid any, opt *ProtectRepositoryEnvironmentsOptions, options ...RequestOptionFunc) (*ProtectedEnvironment, *Response, error)
		UpdateProtectedEnvironments(pid any, environment string, opt *UpdateProtectedEnvironmentsOptions, options ...RequestOptionFunc) (*ProtectedEnvironment, *Response, error)
		UnprotectEnvironment(pid any, environment string, options ...RequestOptionFunc) (*Response, error)
	}

	// ProtectedEnvironmentsService handles communication with the protected
	// environment methods of the GitLab API.
	//
	// GitLab API docs:
	// https://docs.gitlab.com/api/protected_environments/
	ProtectedEnvironmentsService struct {
		client *Client
	}
)

var _ ProtectedEnvironmentsServiceInterface = (*ProtectedEnvironmentsService)(nil)

// ProtectedEnvironment represents a protected environment.
//
// GitLab API docs:
// https://docs.gitlab.com/api/protected_environments/
type ProtectedEnvironment struct {
	Name                  string                          `json:"name"`
	DeployAccessLevels    []*EnvironmentAccessDescription `json:"deploy_access_levels"`
	RequiredApprovalCount int                             `json:"required_approval_count"`
	ApprovalRules         []*EnvironmentApprovalRule      `json:"approval_rules"`
}

// EnvironmentAccessDescription represents the access description for a protected
// environment.
//
// GitLab API docs:
// https://docs.gitlab.com/api/protected_environments/
type EnvironmentAccessDescription struct {
	ID                     int              `json:"id"`
	AccessLevel            AccessLevelValue `json:"access_level"`
	AccessLevelDescription string           `json:"access_level_description"`
	UserID                 int              `json:"user_id"`
	GroupID                int              `json:"group_id"`
	GroupInheritanceType   int              `json:"group_inheritance_type"`
}

// EnvironmentApprovalRule represents the approval rules for a protected
// environment.
//
// GitLab API docs:
// https://docs.gitlab.com/api/protected_environments/#protect-a-single-environment
type EnvironmentApprovalRule struct {
	ID                     int              `json:"id"`
	UserID                 int              `json:"user_id"`
	GroupID                int              `json:"group_id"`
	AccessLevel            AccessLevelValue `json:"access_level"`
	AccessLevelDescription string           `json:"access_level_description"`
	RequiredApprovalCount  int              `json:"required_approvals"`
	GroupInheritanceType   int              `json:"group_inheritance_type"`
}

// ListProtectedEnvironmentsOptions represents the available
// ListProtectedEnvironments() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/protected_environments/#list-protected-environments
type ListProtectedEnvironmentsOptions ListOptions

// ListProtectedEnvironments returns a list of protected environments from a
// project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/protected_environments/#list-protected-environments
func (s *ProtectedEnvironmentsService) ListProtectedEnvironments(pid any, opt *ListProtectedEnvironmentsOptions, options ...RequestOptionFunc) ([]*ProtectedEnvironment, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/protected_environments", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var pes []*ProtectedEnvironment
	resp, err := s.client.Do(req, &pes)
	if err != nil {
		return nil, resp, err
	}

	return pes, resp, nil
}

// GetProtectedEnvironment returns a single protected environment or wildcard
// protected environment.
//
// GitLab API docs:
// https://docs.gitlab.com/api/protected_environments/#get-a-single-protected-environment
func (s *ProtectedEnvironmentsService) GetProtectedEnvironment(pid any, environment string, options ...RequestOptionFunc) (*ProtectedEnvironment, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/protected_environments/%s", PathEscape(project), PathEscape(environment))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	pe := new(ProtectedEnvironment)
	resp, err := s.client.Do(req, pe)
	if err != nil {
		return nil, resp, err
	}

	return pe, resp, nil
}

// ProtectRepositoryEnvironmentsOptions represents the available
// ProtectRepositoryEnvironments() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/protected_environments/#protect-a-single-environment
type ProtectRepositoryEnvironmentsOptions struct {
	Name                  *string                            `url:"name,omitempty" json:"name,omitempty"`
	DeployAccessLevels    *[]*EnvironmentAccessOptions       `url:"deploy_access_levels,omitempty" json:"deploy_access_levels,omitempty"`
	RequiredApprovalCount *int                               `url:"required_approval_count,omitempty" json:"required_approval_count,omitempty"`
	ApprovalRules         *[]*EnvironmentApprovalRuleOptions `url:"approval_rules,omitempty" json:"approval_rules,omitempty"`
}

// EnvironmentAccessOptions represents the options for an access description for
// a protected environment.
//
// GitLab API docs:
// https://docs.gitlab.com/api/protected_environments/#protect-a-single-environment
type EnvironmentAccessOptions struct {
	AccessLevel          *AccessLevelValue `url:"access_level,omitempty" json:"access_level,omitempty"`
	UserID               *int              `url:"user_id,omitempty" json:"user_id,omitempty"`
	GroupID              *int              `url:"group_id,omitempty" json:"group_id,omitempty"`
	GroupInheritanceType *int              `url:"group_inheritance_type,omitempty" json:"group_inheritance_type,omitempty"`
}

// EnvironmentApprovalRuleOptions represents the approval rules for a protected
// environment.
//
// GitLab API docs:
// https://docs.gitlab.com/api/protected_environments/#protect-a-single-environment
type EnvironmentApprovalRuleOptions struct {
	UserID                 *int              `url:"user_id,omitempty" json:"user_id,omitempty"`
	GroupID                *int              `url:"group_id,omitempty" json:"group_id,omitempty"`
	AccessLevel            *AccessLevelValue `url:"access_level,omitempty" json:"access_level,omitempty"`
	AccessLevelDescription *string           `url:"access_level_description,omitempty" json:"access_level_description,omitempty"`
	RequiredApprovalCount  *int              `url:"required_approvals,omitempty" json:"required_approvals,omitempty"`
	GroupInheritanceType   *int              `url:"group_inheritance_type,omitempty" json:"group_inheritance_type,omitempty"`
}

// ProtectRepositoryEnvironments protects a single repository environment or
// several project repository environments using wildcard protected environment.
//
// GitLab API docs:
// https://docs.gitlab.com/api/protected_environments/#protect-a-single-environment
func (s *ProtectedEnvironmentsService) ProtectRepositoryEnvironments(pid any, opt *ProtectRepositoryEnvironmentsOptions, options ...RequestOptionFunc) (*ProtectedEnvironment, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/protected_environments", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	pe := new(ProtectedEnvironment)
	resp, err := s.client.Do(req, pe)
	if err != nil {
		return nil, resp, err
	}

	return pe, resp, nil
}

// UpdateProtectedEnvironmentsOptions represents the available
// UpdateProtectedEnvironments() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/protected_environments/#update-a-protected-environment
type UpdateProtectedEnvironmentsOptions struct {
	Name                  *string                                  `url:"name,omitempty" json:"name,omitempty"`
	DeployAccessLevels    *[]*UpdateEnvironmentAccessOptions       `url:"deploy_access_levels,omitempty" json:"deploy_access_levels,omitempty"`
	RequiredApprovalCount *int                                     `url:"required_approval_count,omitempty" json:"required_approval_count,omitempty"`
	ApprovalRules         *[]*UpdateEnvironmentApprovalRuleOptions `url:"approval_rules,omitempty" json:"approval_rules,omitempty"`
}

// UpdateEnvironmentAccessOptions represents the options for updates to an
// access description for a protected environment.
//
// GitLab API docs:
// https://docs.gitlab.com/api/protected_environments/#update-a-protected-environment
type UpdateEnvironmentAccessOptions struct {
	AccessLevel          *AccessLevelValue `url:"access_level,omitempty" json:"access_level,omitempty"`
	ID                   *int              `url:"id,omitempty" json:"id,omitempty"`
	UserID               *int              `url:"user_id,omitempty" json:"user_id,omitempty"`
	GroupID              *int              `url:"group_id,omitempty" json:"group_id,omitempty"`
	GroupInheritanceType *int              `url:"group_inheritance_type,omitempty" json:"group_inheritance_type,omitempty"`
	Destroy              *bool             `url:"_destroy,omitempty" json:"_destroy,omitempty"`
}

// UpdateEnvironmentApprovalRuleOptions represents the updates to the approval
// rules for a protected environment.
//
// GitLab API docs:
// https://docs.gitlab.com/api/protected_environments/#update-a-protected-environment
type UpdateEnvironmentApprovalRuleOptions struct {
	ID                     *int              `url:"id,omitempty" json:"id,omitempty"`
	UserID                 *int              `url:"user_id,omitempty" json:"user_id,omitempty"`
	GroupID                *int              `url:"group_id,omitempty" json:"group_id,omitempty"`
	AccessLevel            *AccessLevelValue `url:"access_level,omitempty" json:"access_level,omitempty"`
	AccessLevelDescription *string           `url:"access_level_description,omitempty" json:"access_level_description,omitempty"`
	RequiredApprovalCount  *int              `url:"required_approvals,omitempty" json:"required_approvals,omitempty"`
	GroupInheritanceType   *int              `url:"group_inheritance_type,omitempty" json:"group_inheritance_type,omitempty"`
	Destroy                *bool             `url:"_destroy,omitempty" json:"_destroy,omitempty"`
}

// UpdateProtectedEnvironments updates a single repository environment or
// several project repository environments using wildcard protected environment.
//
// GitLab API docs:
// https://docs.gitlab.com/api/protected_environments/#update-a-protected-environment
func (s *ProtectedEnvironmentsService) UpdateProtectedEnvironments(pid any, environment string, opt *UpdateProtectedEnvironmentsOptions, options ...RequestOptionFunc) (*ProtectedEnvironment, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/protected_environments/%s", PathEscape(project), PathEscape(environment))

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	pe := new(ProtectedEnvironment)
	resp, err := s.client.Do(req, pe)
	if err != nil {
		return nil, resp, err
	}

	return pe, resp, nil
}

// UnprotectEnvironment unprotects the given protected environment or wildcard
// protected environment.
//
// GitLab API docs:
// https://docs.gitlab.com/api/protected_environments/#unprotect-a-single-environment
func (s *ProtectedEnvironmentsService) UnprotectEnvironment(pid any, environment string, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/protected_environments/%s", PathEscape(project), PathEscape(environment))

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
