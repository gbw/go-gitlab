package gitlab

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"time"
)

type (
	WorkItemsServiceInterface interface {
		GetWorkItemByID(gid any, options ...RequestOptionFunc) (*WorkItem, *Response, error)
		GetWorkItem(fullPath string, iid int64, options ...RequestOptionFunc) (*WorkItem, *Response, error)
		ListWorkItems(fullPath string, opt *ListWorkItemsOptions, options ...RequestOptionFunc) ([]*WorkItem, *Response, error)
	}

	// WorkItemsService handles communication with the work item related methods
	// of the GitLab API.
	//
	// GitLab API docs: https://docs.gitlab.com/api/graphql/reference/#workitem
	WorkItemsService struct {
		client *Client
	}
)

var _ WorkItemsServiceInterface = (*WorkItemsService)(nil)

// WorkItem represents a GitLab work item.
//
// GitLab API docs: https://docs.gitlab.com/api/graphql/reference/#workitem
type WorkItem struct {
	ID          int64
	IID         int64
	Type        string
	State       string
	Status      string
	Title       string
	Description string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	ClosedAt    *time.Time
	WebURL      string
	Author      *BasicUser
	Assignees   []*BasicUser
}

func (wi WorkItem) GID() string {
	return gidGQL{
		Type: "WorkItem",
		ID:   wi.ID,
	}.String()
}

const workItemQuery = `
id
iid
workItemType {
  name
}
state
title
description
author {
  id
  username
  name
  state
  createdAt
  avatarUrl
  webUrl
}
createdAt
updatedAt
closedAt
webUrl
features {
  assignees {
    assignees {
      nodes {
        id
        username
        name
        state
        createdAt
        avatarUrl
        webUrl
      }
    }
  }
  status {
    status {
      name
    }
  }
}
`

// GetWorkItemByID gets a single work item identified by its global ID.
//
// gid is either a string in the form of "gid://gitlab/WorkItem/<ID>", or an integer.
//
// GitLab API docs: https://docs.gitlab.com/api/graphql/reference/#queryworkitem
func (s *WorkItemsService) GetWorkItemByID(gid any, options ...RequestOptionFunc) (*WorkItem, *Response, error) {
	q := GraphQLQuery{
		Query: fmt.Sprintf(`
query ($id: WorkItemID!) {
	workItem(id: $id) {
		%s
	}
}
		`, workItemQuery),
		Variables: map[string]any{},
	}

	switch v := gid.(type) {
	case string:
		q.Variables["id"] = v

	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		q.Variables["id"] = fmt.Sprintf("gid://gitlab/WorkItem/%d", v)

	default:
		return nil, nil, fmt.Errorf("invalid GID type: %T", gid)
	}

	var result struct {
		Data struct {
			WorkItem workItemGQL `json:"workItem"`
		}
		GenericGraphQLErrors
	}

	resp, err := s.client.GraphQL.Do(q, &result, options...)
	if err != nil {
		return nil, resp, err
	}

	if len(result.Errors) != 0 {
		return nil, resp, &GraphQLResponseError{
			Err:    errors.New("GraphQL query failed"),
			Errors: result.GenericGraphQLErrors,
		}
	}

	if result.Data.WorkItem.ID.ID == 0 {
		return nil, resp, ErrNotFound
	}

	return result.Data.WorkItem.unwrap(), resp, nil
}

// GetWorkItem gets a single work item.
//
// fullPath is the full path to either a group or project.
// iid is the internal ID of the work item.
//
// GitLab API docs:
func (s *WorkItemsService) GetWorkItem(fullPath string, iid int64, options ...RequestOptionFunc) (*WorkItem, *Response, error) {
	q := GraphQLQuery{
		Query: fmt.Sprintf(`
query ($fullPath: ID!, $iid: String) {
  namespace(fullPath: $fullPath) {
    workItems(iid: $iid) {
      nodes {
	    %s
      }
    }
  }
}
		`, workItemQuery),
		Variables: map[string]any{
			"fullPath": fullPath,
			"iid":      strconv.FormatInt(iid, 10),
		},
	}

	var result struct {
		Data struct {
			Namespace struct {
				WorkItems struct {
					Nodes []workItemGQL `json:"nodes"`
				} `json:"workItems"`
			} `json:"namespace"`
		}
		GenericGraphQLErrors
	}

	resp, err := s.client.GraphQL.Do(q, &result, options...)
	if err != nil {
		return nil, resp, err
	}

	if len(result.Errors) != 0 {
		return nil, resp, &GraphQLResponseError{
			Err:    errors.New("GraphQL query failed"),
			Errors: result.GenericGraphQLErrors,
		}
	}

	if len(result.Data.Namespace.WorkItems.Nodes) == 0 {
		return nil, resp, ErrNotFound
	}

	wiQL := result.Data.Namespace.WorkItems.Nodes[0]

	return wiQL.unwrap(), resp, nil
}

/*
workItems(

	search: String
	in: [IssuableSearchableField!]
	ids: [WorkItemID!]
	authorUsername: String
	confidential: Boolean
	assigneeUsernames: [String!]
	assigneeWildcardId: AssigneeWildcardId
	labelName: [String!]
	milestoneTitle: [String!]
	milestoneWildcardId: MilestoneWildcardId
	myReactionEmoji: String
	iids: [String!]
	state: IssuableState
	types: [IssueType!]
	createdBefore: Time
	createdAfter: Time
	updatedBefore: Time
	updatedAfter: Time
	dueBefore: Time
	dueAfter: Time
	closedBefore: Time
	closedAfter: Time
	subscribed: SubscriptionStatus
	not: NegatedWorkItemFilterInput
	or: UnionedWorkItemFilterInput
	parentIds: [WorkItemID!]
	releaseTag: [String!]
	releaseTagWildcardId: ReleaseTagWildcardId
	crmContactId: String
	crmOrganizationId: String
	iid: String
	sort: WorkItemSort = CREATED_DESC
	verificationStatusWidget: VerificationStatusFilterInput
	healthStatusFilter: HealthStatusFilter
	weight: String
	weightWildcardId: WeightWildcardId
	iterationId: [ID]
	iterationWildcardId: IterationWildcardId
	iterationCadenceId: [IterationsCadenceID!]
	includeAncestors: Boolean = false
	includeDescendants: Boolean = false
	timeframe: Timeframe
	after: String
	before: String
	first: Int
	last: Int

): WorkItemConnection
*/

// ListWorkItemsOptions represents the available ListWorkItems() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/graphql/reference/#queryworkitems
type ListWorkItemsOptions struct {
	State          *string
	AuthorUsername *string
}

var workItemFieldTypes = map[string]string{
	"state":          "IssuableState",
	"authorUsername": "String",
}

var listWorkItemsTemplate = template.Must(template.New("ListWorkItems").Parse(`
query ListWorkItems($fullPath: ID!{{ range .Fields }}, ${{ .Name }}: {{ .Type }}{{ end }}) {
  namespace(fullPath: $fullPath) {
    workItems({{ range $i, $f := .Fields }}{{ if ne $i 0 }}, {{ end }}{{ $f.Name }}: ${{ $f.Name }}{{ end }}) {
      nodes {
        id
        iid
        title
      }
    }
  }
}
`))

// ListWorkItems lists workitems in a given namespace (group or project).
func (s *WorkItemsService) ListWorkItems(fullPath string, opt *ListWorkItemsOptions, options ...RequestOptionFunc) ([]*WorkItem, *Response, error) {
	type fieldGQL struct {
		Name string
		Type string
	}

	var (
		queryFields    []fieldGQL
		queryVariables = map[string]any{
			"fullPath": fullPath,
		}
	)

	if opt != nil {
		if opt.State != nil {
			queryFields = append(queryFields, fieldGQL{"state", workItemFieldTypes["state"]})
			queryVariables["state"] = opt.State
		}
		if opt.AuthorUsername != nil {
			queryFields = append(queryFields, fieldGQL{"authorUsername", workItemFieldTypes["authorUsername"]})
			queryVariables["authorUsername"] = opt.AuthorUsername
		}
	}

	var queryBuilder strings.Builder

	if err := listWorkItemsTemplate.Execute(&queryBuilder, map[string]any{"Fields": queryFields}); err != nil {
		return nil, nil, err
	}

	query := GraphQLQuery{
		Query:     queryBuilder.String(),
		Variables: queryVariables,
	}

	var result struct {
		Data struct {
			Namespace struct {
				WorkItems struct {
					Nodes []workItemGQL `json:"nodes"`
				} `json:"workItems"`
			} `json:"namespace"`
		}
		GenericGraphQLErrors
	}

	resp, err := s.client.GraphQL.Do(query, &result, options...)
	if err != nil {
		return nil, resp, err
	}

	if len(result.Errors) != 0 {
		return nil, resp, &GraphQLResponseError{
			Err:    errors.New("GraphQL query failed"),
			Errors: result.GenericGraphQLErrors,
		}
	}

	var ret []*WorkItem

	for _, wi := range result.Data.Namespace.WorkItems.Nodes {
		ret = append(ret, wi.unwrap())
	}

	return ret, resp, nil
}

// workItemGQL represents the JSON structure returned by the GraphQL query.
// It is used to parse the response and convert it to the more user-friendly WorkItem type.
type workItemGQL struct {
	ID           gidGQL `json:"id"`
	IID          iidGQL `json:"iid"`
	WorkItemType struct {
		Name string `json:"name"`
	} `json:"workItemType"`
	State       string              `json:"state"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	CreatedAt   *time.Time          `json:"createdAt"`
	UpdatedAt   *time.Time          `json:"updatedAt"`
	ClosedAt    *time.Time          `json:"closedAt"`
	Author      userCoreGQL         `json:"author"`
	Features    workItemFeaturesGQL `json:"features"`
	WebURL      string              `json:"webUrl"`
}

func (w workItemGQL) unwrap() *WorkItem {
	var assignees []*BasicUser

	for _, a := range w.Features.Assignees.Assignees.Nodes {
		assignees = append(assignees, a.unwrap())
	}

	return &WorkItem{
		ID:          w.ID.ID,
		IID:         int64(w.IID),
		Type:        w.WorkItemType.Name,
		State:       w.State,
		Status:      w.Features.Status.Status.Name,
		Title:       w.Title,
		Description: w.Description,
		CreatedAt:   w.CreatedAt,
		UpdatedAt:   w.UpdatedAt,
		ClosedAt:    w.ClosedAt,
		WebURL:      w.WebURL,
		Author:      w.Author.unwrap(),
		Assignees:   assignees,
	}
}

type workItemFeaturesGQL struct {
	Assignees struct {
		Assignees struct {
			Nodes []userCoreGQL `json:"nodes"`
		} `json:"assignees"`
	} `json:"assignees"`
	Status struct {
		Status struct {
			Name string
		}
	}
}

type userCoreGQL struct {
	ID        gidGQL     `json:"id"`
	Username  string     `json:"username"`
	Name      string     `json:"name"`
	State     string     `json:"state"`
	CreatedAt *time.Time `json:"createdAt"`
	AvatarURL string     `json:"avatarUrl"`
	WebURL    string     `json:"webUrl"`
}

func (u userCoreGQL) unwrap() *BasicUser {
	if u.Username == "" {
		return nil
	}

	return &BasicUser{
		ID:        u.ID.ID,
		Username:  u.Username,
		Name:      u.Name,
		State:     u.State,
		Locked:    u.State != "active",
		CreatedAt: u.CreatedAt,
		AvatarURL: u.AvatarURL,
		WebURL:    u.WebURL,
	}
}

// gidGQL is a global ID. It is used by GraphQL to uniquely identify resources.
type gidGQL struct {
	Type string
	ID   int64
}

var gidGQLRegex = regexp.MustCompile(`^gid://gitlab/([^/]+)/(\d+)$`)

func (id *gidGQL) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	m := gidGQLRegex.FindStringSubmatch(s)
	if len(m) != 3 {
		return fmt.Errorf("invalid global ID format: %q", s)
	}

	i, err := strconv.ParseInt(m[2], 10, 64)
	if err != nil {
		return fmt.Errorf("failed parsing %q as numeric ID: %w", s, err)
	}

	id.Type = m[1]
	id.ID = i

	return nil
}

func (id gidGQL) String() string {
	return fmt.Sprintf("gid://gitlab/%s/%d", id.Type, id.ID)
}

// iidGQL represents an int64 ID that is encoded by GraphQL as a string.
// This type is used unmarshal the string response into an int64 type.
type iidGQL int64

func (id *iidGQL) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return fmt.Errorf("failed parsing %q as numeric ID: %w", s, err)
	}

	*id = iidGQL(i)
	return nil
}
