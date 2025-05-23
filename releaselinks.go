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
	ReleaseLinksServiceInterface interface {
		ListReleaseLinks(pid any, tagName string, opt *ListReleaseLinksOptions, options ...RequestOptionFunc) ([]*ReleaseLink, *Response, error)
		GetReleaseLink(pid any, tagName string, link int, options ...RequestOptionFunc) (*ReleaseLink, *Response, error)
		CreateReleaseLink(pid any, tagName string, opt *CreateReleaseLinkOptions, options ...RequestOptionFunc) (*ReleaseLink, *Response, error)
		UpdateReleaseLink(pid any, tagName string, link int, opt *UpdateReleaseLinkOptions, options ...RequestOptionFunc) (*ReleaseLink, *Response, error)
		DeleteReleaseLink(pid any, tagName string, link int, options ...RequestOptionFunc) (*ReleaseLink, *Response, error)
	}

	// ReleaseLinksService handles communication with the release link methods
	// of the GitLab API.
	//
	// GitLab API docs: https://docs.gitlab.com/api/releases/links/
	ReleaseLinksService struct {
		client *Client
	}
)

var _ ReleaseLinksServiceInterface = (*ReleaseLinksService)(nil)

// ReleaseLink represents a release link.
//
// GitLab API docs: https://docs.gitlab.com/api/releases/links/
type ReleaseLink struct {
	ID             int           `json:"id"`
	Name           string        `json:"name"`
	URL            string        `json:"url"`
	DirectAssetURL string        `json:"direct_asset_url"`
	External       bool          `json:"external"`
	LinkType       LinkTypeValue `json:"link_type"`
}

// ListReleaseLinksOptions represents ListReleaseLinks() options.
//
// GitLab API docs: https://docs.gitlab.com/api/releases/links/#list-links-of-a-release
type ListReleaseLinksOptions ListOptions

// ListReleaseLinks gets assets as links from a Release.
//
// GitLab API docs: https://docs.gitlab.com/api/releases/links/#list-links-of-a-release
func (s *ReleaseLinksService) ListReleaseLinks(pid any, tagName string, opt *ListReleaseLinksOptions, options ...RequestOptionFunc) ([]*ReleaseLink, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/releases/%s/assets/links", PathEscape(project), PathEscape(tagName))

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var rls []*ReleaseLink
	resp, err := s.client.Do(req, &rls)
	if err != nil {
		return nil, resp, err
	}

	return rls, resp, nil
}

// GetReleaseLink returns a link from release assets.
//
// GitLab API docs: https://docs.gitlab.com/api/releases/links/#get-a-release-link
func (s *ReleaseLinksService) GetReleaseLink(pid any, tagName string, link int, options ...RequestOptionFunc) (*ReleaseLink, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/releases/%s/assets/links/%d",
		PathEscape(project),
		PathEscape(tagName),
		link)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	rl := new(ReleaseLink)
	resp, err := s.client.Do(req, rl)
	if err != nil {
		return nil, resp, err
	}

	return rl, resp, nil
}

// CreateReleaseLinkOptions represents CreateReleaseLink() options.
//
// GitLab API docs: https://docs.gitlab.com/api/releases/links/#create-a-release-link
type CreateReleaseLinkOptions struct {
	Name            *string        `url:"name,omitempty" json:"name,omitempty"`
	URL             *string        `url:"url,omitempty" json:"url,omitempty"`
	FilePath        *string        `url:"filepath,omitempty" json:"filepath,omitempty"`
	DirectAssetPath *string        `url:"direct_asset_path,omitempty" json:"direct_asset_path,omitempty"`
	LinkType        *LinkTypeValue `url:"link_type,omitempty" json:"link_type,omitempty"`
}

// CreateReleaseLink creates a link.
//
// GitLab API docs: https://docs.gitlab.com/api/releases/links/#create-a-release-link
func (s *ReleaseLinksService) CreateReleaseLink(pid any, tagName string, opt *CreateReleaseLinkOptions, options ...RequestOptionFunc) (*ReleaseLink, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/releases/%s/assets/links", PathEscape(project), PathEscape(tagName))

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	rl := new(ReleaseLink)
	resp, err := s.client.Do(req, rl)
	if err != nil {
		return nil, resp, err
	}

	return rl, resp, nil
}

// UpdateReleaseLinkOptions represents UpdateReleaseLink() options.
//
// You have to specify at least one of Name of URL.
//
// GitLab API docs: https://docs.gitlab.com/api/releases/links/#update-a-release-link
type UpdateReleaseLinkOptions struct {
	Name            *string        `url:"name,omitempty" json:"name,omitempty"`
	URL             *string        `url:"url,omitempty" json:"url,omitempty"`
	FilePath        *string        `url:"filepath,omitempty" json:"filepath,omitempty"`
	DirectAssetPath *string        `url:"direct_asset_path,omitempty" json:"direct_asset_path,omitempty"`
	LinkType        *LinkTypeValue `url:"link_type,omitempty" json:"link_type,omitempty"`
}

// UpdateReleaseLink updates an asset link.
//
// GitLab API docs: https://docs.gitlab.com/api/releases/links/#update-a-release-link
func (s *ReleaseLinksService) UpdateReleaseLink(pid any, tagName string, link int, opt *UpdateReleaseLinkOptions, options ...RequestOptionFunc) (*ReleaseLink, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/releases/%s/assets/links/%d",
		PathEscape(project),
		PathEscape(tagName),
		link)

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	rl := new(ReleaseLink)
	resp, err := s.client.Do(req, rl)
	if err != nil {
		return nil, resp, err
	}

	return rl, resp, nil
}

// DeleteReleaseLink deletes a link from release.
//
// GitLab API docs: https://docs.gitlab.com/api/releases/links/#delete-a-release-link
func (s *ReleaseLinksService) DeleteReleaseLink(pid any, tagName string, link int, options ...RequestOptionFunc) (*ReleaseLink, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/releases/%s/assets/links/%d",
		PathEscape(project),
		PathEscape(tagName),
		link,
	)

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	rl := new(ReleaseLink)
	resp, err := s.client.Do(req, rl)
	if err != nil {
		return nil, resp, err
	}

	return rl, resp, nil
}
