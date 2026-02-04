package gitlab

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed testdata/get_workitem.json
var getWorkItemResponse []byte

//go:embed testdata/get_workitem_by_id.json
var getWorkItemByIDResponse []byte

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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mux, client := setup(t)

			mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodPost)

				io.Copy(io.Discard, r.Body)
				r.Body.Close()

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

func TestGetWorkItemByID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		gid      int64
		response io.WriterTo
		want     *WorkItem
		wantErr  error // Must be a concrete error type!
	}{
		{
			name:     "successful response",
			gid:      int64(179785913),
			response: bytes.NewReader(getWorkItemByIDResponse),
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
			name: "using a work item ID that doesn't exist returns an error",
			gid:  9999,
			response: strings.NewReader(`
{
  "errors": [
    {
      "message": "The resource that you are attempting to access does not exist or you don't have permission to perform this action",
      "locations": [
        {
          "line": 3,
          "column": 2
        }
      ],
      "path": [
        "workItem"
      ]
    }
  ],
  "data": {
    "workItem": null
  }
}
`),
			want:    nil,
			wantErr: &GraphQLResponseError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mux, client := setup(t)

			mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodPost)

				w.Header().Set("Content-Type", "application/json")
				tt.response.WriteTo(w)
			})

			got, _, err := client.WorkItems.GetWorkItemByID(tt.gid)

			if tt.wantErr != nil {
				require.ErrorAs(t, err, &tt.wantErr) //nolint:testifylint
				assert.Nil(t, got)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestListWorkItems(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		fullPath        string
		opt             *ListWorkItemsOptions
		response        io.WriterTo
		wantQuerySubstr []string
		want            []*WorkItem
		wantErr         error
	}{
		{
			name:     "successful query with authorUsername",
			fullPath: "gitlab-com/gl-infra/platform/runway/team",
			opt: &ListWorkItemsOptions{
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
				          },
				          {
				            "id": "gid://gitlab/WorkItem/181297779",
				            "iid": "38",
				            "title": "Phase 5: Dedicated Integration"
				          }
				        ]
				      }
				    }
				  },
				  "correlationId": "9c88d56b0061dfef-IAD"
				}
			`),
			wantQuerySubstr: []string{
				`query ListWorkItems($fullPath: ID!, $authorUsername: String)`,
				`workItems(authorUsername: $authorUsername) {`,
			},
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
			wantQuerySubstr: []string{
				`query ListWorkItems($fullPath: ID!, $authorUsername: String, $state: IssuableState)`,
				`workItems(authorUsername: $authorUsername, state: $state) {`,
			},
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
			wantQuerySubstr: []string{
				`query ListWorkItems($fullPath: ID!, $authorUsername: String, $state: IssuableState)`,
				`workItems(authorUsername: $authorUsername, state: $state) {`,
			},
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
				Subscribed:           Ptr("SUBSCRIBED"),
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
			wantQuerySubstr: []string{
				// Main filters
				`$assigneeUsernames: [String!]`, `assigneeUsernames: $assigneeUsernames`,
				`$assigneeWildcardId: AssigneeWildcardId`, `assigneeWildcardId: $assigneeWildcardId`,
				`$authorUsername: String`, `authorUsername: $authorUsername`,
				`$confidential: Boolean`, `confidential: $confidential`,
				`$crmContactId: String`, `crmContactId: $crmContactId`,
				`$crmOrganizationId: String`, `crmOrganizationId: $crmOrganizationId`,
				`$healthStatusFilter: HealthStatusFilter`, `healthStatusFilter: $healthStatusFilter`,
				`$ids: [WorkItemID!]`, `ids: $ids`,
				`$iids: [String!]`, `iids: $iids`,
				`$includeAncestors: Boolean`, `includeAncestors: $includeAncestors`,
				`$includeDescendants: Boolean`, `includeDescendants: $includeDescendants`,
				`$iterationCadenceId: [IterationsCadenceID!]`, `iterationCadenceId: $iterationCadenceId`,
				`$iterationId: [ID]`, `iterationId: $iterationId`,
				`$iterationWildcardId: IterationWildcardId`, `iterationWildcardId: $iterationWildcardId`,
				`$labelName: [String!]`, `labelName: $labelName`,
				`$milestoneTitle: [String!]`, `milestoneTitle: $milestoneTitle`,
				`$milestoneWildcardId: MilestoneWildcardId`, `milestoneWildcardId: $milestoneWildcardId`,
				`$myReactionEmoji: String`, `myReactionEmoji: $myReactionEmoji`,
				`$parentIds: [WorkItemID!]`, `parentIds: $parentIds`,
				`$releaseTag: [String!]`, `releaseTag: $releaseTag`,
				`$releaseTagWildcardId: ReleaseTagWildcardId`, `releaseTagWildcardId: $releaseTagWildcardId`,
				`$state: IssuableState`, `state: $state`,
				`$subscribed: SubscriptionStatus`, `subscribed: $subscribed`,
				`$types: [IssueType!]`, `types: $types`,
				`$weight: String`, `weight: $weight`,
				`$weightWildcardId: WeightWildcardId`, `weightWildcardId: $weightWildcardId`,
				// Time filters
				`$closedAfter: Time`, `closedAfter: $closedAfter`,
				`$closedBefore: Time`, `closedBefore: $closedBefore`,
				`$createdAfter: Time`, `createdAfter: $createdAfter`,
				`$createdBefore: Time`, `createdBefore: $createdBefore`,
				`$dueAfter: Time`, `dueAfter: $dueAfter`,
				`$dueBefore: Time`, `dueBefore: $dueBefore`,
				`$updatedAfter: Time`, `updatedAfter: $updatedAfter`,
				`$updatedBefore: Time`, `updatedBefore: $updatedBefore`,
				// Sorting
				`$sort: WorkItemSort`, `sort: $sort`,
				// Search
				`$search: String`, `search: $search`,
				`$in: [IssuableSearchableField!]`, `in: $in`,
				// Pagination
				`$after: String`, `after: $after`,
				`$before: String`, `before: $before`,
				`$first: Int`, `first: $first`,
				`$last: Int`, `last: $last`,
			},
			want: nil,
		},
	}

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

				for _, ss := range tt.wantQuerySubstr {
					if !strings.Contains(q.Query, ss) {
						http.Error(w, fmt.Sprintf("want substring %q, got query %q", ss, q.Query), http.StatusBadRequest)
						return
					}
				}

				w.Header().Set("Content-Type", "application/json")
				tt.response.WriteTo(w)
			})

			got, _, err := client.WorkItems.ListWorkItems(tt.fullPath, tt.opt)

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
