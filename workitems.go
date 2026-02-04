package gitlab

import (
	"errors"
	"strconv"
	"strings"
	"text/template"
	"time"
)

type (
	WorkItemsServiceInterface interface {
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
		Type:  "WorkItem",
		Int64: wi.ID,
	}.String()
}

// workItemTemplate defines the common fields for a work item in GraphQL queries.
// It's chained from userCoreBasicTemplate so nested templates work.
var workItemTemplate = template.Must(template.Must(userCoreBasicTemplate.Clone()).New("WorkItem").Parse(`
	id
	iid
	workItemType {
	  name
	}
	state
	title
	description
	author {
	  {{ template "UserCoreBasic" }}
	}
	createdAt
	updatedAt
	closedAt
	webUrl
	features {
	  assignees {
	    assignees {
	      nodes {
	        {{ template "UserCoreBasic" }}
	      }
	    }
	  }
	  status {
	    status {
	      name
	    }
	  }
	}
`))

// getWorkItemTemplate is chained from workItemTemplate so it has access to both
// UserCoreBasic and WorkItem templates.
var getWorkItemTemplate = template.Must(template.Must(workItemTemplate.Clone()).New("GetWorkItem").Parse(`
	query GetWorkItem($fullPath: ID!, $iid: String!) {
	  namespace(fullPath: $fullPath) {
	    workItem(iid: $iid) {
	      {{ template "WorkItem" }}
	    }
	  }
	}
`))

// GetWorkItem gets a single work item.
//
// fullPath is the full path to either a group or project.
// iid is the internal ID of the work item.
//
// GitLab API docs:
func (s *WorkItemsService) GetWorkItem(fullPath string, iid int64, options ...RequestOptionFunc) (*WorkItem, *Response, error) {
	var queryBuilder strings.Builder
	if err := getWorkItemTemplate.Execute(&queryBuilder, nil); err != nil {
		return nil, nil, err
	}

	q := GraphQLQuery{
		Query: queryBuilder.String(),
		Variables: map[string]any{
			"fullPath": fullPath,
			"iid":      strconv.FormatInt(iid, 10),
		},
	}

	var result struct {
		Data struct {
			Namespace struct {
				WorkItem *workItemGQL `json:"workItem"`
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

	wiQL := result.Data.Namespace.WorkItem
	if wiQL == nil {
		return nil, resp, ErrNotFound
	}

	return wiQL.unwrap(), resp, nil
}

// ListWorkItemsOptions represents the available ListWorkItems() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/graphql/reference/#queryworkitems
type ListWorkItemsOptions struct {
	AssigneeUsernames    []string `gql:"assigneeUsernames [String!]"`
	AssigneeWildcardID   *string  `gql:"assigneeWildcardId AssigneeWildcardId"`
	AuthorUsername       *string  `gql:"authorUsername String"`
	Confidential         *bool    `gql:"confidential Boolean"`
	CRMContactID         *string  `gql:"crmContactId String"`
	CRMOrganizationID    *string  `gql:"crmOrganizationId String"`
	HealthStatusFilter   *string  `gql:"healthStatusFilter HealthStatusFilter"`
	IDs                  []string `gql:"ids [WorkItemID!]"`
	IIDs                 []string `gql:"iids [String!]"`
	IncludeAncestors     *bool    `gql:"includeAncestors Boolean"`
	IncludeDescendants   *bool    `gql:"includeDescendants Boolean"`
	IterationCadenceID   []string `gql:"iterationCadenceId [IterationsCadenceID!]"`
	IterationID          []string `gql:"iterationId [ID]"`
	IterationWildcardID  *string  `gql:"iterationWildcardId IterationWildcardId"`
	LabelName            []string `gql:"labelName [String!]"`
	MilestoneTitle       []string `gql:"milestoneTitle [String!]"`
	MilestoneWildcardID  *string  `gql:"milestoneWildcardId MilestoneWildcardId"`
	MyReactionEmoji      *string  `gql:"myReactionEmoji String"`
	ParentIDs            []string `gql:"parentIds [WorkItemID!]"`
	ReleaseTag           []string `gql:"releaseTag [String!]"`
	ReleaseTagWildcardID *string  `gql:"releaseTagWildcardId ReleaseTagWildcardId"`
	State                *string  `gql:"state IssuableState"`
	Subscribed           *string  `gql:"subscribed SubscriptionStatus"`
	Types                []string `gql:"types [IssueType!]"`
	Weight               *string  `gql:"weight String"`
	WeightWildcardID     *string  `gql:"weightWildcardId WeightWildcardId"`

	// Time filters
	ClosedAfter   *time.Time `gql:"closedAfter Time"`
	ClosedBefore  *time.Time `gql:"closedBefore Time"`
	CreatedAfter  *time.Time `gql:"createdAfter Time"`
	CreatedBefore *time.Time `gql:"createdBefore Time"`
	DueAfter      *time.Time `gql:"dueAfter Time"`
	DueBefore     *time.Time `gql:"dueBefore Time"`
	UpdatedAfter  *time.Time `gql:"updatedAfter Time"`
	UpdatedBefore *time.Time `gql:"updatedBefore Time"`

	// Sorting
	Sort *string `gql:"sort WorkItemSort"`

	// Search
	Search *string  `gql:"search String"`
	In     []string `gql:"in [IssuableSearchableField!]"`

	// Pagination
	After  *string `gql:"after String"`
	Before *string `gql:"before String"`
	First  *int64  `gql:"first Int"`
	Last   *int64  `gql:"last Int"`
}

// listWorkItemsTemplate is chained from workItemTemplate so it has access to both
// UserCoreBasic and WorkItem templates.
var listWorkItemsTemplate = template.Must(template.Must(workItemTemplate.Clone()).New("ListWorkItems").Parse(`
	query ListWorkItems($fullPath: ID!, {{ .Variables.Definitions }}) {
	  namespace(fullPath: $fullPath) {
	    workItems({{ .Variables.Arguments }}) {
	      nodes {
	        {{ template "WorkItem" }}
	      }
	    }
	  }
	}
`))

// ListWorkItems lists workitems in a given namespace (group or project).
func (s *WorkItemsService) ListWorkItems(fullPath string, opt *ListWorkItemsOptions, options ...RequestOptionFunc) ([]*WorkItem, *Response, error) {
	vars, err := gqlVariables(opt)
	if err != nil {
		return nil, nil, err
	}

	var queryBuilder strings.Builder

	if err := listWorkItemsTemplate.Execute(&queryBuilder, map[string]any{
		"Variables": vars,
	}); err != nil {
		return nil, nil, err
	}

	query := GraphQLQuery{
		Query:     queryBuilder.String(),
		Variables: vars.asMap(map[string]any{"fullPath": fullPath}),
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
	Author      userCoreBasicGQL    `json:"author"`
	Features    workItemFeaturesGQL `json:"features"`
	WebURL      string              `json:"webUrl"`
}

func (w workItemGQL) unwrap() *WorkItem {
	var assignees []*BasicUser

	for _, a := range w.Features.Assignees.Assignees.Nodes {
		assignees = append(assignees, a.unwrap())
	}

	return &WorkItem{
		ID:          w.ID.Int64,
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
			Nodes []userCoreBasicGQL `json:"nodes"`
		} `json:"assignees"`
	} `json:"assignees"`
	Status struct {
		Status struct {
			Name string
		}
	}
}
