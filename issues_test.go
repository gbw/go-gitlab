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
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetIssue(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/issues/5", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":1, "description": "This is test project", "author" : {"id" : 1, "name": "snehal"}, "assignees":[{"id":1}],"merge_requests_count": 1}`)
	})

	issue, _, err := client.Issues.GetIssue("1", 5)
	require.NoError(t, err)

	want := &Issue{
		ID:                1,
		Description:       "This is test project",
		Author:            &IssueAuthor{ID: 1, Name: "snehal"},
		Assignees:         []*IssueAssignee{{ID: 1}},
		MergeRequestCount: 1,
	}

	assert.Equal(t, want, issue)
}

func TestGetIssueByID(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/issues/5", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":5, "description": "This is test project", "author" : {"id" : 1, "name": "snehal"}, "assignees":[{"id":1}],"merge_requests_count": 1}`)
	})

	issue, _, err := client.Issues.GetIssueByID(5)
	require.NoError(t, err)

	want := &Issue{
		ID:                5,
		Description:       "This is test project",
		Author:            &IssueAuthor{ID: 1, Name: "snehal"},
		Assignees:         []*IssueAssignee{{ID: 1}},
		MergeRequestCount: 1,
	}

	assert.Equal(t, want, issue)
}

func TestDeleteIssue(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/issues/5", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		fmt.Fprint(w, `{"id":1, "description": "This is test project", "author" : {"id" : 1, "name": "snehal"}, "assignees":[{"id":1}]}`)
	})

	_, err := client.Issues.DeleteIssue("1", 5)
	require.NoError(t, err)
}

func TestReorderIssue(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/issues/5/reorder", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprint(w, `{"id":1, "title" : "Reordered issue", "description": "This is the description of a reordered issue", "author" : {"id" : 1, "name": "corrie"}, "assignees":[{"id":1}]}`)
	})

	afterID := int64(100)
	opt := ReorderIssueOptions{MoveAfterID: &afterID}

	issue, _, err := client.Issues.ReorderIssue("1", 5, &opt)
	require.NoError(t, err)

	want := &Issue{
		ID:          1,
		Title:       "Reordered issue",
		Description: "This is the description of a reordered issue",
		Author:      &IssueAuthor{ID: 1, Name: "corrie"},
		Assignees:   []*IssueAssignee{{ID: 1}},
	}

	assert.Equal(t, want, issue)
}

func TestMoveIssue(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/issues/11/move", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		mustWriteHTTPResponse(t, w, "testdata/issue_move.json")
	})

	issue, _, err := client.Issues.MoveIssue("1", 11, &MoveIssueOptions{ToProjectID: Ptr(int64(5))})
	require.NoError(t, err)

	want := &Issue{
		ID:        92,
		IID:       11,
		ProjectID: 5,
		MovedToID: 0,
	}

	assert.Equal(t, want.IID, issue.IID)
	assert.Equal(t, want.ProjectID, issue.ProjectID)

	mux.HandleFunc("/api/v4/projects/1/issues/11", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `
		{
				"id": 1,
				"iid": 11,
				"project_id": 1,
				"moved_to_id": 92
		}`,
		)
	})
	movedIssue, _, err := client.Issues.GetIssue("1", 11)
	require.NoError(t, err)

	wantedMovedIssue := &Issue{
		ID:        1,
		IID:       11,
		ProjectID: 1,
		MovedToID: 92,
	}

	assert.Equal(t, wantedMovedIssue, movedIssue)
}

func TestListIssues(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/issues", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, "/api/v4/issues?assignee_id=2&author_id=1")
		fmt.Fprint(w, `
			[
				{
					"id": 1,
					"description": "This is test project",
					"author": {
						"id": 1,
						"name": "snehal"
					},
					"assignees": [
						{
							"id": 1
						}
					],
					"labels": [
						"foo",
						"bar"
					]
			  }
			]`,
		)
	})

	listProjectIssue := &ListIssuesOptions{
		AuthorID:   Ptr(int64(0o1)),
		AssigneeID: AssigneeID(0o2),
	}

	issues, _, err := client.Issues.ListIssues(listProjectIssue)
	require.NoError(t, err)

	want := []*Issue{{
		ID:          1,
		Description: "This is test project",
		Author:      &IssueAuthor{ID: 1, Name: "snehal"},
		Assignees:   []*IssueAssignee{{ID: 1}},
		Labels:      []string{"foo", "bar"},
	}}

	assert.Equal(t, want, issues)
}

func TestListIssuesWithLabelDetails(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/issues", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, "/api/v4/issues?assignee_id=2&author_id=1")
		fmt.Fprint(w, `
			[
				{
					"id": 1,
					"description": "This is test project",
					"author": {
						"id": 1,
						"name": "snehal"
					},
					"assignees": [
						{
							"id": 1
						}
					],
					"labels": [
						{
							"id": 1,
							"name": "foo",
							"color": "green",
							"description": "Issue",
							"description_html": "Issue Label",
							"text_color": "black"
						},
						{
							"id": 2,
							"name": "bar",
							"color": "red",
							"description": "Bug",
							"description_html": "Bug Label",
							"text_color": "black"
						}
			    ]
			  }
			]`,
		)
	})

	listProjectIssue := &ListIssuesOptions{
		AuthorID:   Ptr(int64(0o1)),
		AssigneeID: AssigneeID(0o2),
	}

	issues, _, err := client.Issues.ListIssues(listProjectIssue)
	require.NoError(t, err)

	want := []*Issue{{
		ID:          1,
		Description: "This is test project",
		Author:      &IssueAuthor{ID: 1, Name: "snehal"},
		Assignees:   []*IssueAssignee{{ID: 1}},
		Labels:      []string{"foo", "bar"},
		LabelDetails: []*LabelDetails{
			{ID: 1, Name: "foo", Color: "green", Description: "Issue", DescriptionHTML: "Issue Label", TextColor: "black"},
			{ID: 2, Name: "bar", Color: "red", Description: "Bug", DescriptionHTML: "Bug Label", TextColor: "black"},
		},
	}}

	assert.Equal(t, want, issues)
}

func TestListIssuesSearchInTitle(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/issues", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, "/api/v4/issues?in=title&search=Title")
		fmt.Fprint(w, `
			[
				{
					"id": 1,
					"title": "A Test Issue Title",
					"description": "This is the description for the issue"
			  }
			]`,
		)
	})

	listProjectIssue := &ListIssuesOptions{
		Search: Ptr("Title"),
		In:     Ptr("title"),
	}

	issues, _, err := client.Issues.ListIssues(listProjectIssue)
	require.NoError(t, err)

	want := []*Issue{{
		ID:          1,
		Title:       "A Test Issue Title",
		Description: "This is the description for the issue",
	}}

	assert.Equal(t, want, issues)
}

func TestListIssuesSearchInDescription(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/issues", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, "/api/v4/issues?in=description&search=description")
		fmt.Fprint(w, `
			[
				{
					"id": 1,
					"title": "A Test Issue Title",
					"description": "This is the description for the issue"
			  }
			]`,
		)
	})

	listProjectIssue := &ListIssuesOptions{
		Search: Ptr("description"),
		In:     Ptr("description"),
	}

	issues, _, err := client.Issues.ListIssues(listProjectIssue)
	require.NoError(t, err)

	want := []*Issue{{
		ID:          1,
		Title:       "A Test Issue Title",
		Description: "This is the description for the issue",
	}}

	assert.Equal(t, want, issues)
}

func TestListIssuesSearchByIterationID(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/issues", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, "/api/v4/issues?iteration_id=90")
		fmt.Fprint(w, `
			[
				{
					"id": 1,
					"title": "A Test Issue Title",
					"description": "This is the description for the issue",
					"iteration": {
						"id":90,
						"iid":4,
						"sequence":2,
						"group_id":162,
						"state":2,
						"web_url":"https://gitlab.com/groups/my-group/-/iterations/90"
					}
				}
			]`,
		)
	})

	listProjectIssue := &ListIssuesOptions{
		IterationID: Ptr(int64(90)),
	}

	issues, _, err := client.Issues.ListIssues(listProjectIssue)
	require.NoError(t, err)

	want := []*Issue{{
		ID:          1,
		Title:       "A Test Issue Title",
		Description: "This is the description for the issue",
		Iteration: &GroupIteration{
			ID:       90,
			IID:      4,
			Sequence: 2,
			GroupID:  162,
			State:    2,
			WebURL:   "https://gitlab.com/groups/my-group/-/iterations/90",
		},
	}}

	assert.Equal(t, want, issues)
}

func TestListProjectIssues(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/issues", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, "/api/v4/projects/1/issues?assignee_id=2&author_id=1")
		fmt.Fprint(w, `[{"id":1, "description": "This is test project", "author" : {"id" : 1, "name": "snehal"}, "assignees":[{"id":1}]}]`)
	})

	listProjectIssue := &ListProjectIssuesOptions{
		AuthorID:   Ptr(int64(0o1)),
		AssigneeID: AssigneeID(0o2),
	}
	issues, _, err := client.Issues.ListProjectIssues("1", listProjectIssue)
	require.NoError(t, err)

	want := []*Issue{{
		ID:          1,
		Description: "This is test project",
		Author:      &IssueAuthor{ID: 1, Name: "snehal"},
		Assignees:   []*IssueAssignee{{ID: 1}},
	}}

	assert.Equal(t, want, issues)
}

func TestListProjectIssuesSearchByIterationID(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/issues", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, "/api/v4/projects/1/issues?iteration_id=90")
		fmt.Fprint(w, `
			[
				{
					"id": 1,
					"title": "A Test Issue Title",
					"description": "This is the description for the issue",
					"iteration": {
						"id":90,
						"iid":4,
						"sequence":2,
						"group_id":162,
						"state":2,
						"web_url":"https://gitlab.com/groups/my-group/-/iterations/90"
					}
				}
			]`,
		)
	})

	listProjectIssue := &ListProjectIssuesOptions{
		IterationID: Ptr(int64(90)),
	}

	issues, _, err := client.Issues.ListProjectIssues(1, listProjectIssue)
	require.NoError(t, err)

	want := []*Issue{{
		ID:          1,
		Title:       "A Test Issue Title",
		Description: "This is the description for the issue",
		Iteration: &GroupIteration{
			ID:       90,
			IID:      4,
			Sequence: 2,
			GroupID:  162,
			State:    2,
			WebURL:   "https://gitlab.com/groups/my-group/-/iterations/90",
		},
	}}

	assert.Equal(t, want, issues)
}

func TestListGroupIssues(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/issues", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, "/api/v4/groups/1/issues?assignee_id=2&author_id=1&confidential=false&state=Open")
		fmt.Fprint(w, `[{"id":1, "description": "This is test project", "author" : {"id" : 1, "name": "snehal"}, "assignees":[{"id":1}]}]`)
	})

	listGroupIssue := &ListGroupIssuesOptions{
		State:        Ptr("Open"),
		AuthorID:     Ptr(int64(0o1)),
		AssigneeID:   AssigneeID(0o2),
		Confidential: Ptr(false),
	}

	issues, _, err := client.Issues.ListGroupIssues("1", listGroupIssue)
	require.NoError(t, err)

	want := []*Issue{{
		ID:          1,
		Description: "This is test project",
		Author:      &IssueAuthor{ID: 1, Name: "snehal"},
		Assignees:   []*IssueAssignee{{ID: 1}},
	}}

	assert.Equal(t, want, issues)
}

func TestListGroupIssuesSearchByIterationID(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/issues", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, "/api/v4/groups/1/issues?iteration_id=90")
		fmt.Fprint(w, `
			[
				{
					"id": 1,
					"title": "A Test Issue Title",
					"description": "This is the description for the issue",
					"iteration": {
						"id":90,
						"iid":4,
						"sequence":2,
						"group_id":162,
						"state":2,
						"web_url":"https://gitlab.com/groups/my-group/-/iterations/90"
					}
				}
			]`,
		)
	})

	listProjectIssue := &ListGroupIssuesOptions{
		IterationID: Ptr(int64(90)),
	}

	issues, _, err := client.Issues.ListGroupIssues(1, listProjectIssue)
	require.NoError(t, err)

	want := []*Issue{{
		ID:          1,
		Title:       "A Test Issue Title",
		Description: "This is the description for the issue",
		Iteration: &GroupIteration{
			ID:       90,
			IID:      4,
			Sequence: 2,
			GroupID:  162,
			State:    2,
			WebURL:   "https://gitlab.com/groups/my-group/-/iterations/90",
		},
	}}

	assert.Equal(t, want, issues)
}

func TestCreateIssue(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/issues", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"id":1, "title" : "Title of issue", "description": "This is description of an issue", "author" : {"id" : 1, "name": "snehal"}, "assignees":[{"id":1}]}`)
	})

	createIssueOptions := &CreateIssueOptions{
		Title:       Ptr("Title of issue"),
		Description: Ptr("This is description of an issue"),
	}

	issue, _, err := client.Issues.CreateIssue("1", createIssueOptions)
	require.NoError(t, err)

	want := &Issue{
		ID:          1,
		Title:       "Title of issue",
		Description: "This is description of an issue",
		Author:      &IssueAuthor{ID: 1, Name: "snehal"},
		Assignees:   []*IssueAssignee{{ID: 1}},
	}

	assert.Equal(t, want, issue)
}

func TestUpdateIssue(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/issues/5", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		testBodyJSON(t, r, map[string]any{
			"title":       "Title of issue",
			"description": "This is description of an issue",
		})
		mustWriteJSONResponse(t, w, map[string]any{
			"id":          1,
			"title":       "Title of issue",
			"description": "This is description of an issue",
			"author": map[string]any{
				"id":   1,
				"name": "snehal",
			},
			"assignees": []map[string]any{
				{"id": 1},
			},
		})
	})

	updateIssueOpt := &UpdateIssueOptions{
		Title:       Ptr("Title of issue"),
		Description: Ptr("This is description of an issue"),
	}
	issue, _, err := client.Issues.UpdateIssue(1, 5, updateIssueOpt)
	require.NoError(t, err)

	want := &Issue{
		ID:          1,
		Title:       "Title of issue",
		Description: "This is description of an issue",
		Author:      &IssueAuthor{ID: 1, Name: "snehal"},
		Assignees:   []*IssueAssignee{{ID: 1}},
	}

	assert.Equal(t, want, issue)
}

func TestUpdateIssue_ResetFields(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		opts     UpdateIssueOptions
		wantBody map[string]any
	}{
		{
			name:     "set due date",
			opts:     UpdateIssueOptions{DueDate: Ptr(ISOTime(time.Date(2025, time.May, 12, 0, 0, 0, 0, time.UTC)))},
			wantBody: map[string]any{"due_date": "2025-05-12"},
		},
		{
			name:     "unset due date",
			opts:     UpdateIssueOptions{ResetDueDate: true},
			wantBody: map[string]any{"due_date": nil},
		},
		{
			name:     "set epic",
			opts:     UpdateIssueOptions{EpicID: Ptr(int64(42))},
			wantBody: map[string]any{"epic_id": float64(42)},
		},
		{
			name:     "unset epic",
			opts:     UpdateIssueOptions{ResetEpicID: true},
			wantBody: map[string]any{"epic_id": nil},
		},
		{
			name:     "set milestone",
			opts:     UpdateIssueOptions{MilestoneID: Ptr(int64(42))},
			wantBody: map[string]any{"milestone_id": float64(42)},
		},
		{
			name:     "unset milestone",
			opts:     UpdateIssueOptions{ResetMilestoneID: true},
			wantBody: map[string]any{"milestone_id": nil},
		},
		{
			name:     "set weight",
			opts:     UpdateIssueOptions{Weight: Ptr(int64(42))},
			wantBody: map[string]any{"weight": float64(42)},
		},
		{
			name:     "unset weight",
			opts:     UpdateIssueOptions{ResetWeight: true},
			wantBody: map[string]any{"weight": nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mux, client := setup(t)

			mux.HandleFunc("/api/v4/projects/1/issues/5", func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodPut)
				testBodyJSON(t, r, tt.wantBody)
				mustWriteJSONResponse(t, w, map[string]any{"id": 5})
			})

			issue, _, err := client.Issues.UpdateIssue(1, 5, &tt.opts)
			require.NoError(t, err)
			assert.Equal(t, int64(5), issue.ID)
		})
	}
}

func TestSubscribeToIssue(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/issues/5/subscribe", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"id":1, "title" : "Title of issue", "description": "This is description of an issue", "author" : {"id" : 1, "name": "snehal"}, "assignees":[{"id":1}]}`)
	})

	issue, _, err := client.Issues.SubscribeToIssue("1", 5)
	require.NoError(t, err)

	want := &Issue{
		ID:          1,
		Title:       "Title of issue",
		Description: "This is description of an issue",
		Author:      &IssueAuthor{ID: 1, Name: "snehal"},
		Assignees:   []*IssueAssignee{{ID: 1}},
	}

	assert.Equal(t, want, issue)
}

func TestUnsubscribeFromIssue(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/issues/5/unsubscribe", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"id":1, "title" : "Title of issue", "description": "This is description of an issue", "author" : {"id" : 1, "name": "snehal"}, "assignees":[{"id":1}]}`)
	})

	issue, _, err := client.Issues.UnsubscribeFromIssue("1", 5)
	require.NoError(t, err)

	want := &Issue{
		ID:          1,
		Title:       "Title of issue",
		Description: "This is description of an issue",
		Author:      &IssueAuthor{ID: 1, Name: "snehal"},
		Assignees:   []*IssueAssignee{{ID: 1}},
	}

	assert.Equal(t, want, issue)
}

func TestListMergeRequestsClosingIssue(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/issues/5/closed_by", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, "/api/v4/projects/1/issues/5/closed_by?page=1&per_page=10")

		fmt.Fprint(w, `[{"id":1, "title" : "test merge one"},{"id":2, "title" : "test merge two"}]`)
	})

	listMergeRequestsClosingIssueOpt := &ListMergeRequestsClosingIssueOptions{
		ListOptions: ListOptions{
			Page:    1,
			PerPage: 10,
		},
	}
	mergeRequest, _, err := client.Issues.ListMergeRequestsClosingIssue("1", 5, listMergeRequestsClosingIssueOpt)
	require.NoError(t, err)

	want := []*BasicMergeRequest{{ID: 1, Title: "test merge one"}, {ID: 2, Title: "test merge two"}}

	assert.Equal(t, want, mergeRequest)
}

func TestListMergeRequestsRelatedToIssue(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/issues/5/related_merge_requests", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, "/api/v4/projects/1/issues/5/related_merge_requests?page=1&per_page=10")

		fmt.Fprint(w, `[{"id":1, "title" : "test merge one"},{"id":2, "title" : "test merge two"}]`)
	})

	listMergeRequestsRelatedToIssueOpt := &ListMergeRequestsRelatedToIssueOptions{
		ListOptions: ListOptions{
			Page:    1,
			PerPage: 10,
		},
	}
	mergeRequest, _, err := client.Issues.ListMergeRequestsRelatedToIssue("1", 5, listMergeRequestsRelatedToIssueOpt)
	require.NoError(t, err)

	want := []*BasicMergeRequest{{ID: 1, Title: "test merge one"}, {ID: 2, Title: "test merge two"}}

	assert.Equal(t, want, mergeRequest)
}

func TestSetTimeEstimate(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/issues/5/time_estimate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"human_time_estimate": "3h 30m", "human_total_time_spent": null, "time_estimate": 12600, "total_time_spent": 0}`)
	})

	setTimeEstiOpt := &SetTimeEstimateOptions{
		Duration: Ptr("3h 30m"),
	}

	timeState, _, err := client.Issues.SetTimeEstimate("1", 5, setTimeEstiOpt)
	require.NoError(t, err)

	want := &TimeStats{HumanTimeEstimate: "3h 30m", HumanTotalTimeSpent: "", TimeEstimate: 12600, TotalTimeSpent: 0}

	assert.Equal(t, want, timeState)
}

func TestResetTimeEstimate(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/issues/5/reset_time_estimate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"human_time_estimate": null, "human_total_time_spent": null, "time_estimate": 0, "total_time_spent": 0}`)
	})

	timeState, _, err := client.Issues.ResetTimeEstimate("1", 5)
	require.NoError(t, err)

	want := &TimeStats{HumanTimeEstimate: "", HumanTotalTimeSpent: "", TimeEstimate: 0, TotalTimeSpent: 0}

	assert.Equal(t, want, timeState)
}

func TestAddSpentTime(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/issues/5/add_spent_time", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, "/api/v4/projects/1/issues/5/add_spent_time")
		testBodyJSON(t, r, map[string]string{
			"duration": "1h",
			"summary":  "test",
		})
		fmt.Fprint(w, `{"human_time_estimate": null, "human_total_time_spent": "1h", "time_estimate": 0, "total_time_spent": 3600}`)
	})
	addSpentTimeOpt := &AddSpentTimeOptions{
		Duration: Ptr("1h"),
		Summary:  Ptr("test"),
	}

	timeState, _, err := client.Issues.AddSpentTime("1", 5, addSpentTimeOpt)
	require.NoError(t, err)

	want := &TimeStats{HumanTimeEstimate: "", HumanTotalTimeSpent: "1h", TimeEstimate: 0, TotalTimeSpent: 3600}

	assert.Equal(t, want, timeState)
}

func TestResetSpentTime(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/issues/5/reset_spent_time", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testURL(t, r, "/api/v4/projects/1/issues/5/reset_spent_time")
		fmt.Fprint(w, `{"human_time_estimate": null, "human_total_time_spent": "", "time_estimate": 0, "total_time_spent": 0}`)
	})

	timeState, _, err := client.Issues.ResetSpentTime("1", 5)
	require.NoError(t, err)

	want := &TimeStats{HumanTimeEstimate: "", HumanTotalTimeSpent: "", TimeEstimate: 0, TotalTimeSpent: 0}

	assert.Equal(t, want, timeState)
}

func TestGetTimeSpent(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/issues/5/time_stats", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, "/api/v4/projects/1/issues/5/time_stats")
		fmt.Fprint(w, `{"human_time_estimate": "2h", "human_total_time_spent": "1h", "time_estimate": 7200, "total_time_spent": 3600}`)
	})

	timeState, _, err := client.Issues.GetTimeSpent("1", 5)
	require.NoError(t, err)

	want := &TimeStats{HumanTimeEstimate: "2h", HumanTotalTimeSpent: "1h", TimeEstimate: 7200, TotalTimeSpent: 3600}

	assert.Equal(t, want, timeState)
}

func TestGetIssueParticipants(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/issues/5/participants", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, "/api/v4/projects/1/issues/5/participants")

		fmt.Fprint(w, `[{"id":1,"name":"User1","username":"User1","state":"active","avatar_url":"","web_url":"https://localhost/User1"},
		{"id":2,"name":"User2","username":"User2","state":"active","avatar_url":"https://localhost/uploads/-/system/user/avatar/2/avatar.png","web_url":"https://localhost/User2"}]`)
	})

	issueParticipants, _, err := client.Issues.GetParticipants("1", 5)
	require.NoError(t, err)

	want := []*BasicUser{
		{ID: 1, Name: "User1", Username: "User1", State: "active", AvatarURL: "", WebURL: "https://localhost/User1"},
		{ID: 2, Name: "User2", Username: "User2", State: "active", AvatarURL: "https://localhost/uploads/-/system/user/avatar/2/avatar.png", WebURL: "https://localhost/User2"},
	}

	assert.Equal(t, want, issueParticipants)
}

func TestGetIssueMilestone(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/issues/5", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":1, "description": "This is test project", "author" : {"id" : 1, "name": "snehal"}, "assignees":[{"id":1}],"merge_requests_count": 1,
			"milestone": {"due_date": null, "project_id": 1, "state": "closed", "description": "test", "iid": 3, "id": 11, "title": "v3.0"}}`)
	})

	issue, _, err := client.Issues.GetIssue("1", 5)
	require.NoError(t, err)

	want := &Issue{
		ID:                1,
		Description:       "This is test project",
		Author:            &IssueAuthor{ID: 1, Name: "snehal"},
		Assignees:         []*IssueAssignee{{ID: 1}},
		MergeRequestCount: 1,
		Milestone: &Milestone{
			DueDate:     nil,
			ProjectID:   1,
			State:       "closed",
			Description: "test",
			IID:         3,
			ID:          11,
			Title:       "v3.0",
		},
	}

	assert.Equal(t, want, issue)
}

func TestGetIssueGroupMilestone(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/issues/5", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":1, "description": "This is test project", "author" : {"id" : 1, "name": "snehal"}, "assignees":[{"id":1}],"merge_requests_count": 1,
			"milestone": {"due_date": null, "group_id": 13, "state": "closed", "description": "test", "iid": 3, "id": 11, "title": "v3.0"}}`)
	})

	issue, _, err := client.Issues.GetIssue("1", 5)
	require.NoError(t, err)

	want := &Issue{
		ID:                1,
		Description:       "This is test project",
		Author:            &IssueAuthor{ID: 1, Name: "snehal"},
		Assignees:         []*IssueAssignee{{ID: 1}},
		MergeRequestCount: 1,
		Milestone: &Milestone{
			DueDate:     nil,
			GroupID:     13,
			State:       "closed",
			Description: "test",
			IID:         3,
			ID:          11,
			Title:       "v3.0",
		},
	}

	assert.Equal(t, want, issue)
}

func TestGetIssueWithServiceDesk(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/issues/5", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":1, "description": "This is test project", "author" : {"id" : 1, "name": "snehal"}, "assignees":[{"id":1}],"service_desk_reply_to": "snehal@test.com"}`)
	})

	issue, _, err := client.Issues.GetIssue("1", 5)
	require.NoError(t, err)

	want := &Issue{
		ID:                 1,
		Description:        "This is test project",
		Author:             &IssueAuthor{ID: 1, Name: "snehal"},
		Assignees:          []*IssueAssignee{{ID: 1}},
		ServiceDeskReplyTo: "snehal@test.com",
	}

	assert.Equal(t, want, issue)
}
