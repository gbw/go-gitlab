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
	ReleasesServiceInterface interface {
		ListReleases(pid any, opt *ListReleasesOptions, options ...RequestOptionFunc) ([]*Release, *Response, error)
		GetRelease(pid any, tagName string, options ...RequestOptionFunc) (*Release, *Response, error)
		GetLatestRelease(pid any, options ...RequestOptionFunc) (*Release, *Response, error)
		CreateRelease(pid any, opts *CreateReleaseOptions, options ...RequestOptionFunc) (*Release, *Response, error)
		UpdateRelease(pid any, tagName string, opts *UpdateReleaseOptions, options ...RequestOptionFunc) (*Release, *Response, error)
		DeleteRelease(pid any, tagName string, options ...RequestOptionFunc) (*Release, *Response, error)
	}

	// ReleasesService handles communication with the releases methods
	// of the GitLab API.
	//
	// GitLab API docs: https://docs.gitlab.com/api/releases/
	ReleasesService struct {
		client *Client
	}
)

var _ ReleasesServiceInterface = (*ReleasesService)(nil)

// Release represents a project release.
//
// GitLab API docs:
// https://docs.gitlab.com/api/releases/#list-releases
type Release struct {
	TagName         string     `json:"tag_name"`
	Name            string     `json:"name"`
	Description     string     `json:"description"`
	DescriptionHTML string     `json:"description_html"`
	CreatedAt       *time.Time `json:"created_at"`
	ReleasedAt      *time.Time `json:"released_at"`
	Author          struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		State     string `json:"state"`
		AvatarURL string `json:"avatar_url"`
		WebURL    string `json:"web_url"`
	} `json:"author"`
	Commit          Commit              `json:"commit"`
	Milestones      []*ReleaseMilestone `json:"milestones"`
	UpcomingRelease bool                `json:"upcoming_release"`
	CommitPath      string              `json:"commit_path"`
	TagPath         string              `json:"tag_path"`
	Assets          struct {
		Count   int `json:"count"`
		Sources []struct {
			Format string `json:"format"`
			URL    string `json:"url"`
		} `json:"sources"`
		Links            []*ReleaseLink `json:"links"`
		EvidenceFilePath string         `json:"evidence_file_path"`
	} `json:"assets"`
	Evidences []*ReleaseEvidence `json:"evidences"`
	Links     struct {
		ClosedIssueURL     string `json:"closed_issues_url"`
		ClosedMergeRequest string `json:"closed_merge_requests_url"`
		EditURL            string `json:"edit_url"`
		MergedMergeRequest string `json:"merged_merge_requests_url"`
		OpenedIssues       string `json:"opened_issues_url"`
		OpenedMergeRequest string `json:"opened_merge_requests_url"`
		Self               string `json:"self"`
	} `json:"_links"`
}

// ReleaseMilestone represents a project release milestone.
//
// GitLab API docs:
// https://docs.gitlab.com/api/releases/#list-releases
type ReleaseMilestone struct {
	ID          int                         `json:"id"`
	IID         int                         `json:"iid"`
	ProjectID   int                         `json:"project_id"`
	Title       string                      `json:"title"`
	Description string                      `json:"description"`
	State       string                      `json:"state"`
	CreatedAt   *time.Time                  `json:"created_at"`
	UpdatedAt   *time.Time                  `json:"updated_at"`
	DueDate     *ISOTime                    `json:"due_date"`
	StartDate   *ISOTime                    `json:"start_date"`
	WebURL      string                      `json:"web_url"`
	IssueStats  *ReleaseMilestoneIssueStats `json:"issue_stats"`
}

// ReleaseMilestoneIssueStats represents a project release milestone's
// related issues statistics.
//
// GitLab API docs:
// https://docs.gitlab.com/api/releases/#list-releases
type ReleaseMilestoneIssueStats struct {
	Total  int `json:"total"`
	Closed int `json:"closed"`
}

// ReleaseEvidence represents a project release's evidence.
//
// GitLab API docs:
// https://docs.gitlab.com/api/releases/#list-releases
type ReleaseEvidence struct {
	SHA         string     `json:"sha"`
	Filepath    string     `json:"filepath"`
	CollectedAt *time.Time `json:"collected_at"`
}

// ListReleasesOptions represents ListReleases() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/releases/#list-releases
type ListReleasesOptions struct {
	ListOptions
	OrderBy                *string `url:"order_by,omitempty" json:"order_by,omitempty"`
	Sort                   *string `url:"sort,omitempty" json:"sort,omitempty"`
	IncludeHTMLDescription *bool   `url:"include_html_description,omitempty" json:"include_html_description,omitempty"`
}

// ListReleases gets a paginated list of releases accessible by the authenticated user.
//
// GitLab API docs:
// https://docs.gitlab.com/api/releases/#list-releases
func (s *ReleasesService) ListReleases(pid any, opt *ListReleasesOptions, options ...RequestOptionFunc) ([]*Release, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/releases", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var rs []*Release
	resp, err := s.client.Do(req, &rs)
	if err != nil {
		return nil, resp, err
	}

	return rs, resp, nil
}

// GetRelease returns a single release, identified by a tag name.
//
// GitLab API docs:
// https://docs.gitlab.com/api/releases/#get-a-release-by-a-tag-name
func (s *ReleasesService) GetRelease(pid any, tagName string, options ...RequestOptionFunc) (*Release, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/releases/%s", PathEscape(project), PathEscape(tagName))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	r := new(Release)
	resp, err := s.client.Do(req, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// GetLatestRelease returns the latest release for the project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/releases/#get-the-latest-release
func (s *ReleasesService) GetLatestRelease(pid any, options ...RequestOptionFunc) (*Release, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/releases/permalink/latest", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	r := new(Release)
	resp, err := s.client.Do(req, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, err
}

// CreateReleaseOptions represents CreateRelease() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/releases/#create-a-release
type CreateReleaseOptions struct {
	Name        *string               `url:"name,omitempty" json:"name,omitempty"`
	TagName     *string               `url:"tag_name,omitempty" json:"tag_name,omitempty"`
	TagMessage  *string               `url:"tag_message,omitempty" json:"tag_message,omitempty"`
	Description *string               `url:"description,omitempty" json:"description,omitempty"`
	Ref         *string               `url:"ref,omitempty" json:"ref,omitempty"`
	Milestones  *[]string             `url:"milestones,omitempty" json:"milestones,omitempty"`
	Assets      *ReleaseAssetsOptions `url:"assets,omitempty" json:"assets,omitempty"`
	ReleasedAt  *time.Time            `url:"released_at,omitempty" json:"released_at,omitempty"`
}

// ReleaseAssetsOptions represents release assets in CreateRelease() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/releases/#create-a-release
type ReleaseAssetsOptions struct {
	Links []*ReleaseAssetLinkOptions `url:"links,omitempty" json:"links,omitempty"`
}

// ReleaseAssetLinkOptions represents release asset link in CreateRelease()
// options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/releases/#create-a-release
type ReleaseAssetLinkOptions struct {
	Name            *string        `url:"name,omitempty" json:"name,omitempty"`
	URL             *string        `url:"url,omitempty" json:"url,omitempty"`
	FilePath        *string        `url:"filepath,omitempty" json:"filepath,omitempty"`
	DirectAssetPath *string        `url:"direct_asset_path,omitempty" json:"direct_asset_path,omitempty"`
	LinkType        *LinkTypeValue `url:"link_type,omitempty" json:"link_type,omitempty"`
}

// CreateRelease creates a release.
//
// GitLab API docs:
// https://docs.gitlab.com/api/releases/#create-a-release
func (s *ReleasesService) CreateRelease(pid any, opts *CreateReleaseOptions, options ...RequestOptionFunc) (*Release, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/releases", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPost, u, opts, options)
	if err != nil {
		return nil, nil, err
	}

	r := new(Release)
	resp, err := s.client.Do(req, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// UpdateReleaseOptions represents UpdateRelease() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/releases/#update-a-release
type UpdateReleaseOptions struct {
	Name        *string    `url:"name" json:"name"`
	Description *string    `url:"description" json:"description"`
	Milestones  *[]string  `url:"milestones,omitempty" json:"milestones,omitempty"`
	ReleasedAt  *time.Time `url:"released_at,omitempty" json:"released_at,omitempty"`
}

// UpdateRelease updates a release.
//
// GitLab API docs:
// https://docs.gitlab.com/api/releases/#update-a-release
func (s *ReleasesService) UpdateRelease(pid any, tagName string, opts *UpdateReleaseOptions, options ...RequestOptionFunc) (*Release, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/releases/%s", PathEscape(project), PathEscape(tagName))

	req, err := s.client.NewRequest(http.MethodPut, u, opts, options)
	if err != nil {
		return nil, nil, err
	}

	r := new(Release)
	resp, err := s.client.Do(req, &r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}

// DeleteRelease deletes a release.
//
// GitLab API docs:
// https://docs.gitlab.com/api/releases/#delete-a-release
func (s *ReleasesService) DeleteRelease(pid any, tagName string, options ...RequestOptionFunc) (*Release, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/releases/%s", PathEscape(project), PathEscape(tagName))

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	r := new(Release)
	resp, err := s.client.Do(req, r)
	if err != nil {
		return nil, resp, err
	}

	return r, resp, nil
}
