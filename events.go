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
	// EventsServiceInterface defines all the API methods for the EventsService
	EventsServiceInterface interface {
		ListCurrentUserContributionEvents(opt *ListContributionEventsOptions, options ...RequestOptionFunc) ([]*ContributionEvent, *Response, error)
		ListProjectVisibleEvents(pid any, opt *ListProjectVisibleEventsOptions, options ...RequestOptionFunc) ([]*ProjectEvent, *Response, error)
	}

	// EventsService handles communication with the event related methods of
	// the GitLab API.
	//
	// GitLab API docs: https://docs.gitlab.com/api/events/
	EventsService struct {
		client *Client
	}
)

var _ EventsServiceInterface = (*EventsService)(nil)

// ContributionEvent represents a user's contribution
//
// GitLab API docs:
// https://docs.gitlab.com/api/events/#get-user-contribution-events
type ContributionEvent struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	ProjectID   int        `json:"project_id"`
	ActionName  string     `json:"action_name"`
	TargetID    int        `json:"target_id"`
	TargetIID   int        `json:"target_iid"`
	TargetType  string     `json:"target_type"`
	AuthorID    int        `json:"author_id"`
	TargetTitle string     `json:"target_title"`
	CreatedAt   *time.Time `json:"created_at"`
	PushData    struct {
		CommitCount int    `json:"commit_count"`
		Action      string `json:"action"`
		RefType     string `json:"ref_type"`
		CommitFrom  string `json:"commit_from"`
		CommitTo    string `json:"commit_to"`
		Ref         string `json:"ref"`
		CommitTitle string `json:"commit_title"`
	} `json:"push_data"`
	Note   *Note `json:"note"`
	Author struct {
		Name      string `json:"name"`
		Username  string `json:"username"`
		ID        int    `json:"id"`
		State     string `json:"state"`
		AvatarURL string `json:"avatar_url"`
		WebURL    string `json:"web_url"`
	} `json:"author"`
	AuthorUsername string `json:"author_username"`
}

// ListContributionEventsOptions represents the options for GetUserContributionEvents
//
// GitLab API docs:
// https://docs.gitlab.com/api/events/#get-user-contribution-events
type ListContributionEventsOptions struct {
	ListOptions
	Action     *EventTypeValue       `url:"action,omitempty" json:"action,omitempty"`
	TargetType *EventTargetTypeValue `url:"target_type,omitempty" json:"target_type,omitempty"`
	Before     *ISOTime              `url:"before,omitempty" json:"before,omitempty"`
	After      *ISOTime              `url:"after,omitempty" json:"after,omitempty"`
	Sort       *string               `url:"sort,omitempty" json:"sort,omitempty"`
}

// ListUserContributionEvents retrieves user contribution events
// for the specified user, sorted from newest to oldest.
//
// GitLab API docs:
// https://docs.gitlab.com/api/events/#get-user-contribution-events
func (s *UsersService) ListUserContributionEvents(uid any, opt *ListContributionEventsOptions, options ...RequestOptionFunc) ([]*ContributionEvent, *Response, error) {
	user, err := parseID(uid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("users/%s/events", user)

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var cs []*ContributionEvent
	resp, err := s.client.Do(req, &cs)
	if err != nil {
		return nil, resp, err
	}

	return cs, resp, nil
}

// ListCurrentUserContributionEvents gets a list currently authenticated user's events
//
// GitLab API docs: https://docs.gitlab.com/api/events/#list-currently-authenticated-users-events
func (s *EventsService) ListCurrentUserContributionEvents(opt *ListContributionEventsOptions, options ...RequestOptionFunc) ([]*ContributionEvent, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "events", opt, options)
	if err != nil {
		return nil, nil, err
	}

	var cs []*ContributionEvent
	resp, err := s.client.Do(req, &cs)
	if err != nil {
		return nil, resp, err
	}

	return cs, resp, nil
}

// ProjectEvent represents a GitLab project event.
//
// GitLab API docs:
// https://docs.gitlab.com/api/events/#list-a-projects-visible-events
type ProjectEvent struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	ProjectID   int    `json:"project_id"`
	ActionName  string `json:"action_name"`
	TargetID    int    `json:"target_id"`
	TargetIID   int    `json:"target_iid"`
	TargetType  string `json:"target_type"`
	AuthorID    int    `json:"author_id"`
	TargetTitle string `json:"target_title"`
	CreatedAt   string `json:"created_at"`
	Author      struct {
		Name      string `json:"name"`
		Username  string `json:"username"`
		ID        int    `json:"id"`
		State     string `json:"state"`
		AvatarURL string `json:"avatar_url"`
		WebURL    string `json:"web_url"`
	} `json:"author"`
	AuthorUsername string `json:"author_username"`
	Data           struct {
		Before            string      `json:"before"`
		After             string      `json:"after"`
		Ref               string      `json:"ref"`
		UserID            int         `json:"user_id"`
		UserName          string      `json:"user_name"`
		Repository        *Repository `json:"repository"`
		Commits           []*Commit   `json:"commits"`
		TotalCommitsCount int         `json:"total_commits_count"`
	} `json:"data"`
	Note struct {
		ID         int    `json:"id"`
		Body       string `json:"body"`
		Attachment string `json:"attachment"`
		Author     struct {
			ID        int    `json:"id"`
			Username  string `json:"username"`
			Email     string `json:"email"`
			Name      string `json:"name"`
			State     string `json:"state"`
			AvatarURL string `json:"avatar_url"`
			WebURL    string `json:"web_url"`
		} `json:"author"`
		CreatedAt    *time.Time `json:"created_at"`
		System       bool       `json:"system"`
		NoteableID   int        `json:"noteable_id"`
		NoteableType string     `json:"noteable_type"`
		NoteableIID  int        `json:"noteable_iid"`
	} `json:"note"`
	PushData struct {
		CommitCount int    `json:"commit_count"`
		Action      string `json:"action"`
		RefType     string `json:"ref_type"`
		CommitFrom  string `json:"commit_from"`
		CommitTo    string `json:"commit_to"`
		Ref         string `json:"ref"`
		CommitTitle string `json:"commit_title"`
	} `json:"push_data"`
}

func (s ProjectEvent) String() string {
	return Stringify(s)
}

// ListProjectVisibleEventsOptions represents the available
// ListProjectVisibleEvents() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/events/#list-a-projects-visible-events
type ListProjectVisibleEventsOptions struct {
	ListOptions
	Action     *EventTypeValue       `url:"action,omitempty" json:"action,omitempty"`
	TargetType *EventTargetTypeValue `url:"target_type,omitempty" json:"target_type,omitempty"`
	Before     *ISOTime              `url:"before,omitempty" json:"before,omitempty"`
	After      *ISOTime              `url:"after,omitempty" json:"after,omitempty"`
	Sort       *string               `url:"sort,omitempty" json:"sort,omitempty"`
}

// ListProjectVisibleEvents gets the events for the specified project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/events/#list-a-projects-visible-events
func (s *EventsService) ListProjectVisibleEvents(pid any, opt *ListProjectVisibleEventsOptions, options ...RequestOptionFunc) ([]*ProjectEvent, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/events", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var p []*ProjectEvent
	resp, err := s.client.Do(req, &p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}
