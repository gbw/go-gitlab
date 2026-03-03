// EXPERIMENTAL(#2213): The Work Items API is a work in progress and may introduce breaking changes even between minor versions.

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
	//
	// Experimental: The Work Items API is a work in progress and may introduce breaking changes even between minor versions.
	WorkItemsService struct {
		client *Client
	}
)

var _ WorkItemsServiceInterface = (*WorkItemsService)(nil)

// WorkItem represents a GitLab work item.
//
// GitLab API docs: https://docs.gitlab.com/api/graphql/reference/#workitem
//
// Experimental: The Work Items API is a work in progress and may introduce breaking changes even between minor versions.
type WorkItem struct {
	ID          int64
	IID         int64
	Type        string
	State       string
	Status      *string
	Title       string
	Description string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	ClosedAt    *time.Time
	WebURL      string
	Author      *BasicUser
	Assignees   []*BasicUser

	Color        *string
	Confidential bool
	DueDate      *ISOTime
	HealthStatus *string
	IterationID  *int64
	Labels       []LabelDetails
	LinkedItems  []LinkedWorkItem
	MilestoneID  *int64
	Parent       *WorkItemIID
	StartDate    *ISOTime
	Weight       *int64
}

func (wi WorkItem) GID() string {
	return gidGQL{
		Type:  "WorkItem",
		Int64: wi.ID,
	}.String()
}

// WorkItemIID identifies a work item by its namespace path and internal ID.
//
// Experimental: The Work Items API is a work in progress and may introduce breaking changes even between minor versions.
type WorkItemIID struct {
	NamespacePath string
	IID           int64
}

// LinkedWorkItem represents a linked work item with its relationship type.
//
// Experimental: The Work Items API is a work in progress and may introduce breaking changes even between minor versions.
type LinkedWorkItem struct {
	WorkItemIID

	// LinkType is the type of relationship between the work items.
	// Possible values: blocks, is_blocked_by, relates_to
	LinkType string
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
	confidential
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
		color {
			color
			textColor
		}
		healthStatus {
			healthStatus
		}
		hierarchy {
			hasParent
			parent {
				iid
				namespace {
					fullPath
				}
			}
		}
		iteration {
			iteration {
				id
			}
		}
		labels {
			labels {
				nodes {
					id
					title
					color
					description
					descriptionHtml
					textColor
				}
			}
		}
		linkedItems {
			linkedItems {
				nodes {
					workItem {
						iid
						namespace {
							fullPath
						}
					}
					linkType
				}
			}
		}
		milestone {
			milestone {
				id
			}
		}
		startAndDueDate {
			startDate
			dueDate
		}
		status {
			status {
				name
			}
		}
		weight {
			weight
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
// GitLab API docs: https://docs.gitlab.com/api/graphql/reference/#namespaceworkitem
//
// Experimental: The Work Items API is a work in progress and may introduce breaking changes even between minor versions.
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
// GitLab API docs: https://docs.gitlab.com/api/graphql/reference/#namespaceworkitems
//
// Experimental: The Work Item API is work in progress and subject to change even between minor versions.
type ListWorkItemsOptions struct {
	AssigneeUsernames    []string
	AssigneeWildcardID   *string
	AuthorUsername       *string
	Confidential         *bool
	CRMContactID         *string
	CRMOrganizationID    *string
	HealthStatusFilter   *string
	IDs                  []string
	IIDs                 []string
	IncludeAncestors     *bool
	IncludeDescendants   *bool
	IterationCadenceID   []string
	IterationID          []string
	IterationWildcardID  *string
	LabelName            []string
	MilestoneTitle       []string
	MilestoneWildcardID  *string
	MyReactionEmoji      *string
	ParentIDs            []string
	ReleaseTag           []string
	ReleaseTagWildcardID *string
	State                *string
	Subscribed           *string
	Types                []string
	Weight               *string
	WeightWildcardID     *string

	// Time filters
	ClosedAfter   *time.Time
	ClosedBefore  *time.Time
	CreatedAfter  *time.Time
	CreatedBefore *time.Time
	DueAfter      *time.Time
	DueBefore     *time.Time
	UpdatedAfter  *time.Time
	UpdatedBefore *time.Time

	// Sorting
	Sort *string

	// Search
	Search *string
	In     []string

	// Pagination
	After  *string
	Before *string
	First  *int64
	Last   *int64
}

// listWorkItemsTemplate is chained from workItemTemplate so it has access to both
// UserCoreBasic and WorkItem templates.
var listWorkItemsTemplate = template.Must(template.Must(workItemTemplate.Clone()).New("ListWorkItems").Parse(`
	query ListWorkItems(
		$fullPath: ID!
		$assigneeUsernames: [String!]
		$assigneeWildcardId: AssigneeWildcardId
		$authorUsername: String
		$confidential: Boolean
		$crmContactId: String
		$crmOrganizationId: String
		$healthStatusFilter: HealthStatusFilter
		$ids: [WorkItemID!]
		$iids: [String!]
		$includeAncestors: Boolean
		$includeDescendants: Boolean
		$iterationCadenceId: [IterationsCadenceID!]
		$iterationId: [ID]
		$iterationWildcardId: IterationWildcardId
		$labelName: [String!]
		$milestoneTitle: [String!]
		$milestoneWildcardId: MilestoneWildcardId
		$myReactionEmoji: String
		$parentIds: [WorkItemID!]
		$releaseTag: [String!]
		$releaseTagWildcardId: ReleaseTagWildcardId
		$state: IssuableState
		$subscribed: SubscriptionStatus
		$types: [IssueType!]
		$weight: String
		$weightWildcardId: WeightWildcardId
		$closedAfter: Time
		$closedBefore: Time
		$createdAfter: Time
		$createdBefore: Time
		$dueAfter: Time
		$dueBefore: Time
		$updatedAfter: Time
		$updatedBefore: Time
		$sort: WorkItemSort
		$search: String
		$in: [IssuableSearchableField!]
		$after: String
		$before: String
		$first: Int
		$last: Int
	) {
		namespace(fullPath: $fullPath) {
			workItems(
				assigneeUsernames: $assigneeUsernames
				assigneeWildcardId: $assigneeWildcardId
				authorUsername: $authorUsername
				confidential: $confidential
				crmContactId: $crmContactId
				crmOrganizationId: $crmOrganizationId
				healthStatusFilter: $healthStatusFilter
				ids: $ids
				iids: $iids
				includeAncestors: $includeAncestors
				includeDescendants: $includeDescendants
				iterationCadenceId: $iterationCadenceId
				iterationId: $iterationId
				iterationWildcardId: $iterationWildcardId
				labelName: $labelName
				milestoneTitle: $milestoneTitle
				milestoneWildcardId: $milestoneWildcardId
				myReactionEmoji: $myReactionEmoji
				parentIds: $parentIds
				releaseTag: $releaseTag
				releaseTagWildcardId: $releaseTagWildcardId
				state: $state
				subscribed: $subscribed
				types: $types
				weight: $weight
				weightWildcardId: $weightWildcardId
				closedAfter: $closedAfter
				closedBefore: $closedBefore
				createdAfter: $createdAfter
				createdBefore: $createdBefore
				dueAfter: $dueAfter
				dueBefore: $dueBefore
				updatedAfter: $updatedAfter
				updatedBefore: $updatedBefore
				sort: $sort
				search: $search
				in: $in
				after: $after
				before: $before
				first: $first
				last: $last
			) {
				nodes {
					{{ template "WorkItem" }}
				}
				pageInfo {
					endCursor
					hasNextPage
					startCursor
					hasPreviousPage
				}
			}
		}
	}
`))

// ListWorkItems lists workitems in a given namespace (group or project).
//
// GitLab API docs: https://docs.gitlab.com/api/graphql/reference/#namespaceworkitems
//
// Experimental: The Work Items API is a work in progress and may introduce breaking changes even between minor versions.
func (s *WorkItemsService) ListWorkItems(fullPath string, opt *ListWorkItemsOptions, options ...RequestOptionFunc) ([]*WorkItem, *Response, error) {
	var queryBuilder strings.Builder

	if err := listWorkItemsTemplate.Execute(&queryBuilder, nil); err != nil {
		return nil, nil, err
	}

	vars := map[string]any{
		"fullPath":             fullPath,
		"assigneeUsernames":    opt.AssigneeUsernames,
		"assigneeWildcardId":   opt.AssigneeWildcardID,
		"authorUsername":       opt.AuthorUsername,
		"confidential":         opt.Confidential,
		"crmContactId":         opt.CRMContactID,
		"crmOrganizationId":    opt.CRMOrganizationID,
		"healthStatusFilter":   opt.HealthStatusFilter,
		"ids":                  opt.IDs,
		"iids":                 opt.IIDs,
		"includeAncestors":     opt.IncludeAncestors,
		"includeDescendants":   opt.IncludeDescendants,
		"iterationCadenceId":   opt.IterationCadenceID,
		"iterationId":          opt.IterationID,
		"iterationWildcardId":  opt.IterationWildcardID,
		"labelName":            opt.LabelName,
		"milestoneTitle":       opt.MilestoneTitle,
		"milestoneWildcardId":  opt.MilestoneWildcardID,
		"myReactionEmoji":      opt.MyReactionEmoji,
		"parentIds":            opt.ParentIDs,
		"releaseTag":           opt.ReleaseTag,
		"releaseTagWildcardId": opt.ReleaseTagWildcardID,
		"state":                opt.State,
		"subscribed":           opt.Subscribed,
		"types":                opt.Types,
		"weight":               opt.Weight,
		"weightWildcardId":     opt.WeightWildcardID,
		"closedAfter":          opt.ClosedAfter,
		"closedBefore":         opt.ClosedBefore,
		"createdAfter":         opt.CreatedAfter,
		"createdBefore":        opt.CreatedBefore,
		"dueAfter":             opt.DueAfter,
		"dueBefore":            opt.DueBefore,
		"updatedAfter":         opt.UpdatedAfter,
		"updatedBefore":        opt.UpdatedBefore,
		"sort":                 opt.Sort,
		"search":               opt.Search,
		"in":                   opt.In,
		"after":                opt.After,
		"before":               opt.Before,
		"first":                opt.First,
		"last":                 opt.Last,
	}

	query := GraphQLQuery{
		Query:     queryBuilder.String(),
		Variables: vars,
	}

	var result struct {
		Data struct {
			Namespace struct {
				WorkItems connectionGQL[workItemGQL] `json:"workItems"`
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

	resp.PageInfo = &result.Data.Namespace.WorkItems.PageInfo

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
	State        string              `json:"state"`
	Title        string              `json:"title"`
	Description  string              `json:"description"`
	CreatedAt    *time.Time          `json:"createdAt"`
	UpdatedAt    *time.Time          `json:"updatedAt"`
	ClosedAt     *time.Time          `json:"closedAt"`
	Author       *userCoreBasicGQL   `json:"author"`
	Features     workItemFeaturesGQL `json:"features"`
	WebURL       string              `json:"webUrl"`
	Confidential bool                `json:"confidential"`
}

func (w workItemGQL) unwrap() *WorkItem {
	var assignees []*BasicUser

	wi := &WorkItem{
		ID:           w.ID.Int64,
		IID:          int64(w.IID),
		Type:         w.WorkItemType.Name,
		State:        w.State,
		Title:        w.Title,
		Description:  w.Description,
		CreatedAt:    w.CreatedAt,
		UpdatedAt:    w.UpdatedAt,
		ClosedAt:     w.ClosedAt,
		WebURL:       w.WebURL,
		Author:       w.Author.unwrap(),
		Assignees:    assignees,
		Confidential: w.Confidential,
	}

	w.Features.unwrap(wi)

	return wi
}

// workItemFeaturesGQL represents the optional features of the work item.
//
// While the "features" field in the "WorkItem" type is not nullable, each
// feature inside the struct is.
//
// API docs: https://docs.gitlab.com/api/graphql/reference/#workitemfeatures
type workItemFeaturesGQL struct {
	Assignees       *workItemWidgetAssigneesGQL       `json:"assignees"`
	Color           *workItemWidgetColorGQL           `json:"color"`
	HealthStatus    *workItemWidgetHealthStatusGQL    `json:"healthStatus"`
	Hierarchy       *workItemWidgetHierarchyGQL       `json:"hierarchy"`
	Iteration       *workItemWidgetIterationGQL       `json:"iteration"`
	Labels          *workItemWidgetLabelsGQL          `json:"labels"`
	LinkedItems     *workItemWidgetLinkedItemsGQL     `json:"linkedItems"`
	Milestone       *workItemWidgetMilestoneGQL       `json:"milestone"`
	StartAndDueDate *workItemWidgetStartAndDueDateGQL `json:"startAndDueDate"`
	Status          *workItemWidgetStatusGQL          `json:"status"`
	Weight          *workItemWidgetWeightGQL          `json:"weight"`
}

func (f workItemFeaturesGQL) unwrap(wi *WorkItem) {
	wi.Assignees = f.Assignees.unwrap()
	wi.Color = f.Color.unwrap()
	wi.HealthStatus = f.HealthStatus.unwrap()
	wi.Parent = f.Hierarchy.unwrap()
	wi.IterationID = f.Iteration.unwrap()
	wi.Labels = f.Labels.unwrap()
	wi.LinkedItems = f.LinkedItems.unwrap()
	wi.MilestoneID = f.Milestone.unwrap()
	wi.StartDate, wi.DueDate = f.StartAndDueDate.unwrap()
	wi.Status = f.Status.unwrap()
	wi.Weight = f.Weight.unwrap()
}

// workItemWidgetAssigneesGQL represents the assignees widget.
//
// API docs: https://docs.gitlab.com/api/graphql/reference/#workitemwidgetassignees
type workItemWidgetAssigneesGQL struct {
	Assignees connectionGQL[userCoreBasicGQL] `json:"assignees"`
}

func (a *workItemWidgetAssigneesGQL) unwrap() []*BasicUser {
	if a == nil {
		return nil
	}

	ret := make([]*BasicUser, 0, len(a.Assignees.Nodes))

	for _, assignee := range a.Assignees.Nodes {
		ret = append(ret, assignee.unwrap())
	}

	return ret
}

// workItemWidgetColorGQL represents a color widget.
//
// API docs: https://docs.gitlab.com/api/graphql/reference/#workitemwidgetcolor
type workItemWidgetColorGQL struct {
	Color     *string `json:"color"`
	TextColor *string `json:"textColor"`
}

func (c *workItemWidgetColorGQL) unwrap() *string {
	if c == nil {
		return nil
	}

	return c.Color
}

// workItemWidgetHealthStatusGQL represents a health status widget.
//
// API docs: https://docs.gitlab.com/api/graphql/reference/#workitemwidgethealthstatus
type workItemWidgetHealthStatusGQL struct {
	HealthStatus *string `json:"healthStatus"`
}

func (h *workItemWidgetHealthStatusGQL) unwrap() *string {
	if h == nil {
		return nil
	}

	return h.HealthStatus
}

// workItemWidgetHierarchyGQL represents a hierarchy widget.
//
// API docs: https://docs.gitlab.com/api/graphql/reference/#workitemwidgethierarchy
type workItemWidgetHierarchyGQL struct {
	HasParent bool `json:"hasParent"`
	Parent    *struct {
		IID       string `json:"iid"`
		Namespace struct {
			FullPath string `json:"fullPath"`
		} `json:"namespace"`
	} `json:"parent"`
}

func (h *workItemWidgetHierarchyGQL) unwrap() *WorkItemIID {
	if h == nil || !h.HasParent || h.Parent == nil {
		return nil
	}

	iid, err := strconv.ParseInt(h.Parent.IID, 10, 64)
	if err != nil {
		return nil
	}

	return &WorkItemIID{
		NamespacePath: h.Parent.Namespace.FullPath,
		IID:           iid,
	}
}

// workItemWidgetIterationGQL represents a iteration widget.
//
// API docs: https://docs.gitlab.com/api/graphql/reference/#workitemwidgetiteration
type workItemWidgetIterationGQL struct {
	Iteration *struct {
		ID gidGQL `json:"id"`
	} `json:"iteration"`
}

func (c *workItemWidgetIterationGQL) unwrap() *int64 {
	if c == nil || c.Iteration == nil {
		return nil
	}

	return Ptr(c.Iteration.ID.Int64)
}

// workItemWidgetLabelsGQL represents the labels widget.
//
// API docs: https://docs.gitlab.com/api/graphql/reference/#workitemwidgetlabels
type workItemWidgetLabelsGQL struct {
	Labels *connectionGQL[labelGQL] `json:"labels"`
}

func (l *workItemWidgetLabelsGQL) unwrap() []LabelDetails {
	if l == nil || l.Labels == nil {
		return nil
	}

	ret := make([]LabelDetails, 0, len(l.Labels.Nodes))

	for _, label := range l.Labels.Nodes {
		ret = append(ret, label.unwrap())
	}

	return ret
}

// workItemWidgetLinkedItemsGQL represents the linked items widget.
//
// API docs: https://docs.gitlab.com/api/graphql/reference/#workitemwidgetlinkeditems
type workItemWidgetLinkedItemsGQL struct {
	LinkedItems *connectionGQL[struct {
		WorkItem *struct {
			IID       string `json:"iid"`
			Namespace struct {
				FullPath string `json:"fullPath"`
			} `json:"namespace"`
		} `json:"workItem"`
		LinkType string `json:"linkType"`
	}] `json:"linkedItems"`
}

func (li *workItemWidgetLinkedItemsGQL) unwrap() []LinkedWorkItem {
	if li == nil || li.LinkedItems == nil {
		return nil
	}

	var ret []LinkedWorkItem

	for _, item := range li.LinkedItems.Nodes {
		if item.WorkItem == nil {
			continue
		}

		iid, err := strconv.ParseInt(item.WorkItem.IID, 10, 64)
		if err != nil {
			continue
		}

		ret = append(ret, LinkedWorkItem{
			WorkItemIID: WorkItemIID{
				NamespacePath: item.WorkItem.Namespace.FullPath,
				IID:           iid,
			},
			LinkType: item.LinkType,
		})
	}

	return ret
}

// workItemWidgetMilestoneGQL represents the milestone widget.
//
// API docs: https://docs.gitlab.com/api/graphql/reference/#workitemwidgetmilestone
type workItemWidgetMilestoneGQL struct {
	Milestone *struct {
		ID gidGQL `json:"id"`
	} `json:"milestone"`
}

func (m *workItemWidgetMilestoneGQL) unwrap() *int64 {
	if m == nil || m.Milestone == nil {
		return nil
	}

	return Ptr(m.Milestone.ID.Int64)
}

// workItemWidgetStartAndDueDateGQL represents a start and due date widget.
//
// API docs: https://docs.gitlab.com/api/graphql/reference/#workitemwidgetstartandduedate
type workItemWidgetStartAndDueDateGQL struct {
	DueDate   *ISOTime `json:"dueDate"`
	IsFixed   bool     `json:"isFixed"`
	StartDate *ISOTime `json:"startDate"`
}

func (du *workItemWidgetStartAndDueDateGQL) unwrap() (start, due *ISOTime) {
	if du == nil {
		return nil, nil
	}

	return du.StartDate, du.DueDate
}

// workItemWidgetStatusGQL represents the status widget.
//
// API docs: https://docs.gitlab.com/api/graphql/reference/#workitemwidgetstatus
type workItemWidgetStatusGQL struct {
	Status *struct {
		Name *string `json:"name"`
	} `json:"status"`
}

func (s *workItemWidgetStatusGQL) unwrap() *string {
	if s == nil || s.Status == nil {
		return nil
	}

	return s.Status.Name
}

// workItemWidgetWeightGQL represents the weight widget.
//
// API docs: https://docs.gitlab.com/api/graphql/reference/#workitemwidgetweight
type workItemWidgetWeightGQL struct {
	Weight *int64 `json:"weight"`
}

func (w *workItemWidgetWeightGQL) unwrap() *int64 {
	if w == nil {
		return nil
	}

	return w.Weight
}
