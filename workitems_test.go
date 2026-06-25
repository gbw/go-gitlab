package gitlab

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed testdata/get_workitem.json
var getWorkItemResponse []byte

//go:embed testdata/update_workitem.json
var updateWorkItemResponse []byte

//go:embed testdata/list_work_items.json
var listWorkItemResponse []byte

// setupGraphQLHandler registers a handler on mux for /api/graphql that decodes
// the incoming GraphQL request, calls handleFn with the decoded query, and
// writes the response produced by handleFn.
func setupGraphQLHandler(t *testing.T, mux *http.ServeMux, handleFn func(w http.ResponseWriter, q GraphQLQuery)) {
	t.Helper()

	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		testMethod(t, r, http.MethodPost)

		var q GraphQLQuery
		if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		handleFn(w, q)
	})
}

// writeResponse writes the contents of response to w.
func writeResponse(t *testing.T, w http.ResponseWriter, response io.WriterTo) {
	t.Helper()

	if _, err := response.WriteTo(w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// setupWorkItemTypesHandler registers a /api/graphql handler that validates the
// ListWorkItemTypes query and writes response.
func setupWorkItemTypesHandler(t *testing.T, mux *http.ServeMux, response io.Reader) {
	t.Helper()

	setupGraphQLHandler(t, mux, func(w http.ResponseWriter, q GraphQLQuery) {
		assertQueryMatches(t, q.Query, "testdata/query_list_work_item_types.graphql")
		_, _ = io.Copy(w, response)
	})
}

// assertQueryMatches reads the exact query file from `testdata` and asserts that the
// received query string matches it (after trimming leading/trailing whitespace).
// Exact query files in testdata/ contain the exact GraphQL query strings the client
// is expected to produce, making query validation always active and intentional.
func assertQueryMatches(t *testing.T, got, goldenFile string) {
	t.Helper()

	data, err := os.ReadFile(goldenFile)
	require.NoError(t, err, "reading exact file %s", goldenFile)

	assert.Equal(t, strings.TrimSpace(string(data)), strings.TrimSpace(got))
}

// ---------------------------------------------------------------------------
// GetWorkItem tests
// ---------------------------------------------------------------------------

func TestGetWorkItem_SuccessfulResponse(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	// GIVEN a GetWorkItem request for an existing work item
	setupGraphQLHandler(t, mux, func(w http.ResponseWriter, q GraphQLQuery) {
		assertQueryMatches(t, q.Query, "testdata/query_get_work_item.graphql")
		writeResponse(t, w, bytes.NewReader(getWorkItemResponse))
	})

	// WHEN GetWorkItem is called
	got, _, err := client.WorkItems.GetWorkItem("gitlab-com/gl-infra/platform/runway/team", 756)

	// THEN the work item is returned without error
	require.NoError(t, err)
	assert.Equal(t, &WorkItem{
		ID:          179785913,
		IID:         756,
		Type:        "Task",
		State:       "OPEN",
		Status:      Ptr("New"),
		Title:       "Update Helm charts to use Argo Rollouts for progressive deployments",
		Description: "## Overview\n\nUpdate Runway Helm charts to generate Argo Rollout resources ...",
		CreatedAt:   Ptr(time.Date(2026, time.January, 6, 15, 9, 24, 0, time.UTC)),
		UpdatedAt:   Ptr(time.Date(2026, time.January, 9, 13, 6, 8, 0, time.UTC)),
		WebURL:      "https://gitlab.com/gitlab-com/gl-infra/platform/runway/team/-/work_items/756",
		Author: &BasicUser{
			ID:        5532616,
			Username:  "swainaina",
			Name:      "Silvester Wainaina",
			State:     "active",
			CreatedAt: Ptr(time.Date(2020, time.March, 2, 6, 29, 14, 0, time.UTC)),
			AvatarURL: "/uploads/-/system/user/avatar/5532616/avatar.png",
			WebURL:    "https://gitlab.com/swainaina",
		},
		Assignees: []*BasicUser{
			{
				ID:        5532616,
				Username:  "swainaina",
				Name:      "Silvester Wainaina",
				State:     "active",
				CreatedAt: Ptr(time.Date(2020, time.March, 2, 6, 29, 14, 0, time.UTC)),
				AvatarURL: "/uploads/-/system/user/avatar/5532616/avatar.png",
				WebURL:    "https://gitlab.com/swainaina",
			},
		},
		DueDate:      Ptr(ISOTime(time.Date(2026, time.July, 31, 0, 0, 0, 0, time.UTC))),
		HealthStatus: Ptr("onTrack"),
		IterationID:  Ptr(int64(2748074)),
		Labels: []LabelDetails{
			{
				ID:              32754251,
				Name:            "Category:Runway",
				Color:           "#6699cc",
				Description:     "",
				DescriptionHTML: "",
				TextColor:       "#FFFFFF",
			},
			{
				ID:              32832335,
				Name:            "Service::Runway",
				Color:           "#d1d100",
				Description:     "",
				DescriptionHTML: "",
				TextColor:       "#1F1E24",
			},
			{
				ID:              12970969,
				Name:            "workflow-infra::Triage",
				Color:           "#FEAF09",
				Description:     "For @gitlab-com/gl-infra/managers to triage, prioritize, and assign.",
				DescriptionHTML: "For <a href=\"/gitlab-com/gl-infra/managers\" data-reference-type=\"user\" data-group=\"4684757\" data-container=\"body\" data-placement=\"top\" class=\"gfm gfm-project_member js-user-link\" title=\"GitLab.com / GitLab Infrastructure Team / Infrastructure Managers\">@gitlab-com/gl-infra/managers</a> to triage, prioritize, and assign.",
				TextColor:       "#1F1E24",
			},
		},
		LinkedItems: []LinkedWorkItem{
			{
				WorkItemIID: WorkItemIID{
					NamespacePath: "gitlab-com/gl-infra/platform/runway/team",
					IID:           774,
				},
				LinkType: "relates_to",
			},
		},
		MilestoneID: Ptr(int64(6161376)),
		Parent: &WorkItemIID{
			NamespacePath: "gitlab-com/gl-infra/platform/runway/team",
			IID:           673,
		},
		StartDate: Ptr(ISOTime(time.Date(2025, time.August, 1, 0, 0, 0, 0, time.UTC))),
		Weight:    Ptr(int64(8)),
	}, got)
}

func TestGetWorkItem_EmptyNodesReturnsNotFound(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	// GIVEN a GetWorkItem request where the server returns an empty nodes list
	setupGraphQLHandler(t, mux, func(w http.ResponseWriter, q GraphQLQuery) {
		assertQueryMatches(t, q.Query, "testdata/query_get_work_item.graphql")
		writeResponse(t, w, strings.NewReader(`
			{
			    "data": {
			        "project": {
			            "workItems": {
			                "nodes": []
			            }
			        }
			    },
			    "correlationId": "9c5818b053a3354c-IAD"
			}
		`))
	})

	// WHEN GetWorkItem is called for a non-existent IID
	got, _, err := client.WorkItems.GetWorkItem("gitlab-com/gl-infra/platform/runway/team", 999)

	// THEN ErrNotFound is returned and the result is nil
	require.ErrorIs(t, err, ErrNotFound)
	assert.Nil(t, got)
}

func TestGetWorkItem_NullProjectReturnsNotFound(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	// GIVEN a GetWorkItem request where the server returns a null project
	setupGraphQLHandler(t, mux, func(w http.ResponseWriter, q GraphQLQuery) {
		assertQueryMatches(t, q.Query, "testdata/query_get_work_item.graphql")
		writeResponse(t, w, strings.NewReader(`
			{
			"data": {
				"project": null
			},
			"correlationId": "9c59850aa3cdf515-IAD"
			}
		`))
	})

	// WHEN GetWorkItem is called for a non-existent project
	got, _, err := client.WorkItems.GetWorkItem("does/not/exist", 1)

	// THEN ErrNotFound is returned and the result is nil
	require.ErrorIs(t, err, ErrNotFound)
	assert.Nil(t, got)
}

// ---------------------------------------------------------------------------
// ListWorkItems tests
// ---------------------------------------------------------------------------

func TestListWorkItems_SuccessfulQueryWithAuthorUsername(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	// GIVEN a ListWorkItems request filtered by authorUsername
	setupGraphQLHandler(t, mux, func(w http.ResponseWriter, q GraphQLQuery) {
		writeResponse(t, w, bytes.NewReader(listWorkItemResponse))
	})

	// WHEN ListWorkItems is called
	got, resp, err := client.WorkItems.ListWorkItems(
		"gitlab-com/gl-infra/platform/runway/team",
		&ListWorkItemsOptions{AuthorUsername: Ptr("fforster")},
	)

	// THEN the work items and page info are returned without error
	require.NoError(t, err)
	assert.Equal(t, []*WorkItem{
		{
			ID:    181297786,
			IID:   39,
			Title: "Phase 6: Rollout to Additional Services",
		},
		{
			ID:    181297779,
			IID:   38,
			Title: "Phase 5: Dedicated Integration",
		},
	}, got)
	assert.Equal(t, &PageInfo{
		EndCursor:       "eyJjcmVhdGVkX2F0IjoiMjAyNi0wMS0xNiAxMzozMjo0Ny44NTEyMTUwMDAgKzAwMDAiLCJpZCI6IjE4MTI5Nzc3OSJ9",
		HasNextPage:     true,
		StartCursor:     "eyJjcmVhdGVkX2F0IjoiMjAyNi0wMS0xNiAxMzozMjo1Ny43NjgxNzYwMDAgKzAwMDAiLCJpZCI6IjE4MTI5Nzc4NiJ9",
		HasPreviousPage: false,
	}, resp.PageInfo)
}

func TestListWorkItems_SuccessfulResponseWithSingleWorkItem(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	// GIVEN a ListWorkItems request that returns a single work item
	setupGraphQLHandler(t, mux, func(w http.ResponseWriter, q GraphQLQuery) {
		writeResponse(t, w, strings.NewReader(`
			{
			  "data": {
			    "namespace": {
			      "workItems": {
			        "nodes": [
			          {
			            "id": "gid://gitlab/WorkItem/181297786",
			            "iid": "39",
			            "title": "Phase 6: Rollout to Additional Services"
			          }
			        ]
			      }
			    }
			  },
			  "correlationId": "9c88d56b0061dfef-IAD"
			}
		`))
	})

	// WHEN ListWorkItems is called with state and authorUsername filters
	got, _, err := client.WorkItems.ListWorkItems(
		"gitlab-com/gl-infra/platform/runway/team",
		&ListWorkItemsOptions{
			State:          Ptr("opened"),
			AuthorUsername: Ptr("fforster"),
		},
	)

	// THEN the single work item is returned without error
	require.NoError(t, err)
	assert.Equal(t, []*WorkItem{
		{
			ID:    181297786,
			IID:   39,
			Title: "Phase 6: Rollout to Additional Services",
		},
	}, got)
}

func TestListWorkItems_EmptyResponseIsNotAnError(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	// GIVEN a ListWorkItems request that returns an empty nodes list
	setupGraphQLHandler(t, mux, func(w http.ResponseWriter, q GraphQLQuery) {
		writeResponse(t, w, strings.NewReader(`
			{
			  "data": {
			    "namespace": {
			      "workItems": {
			        "nodes": []
			      }
			    }
			  }
			}
		`))
	})

	// WHEN ListWorkItems is called
	got, _, err := client.WorkItems.ListWorkItems(
		"gitlab-com/gl-infra/platform/runway/team",
		&ListWorkItemsOptions{
			State:          Ptr("opened"),
			AuthorUsername: Ptr("fforster"),
		},
	)

	// THEN no error is returned and the result is nil (not an empty slice)
	require.NoError(t, err)
	assert.Nil(t, got)
}

func TestListWorkItems_AllOptionsFieldsIncludedInQuery(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	// GIVEN a ListWorkItems request with all available options set
	setupGraphQLHandler(t, mux, func(w http.ResponseWriter, q GraphQLQuery) {
		writeResponse(t, w, strings.NewReader(`
			{
			  "data": {
			    "namespace": {
			      "workItems": {
			        "nodes": []
			      }
			    }
			  }
			}
		`))
	})

	opt := &ListWorkItemsOptions{
		// Main filters
		AssigneeUsernames:    []string{"user1", "user2"},
		AssigneeWildcardID:   Ptr("NONE"),
		AuthorUsername:       Ptr("fforster"),
		Confidential:         Ptr(true),
		CRMContactID:         Ptr("contact123"),
		CRMOrganizationID:    Ptr("org456"),
		HealthStatusFilter:   Ptr("onTrack"),
		IDs:                  []string{"gid://gitlab/WorkItem/1", "gid://gitlab/WorkItem/2"},
		IIDs:                 []string{"1", "2", "3"},
		IncludeAncestors:     Ptr(true),
		IncludeDescendants:   Ptr(false),
		IterationCadenceID:   []string{"cadence1"},
		IterationID:          []string{"iter1", "iter2"},
		IterationWildcardID:  Ptr("CURRENT"),
		LabelName:            []string{"bug", "urgent"},
		MilestoneTitle:       []string{"v1.0", "v2.0"},
		MilestoneWildcardID:  Ptr("STARTED"),
		MyReactionEmoji:      Ptr("thumbsup"),
		ParentIDs:            []string{"gid://gitlab/WorkItem/100"},
		ReleaseTag:           []string{"v1.0.0"},
		ReleaseTagWildcardID: Ptr("ANY"),
		State:                Ptr("opened"),
		Subscribed:           Ptr("EXPLICITLY_SUBSCRIBED"),
		Types:                []string{"ISSUE", "TASK"},
		Weight:               Ptr("5"),
		WeightWildcardID:     Ptr("NONE"),
		// Time filters
		ClosedAfter:   Ptr(time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC)),
		ClosedBefore:  Ptr(time.Date(2026, time.February, 1, 0, 0, 0, 0, time.UTC)),
		CreatedAfter:  Ptr(time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC)),
		CreatedBefore: Ptr(time.Date(2026, time.February, 1, 0, 0, 0, 0, time.UTC)),
		DueAfter:      Ptr(time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC)),
		DueBefore:     Ptr(time.Date(2026, time.February, 1, 0, 0, 0, 0, time.UTC)),
		UpdatedAfter:  Ptr(time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC)),
		UpdatedBefore: Ptr(time.Date(2026, time.February, 1, 0, 0, 0, 0, time.UTC)),
		// Sorting
		Sort: Ptr("CREATED_DESC"),
		// Search
		Search: Ptr("bug"),
		In:     []string{"TITLE", "DESCRIPTION"},
		// Pagination
		After:  Ptr("cursor123"),
		Before: Ptr("cursor456"),
		First:  Ptr(int64(10)),
		Last:   Ptr(int64(5)),
	}

	// WHEN ListWorkItems is called with all options
	got, _, err := client.WorkItems.ListWorkItems("gitlab-com/gl-infra/platform/runway/team", opt)

	// THEN no error is returned and the result is nil (empty response)
	require.NoError(t, err)
	assert.Nil(t, got)
}

func TestListWorkItems_Pagination(t *testing.T) {
	t.Parallel()

	responses := map[string]string{
		/* page 0 */ "": `
			{
			  "data": {
			    "namespace": {
			      "workItems": {
			        "nodes": [
			          {
			            "id": "gid://gitlab/WorkItem/181297786",
			            "iid": "39",
			            "title": "Phase 6: Rollout to Additional Services"
			          },
			          {
			            "id": "gid://gitlab/WorkItem/181297779",
			            "iid": "38",
			            "title": "Phase 5: Dedicated Integration"
			          }
			        ],
			        "pageInfo": {
			          "endCursor": "eyJjcmVhdGVkX2F0IjoiMjAyNi0wMS0xNiAxMzozMjo0Ny44NTEyMTUwMDAgKzAwMDAiLCJpZCI6IjE4MTI5Nzc3OSJ9",
			          "hasNextPage": true,
			          "startCursor": "eyJjcmVhdGVkX2F0IjoiMjAyNi0wMS0xNiAxMzozMjo1Ny43NjgxNzYwMDAgKzAwMDAiLCJpZCI6IjE4MTI5Nzc4NiJ9",
			          "hasPreviousPage": false
			        }
			      }
			    }
			  },
			  "correlationId": "9ccb04130038971c-IAD"
			}
		`,
		/* page 1 */ "eyJjcmVhdGVkX2F0IjoiMjAyNi0wMS0xNiAxMzozMjo0Ny44NTEyMTUwMDAgKzAwMDAiLCJpZCI6IjE4MTI5Nzc3OSJ9": `
			{
			  "data": {
			    "namespace": {
			      "workItems": {
			        "nodes": [
			          {
			            "id": "gid://gitlab/WorkItem/181297773",
			            "iid": "37",
			            "title": "Phase 4: Pilot Service Migration & Validation"
			          },
			          {
			            "id": "gid://gitlab/WorkItem/181297769",
			            "iid": "36",
			            "title": "Phase 3: GitLab Helm Chart Integration & Values Management"
			          }
			        ],
			        "pageInfo": {
			          "endCursor": "eyJjcmVhdGVkX2F0IjoiMjAyNi0wMS0xNiAxMzozMjozMS4yNTcxNzIwMDAgKzAwMDAiLCJpZCI6IjE4MTI5Nzc2OSJ9",
			          "hasNextPage": true,
			          "startCursor": "eyJjcmVhdGVkX2F0IjoiMjAyNi0wMS0xNiAxMzozMjozOS4xMzMwOTEwMDAgKzAwMDAiLCJpZCI6IjE4MTI5Nzc3MyJ9",
			          "hasPreviousPage": true
			        }
			      }
			    }
			  },
			  "correlationId": "9ccb232d6071931b-IAD"
			}
		`,
		/* page 2 */ "eyJjcmVhdGVkX2F0IjoiMjAyNi0wMS0xNiAxMzozMjozMS4yNTcxNzIwMDAgKzAwMDAiLCJpZCI6IjE4MTI5Nzc2OSJ9": `
			{
			  "data": {
			    "namespace": {
			      "workItems": {
			        "nodes": [
			          {
			            "id": "gid://gitlab/WorkItem/181297761",
			            "iid": "35",
			            "title": "Phase 2: Dual-Variant Chart Generation"
			          },
			          {
			            "id": "gid://gitlab/WorkItem/181286354",
			            "iid": "34",
			            "title": "Phase 1: Foundation & Library Chart"
			          }
			        ],
			        "pageInfo": {
			          "endCursor": "eyJjcmVhdGVkX2F0IjoiMjAyNi0wMS0xNiAxMDo0MDo1My42MTIyOTYwMDAgKzAwMDAiLCJpZCI6IjE4MTI4NjM1NCJ9",
			          "hasNextPage": false,
			          "startCursor": "eyJjcmVhdGVkX2F0IjoiMjAyNi0wMS0xNiAxMzozMjoyMi41MDUyNTMwMDAgKzAwMDAiLCJpZCI6IjE4MTI5Nzc2MSJ9",
			          "hasPreviousPage": true
			        }
			      }
			    }
			  },
			  "correlationId": "9ccb265ff56b931b-IAD"
			}
		`,
	}

	mux, client := setup(t)

	// GIVEN a server that returns paginated work items across three pages
	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		testMethod(t, r, http.MethodPost)

		var q GraphQLQuery

		if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var after string
		if a, ok := q.Variables["after"]; ok && a != nil {
			after = a.(string)
		}

		resp, ok := responses[after]
		if !ok {
			http.Error(w, fmt.Sprintf("unexpected after cursor: %q", after), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		_, err := io.WriteString(w, resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	opt := ListWorkItemsOptions{
		State:          Ptr("opened"),
		AuthorUsername: Ptr("fforster"),
		First:          Ptr(int64(2)),
	}

	// WHEN ScanAndCollect is used to paginate through all work items
	got, err := ScanAndCollect(func(p PaginationOptionFunc) ([]*WorkItem, *Response, error) {
		return client.WorkItems.ListWorkItems("unit/test", &opt, p)
	})

	// THEN all six work items across three pages are returned without error
	require.NoError(t, err)
	assert.Equal(t, []*WorkItem{
		{
			ID:    181297786,
			IID:   39,
			Title: "Phase 6: Rollout to Additional Services",
		},
		{
			ID:    181297779,
			IID:   38,
			Title: "Phase 5: Dedicated Integration",
		},
		{
			ID:    181297773,
			IID:   37,
			Title: "Phase 4: Pilot Service Migration & Validation",
		},
		{
			ID:    181297769,
			IID:   36,
			Title: "Phase 3: GitLab Helm Chart Integration & Values Management",
		},
		{
			ID:    181297761,
			IID:   35,
			Title: "Phase 2: Dual-Variant Chart Generation",
		},
		{
			ID:    181286354,
			IID:   34,
			Title: "Phase 1: Foundation & Library Chart",
		},
	}, got)
}

// ---------------------------------------------------------------------------
// CreateWorkItem tests
// ---------------------------------------------------------------------------

func TestCreateWorkItem_SuccessfulCreationWithTitleOnly(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	// GIVEN a CreateWorkItem mutation request with only a title
	setupGraphQLHandler(t, mux, func(w http.ResponseWriter, q GraphQLQuery) {
		assertQueryMatches(t, q.Query, "testdata/mutation_create_work_item.graphql")

		gotInputs := q.Variables["input"].(map[string]any)
		assert.Equal(t, "New Task", gotInputs["title"], "input %q mismatch", "title")

		writeResponse(t, w, strings.NewReader(`
			{
			  "data": {
			    "workItemCreate": {
			      "workItem": {
			        "id": "gid://gitlab/WorkItem/181297786",
			        "iid": "40",
			        "workItemType": {
			          "name": "Task"
			        },
			        "state": "OPEN",
			        "title": "New Task",
			        "description": "",
			        "author": {
			          "id": "gid://gitlab/User/5532616",
			          "username": "fforster",
			          "name": "Florian Forster",
			          "state": "active",
			          "locked": false,
			          "createdAt": "2020-03-02T06:29:14Z",
			          "avatarUrl": "/uploads/-/system/user/avatar/5532616/avatar.png",
			          "webUrl": "https://gitlab.com/fforster"
			        },
			        "createdAt": "2026-02-06T10:00:00Z",
			        "updatedAt": "2026-02-06T10:00:00Z",
			        "closedAt": null,
			        "webUrl": "https://gitlab.com/gitlab-com/gl-infra/platform/runway/team/-/work_items/40",
			        "features": {
			          "assignees": {
			            "assignees": {
			              "nodes": []
			            }
			          },
			          "status": {
			            "status": {
			              "name": "New"
			            }
			          }
			        }
			      }
			    }
			  },
			  "correlationId": "9c88d56b0061dfef-IAD"
			}
		`))
	})

	// WHEN CreateWorkItem is called with a title-only option
	got, _, err := client.WorkItems.CreateWorkItem(
		"gitlab-com/gl-infra/platform/runway/team",
		WorkItemTypeIssue,
		&CreateWorkItemOptions{Title: "New Task"},
	)

	// THEN the created work item is returned without error
	require.NoError(t, err)
	assert.Equal(t, &WorkItem{
		ID:          181297786,
		IID:         40,
		Type:        "Task",
		State:       "OPEN",
		Status:      Ptr("New"),
		Title:       "New Task",
		Description: "",
		CreatedAt:   Ptr(time.Date(2026, time.February, 6, 10, 0, 0, 0, time.UTC)),
		UpdatedAt:   Ptr(time.Date(2026, time.February, 6, 10, 0, 0, 0, time.UTC)),
		WebURL:      "https://gitlab.com/gitlab-com/gl-infra/platform/runway/team/-/work_items/40",
		Author: &BasicUser{
			ID:        5532616,
			Username:  "fforster",
			Name:      "Florian Forster",
			State:     "active",
			CreatedAt: Ptr(time.Date(2020, time.March, 2, 6, 29, 14, 0, time.UTC)),
			AvatarURL: "/uploads/-/system/user/avatar/5532616/avatar.png",
			WebURL:    "https://gitlab.com/fforster",
		},
		Assignees: nil,
	}, got)
}

func TestCreateWorkItem_SuccessfulCreationWithAllOptions(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	wantInputs := map[string]any{
		"title":        "New Issue",
		"confidential": true,
		"createSource": "api",
		"createdAt":    "2026-02-05T10:00:00Z",
		"descriptionWidget": map[string]any{
			"description": "This is a detailed description",
		},
		"assigneesWidget": map[string]any{
			"assigneeIds": []any{"gid://gitlab/User/123", "gid://gitlab/User/456"},
		},
		"milestoneWidget": map[string]any{
			"milestoneId": "gid://gitlab/Milestone/234",
		},
		"crmContactsWidget": map[string]any{
			"contactIds": []any{"gid://gitlab/CustomerRelations::Contact/1001", "gid://gitlab/CustomerRelations::Contact/1002"},
		},
		"hierarchyWidget": map[string]any{
			"parentId": "gid://gitlab/WorkItem/100",
		},
		"labelsWidget": map[string]any{
			"labelIds": []any{"gid://gitlab/Label/789", "gid://gitlab/Label/790"},
		},
		"linkedItemsWidget": map[string]any{
			"linkType":     "RELATED",
			"workItemsIds": []any{"gid://gitlab/WorkItem/1101"},
		},
		"startAndDueDateWidget": map[string]any{
			"startDate": "2026-02-01",
			"dueDate":   "2026-03-01",
		},
		"weightWidget": map[string]any{
			"weight": float64(5),
		},
		"healthStatusWidget": map[string]any{
			"healthStatus": "onTrack",
		},
		"iterationWidget": map[string]any{
			"iterationId": "gid://gitlab/Iteration/567",
		},
		"colorWidget": map[string]any{
			"color": "#FF0000",
		},
	}

	// GIVEN a CreateWorkItem mutation request with all options set
	setupGraphQLHandler(t, mux, func(w http.ResponseWriter, q GraphQLQuery) {
		assertQueryMatches(t, q.Query, "testdata/mutation_create_work_item.graphql")

		gotInputs := q.Variables["input"].(map[string]any)
		for k, v := range wantInputs {
			assert.Equal(t, v, gotInputs[k], "input %q mismatch", k)
		}

		writeResponse(t, w, strings.NewReader(`
			{
			  "data": {
			    "workItemCreate": {
			      "workItem": {
			        "id": "gid://gitlab/WorkItem/181297787",
			        "iid": "41",
			        "workItemType": {
			          "name": "Issue"
			        },
			        "state": "OPEN",
			        "title": "New Issue",
			        "description": "This is a detailed description",
			        "author": {
			          "id": "gid://gitlab/User/5532616",
			          "username": "fforster",
			          "name": "Florian Forster",
			          "state": "active",
			          "locked": false,
			          "createdAt": "2020-03-02T06:29:14Z",
			          "avatarUrl": "/uploads/-/system/user/avatar/5532616/avatar.png",
			          "webUrl": "https://gitlab.com/fforster"
			        },
			        "createdAt": "2026-02-06T10:00:00Z",
			        "updatedAt": "2026-02-06T10:00:00Z",
			        "closedAt": null,
			        "webUrl": "https://gitlab.com/gitlab-com/gl-infra/platform/runway/team/-/work_items/41",
			        "features": {
			          "assignees": {
			            "assignees": {
			              "nodes": []
			            }
			          },
			          "status": {
			            "status": {
			              "name": "New"
			            }
			          }
			        }
			      }
			    }
			  },
			  "correlationId": "9c88d56b0061dfef-IAD"
			}
		`))
	})

	opt := &CreateWorkItemOptions{
		// Required
		Title: "New Issue",
		// Optional
		Description:   Ptr("This is a detailed description"),
		Confidential:  Ptr(true),
		AssigneeIDs:   []int64{123, 456},
		MilestoneID:   Ptr(int64(234)),
		CreateSource:  Ptr("api"),
		CreatedAt:     Ptr(time.Date(2026, time.February, 5, 10, 0, 0, 0, time.UTC)),
		CRMContactIDs: []int64{1001, 1002},
		ParentID:      Ptr(int64(100)),
		LabelIDs:      []int64{789, 790},
		LinkedItems: &CreateWorkItemOptionsLinkedItems{
			LinkType:    Ptr("RELATED"),
			WorkItemIDs: []int64{1101},
		},
		StartDate:    Ptr(ISOTime(time.Date(2026, time.February, 1, 0, 0, 0, 0, time.UTC))),
		DueDate:      Ptr(ISOTime(time.Date(2026, time.March, 1, 0, 0, 0, 0, time.UTC))),
		Weight:       Ptr(int64(5)),
		HealthStatus: Ptr("onTrack"),
		IterationID:  Ptr(int64(567)),
		Color:        Ptr("#FF0000"),
	}

	// WHEN CreateWorkItem is called with all options
	got, _, err := client.WorkItems.CreateWorkItem(
		"gitlab-com/gl-infra/platform/runway/team",
		WorkItemTypeIssue,
		opt,
	)

	// THEN the created work item is returned without error
	require.NoError(t, err)
	assert.Equal(t, &WorkItem{
		ID:          181297787,
		IID:         41,
		Type:        "Issue",
		State:       "OPEN",
		Status:      Ptr("New"),
		Title:       "New Issue",
		Description: "This is a detailed description",
		CreatedAt:   Ptr(time.Date(2026, time.February, 6, 10, 0, 0, 0, time.UTC)),
		UpdatedAt:   Ptr(time.Date(2026, time.February, 6, 10, 0, 0, 0, time.UTC)),
		WebURL:      "https://gitlab.com/gitlab-com/gl-infra/platform/runway/team/-/work_items/41",
		Author: &BasicUser{
			ID:        5532616,
			Username:  "fforster",
			Name:      "Florian Forster",
			State:     "active",
			CreatedAt: Ptr(time.Date(2020, time.March, 2, 6, 29, 14, 0, time.UTC)),
			AvatarURL: "/uploads/-/system/user/avatar/5532616/avatar.png",
			WebURL:    "https://gitlab.com/fforster",
		},
		Assignees: nil,
	}, got)
}

func TestCreateWorkItem_MutationErrorReturnsError(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	// GIVEN a CreateWorkItem mutation that returns a mutation-level error
	setupGraphQLHandler(t, mux, func(w http.ResponseWriter, q GraphQLQuery) {
		assertQueryMatches(t, q.Query, "testdata/mutation_create_work_item.graphql")
		writeResponse(t, w, strings.NewReader(`
			{
			  "data": {
			    "workItemCreate": {
			      "workItem": null,
			      "errors": ["Title can't be blank"]
			    }
			  },
			  "correlationId": "9c88d56b0061dfef-IAD"
			}
		`))
	})

	// WHEN CreateWorkItem is called with an empty title
	got, _, err := client.WorkItems.CreateWorkItem(
		"gitlab-com/gl-infra/platform/runway/team",
		WorkItemTypeIssue,
		&CreateWorkItemOptions{Title: ""},
	)

	// THEN an error containing the mutation message is returned
	require.ErrorContains(t, err, "Title can't be blank")
	assert.Nil(t, got)
}

func TestCreateWorkItem_GraphQLErrorReturnsError(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	// GIVEN a CreateWorkItem request where the server returns a top-level GraphQL error
	setupGraphQLHandler(t, mux, func(w http.ResponseWriter, q GraphQLQuery) {
		assertQueryMatches(t, q.Query, "testdata/mutation_create_work_item.graphql")
		writeResponse(t, w, strings.NewReader(`
			{
			  "errors": [
			    {
			      "message": "Your GraphQL is bad and you should feel bad"
			    }
			  ]
			}
		`))
	})

	// WHEN CreateWorkItem is called
	got, _, err := client.WorkItems.CreateWorkItem(
		"gitlab-com/gl-infra/platform/runway/team",
		WorkItemTypeIssue,
		&CreateWorkItemOptions{Title: "New Issue"},
	)

	// THEN an error containing the GraphQL error message is returned
	require.ErrorContains(t, err, "Your GraphQL is bad and you should feel bad")
	assert.Nil(t, got)
}

func TestCreateWorkItem_NullWorkItemReturnsEmptyResponseError(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	// GIVEN a CreateWorkItem mutation that returns a null work item with no errors
	setupGraphQLHandler(t, mux, func(w http.ResponseWriter, q GraphQLQuery) {
		assertQueryMatches(t, q.Query, "testdata/mutation_create_work_item.graphql")
		writeResponse(t, w, strings.NewReader(`
			{
			  "data": {
			    "workItemCreate": {
			      "workItem": null
			    }
			  },
			  "correlationId": "9c88d56b0061dfef-IAD"
			}
		`))
	})

	// WHEN CreateWorkItem is called
	got, _, err := client.WorkItems.CreateWorkItem(
		"gitlab-com/gl-infra/platform/runway/team",
		WorkItemTypeIssue,
		&CreateWorkItemOptions{Title: "New Issue"},
	)

	// THEN ErrEmptyResponse is returned
	require.ErrorContains(t, err, ErrEmptyResponse.Error())
	assert.Nil(t, got)
}

// ---------------------------------------------------------------------------
// UpdateWorkItem tests
// ---------------------------------------------------------------------------

func TestUpdateWorkItem_SuccessfulUpdateWithTitleOnly(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	// GIVEN an UpdateWorkItem request that updates only the title
	setupGraphQLHandler(t, mux, func(w http.ResponseWriter, q GraphQLQuery) {
		// WHEN the query is for the work item ID lookup
		if strings.Contains(q.Query, "GetWorkItemID") {
			assertQueryMatches(t, q.Query, "testdata/query_get_work_item_id.graphql")
			io.WriteString(w, `{"data":{"namespace":{"workItem":{"id":"gid://gitlab/WorkItem/179785913"}}}}`)
			return
		}

		// WHEN the query is the update mutation
		assertQueryMatches(t, q.Query, "testdata/mutation_update_work_item.graphql")

		gotInput := q.Variables["input"].(map[string]any)
		assert.Equal(t, "test title", gotInput["title"], "input %q mismatch", "title")

		writeResponse(t, w, bytes.NewReader(updateWorkItemResponse))
	})

	// WHEN UpdateWorkItem is called with a title update
	got, _, err := client.WorkItems.UpdateWorkItem(
		"testing/unittest",
		117869168,
		&UpdateWorkItemOptions{Title: Ptr("test title")},
	)

	// THEN the updated work item is returned without error
	require.NoError(t, err)
	assert.Equal(t, &WorkItem{
		ID:          179785913,
		IID:         756,
		Type:        "Task",
		State:       "OPEN",
		Title:       "test title update",
		Description: "## Overview\n\nUpdate Runway Helm charts to generate Argo Rollout resources ...",
		Author: &BasicUser{
			ID:        5532616,
			Username:  "swainaina",
			Name:      "Silvester Wainaina",
			State:     "active",
			CreatedAt: Ptr(time.Date(2020, time.March, 2, 6, 29, 14, 0, time.UTC)),
			AvatarURL: "/uploads/-/system/user/avatar/5532616/avatar.png",
			WebURL:    "https://gitlab.com/swainaina",
		},
	}, got)
}

func TestUpdateWorkItem_SuccessfulUpdateWithAllOptions(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	wantInputs := map[string]any{
		"title":      "Updated Title",
		"stateEvent": "CLOSE",
		"descriptionWidget": map[string]any{
			"description": "Updated description",
		},
		"assigneesWidget": map[string]any{
			"assigneeIds": []any{"gid://gitlab/User/123", "gid://gitlab/User/456"},
		},
		"milestoneWidget": map[string]any{
			"milestoneId": "gid://gitlab/Milestone/789",
		},
		"crmContactsWidget": map[string]any{
			"contactIds":    []any{"gid://gitlab/CustomerRelations::Contact/1001", "gid://gitlab/CustomerRelations::Contact/1002"},
			"operationMode": "REPLACE",
		},
		"hierarchyWidget": map[string]any{
			"parentId": "gid://gitlab/WorkItem/100",
		},
		"labelsWidget": map[string]any{
			"addLabelIds":    []any{"gid://gitlab/Label/201", "gid://gitlab/Label/202"},
			"removeLabelIds": []any{"gid://gitlab/Label/203", "gid://gitlab/Label/204"},
		},
		"startAndDueDateWidget": map[string]any{
			"startDate": "2026-03-01",
			"dueDate":   "2026-04-01",
		},
		"weightWidget": map[string]any{
			"weight": float64(8),
		},
		"healthStatusWidget": map[string]any{
			"healthStatus": "needsAttention",
		},
		"iterationWidget": map[string]any{
			"iterationId": "gid://gitlab/Iteration/567",
		},
		"colorWidget": map[string]any{
			"color": "#00FF00",
		},
		"statusWidget": map[string]any{
			"status": "gid://gitlab/WorkItems::Statuses::SystemDefined::Status/2",
		},
	}

	// GIVEN an UpdateWorkItem request with all options set
	setupGraphQLHandler(t, mux, func(w http.ResponseWriter, q GraphQLQuery) {
		// WHEN the query is for the work item ID lookup
		if strings.Contains(q.Query, "GetWorkItemID") {
			assertQueryMatches(t, q.Query, "testdata/query_get_work_item_id.graphql")
			io.WriteString(w, `{"data":{"namespace":{"workItem":{"id":"gid://gitlab/WorkItem/179785913"}}}}`)
			return
		}

		// WHEN the query is the update mutation
		assertQueryMatches(t, q.Query, "testdata/mutation_update_work_item.graphql")

		gotInput := q.Variables["input"].(map[string]any)
		for k, v := range wantInputs {
			assert.Equal(t, v, gotInput[k], "input %q mismatch", k)
		}

		writeResponse(t, w, strings.NewReader(`
			{
			  "data": {
			    "workItemUpdate": {
			      "workItem": {
			        "id": "gid://gitlab/WorkItem/179785913",
			        "iid": "756",
			        "workItemType": {
			          "name": "Task"
			        },
			        "state": "CLOSED",
			        "title": "Updated Title",
			        "description": "Updated description",
			        "author": {
			          "id": "gid://gitlab/User/5532616",
			          "username": "swainaina",
			          "name": "Silvester Wainaina",
			          "state": "active",
			          "locked": false,
			          "createdAt": "2020-03-02T06:29:14Z",
			          "avatarUrl": "/uploads/-/system/user/avatar/5532616/avatar.png",
			          "webUrl": "https://gitlab.com/swainaina"
			        },
			        "createdAt": "2026-01-06T15:09:24Z",
			        "updatedAt": "2026-01-09T13:06:08Z",
			        "closedAt": "2026-01-09T13:06:08Z",
			        "webUrl": "https://gitlab.com/gitlab-com/gl-infra/platform/runway/team/-/work_items/756",
			        "features": {
			          "assignees": {
			            "assignees": {
			              "nodes": [
			                {
			                  "id": "gid://gitlab/User/123",
			                  "username": "user1",
			                  "name": "User One",
			                  "state": "active",
			                  "locked": false,
			                  "createdAt": "2020-01-01T00:00:00Z",
			                  "avatarUrl": "/avatar1.png",
			                  "webUrl": "https://gitlab.com/user1"
			                },
			                {
			                  "id": "gid://gitlab/User/456",
			                  "username": "user2",
			                  "name": "User Two",
			                  "state": "active",
			                  "locked": false,
			                  "createdAt": "2020-01-02T00:00:00Z",
			                  "avatarUrl": "/avatar2.png",
			                  "webUrl": "https://gitlab.com/user2"
			                }
			              ]
			            }
			          },
			          "status": {
			            "status": {
			              "name": "In Progress"
			            }
			          }
			        }
			      }
			    }
			  },
			  "correlationId": "9c88d56b0061dfef-IAD"
			}
		`))
	})

	opt := &UpdateWorkItemOptions{
		Title:          Ptr("Updated Title"),
		StateEvent:     Ptr(WorkItemStateEventClose),
		Description:    Ptr("Updated description"),
		AssigneeIDs:    []int64{123, 456},
		MilestoneID:    Ptr(int64(789)),
		CRMContactIDs:  []int64{1001, 1002},
		ParentID:       Ptr(int64(100)),
		AddLabelIDs:    []int64{201, 202},
		RemoveLabelIDs: []int64{203, 204},
		StartDate:      Ptr(ISOTime(time.Date(2026, time.March, 1, 0, 0, 0, 0, time.UTC))),
		DueDate:        Ptr(ISOTime(time.Date(2026, time.April, 1, 0, 0, 0, 0, time.UTC))),
		Weight:         Ptr(int64(8)),
		HealthStatus:   Ptr("needsAttention"),
		IterationID:    Ptr(int64(567)),
		Color:          Ptr("#00FF00"),
		Status:         Ptr(WorkItemStatusInProgress),
	}

	// WHEN UpdateWorkItem is called with all options
	got, _, err := client.WorkItems.UpdateWorkItem(
		"gitlab-com/gl-infra/platform/runway/team",
		756,
		opt,
	)

	// THEN the updated work item is returned without error
	require.NoError(t, err)
	assert.Equal(t, &WorkItem{
		ID:          179785913,
		IID:         756,
		Type:        "Task",
		State:       "CLOSED",
		Status:      Ptr("In Progress"),
		Title:       "Updated Title",
		Description: "Updated description",
		CreatedAt:   Ptr(time.Date(2026, time.January, 6, 15, 9, 24, 0, time.UTC)),
		UpdatedAt:   Ptr(time.Date(2026, time.January, 9, 13, 6, 8, 0, time.UTC)),
		ClosedAt:    Ptr(time.Date(2026, time.January, 9, 13, 6, 8, 0, time.UTC)),
		WebURL:      "https://gitlab.com/gitlab-com/gl-infra/platform/runway/team/-/work_items/756",
		Author: &BasicUser{
			ID:        5532616,
			Username:  "swainaina",
			Name:      "Silvester Wainaina",
			State:     "active",
			CreatedAt: Ptr(time.Date(2020, time.March, 2, 6, 29, 14, 0, time.UTC)),
			AvatarURL: "/uploads/-/system/user/avatar/5532616/avatar.png",
			WebURL:    "https://gitlab.com/swainaina",
		},
		Assignees: []*BasicUser{
			{
				ID:        123,
				Username:  "user1",
				Name:      "User One",
				State:     "active",
				CreatedAt: Ptr(time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)),
				AvatarURL: "/avatar1.png",
				WebURL:    "https://gitlab.com/user1",
			},
			{
				ID:        456,
				Username:  "user2",
				Name:      "User Two",
				State:     "active",
				CreatedAt: Ptr(time.Date(2020, time.January, 2, 0, 0, 0, 0, time.UTC)),
				AvatarURL: "/avatar2.png",
				WebURL:    "https://gitlab.com/user2",
			},
		},
	}, got)
}

func TestUpdateWorkItem_WorkItemNotFoundDuringIDLookup(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	// GIVEN an UpdateWorkItem request where the ID lookup returns null
	setupGraphQLHandler(t, mux, func(w http.ResponseWriter, q GraphQLQuery) {
		io.WriteString(w, `{"data":{"namespace":{"workItem":null}}}`)
	})

	// WHEN UpdateWorkItem is called
	got, _, err := client.WorkItems.UpdateWorkItem(
		"testing/unittest",
		117869168,
		&UpdateWorkItemOptions{},
	)

	// THEN ErrEmptyResponse is returned
	require.ErrorContains(t, err, ErrEmptyResponse.Error())
	assert.Nil(t, got)
}

func TestUpdateWorkItem_MutationReturnsNullWorkItem(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	// GIVEN an UpdateWorkItem request where the mutation returns a null work item
	setupGraphQLHandler(t, mux, func(w http.ResponseWriter, q GraphQLQuery) {
		// WHEN the query is for the work item ID lookup
		if strings.Contains(q.Query, "GetWorkItemID") {
			io.WriteString(w, `{"data":{"namespace":{"workItem":{"id":"gid://gitlab/WorkItem/179785913"}}}}`)
			return
		}

		// WHEN the query is the update mutation
		writeResponse(t, w, strings.NewReader(`
			{
			  "data": {
			    "workItemUpdate": {
			      "workItem": null,
			      "errors": ["404 Not Found"]
			    }
			  }
			}
		`))
	})

	// WHEN UpdateWorkItem is called
	got, _, err := client.WorkItems.UpdateWorkItem(
		"testing/unittest",
		117869168,
		&UpdateWorkItemOptions{},
	)

	// THEN ErrNotFound is returned
	require.ErrorContains(t, err, ErrNotFound.Error())
	assert.Nil(t, got)
}

func TestUpdateWorkItem_MutationReturnsError(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	// GIVEN an UpdateWorkItem request where the mutation returns a permission error
	setupGraphQLHandler(t, mux, func(w http.ResponseWriter, q GraphQLQuery) {
		// WHEN the query is for the work item ID lookup
		if strings.Contains(q.Query, "GetWorkItemID") {
			io.WriteString(w, `{"data":{"namespace":{"workItem":{"id":"gid://gitlab/WorkItem/179785913"}}}}`)
			return
		}

		// WHEN the query is the update mutation
		writeResponse(t, w, strings.NewReader(`
			{
			  "data": {
			    "workItemUpdate": {
			      "workItem": null,
			      "errors": ["User doesn't have permission to update this work item"]
			    }
			  }
			}
		`))
	})

	// WHEN UpdateWorkItem is called
	got, _, err := client.WorkItems.UpdateWorkItem(
		"testing/unittest",
		117869168,
		&UpdateWorkItemOptions{},
	)

	// THEN an error containing the permission message is returned
	require.ErrorContains(t, err, "User doesn't have permission to update this work item")
	assert.Nil(t, got)
}

// ---------------------------------------------------------------------------
// DeleteWorkItem tests
// ---------------------------------------------------------------------------

func TestDeleteWorkItem_SuccessfullyDeletesWorkItem(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	// GIVEN a DeleteWorkItem request for an existing work item
	setupGraphQLHandler(t, mux, func(w http.ResponseWriter, q GraphQLQuery) {
		switch {
		// WHEN the query is for the work item ID lookup
		case strings.Contains(q.Query, "GetWorkItemID"):
			assertQueryMatches(t, q.Query, "testdata/query_get_work_item_id.graphql")
			writeResponse(t, w, strings.NewReader(`{
				"data": {
					"namespace": {
						"workItem": {
							"id": "gid://gitlab/WorkItem/183771442"
						}
					}
				}
			}`))
		// WHEN the query is the delete mutation
		case strings.Contains(q.Query, "DeleteWorkItem"):
			assertQueryMatches(t, q.Query, "testdata/mutation_delete_work_item.graphql")
			writeResponse(t, w, strings.NewReader(`{
				"data": {
					"workItemDelete": {
						"clientMutationId": null,
						"namespace": {
							"id": "gid://gitlab/Namespaces::ProjectNamespace/124736349"
						}
					}
				}
			}`))
		default:
			assert.Failf(t, "unexpected query: %s", q.Query)
			http.Error(w, "unexpected query", http.StatusBadRequest)
		}
	})

	// WHEN DeleteWorkItem is called
	_, err := client.WorkItems.DeleteWorkItem("test-gitlab-org/gitlab", 123)

	// THEN no error is returned
	require.NoError(t, err)
}

func TestDeleteWorkItem_WorkItemNotFoundReturnsError(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	// GIVEN a DeleteWorkItem request where the work item does not exist
	setupGraphQLHandler(t, mux, func(w http.ResponseWriter, q GraphQLQuery) {
		// WHEN the query is for the work item ID lookup
		if strings.Contains(q.Query, "GetWorkItemID") {
			writeResponse(t, w, strings.NewReader(`{
				 "data": {
					 "namespace": {
						 "workItem": null
					 }
				 }
			 }`))
			return
		}

		assert.Failf(t, "unexpected query: %s", q.Query)
		http.Error(w, "unexpected query", http.StatusBadRequest)
	})

	// WHEN DeleteWorkItem is called for a non-existent IID
	_, err := client.WorkItems.DeleteWorkItem("test-gitlab-org/gitlab", 999)

	// THEN ErrNotFound is returned
	require.ErrorContains(t, err, ErrNotFound.Error())
}

func TestDeleteWorkItem_NamespaceNotFoundReturnsError(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	// GIVEN a DeleteWorkItem request where the namespace does not exist
	setupGraphQLHandler(t, mux, func(w http.ResponseWriter, q GraphQLQuery) {
		// WHEN the query is for the work item ID lookup
		if strings.Contains(q.Query, "GetWorkItemID") {
			writeResponse(t, w, strings.NewReader(`{
				 "data": {
					 "namespace": null
				 }
			 }`))
			return
		}

		assert.Failf(t, "unexpected query: %s", q.Query)
		http.Error(w, "unexpected query", http.StatusBadRequest)
	})

	// WHEN DeleteWorkItem is called for a non-existent namespace
	_, err := client.WorkItems.DeleteWorkItem("does/not/exist", 123)

	// THEN ErrNotFound is returned
	require.ErrorContains(t, err, ErrNotFound.Error())
}

func TestDeleteWorkItem_MutationReturnsErrors(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	// GIVEN a DeleteWorkItem request where the mutation returns an error
	setupGraphQLHandler(t, mux, func(w http.ResponseWriter, q GraphQLQuery) {
		switch {
		// WHEN the query is for the work item ID lookup
		case strings.Contains(q.Query, "GetWorkItemID"):
			writeResponse(t, w, strings.NewReader(`{
				 "data": {
					 "namespace": {
						 "workItem": {
							 "id": "gid://gitlab/WorkItem/123"
						 }
					 }
				 }
			 }`))
		// WHEN the query is the delete mutation
		case strings.Contains(q.Query, "DeleteWorkItem"):
			writeResponse(t, w, strings.NewReader(`{
				 "data": {
					 "workItemDelete": {
						 "clientMutationId": null,
						 "namespace": null,
						 "errors": ["Work item cannot be deleted"]
					 }
				 }
			 }`))
		default:
			assert.Failf(t, "unexpected query: %s", q.Query)
			http.Error(w, "unexpected query", http.StatusBadRequest)
		}
	})

	// WHEN DeleteWorkItem is called
	_, err := client.WorkItems.DeleteWorkItem("test-gitlab-org/gitlab", 123)

	// THEN an error containing the mutation message is returned
	require.ErrorContains(t, err, "Work item cannot be deleted")
}

// ---------------------------------------------------------------------------
// ListWorkItemTypes tests
// ---------------------------------------------------------------------------

func TestListWorkItemTypes_SystemAndCustomTypes(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	// GIVEN a ListWorkItemTypes request that returns system and custom types
	setupWorkItemTypesHandler(t, mux, strings.NewReader(`{
		"data": {
			"namespace": {
				"workItemTypes": {
					"nodes": [
						{"id": "gid://gitlab/WorkItems::Type/1",  "name": "Issue",      "enabled": true},
						{"id": "gid://gitlab/WorkItems::Type/5",  "name": "Task",       "enabled": true},
						{"id": "gid://gitlab/WorkItems::Type/99", "name": "CustomType", "enabled": true}
					],
					"pageInfo": {
						"endCursor": "cursor123", "hasNextPage": false,
						"startCursor": "cursor000", "hasPreviousPage": false
					}
				}
			}
		}
	}`))

	// WHEN ListWorkItemTypes is called
	got, resp, err := client.WorkItems.ListWorkItemTypes("gitlab-org/gitlab", &ListWorkItemTypesOptions{})

	// THEN all three types are returned with correct page info
	require.NoError(t, err)
	assert.Equal(t, []WorkItemType{
		{ID: WorkItemTypeIssue, Name: "Issue", Enabled: true},
		{ID: WorkItemTypeTask, Name: "Task", Enabled: true},
		{ID: "gid://gitlab/WorkItems::Type/99", Name: "CustomType", Enabled: true},
	}, got)
	assert.Equal(t, &PageInfo{
		EndCursor:       "cursor123",
		HasNextPage:     false,
		StartCursor:     "cursor000",
		HasPreviousPage: false,
	}, resp.PageInfo)
}

func TestListWorkItemTypes_FilterByName(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	// GIVEN a ListWorkItemTypes request filtered by name
	setupWorkItemTypesHandler(t, mux, strings.NewReader(`{
		"data": {
			"namespace": {
				"workItemTypes": {
					"nodes": [
						{"id": "gid://gitlab/WorkItems::Type/1", "name": "Issue", "enabled": true}
					],
					"pageInfo": {
						"endCursor": "", "hasNextPage": false,
						"startCursor": "", "hasPreviousPage": false
					}
				}
			}
		}
	}`))

	// WHEN ListWorkItemTypes is called with a name filter
	got, _, err := client.WorkItems.ListWorkItemTypes(
		"gitlab-org/gitlab",
		&ListWorkItemTypesOptions{Name: Ptr("ISSUE")},
	)

	// THEN only the matching type is returned
	require.NoError(t, err)
	assert.Equal(t, []WorkItemType{
		{ID: WorkItemTypeIssue, Name: "Issue", Enabled: true},
	}, got)
}

func TestListWorkItemTypes_EmptyResponse(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	// GIVEN a ListWorkItemTypes request that returns no types
	setupWorkItemTypesHandler(t, mux, strings.NewReader(`{
		"data": {
			"namespace": {
				"workItemTypes": {
					"nodes": [],
					"pageInfo": {
						"endCursor": "", "hasNextPage": false,
						"startCursor": "", "hasPreviousPage": false
					}
				}
			}
		}
	}`))

	// WHEN ListWorkItemTypes is called
	got, _, err := client.WorkItems.ListWorkItemTypes("gitlab-org/gitlab", &ListWorkItemTypesOptions{})

	// THEN an empty (non-nil) slice is returned without error
	require.NoError(t, err)
	assert.Equal(t, []WorkItemType{}, got)
}

func TestListWorkItemTypes_NilOptDoesNotPanic(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	// GIVEN a ListWorkItemTypes request with nil options
	setupWorkItemTypesHandler(t, mux, strings.NewReader(`{
		"data": {
			"namespace": {
				"workItemTypes": {
					"nodes": [],
					"pageInfo": {
						"endCursor": "", "hasNextPage": false,
						"startCursor": "", "hasPreviousPage": false
					}
				}
			}
		}
	}`))

	// WHEN ListWorkItemTypes is called with nil options
	got, _, err := client.WorkItems.ListWorkItemTypes("gitlab-org/gitlab", nil)

	// THEN no panic occurs and an empty slice is returned without error
	require.NoError(t, err)
	assert.Equal(t, []WorkItemType{}, got)
}

// ---------------------------------------------------------------------------
// buildListWorkItemsQuery tests
// ---------------------------------------------------------------------------

func TestBuildListWorkItemsQuery_AllFields(t *testing.T) {
	t.Parallel()

	now := time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC)

	// GIVEN a ListWorkItemsOptions with every field populated
	opt := &ListWorkItemsOptions{
		AssigneeUsernames:    []string{"user1"},
		AssigneeWildcardID:   Ptr("NONE"),
		AuthorUsername:       Ptr("fforster"),
		Confidential:         Ptr(true),
		CRMContactID:         Ptr("contact123"),
		CRMOrganizationID:    Ptr("org456"),
		HealthStatusFilter:   Ptr("onTrack"),
		IDs:                  []string{"gid://gitlab/WorkItem/1"},
		IIDs:                 []string{"1"},
		IncludeAncestors:     Ptr(true),
		IncludeDescendants:   Ptr(false),
		IterationCadenceID:   []string{"cadence1"},
		IterationID:          []string{"iter1"},
		IterationWildcardID:  Ptr("CURRENT"),
		LabelName:            []string{"bug"},
		MilestoneTitle:       []string{"v1.0"},
		MilestoneWildcardID:  Ptr("STARTED"),
		MyReactionEmoji:      Ptr("thumbsup"),
		ParentIDs:            []string{"gid://gitlab/WorkItem/100"},
		ReleaseTag:           []string{"v1.0.0"},
		ReleaseTagWildcardID: Ptr("ANY"),
		State:                Ptr("opened"),
		Subscribed:           Ptr("EXPLICITLY_SUBSCRIBED"),
		Types:                []string{"ISSUE"},
		Weight:               Ptr("5"),
		WeightWildcardID:     Ptr("NONE"),
		ClosedAfter:          &now,
		ClosedBefore:         &now,
		CreatedAfter:         &now,
		CreatedBefore:        &now,
		DueAfter:             &now,
		DueBefore:            &now,
		UpdatedAfter:         &now,
		UpdatedBefore:        &now,
		Sort:                 Ptr("CREATED_DESC"),
		Search:               Ptr("bug"),
		In:                   []string{"TITLE"},
		After:                Ptr("cursor123"),
		Before:               Ptr("cursor456"),
		First:                Ptr(int64(10)),
		Last:                 Ptr(int64(5)),
	}

	// WHEN buildListWorkItemsQuery is called
	query, vars, err := buildListWorkItemsQuery("my/project", opt)

	// THEN every field appears in both the query string and the vars map
	require.NoError(t, err)

	for _, field := range []string{
		"assigneeUsernames", "assigneeWildcardId", "authorUsername",
		"confidential", "crmContactId", "crmOrganizationId",
		"healthStatusFilter", "ids", "iids", "includeAncestors",
		"includeDescendants", "iterationCadenceId", "iterationId",
		"iterationWildcardId", "labelName", "milestoneTitle",
		"milestoneWildcardId", "myReactionEmoji", "parentIds",
		"releaseTag", "releaseTagWildcardId", "state", "subscribed",
		"types", "weight", "weightWildcardId", "closedAfter",
		"closedBefore", "createdAfter", "createdBefore", "dueAfter",
		"dueBefore", "updatedAfter", "updatedBefore", "sort",
		"search", "in", "after", "before", "first", "last",
	} {
		assert.Contains(t, query, "$"+field+":", "field %q missing from declaration in query string", field)
		assert.Contains(t, query, field+": $"+field, "field %q missing from args in query string", field)
		assert.Contains(t, vars, field, "field %q missing from vars map", field)
	}

	assert.Equal(t, "my/project", vars["fullPath"])
}

func TestBuildListWorkItemsQuery_NilOpt(t *testing.T) {
	t.Parallel()

	// GIVEN a nil ListWorkItemsOptions
	// WHEN buildListWorkItemsQuery is called
	query, vars, err := buildListWorkItemsQuery("my/project", nil)

	// THEN the query declares only $fullPath and the vars map contains only fullPath
	require.NoError(t, err)
	assert.Contains(t, query, "$fullPath: ID!")
	assert.NotContains(t, query, "$assigneeUsernames")
	assert.NotContains(t, query, "$state")
	assert.NotContains(t, query, "$first")
	assert.Equal(t, map[string]any{"fullPath": "my/project"}, vars)
}

func TestBuildListWorkItemsQuery_PartialFields(t *testing.T) {
	t.Parallel()

	// GIVEN a ListWorkItemsOptions with only state and first set
	opt := &ListWorkItemsOptions{
		State: Ptr("opened"),
		First: Ptr(int64(10)),
	}

	// WHEN buildListWorkItemsQuery is called
	query, vars, err := buildListWorkItemsQuery("my/project", opt)

	// THEN only the set fields appear in the query and vars map
	require.NoError(t, err)
	assert.Contains(t, query, "$state:")
	assert.Contains(t, query, "$first:")
	assert.NotContains(t, query, "$assigneeUsernames")
	assert.NotContains(t, query, "$authorUsername")
	assert.Contains(t, vars, "state")
	assert.Contains(t, vars, "first")
	assert.NotContains(t, vars, "assigneeUsernames")
}
