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
	"testing"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	ajk = BasicUser{
		ID:        3614858,
		Name:      "Alex Kalderimis",
		Username:  "alexkalderimis",
		State:     "active",
		AvatarURL: "https://assets.gitlab-static.net/uploads/-/system/user/avatar/3614858/avatar.png",
		WebURL:    "https://gitlab.com/alexkalderimis",
	}
	tk = BasicUser{
		ID:        2535118,
		Name:      "Thong Kuah",
		Username:  "tkuah",
		State:     "active",
		AvatarURL: "https://secure.gravatar.com/avatar/f7b51bdd49a4914d29504d7ff4c3f7b9?s=80&d=identicon",
		WebURL:    "https://gitlab.com/tkuah",
	}
	getOpts = GetMergeRequestsOptions{}
	labels  = Labels{
		"GitLab Enterprise Edition",
		"backend",
		"database",
		"database::reviewed",
		"design management",
		"feature",
		"frontend",
		"group::knowledge",
		"missed:12.1",
	}
	pipelineCreation = time.Date(2019, 8, 19, 9, 50, 58, 157000000, time.UTC)
	pipelineUpdate   = time.Date(2019, 8, 19, 19, 22, 29, 647000000, time.UTC)
	pipelineBasic    = PipelineInfo{
		ID:        77056819,
		SHA:       "8e0b45049b6253b8984cde9241830d2851168142",
		Ref:       "delete-designs-v2",
		Status:    "success",
		WebURL:    "https://gitlab.com/gitlab-org/gitlab-ee/pipelines/77056819",
		CreatedAt: &pipelineCreation,
		UpdatedAt: &pipelineUpdate,
	}
	pipelineStarted  = time.Date(2019, 8, 19, 9, 51, 6, 545000000, time.UTC)
	pipelineFinished = time.Date(2019, 8, 19, 19, 22, 29, 632000000, time.UTC)
	pipelineDetailed = Pipeline{
		ID:         77056819,
		SHA:        "8e0b45049b6253b8984cde9241830d2851168142",
		Ref:        "delete-designs-v2",
		Status:     "success",
		WebURL:     "https://gitlab.com/gitlab-org/gitlab-ee/pipelines/77056819",
		BeforeSHA:  "3fe568caacb261b63090886f5b879ca0d9c6f4c3",
		Tag:        false,
		User:       &ajk,
		CreatedAt:  &pipelineCreation,
		UpdatedAt:  &pipelineUpdate,
		StartedAt:  &pipelineStarted,
		FinishedAt: &pipelineFinished,
		Duration:   4916,
		Coverage:   "82.68",
		DetailedStatus: &DetailedStatus{
			Icon:        "status_warning",
			Text:        "passed",
			Label:       "passed with warnings",
			Group:       "success-with-warnings",
			Tooltip:     "passed",
			HasDetails:  true,
			DetailsPath: "/gitlab-org/gitlab-ee/pipelines/77056819",
			Favicon:     "https://gitlab.com/assets/ci_favicons/favicon_status_success-8451333011eee8ce9f2ab25dc487fe24a8758c694827a582f17f42b0a90446a2.png",
		},
	}
)

func TestGetMergeRequest(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := `/api/v4/projects/namespace%2Fname/merge_requests/123`

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_merge_request.json")
	})

	mergeRequest, _, err := client.MergeRequests.GetMergeRequest("namespace/name", 123, &getOpts)

	require.NoError(t, err)

	require.Equal(t, mergeRequest.ID, 33092005)
	require.Equal(t, mergeRequest.SHA, "8e0b45049b6253b8984cde9241830d2851168142")
	require.Equal(t, mergeRequest.IID, 14656)
	require.Equal(t, mergeRequest.ProjectID, 278964)
	require.Equal(t, mergeRequest.SourceBranch, "delete-designs-v2")
	require.Equal(t, mergeRequest.TaskCompletionStatus.Count, 9)
	require.Equal(t, mergeRequest.TaskCompletionStatus.CompletedCount, 8)
	require.Equal(t, mergeRequest.Title, "Add deletion support for designs")
	require.Equal(t, mergeRequest.Description,
		"## What does this MR do?\r\n\r\nThis adds the capability to destroy/hide designs.")
	require.Equal(t, mergeRequest.WebURL,
		"https://gitlab.com/gitlab-org/gitlab-ee/merge_requests/14656")
	require.Equal(t, mergeRequest.DetailedMergeStatus, "mergeable")
	require.Equal(t, mergeRequest.Author, &ajk)
	require.Equal(t, mergeRequest.Assignee, &tk)
	require.Equal(t, mergeRequest.Assignees, []*BasicUser{&tk})
	require.Equal(t, mergeRequest.Reviewers, []*BasicUser{&tk})
	require.Equal(t, mergeRequest.Labels, labels)
	require.Equal(t, mergeRequest.Squash, true)
	require.Equal(t, mergeRequest.UserNotesCount, 245)
	require.Equal(t, mergeRequest.Pipeline, &pipelineBasic)
	require.Equal(t, mergeRequest.HeadPipeline, &pipelineDetailed)
	mrCreation := time.Date(2019, 7, 11, 22, 34, 43, 500000000, time.UTC)
	require.Equal(t, mergeRequest.CreatedAt, &mrCreation)
	mrUpdate := time.Date(2019, 8, 20, 9, 9, 56, 690000000, time.UTC)
	require.Equal(t, mergeRequest.UpdatedAt, &mrUpdate)
	require.Equal(t, mergeRequest.FirstContribution, true)
	require.Equal(t, mergeRequest.HasConflicts, true)
	require.Equal(t, mergeRequest.Draft, true)
}

func TestListProjectMergeRequests(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := "/api/v4/projects/278964/merge_requests"

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testParams(t, r, "assignee_id=Any&with_labels_details=true&with_merge_status_recheck=true")
		mustWriteHTTPResponse(t, w, "testdata/get_merge_requests.json")
	})

	opts := ListProjectMergeRequestsOptions{
		AssigneeID:             AssigneeID(UserIDAny),
		WithLabelsDetails:      Ptr(true),
		WithMergeStatusRecheck: Ptr(true),
	}

	mergeRequests, _, err := client.MergeRequests.ListProjectMergeRequests(278964, &opts)

	require.NoError(t, err)
	require.Equal(t, 3, len(mergeRequests))

	validStates := []string{"opened", "closed", "locked", "merged"}
	detailedMergeStatuses := []string{
		"blocked_status",
		"broken_status",
		"checking",
		"ci_must_pass",
		"ci_still_running",
		"discussions_not_resolved",
		"draft_status",
		"external_status_checks",
		"mergeable",
		"not_approved",
		"not_open",
		"policies_denied",
		"unchecked",
	}
	allCreatedBefore := time.Date(2019, 8, 21, 0, 0, 0, 0, time.UTC)
	allCreatedAfter := time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC)

	for _, mr := range mergeRequests {
		require.Equal(t, 278964, mr.ProjectID)
		require.Contains(t, validStates, mr.State)
		assert.Less(t, mr.CreatedAt.Unix(), allCreatedBefore.Unix())
		assert.Greater(t, mr.CreatedAt.Unix(), allCreatedAfter.Unix())
		assert.LessOrEqual(t, mr.CreatedAt.Unix(), mr.UpdatedAt.Unix())
		assert.LessOrEqual(t, mr.TaskCompletionStatus.CompletedCount, mr.TaskCompletionStatus.Count)
		require.Contains(t, detailedMergeStatuses, mr.DetailedMergeStatus)
	}
}

func TestListProjectMergeRequestsAuthorUsername(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := "/api/v4/projects/278964/merge_requests"

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testParams(t, r, "assignee_id=Any&author_username=hfyngvason&with_labels_details=true&with_merge_status_recheck=true")
		mustWriteHTTPResponse(t, w, "testdata/get_merge_requests_author_username.json")
	})

	opts := ListProjectMergeRequestsOptions{
		AssigneeID:             AssigneeID(UserIDAny),
		AuthorUsername:         Ptr("hfyngvason"),
		WithLabelsDetails:      Ptr(true),
		WithMergeStatusRecheck: Ptr(true),
	}

	mergeRequests, _, err := client.MergeRequests.ListProjectMergeRequests(278964, &opts)

	require.NoError(t, err)
	require.Equal(t, 1, len(mergeRequests))

	validStates := []string{"opened", "closed", "locked", "merged"}
	detailedMergeStatuses := []string{
		"blocked_status",
		"broken_status",
		"checking",
		"ci_must_pass",
		"ci_still_running",
		"discussions_not_resolved",
		"draft_status",
		"external_status_checks",
		"mergeable",
		"not_approved",
		"not_open",
		"policies_denied",
		"unchecked",
	}
	allCreatedBefore := time.Date(2019, 8, 21, 0, 0, 0, 0, time.UTC)
	allCreatedAfter := time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC)

	for _, mr := range mergeRequests {
		require.Equal(t, 278964, mr.ProjectID)
		require.Contains(t, validStates, mr.State)
		assert.Less(t, mr.CreatedAt.Unix(), allCreatedBefore.Unix())
		assert.Greater(t, mr.CreatedAt.Unix(), allCreatedAfter.Unix())
		assert.LessOrEqual(t, mr.CreatedAt.Unix(), mr.UpdatedAt.Unix())
		assert.LessOrEqual(t, mr.TaskCompletionStatus.CompletedCount, mr.TaskCompletionStatus.Count)
		require.Contains(t, detailedMergeStatuses, mr.DetailedMergeStatus)
	}
}

func TestListProjectMergeRequestsNotAuthorUsername(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	path := "/api/v4/projects/278964/merge_requests"

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testParams(t, r, "assignee_id=Any&not%5Bauthor_username%5D=hfyngvason&with_labels_details=true&with_merge_status_recheck=true")
		mustWriteHTTPResponse(t, w, "testdata/get_merge_requests_not_author_username.json")
	})

	opts := ListProjectMergeRequestsOptions{
		AssigneeID:             AssigneeID(UserIDAny),
		NotAuthorUsername:      Ptr("hfyngvason"),
		WithLabelsDetails:      Ptr(true),
		WithMergeStatusRecheck: Ptr(true),
	}

	mergeRequests, _, err := client.MergeRequests.ListProjectMergeRequests(278964, &opts)

	require.NoError(t, err)
	require.Equal(t, 2, len(mergeRequests))

	validStates := []string{"opened", "closed", "locked", "merged"}
	detailedMergeStatuses := []string{
		"blocked_status",
		"broken_status",
		"checking",
		"ci_must_pass",
		"ci_still_running",
		"discussions_not_resolved",
		"draft_status",
		"external_status_checks",
		"mergeable",
		"not_approved",
		"not_open",
		"policies_denied",
		"unchecked",
	}
	allCreatedBefore := time.Date(2019, 8, 21, 0, 0, 0, 0, time.UTC)
	allCreatedAfter := time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC)

	for _, mr := range mergeRequests {
		require.Equal(t, 278964, mr.ProjectID)
		require.Contains(t, validStates, mr.State)
		assert.Less(t, mr.CreatedAt.Unix(), allCreatedBefore.Unix())
		assert.Greater(t, mr.CreatedAt.Unix(), allCreatedAfter.Unix())
		assert.LessOrEqual(t, mr.CreatedAt.Unix(), mr.UpdatedAt.Unix())
		assert.LessOrEqual(t, mr.TaskCompletionStatus.CompletedCount, mr.TaskCompletionStatus.Count)
		require.Contains(t, detailedMergeStatuses, mr.DetailedMergeStatus)
	}
}

func TestCreateMergeRequestPipeline(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/merge_requests/1/pipelines", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"id":1, "status":"pending"}`)
	})

	pipeline, _, err := client.MergeRequests.CreateMergeRequestPipeline(1, 1)
	if err != nil {
		t.Errorf("MergeRequests.CreateMergeRequestPipeline returned error: %v", err)
	}

	assert.Equal(t, 1, pipeline.ID)
	assert.Equal(t, "pending", pipeline.Status)
}

func TestGetMergeRequestParticipants(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/merge_requests/5/participants", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, "/api/v4/projects/1/merge_requests/5/participants")

		fmt.Fprint(w, `[{"id":1,"name":"User1","username":"User1","state":"active","avatar_url":"","web_url":"https://localhost/User1"},
		{"id":2,"name":"User2","username":"User2","state":"active","avatar_url":"https://localhost/uploads/-/system/user/avatar/2/avatar.png","web_url":"https://localhost/User2"}]`)
	})

	mergeRequestParticipants, _, err := client.MergeRequests.GetMergeRequestParticipants("1", 5)
	if err != nil {
		t.Fatal(err)
	}

	want := []*BasicUser{
		{ID: 1, Name: "User1", Username: "User1", State: "active", AvatarURL: "", WebURL: "https://localhost/User1"},
		{ID: 2, Name: "User2", Username: "User2", State: "active", AvatarURL: "https://localhost/uploads/-/system/user/avatar/2/avatar.png", WebURL: "https://localhost/User2"},
	}

	if !reflect.DeepEqual(want, mergeRequestParticipants) {
		t.Errorf("Issues.GetMergeRequestParticipants returned %+v, want %+v", mergeRequestParticipants, want)
	}
}

func TestGetMergeRequestReviewers(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/merge_requests/5/reviewers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, "/api/v4/projects/1/merge_requests/5/reviewers")

		fmt.Fprint(w, `[{"user":{"id":1,"name":"John Doe1","username":"user1","state":"active","avatar_url":"http://www.gravatar.com/avatar/c922747a93b40d1ea88262bf1aebee62?s=80&d=identicon","web_url":"http://localhost/user1"},"state":"unreviewed","created_at":"2022-07-27T17:03:27.684Z"},{"user":{"id":2,"name":"John Doe2","username":"user2","state":"active","avatar_url":"http://www.gravatar.com/avatar/10fc7f102be8de7657fb4d80898bbfe3?s=80&d=identicon","web_url":"http://localhost/user2"},"state":"reviewed","created_at":"2022-07-27T17:03:27.684Z"}]`)
	})

	mergeRequestReviewers, _, err := client.MergeRequests.GetMergeRequestReviewers("1", 5)
	if err != nil {
		t.Fatal(err)
	}

	createdAt := time.Date(2022, 0o7, 27, 17, 3, 27, 684000000, time.UTC)
	user1 := BasicUser{ID: 1, Name: "John Doe1", Username: "user1", State: "active", AvatarURL: "http://www.gravatar.com/avatar/c922747a93b40d1ea88262bf1aebee62?s=80&d=identicon", WebURL: "http://localhost/user1"}
	user2 := BasicUser{ID: 2, Name: "John Doe2", Username: "user2", State: "active", AvatarURL: "http://www.gravatar.com/avatar/10fc7f102be8de7657fb4d80898bbfe3?s=80&d=identicon", WebURL: "http://localhost/user2"}

	assert.Len(t, mergeRequestReviewers, 2)
	assert.Equal(t, "unreviewed", mergeRequestReviewers[0].State)
	require.Equal(t, &user1, mergeRequestReviewers[0].User)
	require.Equal(t, &createdAt, mergeRequestReviewers[0].CreatedAt)
	assert.Equal(t, "reviewed", mergeRequestReviewers[1].State)
	require.Equal(t, &user2, mergeRequestReviewers[1].User)
	require.Equal(t, &createdAt, mergeRequestReviewers[1].CreatedAt)
}

func TestGetIssuesClosedOnMerge_Jira(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)
	mux.HandleFunc("/api/v4/projects/1/merge_requests/1/closes_issues", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":"PROJECT-123","title":"Title of this issue"}]`)
	})

	issues, _, err := client.MergeRequests.GetIssuesClosedOnMerge(1, 1, nil)

	assert.NoError(t, err)
	assert.Len(t, issues, 1)
	assert.Equal(t, "PROJECT-123", issues[0].ExternalID)
	assert.Equal(t, "Title of this issue", issues[0].Title)
}

func TestListRelatedIssues(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)
	mux.HandleFunc("/api/v4/projects/1/merge_requests/1/related_issues", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":"PROJECT-123","title":"Title of this issue"}]`)
	})

	issues, _, err := client.MergeRequests.ListRelatedIssues(1, 1, nil)

	assert.NoError(t, err)
	assert.Len(t, issues, 1)
	assert.Equal(t, "PROJECT-123", issues[0].ExternalID)
	assert.Equal(t, "Title of this issue", issues[0].Title)
}

func TestListMergeRequestDiffs(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/merge_requests/1/diffs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/list_merge_request_diff.json")
	})

	opts := &ListMergeRequestDiffsOptions{
		ListOptions: ListOptions{Page: 1, PerPage: 2},
	}

	diffs, _, err := client.MergeRequests.ListMergeRequestDiffs(1, 1, opts)
	if err != nil {
		t.Errorf("MergeRequests.ListMergeRequestDiffs returned error: %v", err)
	}

	want := []*MergeRequestDiff{
		{
			OldPath:     "README",
			NewPath:     "README",
			AMode:       "100644",
			BMode:       "100644",
			Diff:        "@@ -1 +1 @@ -Title +README",
			NewFile:     false,
			RenamedFile: false,
			DeletedFile: false,
		},
		{
			OldPath:     "VERSION",
			NewPath:     "VERSION",
			AMode:       "100644",
			BMode:       "100644",
			Diff:        "@@ -1.9.7 +1.9.8",
			NewFile:     false,
			RenamedFile: false,
			DeletedFile: false,
		},
	}

	if !reflect.DeepEqual(want, diffs) {
		t.Errorf("MergeRequests.ListMergeRequestDiffs returned %+v, want %+v", diffs, want)
	}
}

func TestShowMergeRequestRawDiffs(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/merge_requests/1/raw_diffs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/show_merge_request_raw_diff.txt")
	})

	opts := &ShowMergeRequestRawDiffsOptions{}

	rawDiff, _, err := client.MergeRequests.ShowMergeRequestRawDiffs(1, 1, opts)
	if err != nil {
		t.Errorf("MergeRequests.ShowMergeRequestRawDiffs returned error: %v", err)
	}

	want := `diff --git a/also_testing b/also_testing
index d4d510b..2a2c3b1 100644
--- a/also_testing
+++ b/also_testing
@@ -1,3 +1,2 @@
 aaaaaaaaaaaaaaaa
-bbbbbbbbbbbbbbbb
 cccccccccccccccc
diff --git a/testing b/testing
index c7c7da3..ce2cd85 100644
--- a/testing
+++ b/testing
@@ -1 +1,2 @@
 hello there
+i'm a test :)
`

	require.Equal(t, []byte(want), rawDiff)
}

func TestIntSliceOrString(t *testing.T) {
	t.Parallel()
	t.Run("any", func(t *testing.T) {
		opts := &ListMergeRequestsOptions{}
		opts.ApprovedByIDs = ApproverIDs(UserIDAny)
		q, err := query.Values(opts)
		assert.NoError(t, err)
		assert.Equal(t, "Any", q.Get("approved_by_ids"))
	})
	t.Run("none", func(t *testing.T) {
		opts := &ListMergeRequestsOptions{}
		opts.ApprovedByIDs = ApproverIDs(UserIDNone)
		q, err := query.Values(opts)
		assert.NoError(t, err)
		assert.Equal(t, "None", q.Get("approved_by_ids"))
	})
	t.Run("ids", func(t *testing.T) {
		opts := &ListMergeRequestsOptions{}
		opts.ApprovedByIDs = ApproverIDs([]int{1, 2, 3})
		q, err := query.Values(opts)
		assert.NoError(t, err)
		includedIDs := q["approved_by_ids[]"]
		assert.Equal(t, []string{"1", "2", "3"}, includedIDs)
	})
}

func TestAssigneeIDMarshalling(t *testing.T) {
	t.Parallel()
	t.Run("any", func(t *testing.T) {
		opts := &ListMergeRequestsOptions{}
		opts.AssigneeID = AssigneeID(UserIDAny)
		q, err := query.Values(opts)
		assert.NoError(t, err)
		assert.Equal(t, "Any", q.Get("assignee_id"))
		js, _ := json.Marshal(opts)
		assert.Equal(t, `{"assignee_id":"Any"}`, string(js))
	})
	t.Run("none", func(t *testing.T) {
		opts := &ListMergeRequestsOptions{}
		opts.AssigneeID = AssigneeID(UserIDNone)
		q, err := query.Values(opts)
		assert.NoError(t, err)
		assert.Equal(t, "None", q.Get("assignee_id"))
		js, _ := json.Marshal(opts)
		assert.Equal(t, `{"assignee_id":"None"}`, string(js))
	})
	t.Run("id", func(t *testing.T) {
		opts := &ListMergeRequestsOptions{}
		opts.AssigneeID = AssigneeID(5)
		q, err := query.Values(opts)
		assert.NoError(t, err)
		assert.Equal(t, "5", q.Get("assignee_id"))
		js, _ := json.Marshal(opts)
		assert.Equal(t, `{"assignee_id":5}`, string(js))
	})
}

func TestCreateMergeRequestDependency(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)
	const project = "12345"
	const mergeRequest = 1
	blockingMergeRequest := int(2)

	path := fmt.Sprintf("/%sprojects/%s/merge_requests/%d/blocks", apiVersionPath, project, mergeRequest)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)

		fmt.Fprint(w, `{"id": 1, "blocking_merge_request": {"iid":12}, "project_id": 7}`)
	})

	opts := CreateMergeRequestDependencyOptions{
		BlockingMergeRequestID: &blockingMergeRequest,
	}
	dependencies, resp, err := client.MergeRequests.CreateMergeRequestDependency(project, mergeRequest, opts, nil)
	if err != nil {
		t.Errorf("MergeRequestDependencies.CreateMergeRequestDependency returned error: %v", err)
	}
	if resp != nil {
		if resp.StatusCode != http.StatusCreated {
			t.Errorf("MergeRequestDependencies.CreateMergeRequestDependency = %v, want %v", resp.StatusCode, http.StatusCreated)
		}
	}

	want := &MergeRequestDependency{
		ID: 1,
		BlockingMergeRequest: BlockingMergeRequest{
			Iid: 12,
		},
		ProjectID: 7,
	}
	if dependencies == nil || !reflect.DeepEqual(*want, *dependencies) {
		t.Fatalf("MergeRequestDependencies.GetMergeRequestDependencies returned %+v, want %+v", dependencies, want)
	}
}

func TestGetMergeRequestDependencies(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)
	const project = "12345"
	const mergeRequest = 1

	path := fmt.Sprintf("/%sprojects/%s/merge_requests/%d/blocks", apiVersionPath, project, mergeRequest)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/get_merge_request_dependencies.json")
	})

	dependencies, resp, err := client.MergeRequests.GetMergeRequestDependencies(project, mergeRequest, nil)
	if err != nil {
		t.Errorf("MergeRequestDependencies.GetMergeRequestDependencies returned error: %v", err)
	}
	if resp == nil {
		t.Error("MergeRequestDependencies.GetMergeRequestDependencies did not return a response")
	}
	if len(dependencies) != 1 {
		t.Errorf("MergeRequestDependencies.GetMergeRequestDependencies returned %d dependencies, want 1", len(dependencies))
	}
}

func TestDeleteMergeRequestDependency(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)
	const project = "12345"
	const mergeRequest = 1
	const blockingMergeRequest = 2

	path := fmt.Sprintf("/%sprojects/%s/merge_requests/%d/blocks/%d", apiVersionPath, project, mergeRequest, blockingMergeRequest)
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := client.MergeRequests.DeleteMergeRequestDependency(project, mergeRequest, blockingMergeRequest, nil)
	if err != nil {
		t.Errorf("MergeRequestDependencies.DeleteMergeRequestDependency returned error: %v", err)
	}
	if resp != nil {
		if resp.StatusCode != http.StatusNoContent {
			t.Errorf("MergeRequestDependencies.DeleteMergeRequestDependency = %v, want %v", resp.StatusCode, http.StatusNoContent)
		}
	}
}
