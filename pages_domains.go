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
	PagesDomainsServiceInterface interface {
		ListPagesDomains(pid any, opt *ListPagesDomainsOptions, options ...RequestOptionFunc) ([]*PagesDomain, *Response, error)
		ListAllPagesDomains(options ...RequestOptionFunc) ([]*PagesDomain, *Response, error)
		GetPagesDomain(pid any, domain string, options ...RequestOptionFunc) (*PagesDomain, *Response, error)
		CreatePagesDomain(pid any, opt *CreatePagesDomainOptions, options ...RequestOptionFunc) (*PagesDomain, *Response, error)
		UpdatePagesDomain(pid any, domain string, opt *UpdatePagesDomainOptions, options ...RequestOptionFunc) (*PagesDomain, *Response, error)
		DeletePagesDomain(pid any, domain string, options ...RequestOptionFunc) (*Response, error)
	}

	// PagesDomainsService handles communication with the pages domains
	// related methods of the GitLab API.
	//
	// GitLab API docs: https://docs.gitlab.com/api/pages_domains/
	PagesDomainsService struct {
		client *Client
	}
)

var _ PagesDomainsServiceInterface = (*PagesDomainsService)(nil)

// PagesDomain represents a pages domain.
//
// GitLab API docs: https://docs.gitlab.com/api/pages_domains/
type PagesDomain struct {
	Domain           string     `json:"domain"`
	AutoSslEnabled   bool       `json:"auto_ssl_enabled"`
	URL              string     `json:"url"`
	ProjectID        int        `json:"project_id"`
	Verified         bool       `json:"verified"`
	VerificationCode string     `json:"verification_code"`
	EnabledUntil     *time.Time `json:"enabled_until"`
	Certificate      struct {
		Subject         string     `json:"subject"`
		Expired         bool       `json:"expired"`
		Expiration      *time.Time `json:"expiration"`
		Certificate     string     `json:"certificate"`
		CertificateText string     `json:"certificate_text"`
	} `json:"certificate"`
}

// ListPagesDomainsOptions represents the available ListPagesDomains() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pages_domains/#list-pages-domains
type ListPagesDomainsOptions ListOptions

// ListPagesDomains gets a list of project pages domains.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pages_domains/#list-pages-domains
func (s *PagesDomainsService) ListPagesDomains(pid any, opt *ListPagesDomainsOptions, options ...RequestOptionFunc) ([]*PagesDomain, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/pages/domains", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var pd []*PagesDomain
	resp, err := s.client.Do(req, &pd)
	if err != nil {
		return nil, resp, err
	}

	return pd, resp, nil
}

// ListAllPagesDomains gets a list of all pages domains.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pages_domains/#list-all-pages-domains
func (s *PagesDomainsService) ListAllPagesDomains(options ...RequestOptionFunc) ([]*PagesDomain, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "pages/domains", nil, options)
	if err != nil {
		return nil, nil, err
	}

	var pd []*PagesDomain
	resp, err := s.client.Do(req, &pd)
	if err != nil {
		return nil, resp, err
	}

	return pd, resp, nil
}

// GetPagesDomain get a specific pages domain for a project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pages_domains/#single-pages-domain
func (s *PagesDomainsService) GetPagesDomain(pid any, domain string, options ...RequestOptionFunc) (*PagesDomain, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/pages/domains/%s", PathEscape(project), domain)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	pd := new(PagesDomain)
	resp, err := s.client.Do(req, pd)
	if err != nil {
		return nil, resp, err
	}

	return pd, resp, nil
}

// CreatePagesDomainOptions represents the available CreatePagesDomain() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pages_domains/#create-new-pages-domain
type CreatePagesDomainOptions struct {
	Domain         *string `url:"domain,omitempty" json:"domain,omitempty"`
	AutoSslEnabled *bool   `url:"auto_ssl_enabled,omitempty" json:"auto_ssl_enabled,omitempty"`
	Certificate    *string `url:"certificate,omitempty" json:"certificate,omitempty"`
	Key            *string `url:"key,omitempty" json:"key,omitempty"`
}

// CreatePagesDomain creates a new project pages domain.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pages_domains/#create-new-pages-domain
func (s *PagesDomainsService) CreatePagesDomain(pid any, opt *CreatePagesDomainOptions, options ...RequestOptionFunc) (*PagesDomain, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/pages/domains", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	pd := new(PagesDomain)
	resp, err := s.client.Do(req, pd)
	if err != nil {
		return nil, resp, err
	}

	return pd, resp, nil
}

// UpdatePagesDomainOptions represents the available UpdatePagesDomain() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pages_domains/#update-pages-domain
type UpdatePagesDomainOptions struct {
	AutoSslEnabled *bool   `url:"auto_ssl_enabled,omitempty" json:"auto_ssl_enabled,omitempty"`
	Certificate    *string `url:"certificate,omitempty" json:"certificate,omitempty"`
	Key            *string `url:"key,omitempty" json:"key,omitempty"`
}

// UpdatePagesDomain updates an existing project pages domain.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pages_domains/#update-pages-domain
func (s *PagesDomainsService) UpdatePagesDomain(pid any, domain string, opt *UpdatePagesDomainOptions, options ...RequestOptionFunc) (*PagesDomain, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/pages/domains/%s", PathEscape(project), domain)

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	pd := new(PagesDomain)
	resp, err := s.client.Do(req, pd)
	if err != nil {
		return nil, resp, err
	}

	return pd, resp, nil
}

// DeletePagesDomain deletes an existing project pages domain.
//
// GitLab API docs:
// https://docs.gitlab.com/api/pages_domains/#delete-pages-domain
func (s *PagesDomainsService) DeletePagesDomain(pid any, domain string, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/pages/domains/%s", PathEscape(project), domain)

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
