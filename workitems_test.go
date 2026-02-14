package gitlab

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed testdata/get_workitem.json
var getWorkItemResponse []byte

func TestGetWorkItem(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		fullPath string
		iid      int64
		response io.WriterTo
		want     *WorkItem
		wantErr  error
	}{
		{
			name:     "successful response with work item",
			fullPath: "gitlab-com/gl-infra/platform/runway/team",
			iid:      756,
			response: bytes.NewReader(getWorkItemResponse),
			want: &WorkItem{
				ID:          179785913,
				IID:         756,
				Type:        "Task",
				State:       "OPEN",
				Status:      "New",
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
					Locked:    false,
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
						Locked:    false,
						CreatedAt: Ptr(time.Date(2020, time.March, 2, 6, 29, 14, 0, time.UTC)),
						AvatarURL: "/uploads/-/system/user/avatar/5532616/avatar.png",
						WebURL:    "https://gitlab.com/swainaina",
					},
				},
			},
		},
		{
			name:     "successful response with zero work items returns not found error",
			fullPath: "gitlab-com/gl-infra/platform/runway/team",
			iid:      999,
			response: strings.NewReader(`
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
			`),
			want:    nil,
			wantErr: ErrNotFound,
		},
		{
			name:     "successful response without project returns not found error",
			fullPath: "does/not/exist",
			iid:      1,
			response: strings.NewReader(`
				{
				"data": {
					"project": null
				},
				"correlationId": "9c59850aa3cdf515-IAD"
				}
			`),
			want:    nil,
			wantErr: ErrNotFound,
		},
	}

	schema := loadSchema(t)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mux, client := setup(t)

			mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
				defer r.Body.Close()

				testMethod(t, r, http.MethodPost)

				var q GraphQLQuery

				if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}

				if err := validateSchema(schema, q); err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}

				w.Header().Set("Content-Type", "application/json")
				tt.response.WriteTo(w)
			})

			got, _, err := client.WorkItems.GetWorkItem(tt.fullPath, tt.iid)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, got)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

//go:embed testdata/list_work_items.json
var listWorkItemResponse []byte

func TestListWorkItems(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		fullPath     string
		opt          *ListWorkItemsOptions
		response     io.WriterTo
		want         []*WorkItem
		wantPageInfo *PageInfo
		wantErr      error
	}{
		{
			name:     "successful query with authorUsername",
			fullPath: "gitlab-com/gl-infra/platform/runway/team",
			opt: &ListWorkItemsOptions{
				AuthorUsername: Ptr("fforster"),
			},
			response: bytes.NewReader(listWorkItemResponse),
			want: []*WorkItem{
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
			},
			wantPageInfo: &PageInfo{
				EndCursor:       "eyJjcmVhdGVkX2F0IjoiMjAyNi0wMS0xNiAxMzozMjo0Ny44NTEyMTUwMDAgKzAwMDAiLCJpZCI6IjE4MTI5Nzc3OSJ9",
				HasNextPage:     true,
				StartCursor:     "eyJjcmVhdGVkX2F0IjoiMjAyNi0wMS0xNiAxMzozMjo1Ny43NjgxNzYwMDAgKzAwMDAiLCJpZCI6IjE4MTI5Nzc4NiJ9",
				HasPreviousPage: false,
			},
		},
		{
			name:     "successful response with work item",
			fullPath: "gitlab-com/gl-infra/platform/runway/team",
			opt: &ListWorkItemsOptions{
				State:          Ptr("opened"),
				AuthorUsername: Ptr("fforster"),
			},
			response: strings.NewReader(`
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
			`),
			want: []*WorkItem{
				{
					ID:    181297786,
					IID:   39,
					Title: "Phase 6: Rollout to Additional Services",
				},
			},
		},
		{
			name:     "empty response is not an error",
			fullPath: "gitlab-com/gl-infra/platform/runway/team",
			opt: &ListWorkItemsOptions{
				State:          Ptr("opened"),
				AuthorUsername: Ptr("fforster"),
			},
			response: strings.NewReader(`
				{
				  "data": {
				    "namespace": {
				      "workItems": {
				        "nodes": []
				      }
				    }
				  }
				}
			`),
			want: nil,
		},
		{
			name:     "all ListWorkItemsOptions fields are included in query",
			fullPath: "gitlab-com/gl-infra/platform/runway/team",
			opt: &ListWorkItemsOptions{
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
			},
			response: strings.NewReader(`
				{
				  "data": {
				    "namespace": {
				      "workItems": {
				        "nodes": []
				      }
				    }
				  }
				}
			`),
			want: nil,
		},
	}

	schema := loadSchema(t)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mux, client := setup(t)

			mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
				defer r.Body.Close()

				testMethod(t, r, http.MethodPost)

				var q GraphQLQuery

				if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}

				if err := validateSchema(schema, q); err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}

				w.Header().Set("Content-Type", "application/json")
				tt.response.WriteTo(w)
			})

			got, resp, err := client.WorkItems.ListWorkItems(tt.fullPath, tt.opt)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, got)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.want, got)

			if tt.wantPageInfo != nil {
				assert.Equal(t, tt.wantPageInfo, resp.PageInfo)
			}
		})
	}
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

	schema := loadSchema(t)

	mux, client := setup(t)

	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		testMethod(t, r, http.MethodPost)

		var q GraphQLQuery

		if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := validateSchema(schema, q); err != nil {
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

	got, err := ScanAndCollect(func(p PaginationOptionFunc) ([]*WorkItem, *Response, error) {
		return client.WorkItems.ListWorkItems("unit/test", &opt, p)
	})

	require.NoError(t, err)

	want := []*WorkItem{
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
	}

	assert.Equal(t, want, got)
}

func loadSchema(t *testing.T) *graphql.Schema {
	t.Helper()

	const filename = "schema/gitlab.graphql"

	fh, err := os.Open(filename)
	switch {
	case errors.Is(err, os.ErrNotExist):
		t.Logf("GraphQL schema file %q is not available", filename)
		return nil

	case err != nil:
		t.Fatalf("opening schema failed: %v", err)
	}

	data, err := io.ReadAll(fh)
	if err != nil {
		t.Fatalf("reading schema failed: %v", err)
	}

	schema, err := graphql.ParseSchema(string(data), nil)
	if err != nil {
		t.Fatalf("parsing schema failed: %v", err)
	}

	return schema
}

func validateSchema(schema *graphql.Schema, query GraphQLQuery) error {
	if schema == nil {
		return nil
	}

	queryErrors := schema.ValidateWithVariables(query.Query, query.Variables)

	var errs error
	for _, err := range queryErrors {
		errs = errors.Join(errs, err)
	}

	return errs
}
