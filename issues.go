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
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"
)

type (
	IssuesServiceInterface interface {
		ListIssues(opt *ListIssuesOptions, options ...RequestOptionFunc) ([]*Issue, *Response, error)
		ListGroupIssues(pid any, opt *ListGroupIssuesOptions, options ...RequestOptionFunc) ([]*Issue, *Response, error)
		ListProjectIssues(pid any, opt *ListProjectIssuesOptions, options ...RequestOptionFunc) ([]*Issue, *Response, error)
		GetIssueByID(issue int, options ...RequestOptionFunc) (*Issue, *Response, error)
		GetIssue(pid any, issue int, options ...RequestOptionFunc) (*Issue, *Response, error)
		CreateIssue(pid any, opt *CreateIssueOptions, options ...RequestOptionFunc) (*Issue, *Response, error)
		UpdateIssue(pid any, issue int, opt *UpdateIssueOptions, options ...RequestOptionFunc) (*Issue, *Response, error)
		DeleteIssue(pid any, issue int, options ...RequestOptionFunc) (*Response, error)
		ReorderIssue(pid any, issue int, opt *ReorderIssueOptions, options ...RequestOptionFunc) (*Issue, *Response, error)
		MoveIssue(pid any, issue int, opt *MoveIssueOptions, options ...RequestOptionFunc) (*Issue, *Response, error)
		SubscribeToIssue(pid any, issue int, options ...RequestOptionFunc) (*Issue, *Response, error)
		UnsubscribeFromIssue(pid any, issue int, options ...RequestOptionFunc) (*Issue, *Response, error)
		CreateTodo(pid any, issue int, options ...RequestOptionFunc) (*Todo, *Response, error)
		ListMergeRequestsClosingIssue(pid any, issue int, opt *ListMergeRequestsClosingIssueOptions, options ...RequestOptionFunc) ([]*BasicMergeRequest, *Response, error)
		ListMergeRequestsRelatedToIssue(pid any, issue int, opt *ListMergeRequestsRelatedToIssueOptions, options ...RequestOptionFunc) ([]*BasicMergeRequest, *Response, error)
		SetTimeEstimate(pid any, issue int, opt *SetTimeEstimateOptions, options ...RequestOptionFunc) (*TimeStats, *Response, error)
		ResetTimeEstimate(pid any, issue int, options ...RequestOptionFunc) (*TimeStats, *Response, error)
		AddSpentTime(pid any, issue int, opt *AddSpentTimeOptions, options ...RequestOptionFunc) (*TimeStats, *Response, error)
		ResetSpentTime(pid any, issue int, options ...RequestOptionFunc) (*TimeStats, *Response, error)
		GetTimeSpent(pid any, issue int, options ...RequestOptionFunc) (*TimeStats, *Response, error)
		GetParticipants(pid any, issue int, options ...RequestOptionFunc) ([]*BasicUser, *Response, error)
	}

	// IssuesService handles communication with the issue related methods
	// of the GitLab API.
	//
	// GitLab API docs: https://docs.gitlab.com/api/issues/
	IssuesService struct {
		client    *Client
		timeStats *timeStatsService
	}
)

var _ IssuesServiceInterface = (*IssuesService)(nil)

// IssueAuthor represents a author of the issue.
type IssueAuthor struct {
	ID        int    `json:"id"`
	State     string `json:"state"`
	WebURL    string `json:"web_url"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
	Username  string `json:"username"`
}

// IssueAssignee represents a assignee of the issue.
type IssueAssignee struct {
	ID        int    `json:"id"`
	State     string `json:"state"`
	WebURL    string `json:"web_url"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
	Username  string `json:"username"`
}

// IssueReferences represents references of the issue.
type IssueReferences struct {
	Short    string `json:"short"`
	Relative string `json:"relative"`
	Full     string `json:"full"`
}

// IssueCloser represents a closer of the issue.
type IssueCloser struct {
	ID        int    `json:"id"`
	State     string `json:"state"`
	WebURL    string `json:"web_url"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
	Username  string `json:"username"`
}

// IssueLinks represents links of the issue.
type IssueLinks struct {
	Self       string `json:"self"`
	Notes      string `json:"notes"`
	AwardEmoji string `json:"award_emoji"`
	Project    string `json:"project"`
}

// Issue represents a GitLab issue.
//
// GitLab API docs: https://docs.gitlab.com/api/issues/
type Issue struct {
	ID                   int                    `json:"id"`
	IID                  int                    `json:"iid"`
	ExternalID           string                 `json:"external_id"`
	State                string                 `json:"state"`
	Description          string                 `json:"description"`
	HealthStatus         string                 `json:"health_status"`
	Author               *IssueAuthor           `json:"author"`
	Milestone            *Milestone             `json:"milestone"`
	ProjectID            int                    `json:"project_id"`
	Assignees            []*IssueAssignee       `json:"assignees"`
	UpdatedAt            *time.Time             `json:"updated_at"`
	ClosedAt             *time.Time             `json:"closed_at"`
	ClosedBy             *IssueCloser           `json:"closed_by"`
	Title                string                 `json:"title"`
	CreatedAt            *time.Time             `json:"created_at"`
	MovedToID            int                    `json:"moved_to_id"`
	Labels               Labels                 `json:"labels"`
	LabelDetails         []*LabelDetails        `json:"label_details"`
	Upvotes              int                    `json:"upvotes"`
	Downvotes            int                    `json:"downvotes"`
	DueDate              *ISOTime               `json:"due_date"`
	WebURL               string                 `json:"web_url"`
	References           *IssueReferences       `json:"references"`
	TimeStats            *TimeStats             `json:"time_stats"`
	Confidential         bool                   `json:"confidential"`
	Weight               int                    `json:"weight"`
	DiscussionLocked     bool                   `json:"discussion_locked"`
	IssueType            *string                `json:"issue_type,omitempty"`
	Subscribed           bool                   `json:"subscribed"`
	UserNotesCount       int                    `json:"user_notes_count"`
	Links                *IssueLinks            `json:"_links"`
	IssueLinkID          int                    `json:"issue_link_id"`
	MergeRequestCount    int                    `json:"merge_requests_count"`
	EpicIssueID          int                    `json:"epic_issue_id"`
	Epic                 *Epic                  `json:"epic"`
	Iteration            *GroupIteration        `json:"iteration"`
	TaskCompletionStatus *TasksCompletionStatus `json:"task_completion_status"`
	ServiceDeskReplyTo   string                 `json:"service_desk_reply_to"`

	// Deprecated: use Assignees instead
	Assignee *IssueAssignee `json:"assignee"`
}

func (i Issue) String() string {
	return Stringify(i)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (i *Issue) UnmarshalJSON(data []byte) error {
	type alias Issue

	raw := make(map[string]any)
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	if reflect.TypeOf(raw["id"]).Kind() == reflect.String {
		raw["external_id"] = raw["id"]
		delete(raw, "id")
	}

	labelDetails, ok := raw["labels"].([]any)
	if ok && len(labelDetails) > 0 {
		// We only want to change anything if we got label details.
		if _, ok := labelDetails[0].(map[string]any); ok {
			labels := make([]any, len(labelDetails))
			for i, details := range labelDetails {
				labels[i] = details.(map[string]any)["name"]
			}

			// Set the correct values
			raw["labels"] = labels
			raw["label_details"] = labelDetails
		}
	}

	data, err = json.Marshal(raw)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, (*alias)(i))
}

// LabelDetails represents detailed label information.
type LabelDetails struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Color           string `json:"color"`
	Description     string `json:"description"`
	DescriptionHTML string `json:"description_html"`
	TextColor       string `json:"text_color"`
}

// ListIssuesOptions represents the available ListIssues() options.
//
// GitLab API docs: https://docs.gitlab.com/api/issues/#list-issues
type ListIssuesOptions struct {
	ListOptions
	State               *string          `url:"state,omitempty" json:"state,omitempty"`
	Labels              *LabelOptions    `url:"labels,comma,omitempty" json:"labels,omitempty"`
	NotLabels           *LabelOptions    `url:"not[labels],comma,omitempty" json:"not[labels],omitempty"`
	WithLabelDetails    *bool            `url:"with_labels_details,omitempty" json:"with_labels_details,omitempty"`
	Milestone           *string          `url:"milestone,omitempty" json:"milestone,omitempty"`
	NotMilestone        *string          `url:"not[milestone],omitempty" json:"not[milestone],omitempty"`
	Scope               *string          `url:"scope,omitempty" json:"scope,omitempty"`
	AuthorID            *int             `url:"author_id,omitempty" json:"author_id,omitempty"`
	AuthorUsername      *string          `url:"author_username,omitempty" json:"author_username,omitempty"`
	NotAuthorUsername   *string          `url:"not[author_username],omitempty" json:"not[author_username],omitempty"`
	NotAuthorID         *[]int           `url:"not[author_id],omitempty" json:"not[author_id],omitempty"`
	AssigneeID          *AssigneeIDValue `url:"assignee_id,omitempty" json:"assignee_id,omitempty"`
	NotAssigneeID       *[]int           `url:"not[assignee_id],omitempty" json:"not[assignee_id],omitempty"`
	AssigneeUsername    *string          `url:"assignee_username,omitempty" json:"assignee_username,omitempty"`
	NotAssigneeUsername *string          `url:"not[assignee_username],omitempty" json:"not[assignee_username],omitempty"`
	MyReactionEmoji     *string          `url:"my_reaction_emoji,omitempty" json:"my_reaction_emoji,omitempty"`
	NotMyReactionEmoji  *[]string        `url:"not[my_reaction_emoji],omitempty" json:"not[my_reaction_emoji],omitempty"`
	IIDs                *[]int           `url:"iids[],omitempty" json:"iids,omitempty"`
	In                  *string          `url:"in,omitempty" json:"in,omitempty"`
	NotIn               *string          `url:"not[in],omitempty" json:"not[in],omitempty"`
	OrderBy             *string          `url:"order_by,omitempty" json:"order_by,omitempty"`
	Sort                *string          `url:"sort,omitempty" json:"sort,omitempty"`
	Search              *string          `url:"search,omitempty" json:"search,omitempty"`
	NotSearch           *string          `url:"not[search],omitempty" json:"not[search],omitempty"`
	CreatedAfter        *time.Time       `url:"created_after,omitempty" json:"created_after,omitempty"`
	CreatedBefore       *time.Time       `url:"created_before,omitempty" json:"created_before,omitempty"`
	DueDate             *string          `url:"due_date,omitempty" json:"due_date,omitempty"`
	UpdatedAfter        *time.Time       `url:"updated_after,omitempty" json:"updated_after,omitempty"`
	UpdatedBefore       *time.Time       `url:"updated_before,omitempty" json:"updated_before,omitempty"`
	Confidential        *bool            `url:"confidential,omitempty" json:"confidential,omitempty"`
	IssueType           *string          `url:"issue_type,omitempty" json:"issue_type,omitempty"`
	IterationID         *int             `url:"iteration_id,omitempty" json:"iteration_id,omitempty"`
}

// ListIssues gets all issues created by authenticated user. This function
// takes pagination parameters page and per_page to restrict the list of issues.
//
// GitLab API docs: https://docs.gitlab.com/api/issues/#list-issues
func (s *IssuesService) ListIssues(opt *ListIssuesOptions, options ...RequestOptionFunc) ([]*Issue, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "issues", opt, options)
	if err != nil {
		return nil, nil, err
	}

	var i []*Issue
	resp, err := s.client.Do(req, &i)
	if err != nil {
		return nil, resp, err
	}

	return i, resp, nil
}

// ListGroupIssuesOptions represents the available ListGroupIssues() options.
//
// GitLab API docs: https://docs.gitlab.com/api/issues/#list-group-issues
type ListGroupIssuesOptions struct {
	ListOptions
	State             *string       `url:"state,omitempty" json:"state,omitempty"`
	Labels            *LabelOptions `url:"labels,comma,omitempty" json:"labels,omitempty"`
	NotLabels         *LabelOptions `url:"not[labels],comma,omitempty" json:"not[labels],omitempty"`
	WithLabelDetails  *bool         `url:"with_labels_details,omitempty" json:"with_labels_details,omitempty"`
	IIDs              *[]int        `url:"iids[],omitempty" json:"iids,omitempty"`
	Milestone         *string       `url:"milestone,omitempty" json:"milestone,omitempty"`
	NotMilestone      *string       `url:"not[milestone],omitempty" json:"not[milestone],omitempty"`
	Scope             *string       `url:"scope,omitempty" json:"scope,omitempty"`
	AuthorID          *int          `url:"author_id,omitempty" json:"author_id,omitempty"`
	NotAuthorID       *int          `url:"not[author_id],omitempty" json:"not[author_id],omitempty"`
	AuthorUsername    *string       `url:"author_username,omitempty" json:"author_username,omitempty"`
	NotAuthorUsername *string       `url:"not[author_username],omitempty" json:"not[author_username],omitempty"`

	// AssigneeID is defined as an int in the documentation, however, the field
	// must be able to accept Assignee IDs and the words 'None' and 'Any'.  Use
	// *AssigneeIDValue instead of *int.
	AssigneeID          *AssigneeIDValue `url:"assignee_id,omitempty" json:"assignee_id,omitempty"`
	NotAssigneeID       *int             `url:"not[assignee_id],omitempty" json:"not[assignee_id],omitempty"`
	AssigneeUsername    *string          `url:"assignee_username,omitempty" json:"assignee_username,omitempty"`
	NotAssigneeUsername *string          `url:"not[assignee_username],omitempty" json:"not[assignee_username],omitempty"`
	MyReactionEmoji     *string          `url:"my_reaction_emoji,omitempty" json:"my_reaction_emoji,omitempty"`
	NotMyReactionEmoji  *string          `url:"not[my_reaction_emoji],omitempty" json:"not[my_reaction_emoji],omitempty"`
	OrderBy             *string          `url:"order_by,omitempty" json:"order_by,omitempty"`
	Sort                *string          `url:"sort,omitempty" json:"sort,omitempty"`
	Search              *string          `url:"search,omitempty" json:"search,omitempty"`
	NotSearch           *string          `url:"not[search],omitempty" json:"not[search],omitempty"`
	In                  *string          `url:"in,omitempty" json:"in,omitempty"`
	NotIn               *string          `url:"not[in],omitempty" json:"not[in],omitempty"`
	CreatedAfter        *time.Time       `url:"created_after,omitempty" json:"created_after,omitempty"`
	CreatedBefore       *time.Time       `url:"created_before,omitempty" json:"created_before,omitempty"`
	DueDate             *string          `url:"due_date,omitempty" json:"due_date,omitempty"`
	UpdatedAfter        *time.Time       `url:"updated_after,omitempty" json:"updated_after,omitempty"`
	UpdatedBefore       *time.Time       `url:"updated_before,omitempty" json:"updated_before,omitempty"`
	Confidential        *bool            `url:"confidential,omitempty" json:"confidential,omitempty"`
	IssueType           *string          `url:"issue_type,omitempty" json:"issue_type,omitempty"`
	IterationID         *int             `url:"iteration_id,omitempty" json:"iteration_id,omitempty"`
}

// ListGroupIssues gets a list of group issues. This function accepts
// pagination parameters page and per_page to return the list of group issues.
//
// GitLab API docs: https://docs.gitlab.com/api/issues/#list-group-issues
func (s *IssuesService) ListGroupIssues(pid any, opt *ListGroupIssuesOptions, options ...RequestOptionFunc) ([]*Issue, *Response, error) {
	group, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("groups/%s/issues", PathEscape(group))

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var i []*Issue
	resp, err := s.client.Do(req, &i)
	if err != nil {
		return nil, resp, err
	}

	return i, resp, nil
}

// ListProjectIssuesOptions represents the available ListProjectIssues() options.
//
// GitLab API docs: https://docs.gitlab.com/api/issues/#list-project-issues
type ListProjectIssuesOptions struct {
	ListOptions
	IIDs                *[]int        `url:"iids[],omitempty" json:"iids,omitempty"`
	State               *string       `url:"state,omitempty" json:"state,omitempty"`
	Labels              *LabelOptions `url:"labels,comma,omitempty" json:"labels,omitempty"`
	NotLabels           *LabelOptions `url:"not[labels],comma,omitempty" json:"not[labels],omitempty"`
	WithLabelDetails    *bool         `url:"with_labels_details,omitempty" json:"with_labels_details,omitempty"`
	Milestone           *string       `url:"milestone,omitempty" json:"milestone,omitempty"`
	NotMilestone        *string       `url:"not[milestone],omitempty" json:"not[milestone],omitempty"`
	Scope               *string       `url:"scope,omitempty" json:"scope,omitempty"`
	AuthorID            *int          `url:"author_id,omitempty" json:"author_id,omitempty"`
	AuthorUsername      *string       `url:"author_username,omitempty" json:"author_username,omitempty"`
	NotAuthorUsername   *string       `url:"not[author_username],omitempty" json:"not[author_username],omitempty"`
	NotAuthorID         *int          `url:"not[author_id],omitempty" json:"not[author_id],omitempty"`
	AssigneeID          *int          `url:"assignee_id,omitempty" json:"assignee_id,omitempty"`
	NotAssigneeID       *int          `url:"not[assignee_id],omitempty" json:"not[assignee_id],omitempty"`
	AssigneeUsername    *string       `url:"assignee_username,omitempty" json:"assignee_username,omitempty"`
	NotAssigneeUsername *string       `url:"not[assignee_username],omitempty" json:"not[assignee_username],omitempty"`
	MyReactionEmoji     *string       `url:"my_reaction_emoji,omitempty" json:"my_reaction_emoji,omitempty"`
	NotMyReactionEmoji  *string       `url:"not[my_reaction_emoji],omitempty" json:"not[my_reaction_emoji],omitempty"`
	OrderBy             *string       `url:"order_by,omitempty" json:"order_by,omitempty"`
	Sort                *string       `url:"sort,omitempty" json:"sort,omitempty"`
	Search              *string       `url:"search,omitempty" json:"search,omitempty"`
	In                  *string       `url:"in,omitempty" json:"in,omitempty"`
	NotIn               *string       `url:"not[in],omitempty" json:"not[in],omitempty"`
	CreatedAfter        *time.Time    `url:"created_after,omitempty" json:"created_after,omitempty"`
	CreatedBefore       *time.Time    `url:"created_before,omitempty" json:"created_before,omitempty"`
	DueDate             *string       `url:"due_date,omitempty" json:"due_date,omitempty"`
	UpdatedAfter        *time.Time    `url:"updated_after,omitempty" json:"updated_after,omitempty"`
	UpdatedBefore       *time.Time    `url:"updated_before,omitempty" json:"updated_before,omitempty"`
	Confidential        *bool         `url:"confidential,omitempty" json:"confidential,omitempty"`
	IssueType           *string       `url:"issue_type,omitempty" json:"issue_type,omitempty"`
	IterationID         *int          `url:"iteration_id,omitempty" json:"iteration_id,omitempty"`
}

// ListProjectIssues gets a list of project issues. This function accepts
// pagination parameters page and per_page to return the list of project issues.
//
// GitLab API docs: https://docs.gitlab.com/api/issues/#list-project-issues
func (s *IssuesService) ListProjectIssues(pid any, opt *ListProjectIssuesOptions, options ...RequestOptionFunc) ([]*Issue, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/issues", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var i []*Issue
	resp, err := s.client.Do(req, &i)
	if err != nil {
		return nil, resp, err
	}

	return i, resp, nil
}

// GetIssueByID gets a single issue.
//
// GitLab API docs: https://docs.gitlab.com/api/issues/#single-issue
func (s *IssuesService) GetIssueByID(issue int, options ...RequestOptionFunc) (*Issue, *Response, error) {
	u := fmt.Sprintf("issues/%d", issue)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	i := new(Issue)
	resp, err := s.client.Do(req, i)
	if err != nil {
		return nil, resp, err
	}

	return i, resp, nil
}

// GetIssue gets a single project issue.
//
// GitLab API docs: https://docs.gitlab.com/api/issues/#single-project-issue
func (s *IssuesService) GetIssue(pid any, issue int, options ...RequestOptionFunc) (*Issue, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/issues/%d", PathEscape(project), issue)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	i := new(Issue)
	resp, err := s.client.Do(req, i)
	if err != nil {
		return nil, resp, err
	}

	return i, resp, nil
}

// CreateIssueOptions represents the available CreateIssue() options.
//
// GitLab API docs: https://docs.gitlab.com/api/issues/#new-issue
type CreateIssueOptions struct {
	IID                                *int          `url:"iid,omitempty" json:"iid,omitempty"`
	Title                              *string       `url:"title,omitempty" json:"title,omitempty"`
	Description                        *string       `url:"description,omitempty" json:"description,omitempty"`
	Confidential                       *bool         `url:"confidential,omitempty" json:"confidential,omitempty"`
	AssigneeIDs                        *[]int        `url:"assignee_ids,omitempty" json:"assignee_ids,omitempty"`
	MilestoneID                        *int          `url:"milestone_id,omitempty" json:"milestone_id,omitempty"`
	Labels                             *LabelOptions `url:"labels,comma,omitempty" json:"labels,omitempty"`
	CreatedAt                          *time.Time    `url:"created_at,omitempty" json:"created_at,omitempty"`
	DueDate                            *ISOTime      `url:"due_date,omitempty" json:"due_date,omitempty"`
	EpicID                             *int          `url:"epic_id,omitempty" json:"epic_id,omitempty"`
	MergeRequestToResolveDiscussionsOf *int          `url:"merge_request_to_resolve_discussions_of,omitempty" json:"merge_request_to_resolve_discussions_of,omitempty"`
	DiscussionToResolve                *string       `url:"discussion_to_resolve,omitempty" json:"discussion_to_resolve,omitempty"`
	Weight                             *int          `url:"weight,omitempty" json:"weight,omitempty"`
	IssueType                          *string       `url:"issue_type,omitempty" json:"issue_type,omitempty"`
}

// CreateIssue creates a new project issue.
//
// GitLab API docs: https://docs.gitlab.com/api/issues/#new-issue
func (s *IssuesService) CreateIssue(pid any, opt *CreateIssueOptions, options ...RequestOptionFunc) (*Issue, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/issues", PathEscape(project))

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	i := new(Issue)
	resp, err := s.client.Do(req, i)
	if err != nil {
		return nil, resp, err
	}

	return i, resp, nil
}

// UpdateIssueOptions represents the available UpdateIssue() options.
//
// To reset the due date, epic, milestone, or weight of the issue, set the
// ResetDueDate, ResetEpic, ResetMilestone, or ResetWeight field to true.
//
// GitLab API docs: https://docs.gitlab.com/api/issues/#edit-an-issue
type UpdateIssueOptions struct {
	Title            *string       `url:"title,omitempty" json:"title,omitempty"`
	Description      *string       `url:"description,omitempty" json:"description,omitempty"`
	Confidential     *bool         `url:"confidential,omitempty" json:"confidential,omitempty"`
	AssigneeIDs      *[]int        `url:"assignee_ids,omitempty" json:"assignee_ids,omitempty"`
	MilestoneID      *int          `url:"milestone_id,omitempty" json:"milestone_id,omitempty"`
	Labels           *LabelOptions `url:"labels,comma,omitempty" json:"labels,omitempty"`
	AddLabels        *LabelOptions `url:"add_labels,comma,omitempty" json:"add_labels,omitempty"`
	RemoveLabels     *LabelOptions `url:"remove_labels,comma,omitempty" json:"remove_labels,omitempty"`
	StateEvent       *string       `url:"state_event,omitempty" json:"state_event,omitempty"`
	UpdatedAt        *time.Time    `url:"updated_at,omitempty" json:"updated_at,omitempty"`
	DueDate          *ISOTime      `url:"due_date,omitempty" json:"due_date,omitempty"`
	EpicID           *int          `url:"epic_id,omitempty" json:"epic_id,omitempty"`
	Weight           *int          `url:"weight,omitempty" json:"weight,omitempty"`
	DiscussionLocked *bool         `url:"discussion_locked,omitempty" json:"discussion_locked,omitempty"`
	IssueType        *string       `url:"issue_type,omitempty" json:"issue_type,omitempty"`

	ResetDueDate     bool `url:"-" json:"-"`
	ResetEpicID      bool `url:"-" json:"-"`
	ResetMilestoneID bool `url:"-" json:"-"`
	ResetWeight      bool `url:"-" json:"-"`
}

// MarshalJSON implements custom JSON marshaling for UpdateIssueOptions.
// This is needed to support emitting a literal `null` when the field needs to be removed.
func (o UpdateIssueOptions) MarshalJSON() ([]byte, error) {
	data := map[string]any{}

	// Use reflection to copy all fields from o to data
	val := reflect.ValueOf(o)
	typ := val.Type()

	for i := range val.NumField() {
		field := val.Field(i)
		fieldName := typ.Field(i).Name

		if field.IsZero() {
			continue
		}

		name := fieldName

		if tag := typ.Field(i).Tag.Get("json"); tag != "" {
			tagFields := strings.Split(tag, ",")
			name = tagFields[0]
		}

		// Skip unexported fields.
		if name == "-" {
			continue
		}

		data[name] = field.Interface()
	}

	// Emit a literal `null` when the field needs to be removed
	if o.ResetDueDate {
		data["due_date"] = nil
	}

	if o.ResetEpicID {
		data["epic_id"] = nil
	}

	if o.ResetMilestoneID {
		data["milestone_id"] = nil
	}

	if o.ResetWeight {
		data["weight"] = nil
	}

	return json.Marshal(data)
}

// UpdateIssue updates an existing project issue. This function is also used
// to mark an issue as closed.
//
// GitLab API docs: https://docs.gitlab.com/api/issues/#edit-an-issue
func (s *IssuesService) UpdateIssue(pid any, issue int, opt *UpdateIssueOptions, options ...RequestOptionFunc) (*Issue, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/issues/%d", PathEscape(project), issue)

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	i := new(Issue)
	resp, err := s.client.Do(req, i)
	if err != nil {
		return nil, resp, err
	}

	return i, resp, nil
}

// DeleteIssue deletes a single project issue.
//
// GitLab API docs: https://docs.gitlab.com/api/issues/#delete-an-issue
func (s *IssuesService) DeleteIssue(pid any, issue int, options ...RequestOptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/issues/%d", PathEscape(project), issue)

	req, err := s.client.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ReorderIssueOptions represents the available ReorderIssue() options.
//
// GitLab API docs: https://docs.gitlab.com/api/issues/#reorder-an-issue
type ReorderIssueOptions struct {
	MoveAfterID  *int `url:"move_after_id,omitempty" json:"move_after_id,omitempty"`
	MoveBeforeID *int `url:"move_before_id,omitempty" json:"move_before_id,omitempty"`
}

// ReorderIssue reorders an issue.
//
// GitLab API docs: https://docs.gitlab.com/api/issues/#reorder-an-issue
func (s *IssuesService) ReorderIssue(pid any, issue int, opt *ReorderIssueOptions, options ...RequestOptionFunc) (*Issue, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/issues/%d/reorder", PathEscape(project), issue)

	req, err := s.client.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	i := new(Issue)
	resp, err := s.client.Do(req, i)
	if err != nil {
		return nil, resp, err
	}

	return i, resp, nil
}

// MoveIssueOptions represents the available MoveIssue() options.
//
// GitLab API docs: https://docs.gitlab.com/api/issues/#move-an-issue
type MoveIssueOptions struct {
	ToProjectID *int `url:"to_project_id,omitempty" json:"to_project_id,omitempty"`
}

// MoveIssue updates an existing project issue. This function is also used
// to mark an issue as closed.
//
// GitLab API docs: https://docs.gitlab.com/api/issues/#move-an-issue
func (s *IssuesService) MoveIssue(pid any, issue int, opt *MoveIssueOptions, options ...RequestOptionFunc) (*Issue, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/issues/%d/move", PathEscape(project), issue)

	req, err := s.client.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	i := new(Issue)
	resp, err := s.client.Do(req, i)
	if err != nil {
		return nil, resp, err
	}

	return i, resp, nil
}

// SubscribeToIssue subscribes the authenticated user to the given issue to
// receive notifications. If the user is already subscribed to the issue, the
// status code 304 is returned.
//
// GitLab API docs:
// https://docs.gitlab.com/api/issues/#subscribe-to-an-issue
func (s *IssuesService) SubscribeToIssue(pid any, issue int, options ...RequestOptionFunc) (*Issue, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/issues/%d/subscribe", PathEscape(project), issue)

	req, err := s.client.NewRequest(http.MethodPost, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	i := new(Issue)
	resp, err := s.client.Do(req, i)
	if err != nil {
		return nil, resp, err
	}

	return i, resp, nil
}

// UnsubscribeFromIssue unsubscribes the authenticated user from the given
// issue to not receive notifications from that merge request. If the user
// is not subscribed to the issue, status code 304 is returned.
//
// GitLab API docs:
// https://docs.gitlab.com/api/issues/#unsubscribe-from-an-issue
func (s *IssuesService) UnsubscribeFromIssue(pid any, issue int, options ...RequestOptionFunc) (*Issue, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/issues/%d/unsubscribe", PathEscape(project), issue)

	req, err := s.client.NewRequest(http.MethodPost, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	i := new(Issue)
	resp, err := s.client.Do(req, i)
	if err != nil {
		return nil, resp, err
	}

	return i, resp, nil
}

// CreateTodo creates a todo for the current user for an issue.
// If there already exists a todo for the user on that issue, status code
// 304 is returned.
//
// GitLab API docs:
// https://docs.gitlab.com/api/issues/#create-a-to-do-item
func (s *IssuesService) CreateTodo(pid any, issue int, options ...RequestOptionFunc) (*Todo, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/issues/%d/todo", PathEscape(project), issue)

	req, err := s.client.NewRequest(http.MethodPost, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	t := new(Todo)
	resp, err := s.client.Do(req, t)
	if err != nil {
		return nil, resp, err
	}

	return t, resp, nil
}

// ListMergeRequestsClosingIssueOptions represents the available
// ListMergeRequestsClosingIssue() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/issues/#list-merge-requests-that-close-a-particular-issue-on-merge
type ListMergeRequestsClosingIssueOptions ListOptions

// ListMergeRequestsClosingIssue gets all the merge requests that will close
// issue when merged.
//
// GitLab API docs:
// https://docs.gitlab.com/api/issues/#list-merge-requests-that-close-a-particular-issue-on-merge
func (s *IssuesService) ListMergeRequestsClosingIssue(pid any, issue int, opt *ListMergeRequestsClosingIssueOptions, options ...RequestOptionFunc) ([]*BasicMergeRequest, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/issues/%d/closed_by", PathEscape(project), issue)

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var m []*BasicMergeRequest
	resp, err := s.client.Do(req, &m)
	if err != nil {
		return nil, resp, err
	}

	return m, resp, nil
}

// ListMergeRequestsRelatedToIssueOptions represents the available
// ListMergeRequestsRelatedToIssue() options.
//
// GitLab API docs:
// https://docs.gitlab.com/api/issues/#list-merge-requests-related-to-issue
type ListMergeRequestsRelatedToIssueOptions ListOptions

// ListMergeRequestsRelatedToIssue gets all the merge requests that are
// related to the issue
//
// GitLab API docs:
// https://docs.gitlab.com/api/issues/#list-merge-requests-related-to-issue
func (s *IssuesService) ListMergeRequestsRelatedToIssue(pid any, issue int, opt *ListMergeRequestsRelatedToIssueOptions, options ...RequestOptionFunc) ([]*BasicMergeRequest, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/issues/%d/related_merge_requests",
		PathEscape(project),
		issue,
	)

	req, err := s.client.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var m []*BasicMergeRequest
	resp, err := s.client.Do(req, &m)
	if err != nil {
		return nil, resp, err
	}

	return m, resp, nil
}

// SetTimeEstimate sets the time estimate for a single project issue.
//
// GitLab API docs:
// https://docs.gitlab.com/api/issues/#set-a-time-estimate-for-an-issue
func (s *IssuesService) SetTimeEstimate(pid any, issue int, opt *SetTimeEstimateOptions, options ...RequestOptionFunc) (*TimeStats, *Response, error) {
	return s.timeStats.setTimeEstimate(pid, "issues", issue, opt, options...)
}

// ResetTimeEstimate resets the time estimate for a single project issue.
//
// GitLab API docs:
// https://docs.gitlab.com/api/issues/#reset-the-time-estimate-for-an-issue
func (s *IssuesService) ResetTimeEstimate(pid any, issue int, options ...RequestOptionFunc) (*TimeStats, *Response, error) {
	return s.timeStats.resetTimeEstimate(pid, "issues", issue, options...)
}

// AddSpentTime adds spent time for a single project issue.
//
// GitLab API docs:
// https://docs.gitlab.com/api/issues/#add-spent-time-for-an-issue
func (s *IssuesService) AddSpentTime(pid any, issue int, opt *AddSpentTimeOptions, options ...RequestOptionFunc) (*TimeStats, *Response, error) {
	return s.timeStats.addSpentTime(pid, "issues", issue, opt, options...)
}

// ResetSpentTime resets the spent time for a single project issue.
//
// GitLab API docs:
// https://docs.gitlab.com/api/issues/#reset-spent-time-for-an-issue
func (s *IssuesService) ResetSpentTime(pid any, issue int, options ...RequestOptionFunc) (*TimeStats, *Response, error) {
	return s.timeStats.resetSpentTime(pid, "issues", issue, options...)
}

// GetTimeSpent gets the spent time for a single project issue.
//
// GitLab API docs:
// https://docs.gitlab.com/api/issues/#get-time-tracking-stats
func (s *IssuesService) GetTimeSpent(pid any, issue int, options ...RequestOptionFunc) (*TimeStats, *Response, error) {
	return s.timeStats.getTimeSpent(pid, "issues", issue, options...)
}

// GetParticipants gets a list of issue participants.
//
// GitLab API docs:
// https://docs.gitlab.com/api/issues/#list-participants-in-an-issue
func (s *IssuesService) GetParticipants(pid any, issue int, options ...RequestOptionFunc) ([]*BasicUser, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/issues/%d/participants", PathEscape(project), issue)

	req, err := s.client.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	var bu []*BasicUser
	resp, err := s.client.Do(req, &bu)
	if err != nil {
		return nil, resp, err
	}

	return bu, resp, nil
}
