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
	ProjectMirrorServiceInterface interface {
		ListProjectMirror(pid any, opt *ListProjectMirrorOptions, options ...RequestOptionFunc) ([]*ProjectMirror, *Response, error)
		GetProjectMirror(pid any, mirror int, options ...RequestOptionFunc) (*ProjectMirror, *Response, error)
		GetProjectMirrorPublicKey(pid any, mirror int, options ...RequestOptionFunc) (*ProjectMirrorPublicKey, *Response, error)
		AddProjectMirror(pid any, opt *AddProjectMirrorOptions, options ...RequestOptionFunc) (*ProjectMirror, *Response, error)
		EditProjectMirror(pid any, mirror int, opt *EditProjectMirrorOptions, options ...RequestOptionFunc) (*ProjectMirror, *Response, error)
		DeleteProjectMirror(pid any, mirror int, options ...RequestOptionFunc) (*Response, error)
	}

	// ProjectMirrorService handles communication with the project mirror
	// related methods of the GitLab API.
	//
	// GitLAb API docs: https://docs.gitlab.com/api/remote_mirrors/
	ProjectMirrorService struct {
		client *Client
	}
)

var _ ProjectMirrorServiceInterface = (*ProjectMirrorService)(nil)

// ProjectMirror represents a project mirror configuration.
//
// GitLAb API docs: https://docs.gitlab.com/api/remote_mirrors/
type ProjectMirror struct {
	Enabled                bool       `json:"enabled"`
	ID                     int        `json:"id"`
	LastError              string     `json:"last_error"`
	LastSuccessfulUpdateAt *time.Time `json:"last_successful_update_at"`
	LastUpdateAt           *time.Time `json:"last_update_at"`
	LastUpdateStartedAt    *time.Time `json:"last_update_started_at"`
	MirrorBranchRegex      string     `json:"mirror_branch_regex"`
	OnlyProtectedBranches  bool       `json:"only_protected_branches"`
	KeepDivergentRefs      bool       `json:"keep_divergent_refs"`
	UpdateStatus           string     `json:"update_status"`
	URL                    string     `json:"url"`
	AuthMethod             string     `json:"auth_method"`
}

type ProjectMirrorPublicKey struct {
	PublicKey string `json:"public_key"`
}

// ListProjectMirrorOptions represents the available ListProjectMirror() options.
type ListProjectMirrorOptions ListOptions

// ListProjectMirror gets a list of mirrors configured on the project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/remote_mirrors/#list-a-projects-remote-mirrors
func (s *ProjectMirrorService) ListProjectMirror(pid any, opt *ListProjectMirrorOptions, options ...RequestOptionFunc) ([]*ProjectMirror, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/remote_mirrors", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var pm []*ProjectMirror
	resp, err := s.client.Do(req, &pm)
	if err != nil {
		return nil, resp, err
	}

	return pm, resp, nil
}

// GetProjectMirror gets a single mirror configured on the project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/remote_mirrors/#get-a-single-projects-remote-mirror
func (s *ProjectMirrorService) GetProjectMirror(pid any, mirror int, options ...RequestOptionFunc) (*ProjectMirror, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/remote_mirrors/%d", PathEscape(project), mirror)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	pm := new(ProjectMirror)
	resp, err := s.client.Do(req, &pm)
	if err != nil {
		return nil, resp, err
	}

	return pm, resp, nil
}

// GetProjectMirrorPublicKey gets the SSH public key for a single mirror configured on the project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/remote_mirrors/#get-a-single-projects-remote-mirror-public-key
func (s *ProjectMirrorService) GetProjectMirrorPublicKey(pid any, mirror int, options ...RequestOptionFunc) (*ProjectMirrorPublicKey, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/remote_mirrors/%d/public_key", PathEscape(project), mirror)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	pmpk := new(ProjectMirrorPublicKey)
	resp, err := s.client.Do(req, &pmpk)
	if err != nil {
		return nil, resp, err
	}

	return pmpk, resp, nil
}

// AddProjectMirrorOptions contains the properties requires to create
// a new project mirror.
//
// GitLab API docs:
// https://docs.gitlab.com/api/remote_mirrors/#create-a-push-mirror
type AddProjectMirrorOptions struct {
	URL                   *string `url:"url,omitempty" json:"url,omitempty"`
	Enabled               *bool   `url:"enabled,omitempty" json:"enabled,omitempty"`
	KeepDivergentRefs     *bool   `url:"keep_divergent_refs,omitempty" json:"keep_divergent_refs,omitempty"`
	OnlyProtectedBranches *bool   `url:"only_protected_branches,omitempty" json:"only_protected_branches,omitempty"`
	MirrorBranchRegex     *string `url:"mirror_branch_regex,omitempty" json:"mirror_branch_regex,omitempty"`
	AuthMethod            *string `url:"auth_method,omitempty" json:"auth_method,omitempty"`
}

// AddProjectMirror creates a new mirror on the project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/remote_mirrors/#create-a-push-mirror
func (s *ProjectMirrorService) AddProjectMirror(pid any, opt *AddProjectMirrorOptions, options ...RequestOptionFunc) (*ProjectMirror, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/remote_mirrors", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	pm := new(ProjectMirror)
	resp, err := s.client.Do(req, pm)
	if err != nil {
		return nil, resp, err
	}

	return pm, resp, nil
}

// EditProjectMirrorOptions contains the properties requires to edit
// an existing project mirror.
//
// GitLab API docs:
// https://docs.gitlab.com/api/remote_mirrors/#update-a-remote-mirrors-attributes
type EditProjectMirrorOptions struct {
	Enabled               *bool   `url:"enabled,omitempty" json:"enabled,omitempty"`
	KeepDivergentRefs     *bool   `url:"keep_divergent_refs,omitempty" json:"keep_divergent_refs,omitempty"`
	OnlyProtectedBranches *bool   `url:"only_protected_branches,omitempty" json:"only_protected_branches,omitempty"`
	MirrorBranchRegex     *string `url:"mirror_branch_regex,omitempty" json:"mirror_branch_regex,omitempty"`
	AuthMethod            *string `url:"auth_method,omitempty" json:"auth_method,omitempty"`
}

// EditProjectMirror updates a project team member to a specified access level..
//
// GitLab API docs:
// https://docs.gitlab.com/api/remote_mirrors/#update-a-remote-mirrors-attributes
func (s *ProjectMirrorService) EditProjectMirror(pid any, mirror int, opt *EditProjectMirrorOptions, options ...RequestOptionFunc) (*ProjectMirror, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/remote_mirrors/%d", PathEscape(project), mirror)

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	pm := new(ProjectMirror)
	resp, err := s.client.Do(req, pm)
	if err != nil {
		return nil, resp, err
	}

	return pm, resp, nil
}

// DeleteProjectMirror deletes a project mirror.
//
// GitLab API docs:
// https://docs.gitlab.com/api/remote_mirrors/#delete-a-remote-mirror
func (s *ProjectMirrorService) DeleteProjectMirror(pid any, mirror int, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/remote_mirrors/%d", PathEscape(project), mirror)

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
