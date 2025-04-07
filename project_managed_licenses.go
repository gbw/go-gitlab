//
// Copyright 2021, Andrea Perizzato
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
	// Deprecated: no matching docs found
	ManagedLicensesServiceInterface interface {
		// Deprecated: no matching docs found
		ListManagedLicenses(pid interface{}, options ...RequestOptionFunc) ([]*ManagedLicense, *Response, error)
		// Deprecated: no matching docs found
		GetManagedLicense(pid, mlid interface{}, options ...RequestOptionFunc) (*ManagedLicense, *Response, error)
		// Deprecated: no matching docs found
		AddManagedLicense(pid interface{}, opt *AddManagedLicenseOptions, options ...RequestOptionFunc) (*ManagedLicense, *Response, error)
		// Deprecated: no matching docs found
		DeleteManagedLicense(pid, mlid interface{}, options ...RequestOptionFunc) (*Response, error)
		// Deprecated: no matching docs found
		EditManagedLicense(pid, mlid interface{}, opt *EditManagedLicenceOptions, options ...RequestOptionFunc) (*ManagedLicense, *Response, error)
	}

	// ManagedLicensesService handles communication with the managed licenses
	// methods of the GitLab API.
	// Deprecated: no matching docs found
	//
	// GitLab API docs: https://docs.gitlab.com/ee/api/managed_licenses.html
	ManagedLicensesService struct {
		client *Client
	}
)

// Deprecated: no matching docs found
var _ ManagedLicensesServiceInterface = (*ManagedLicensesService)(nil)

// ManagedLicense represents a managed license.
// Deprecated: no matching docs found
//
// GitLab API docs: https://docs.gitlab.com/ee/api/managed_licenses.html
type ManagedLicense struct {
	ID             int                        `json:"id"`
	Name           string                     `json:"name"`
	ApprovalStatus LicenseApprovalStatusValue `json:"approval_status"`
}

// ListManagedLicenses returns a list of managed licenses from a project.
// Deprecated: no matching docs found
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/managed_licenses.html#list-managed-licenses
func (s *ManagedLicensesService) ListManagedLicenses(pid interface{}, options ...RequestOptionFunc) ([]*ManagedLicense, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/managed_licenses", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var mls []*ManagedLicense
	resp, err := s.client.Do(req, &mls)
	if err != nil {
		return nil, resp, err
	}

	return mls, resp, nil
}

// GetManagedLicense returns an existing managed license.
// Deprecated: no matching docs found
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/managed_licenses.html#show-an-existing-managed-license
func (s *ManagedLicensesService) GetManagedLicense(pid, mlid interface{}, options ...RequestOptionFunc) (*ManagedLicense, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	license, err := parseID(mlid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/managed_licenses/%s", PathEscape(project), PathEscape(license))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	ml := new(ManagedLicense)
	resp, err := s.client.Do(req, ml)
	if err != nil {
		return nil, resp, err
	}

	return ml, resp, nil
}

// AddManagedLicenseOptions represents the available AddManagedLicense() options.
// Deprecated: no matching docs found
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/managed_licenses.html#create-a-new-managed-license
type AddManagedLicenseOptions struct {
	Name           *string                     `url:"name,omitempty" json:"name,omitempty"`
	ApprovalStatus *LicenseApprovalStatusValue `url:"approval_status,omitempty" json:"approval_status,omitempty"`
}

// AddManagedLicense adds a managed license to a project.
// Deprecated: no matching docs found
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/managed_licenses.html#create-a-new-managed-license
func (s *ManagedLicensesService) AddManagedLicense(pid interface{}, opt *AddManagedLicenseOptions, options ...RequestOptionFunc) (*ManagedLicense, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/managed_licenses", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	ml := new(ManagedLicense)
	resp, err := s.client.Do(req, ml)
	if err != nil {
		return nil, resp, err
	}

	return ml, resp, nil
}

// DeleteManagedLicense deletes a managed license with a given ID.
// Deprecated: no matching docs found
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/managed_licenses.html#delete-a-managed-license
func (s *ManagedLicensesService) DeleteManagedLicense(pid, mlid interface{}, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	license, err := parseID(mlid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/managed_licenses/%s", PathEscape(project), PathEscape(license))

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// EditManagedLicenceOptions represents the available EditManagedLicense() options.
// Deprecated: no matching docs found
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/managed_licenses.html#edit-an-existing-managed-license
type EditManagedLicenceOptions struct {
	ApprovalStatus *LicenseApprovalStatusValue `url:"approval_status,omitempty" json:"approval_status,omitempty"`
}

// EditManagedLicense updates an existing managed license with a new approval
// status.
// Deprecated: no matching docs found
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/managed_licenses.html#edit-an-existing-managed-license
func (s *ManagedLicensesService) EditManagedLicense(pid, mlid interface{}, opt *EditManagedLicenceOptions, options ...RequestOptionFunc) (*ManagedLicense, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	license, err := parseID(mlid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/managed_licenses/%s", PathEscape(project), PathEscape(license))

	req, err := s.client.NewRequest(http.MethodPatch, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	ml := new(ManagedLicense)
	resp, err := s.client.Do(req, ml)
	if err != nil {
		return nil, resp, err
	}

	return ml, resp, nil
}
