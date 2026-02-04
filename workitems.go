package gitlab

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
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

var listWorkItemsTemplate = template.Must(template.New("ListWorkItems").Parse(`
query ListWorkItems($fullPath: ID!, {{ .Variables.Definitions }}) {
  namespace(fullPath: $fullPath) {
    workItems({{ .Variables.Arguments }}) {
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

type variableGQL struct {
	Name  string
	Type  string
	Value any
}

func (v variableGQL) definition() string {
	return fmt.Sprintf("$%s: %s", v.Name, v.Type)
}

func (v variableGQL) argument() string {
	return fmt.Sprintf("%s: $%s", v.Name, v.Name)
}

type variablesGQL []variableGQL

func (vs variablesGQL) asMap(base map[string]any) map[string]any {
	if base == nil {
		base = make(map[string]any)
	}

	for _, f := range vs {
		base[f.Name] = f.Value
	}

	return base
}

// Definitions generates the GraphQL query variable declarations for use in a query definition.
// It returns a comma-separated string of parameter declarations in the format "$name: Type".
// For example, if fieldsGQL contains fields with names "state" and "authorUsername" with types
// "IssuableState" and "String", it returns: "$state: IssuableState, $authorUsername: String".
// This is typically used in the query signature section of a GraphQL query.
func (vs variablesGQL) Definitions() string {
	var args []string

	for _, v := range vs {
		args = append(args, v.definition())
	}

	return strings.Join(args, ", ")
}

// Arguments generates the GraphQL argument assignments for use in a query body.
// It returns a comma-separated string of argument assignments in the format "name: $name".
// For example, if fieldsGQL contains fields with names "state" and "authorUsername", it returns:
// "state: $state, authorUsername: $authorUsername".
// This is typically used when passing variables to a GraphQL field or connection.
func (vs variablesGQL) Arguments() string {
	var args []string

	for _, v := range vs {
		args = append(args, v.argument())
	}

	return strings.Join(args, ", ")
}

// gqlVariables extracts GraphQL variable definitions from a struct's fields.
// It accepts a pointer to a struct where each field is annotated with a `gql:"name type"` tag.
// The tag specifies the GraphQL variable name and type (e.g., `gql:"state IssuableState"`).
//
// Fields can be excluded using `gql:"-"`. Only non-zero fields are included in the result.
//
// Returns a variablesGQL slice containing the variable name, GraphQL type, and value for each field.
// This can be used to generate both variable definitions (for query signatures) and variable
// arguments (for field parameters) in GraphQL queries.
//
// Returns an error if:
//   - s is not a pointer to a struct
//   - any field is missing a `gql` tag
//   - a `gql` tag has invalid format (must be "name type", except those tagged with "-")
//
// Example:
//
//	type Options struct {
//	    State  *string `gql:"state IssuableState"`
//	    Author *string `gql:"authorUsername String"`
//	}
//	fields, err := gqlQueryArgs(&Options{State: Ptr("opened")})
//	// Returns: [{Name: "state", Type: "IssuableState", Value: "opened"}]
func gqlVariables(s any) (variablesGQL, error) {
	if s == nil {
		return nil, nil
	}

	structValue := reflect.ValueOf(s)
	if structValue.Kind() != reflect.Ptr || structValue.Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a pointer to a struct, got %T", s)
	}

	structValue = structValue.Elem() // Dereference the pointer to get the struct value
	structType := structValue.Type()

	var fields variablesGQL

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		gqlTag := field.Tag.Get("gql")

		switch gqlTag {
		case "":
			return nil, fmt.Errorf("field %s.%s is missing a 'gql' tag", structType.Name(), field.Name)
		case "-":
			continue
		}

		parts := strings.Fields(gqlTag)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid 'gql' tag format for field %s.%s: got %q, want \"name type\"", structType.Name(), field.Name, gqlTag)
		}

		fieldValue := structValue.Field(i)
		if fieldValue.IsZero() {
			continue
		}

		fields = append(fields, variableGQL{
			Name:  parts[0],
			Type:  parts[1],
			Value: fieldValue.Interface(),
		})
	}

	return fields, nil
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
