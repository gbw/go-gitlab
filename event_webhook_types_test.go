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
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildEventUnmarshal(t *testing.T) {
	t.Parallel()
	jsonObject := loadFixture(t, "testdata/webhooks/build.json")

	var event *BuildEvent
	err := json.Unmarshal(jsonObject, &event)
	require.NoError(t, err)

	expectedEvent := &BuildEvent{
		ObjectKind:     "build",
		Ref:            "gitlab-script-trigger",
		Tag:            false,
		BeforeSHA:      "2293ada6b400935a1378653304eaf6221e0fdb8f",
		SHA:            "2293ada6b400935a1378653304eaf6221e0fdb8f",
		BuildID:        1977,
		BuildName:      "test",
		BuildStage:     "test",
		BuildStatus:    "created",
		BuildCreatedAt: "2021-02-23T02:41:37.886Z",
		ProjectID:      380,
		ProjectName:    "gitlab-org/gitlab-test",
		User: &EventUser{
			ID:        42,
			Name:      "User1",
			Username:  "user1",
			AvatarURL: "http://www.gravatar.com/avatar/e32bd13e2add097461cb96824b7a829c?s=80\u0026d=identicon",
			Email:     "user1@example.com",
		},
		Commit: BuildEventCommit{
			ID:          2366,
			SHA:         "2293ada6b400935a1378653304eaf6221e0fdb8f",
			Message:     "test\n",
			AuthorName:  "User",
			AuthorEmail: "user@gitlab.com",
			Status:      "created",
		},
		Repository: &Repository{
			Name:        "gitlab_test",
			Description: "Atque in sunt eos similique dolores voluptatem.",
			Homepage:    "http://192.168.64.1:3005/gitlab-org/gitlab-test",
			GitSSHURL:   "git@192.168.64.1:gitlab-org/gitlab-test.git",
			GitHTTPURL:  "http://192.168.64.1:3005/gitlab-org/gitlab-test.git",
		},
	}
	assert.Equal(t, expectedEvent, event)
}

func TestCommitCommentEventUnmarshal(t *testing.T) {
	t.Parallel()
	jsonObject := loadFixture(t, "testdata/webhooks/note_commit.json")

	var event *CommitCommentEvent
	err := json.Unmarshal(jsonObject, &event)
	require.NoError(t, err)

	expectedEvent := &CommitCommentEvent{
		ObjectKind: "note",
		EventType:  "note",
		User: &User{
			ID:        42,
			Username:  "user1",
			Email:     "user1@example.com",
			Name:      "User1",
			AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=40\u0026d=identicon",
		},
		ProjectID: 5,
		Project: CommitCommentEventProject{
			ID:                5,
			Name:              "GitLab Test",
			Description:       "Aut reprehenderit ut est.",
			GitSSHURL:         "git@example.com:gitlabhq/gitlab-test.git",
			GitHTTPURL:        "http://example.com/gitlabhq/gitlab-test.git",
			Namespace:         "GitlabHQ",
			PathWithNamespace: "gitlabhq/gitlab-test",
			DefaultBranch:     "master",
			Homepage:          "http://example.com/gitlabhq/gitlab-test",
			URL:               "http://example.com/gitlabhq/gitlab-test.git",
			SSHURL:            "git@example.com:gitlabhq/gitlab-test.git",
			HTTPURL:           "http://example.com/gitlabhq/gitlab-test.git",
			WebURL:            "http://example.com/gitlabhq/gitlab-test",
		},
		Repository: &Repository{
			Name:        "GitLab Test",
			URL:         "http://example.com/gitlab-org/gitlab-test.git",
			Description: "Aut reprehenderit ut est.",
			Homepage:    "http://example.com/gitlab-org/gitlab-test",
		},
		ObjectAttributes: CommitCommentEventObjectAttributes{
			ID:           1243,
			Note:         "This is a commit comment. How does this work?",
			NoteableType: "Commit",
			AuthorID:     1,
			CreatedAt:    "2015-05-17 18:08:09 UTC",
			UpdatedAt:    "2015-05-17 18:08:09 UTC",
			ProjectID:    5,
			LineCode:     "bec9703f7a456cd2b4ab5fb3220ae016e3e394e3_0_1",
			CommitID:     "cfe32cf61b73a0d5e9f13e774abde7ff789b1660",
			System:       false,
			StDiff: &Diff{
				Diff:        "--- /dev/null\n+++ b/six\n@@ -0,0 +1 @@\n+Subproject commit 409f37c4f05865e4fb208c771485f211a22c4c2d\n",
				NewPath:     "six",
				OldPath:     "six",
				AMode:       "0",
				BMode:       "160000",
				NewFile:     true,
				RenamedFile: false,
				DeletedFile: false,
			},
			Description: "This is a commit comment. How does this work?",
			Action:      CommentEventActionCreate,
			URL:         "http://example.com/gitlab-org/gitlab-test/commit/cfe32cf61b73a0d5e9f13e774abde7ff789b1660#note_1243",
		},
		Commit: &CommitCommentEventCommit{
			ID:      "cfe32cf61b73a0d5e9f13e774abde7ff789b1660",
			Title:   "Add submodule",
			Message: "Add submodule\n\nSigned-off-by: Dmitriy Zaporozhets <dmitriy.zaporozhets@gmail.com>\n",
			Timestamp: func() *time.Time {
				ts, err := time.Parse(time.RFC3339, "2014-02-27T10:06:20+02:00")
				require.NoError(t, err)
				return &ts
			}(),
			URL: "http://example.com/gitlab-org/gitlab-test/commit/cfe32cf61b73a0d5e9f13e774abde7ff789b1660",
			Author: EventCommitAuthor{
				Name:  "Dmitriy Zaporozhets",
				Email: "dmitriy.zaporozhets@gmail.com",
			},
		},
	}
	assert.Equal(t, expectedEvent, event)
}

func TestJobEventUnmarshal(t *testing.T) {
	t.Parallel()
	jsonObject := loadFixture(t, "testdata/webhooks/job.json")

	var event *JobEvent
	err := json.Unmarshal(jsonObject, &event)
	require.NoError(t, err)

	expectedEvent := &JobEvent{
		ObjectKind:          "build",
		Ref:                 "main",
		Tag:                 false,
		BeforeSHA:           "0000000000000000000000000000000000000000",
		SHA:                 "95d49d1efbd941908580e79d65e4b5ecaf4a8305",
		RetriesCount:        1,
		BuildID:             3580121225,
		BuildName:           "auto_deploy:start",
		BuildStage:          "coordinated:tag",
		BuildStatus:         "success",
		BuildCreatedAt:      "2023-01-10 13:50:02 UTC",
		BuildStartedAt:      "2023-01-10 13:50:05 UTC",
		BuildFinishedAt:     "2023-01-10 13:50:54 UTC",
		BuildCreatedAtISO:   "2023-01-10T13:50:02Z",
		BuildStartedAtISO:   "2023-01-10T13:50:05Z",
		BuildFinishedAtISO:  "2023-01-10T13:50:54Z",
		BuildDuration:       49.503592,
		BuildQueuedDuration: 0.193009,
		BuildAllowFailure:   false,
		BuildFailureReason:  "unknown_failure",
		PipelineID:          743121198,
		Runner: JobEventRunner{
			ID:          12270837,
			Description: "4-blue.shared.runners-manager.gitlab.com/default",
			RunnerType:  "instance_type",
			Active:      true,
			IsShared:    true,
			Tags:        []string{"linux", "docker"},
		},
		ProjectID:   31537070,
		ProjectName: "John Smith / release-tools-fake",
		User: &EventUser{
			ID:        2967854,
			Name:      "John Smith",
			Username:  "jsmithy2",
			AvatarURL: "https://gitlab.com/uploads/-/system/user/avatar/2967852/avatar.png",
			Email:     "john@smith.com",
		},
		Commit: JobEventCommit{
			ID:          743121198,
			Name:        "Build pipeline",
			SHA:         "95d49d1efbd941908580e79d65e4b5ecaf4a8305",
			Message:     "Remove test jobs and add back other jobs",
			AuthorName:  "John Smith",
			AuthorEmail: "john@smith.com",
			AuthorURL:   "https://gitlab.com/jsmithy2",
			Status:      "running",
			Duration:    128,
			StartedAt:   "2023-01-10 13:50:05 UTC",
			FinishedAt:  "2022-10-12 08:09:29 UTC",
		},
		Repository: &Repository{
			Name:              "release-tools-fake",
			Description:       "",
			WebURL:            "",
			AvatarURL:         "",
			GitSSHURL:         "git@gitlab.com:jsmithy2/release-tools-fake.git",
			GitHTTPURL:        "https://gitlab.com/jsmithy2/release-tools-fake.git",
			Namespace:         "",
			Visibility:        "",
			PathWithNamespace: "",
			DefaultBranch:     "",
			Homepage:          "https://gitlab.com/jsmithy2/release-tools-fake",
			URL:               "git@gitlab.com:jsmithy2/release-tools-fake.git",
			SSHURL:            "",
			HTTPURL:           "",
		},
		Project: JobEventProject{
			ID:                1,
			Name:              "Gitlab Test",
			Description:       "Atque in sunt eos similique dolores voluptatem.",
			WebURL:            "http://192.168.64.1:3005/gitlab-org/gitlab-test",
			AvatarURL:         "",
			GitSSHURL:         "git@192.168.64.1:gitlab-org/gitlab-test.git",
			GitHTTPURL:        "http://192.168.64.1:3005/gitlab-org/gitlab-test.git",
			Namespace:         "Gitlab Org",
			VisibilityLevel:   20,
			PathWithNamespace: "gitlab-org/gitlab-test",
			DefaultBranch:     "master",
			CIConfigPath:      "",
		},
		Environment: EventEnvironment{
			Name:           "production",
			Action:         "start",
			DeploymentTier: "production",
		},
		SourcePipeline: EventSourcePipeline{
			Project: EventSourcePipelineProject{
				ID:                41,
				WebURL:            "https://gitlab.example.com/gitlab-org/upstream-project",
				PathWithNamespace: "gitlab-org/upstream-project",
			},
			PipelineID: 30,
			JobID:      3401,
		},
	}

	assert.Equal(t, expectedEvent, event, "event should be equal to the expected one")
}

func TestDeploymentEventUnmarshal(t *testing.T) {
	t.Parallel()
	jsonObject := loadFixture(t, "testdata/webhooks/deployment.json")

	var event *DeploymentEvent
	err := json.Unmarshal(jsonObject, &event)
	require.NoError(t, err)

	expectedEvent := &DeploymentEvent{
		ObjectKind:             "deployment",
		Status:                 "success",
		StatusChangedAt:        "2021-04-28 21:50:00 +0200",
		DeploymentID:           15,
		DeployableID:           796,
		DeployableURL:          "http://10.126.0.2:3000/root/test-deployment-webhooks/-/jobs/796",
		Environment:            "staging",
		EnvironmentSlug:        "staging",
		EnvironmentExternalURL: "https://staging.example.com",
		Project: DeploymentEventProject{
			ID:                30,
			Name:              "test-deployment-webhooks",
			Description:       "",
			WebURL:            "http://10.126.0.2:3000/root/test-deployment-webhooks",
			GitSSHURL:         "ssh://vlad@10.126.0.2:2222/root/test-deployment-webhooks.git",
			GitHTTPURL:        "http://10.126.0.2:3000/root/test-deployment-webhooks.git",
			Namespace:         "User1",
			VisibilityLevel:   0,
			PathWithNamespace: "root/test-deployment-webhooks",
			DefaultBranch:     "master",
			CIConfigPath:      "",
			Homepage:          "http://10.126.0.2:3000/root/test-deployment-webhooks",
			URL:               "ssh://vlad@10.126.0.2:2222/root/test-deployment-webhooks.git",
			SSHURL:            "ssh://vlad@10.126.0.2:2222/root/test-deployment-webhooks.git",
			HTTPURL:           "http://10.126.0.2:3000/root/test-deployment-webhooks.git",
		},
		ShortSHA: "279484c0",
		User: &EventUser{
			ID:        42,
			Name:      "User1",
			Username:  "user1",
			AvatarURL: "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80\u0026d=identicon",
			Email:     "admin@example.com",
		},
		UserURL:     "http://10.126.0.2:3000/root",
		CommitURL:   "http://10.126.0.2:3000/root/test-deployment-webhooks/-/commit/279484c09fbe69ededfced8c1bb6e6d24616b468",
		CommitTitle: "Add new file",
		Ref:         "1.0.0",
	}
	assert.Equal(t, expectedEvent, event)
}

func TestFeatureFlagEventUnmarshal(t *testing.T) {
	t.Parallel()
	jsonObject := loadFixture(t, "testdata/webhooks/feature_flag.json")

	var event *FeatureFlagEvent
	err := json.Unmarshal(jsonObject, &event)
	require.NoError(t, err)

	expectedEvent := &FeatureFlagEvent{
		ObjectKind: "feature_flag",
		Project: FeatureFlagEventProject{
			ID:                1,
			Name:              "Gitlab Test",
			Description:       "Aut reprehenderit ut est.",
			WebURL:            "http://example.com/gitlabhq/gitlab-test",
			GitSSHURL:         "git@example.com:gitlabhq/gitlab-test.git",
			GitHTTPURL:        "http://example.com/gitlabhq/gitlab-test.git",
			Namespace:         "GitlabHQ",
			VisibilityLevel:   20,
			PathWithNamespace: "gitlabhq/gitlab-test",
			DefaultBranch:     "master",
			Homepage:          "http://example.com/gitlabhq/gitlab-test",
			URL:               "http://example.com/gitlabhq/gitlab-test.git",
			SSHURL:            "git@example.com:gitlabhq/gitlab-test.git",
			HTTPURL:           "http://example.com/gitlabhq/gitlab-test.git",
		},
		User: &EventUser{
			ID:        1,
			Name:      "Administrator",
			Username:  "root",
			AvatarURL: "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80\u0026d=identicon",
			Email:     "admin@example.com",
		},
		UserURL: "http://example.com/root",
		ObjectAttributes: FeatureFlagEventObjectAttributes{
			ID:          6,
			Name:        "test-feature-flag",
			Description: "test-feature-flag-description",
			Active:      true,
		},
	}
	assert.Equal(t, expectedEvent, event)
}

func TestGroupResourceAccessTokenEventUnmarshal(t *testing.T) {
	t.Parallel()
	jsonObject := loadFixture(t, "testdata/webhooks/resource_access_token_group.json")

	var event *GroupResourceAccessTokenEvent
	err := json.Unmarshal(jsonObject, &event)
	require.NoError(t, err)

	expiresAt, err := ParseISOTime("2024-01-26")
	require.NoError(t, err)

	expectedEvent := &GroupResourceAccessTokenEvent{
		ObjectKind: "access_token",
		EventName:  "expiring_access_token",
		Group: GroupResourceAccessTokenEventGroup{
			GroupID:   35,
			GroupName: "Twitter",
			GroupPath: "twitter",
			FullPath:  "twitter",
		},
		ObjectAttributes: GroupResourceAccessTokenEventObjectAttributes{
			ID:        25,
			UserID:    90,
			Name:      "acd",
			CreatedAt: "2024-01-24 16:27:40 UTC",
			ExpiresAt: &expiresAt,
		},
	}
	assert.Equal(t, expectedEvent, event)
}

func TestIssueCommentEventUnmarshal(t *testing.T) {
	t.Parallel()
	jsonObject := loadFixture(t, "testdata/webhooks/note_issue.json")

	var event *IssueCommentEvent
	err := json.Unmarshal(jsonObject, &event)
	require.NoError(t, err)

	expectedEvent := &IssueCommentEvent{
		ObjectKind: "note",
		EventType:  "note",
		User: &User{
			ID:        42,
			Name:      "User1",
			Username:  "user1",
			Email:     "user1@example.com",
			AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=40\u0026d=identicon",
		},
		ProjectID: 5,
		Project: IssueCommentEventProject{
			Name:              "Gitlab Test",
			Description:       "Aut reprehenderit ut est.",
			GitSSHURL:         "git@example.com:gitlab-org/gitlab-test.git",
			GitHTTPURL:        "http://example.com/gitlab-org/gitlab-test.git",
			Namespace:         "Gitlab Org",
			PathWithNamespace: "gitlab-org/gitlab-test",
			DefaultBranch:     "master",
			Homepage:          "http://example.com/gitlab-org/gitlab-test",
			URL:               "http://example.com/gitlab-org/gitlab-test.git",
			SSHURL:            "git@example.com:gitlab-org/gitlab-test.git",
			HTTPURL:           "http://example.com/gitlab-org/gitlab-test.git",
			WebURL:            "http://example.com/gitlab-org/gitlab-test",
		},
		Repository: &Repository{
			Name:        "diaspora",
			URL:         "git@example.com:mike/diaspora.git",
			Description: "",
			Homepage:    "http://example.com/mike/diaspora",
		},
		ObjectAttributes: IssueCommentEventObjectAttributes{
			ID:           1241,
			Note:         "Hello world",
			NoteableType: "Issue",
			AuthorID:     1,
			CreatedAt:    "2015-05-17 17:06:40 UTC",
			UpdatedAt:    "2015-05-17 17:06:40 UTC",
			ProjectID:    5,
			NoteableID:   92,
			System:       false,
			Description:  "Hello world",
			Action:       CommentEventActionCreate,
			URL:          "http://example.com/gitlab-org/gitlab-test/issues/17#note_1241",
		},
		Issue: IssueCommentEventIssue{
			ID:                  92,
			IID:                 17,
			ProjectID:           5,
			AuthorID:            1,
			Title:               "test_issue",
			Description:         "test issue",
			State:               "closed",
			CreatedAt:           "2016-01-04T15:31:46.176Z",
			UpdatedAt:           "2016-01-04T15:31:46.176Z",
			TimeEstimate:        3600,
			TotalTimeSpent:      600,
			HumanTotalTimeSpent: "10m",
			HumanTimeEstimate:   "1h",
			AssigneeIDs:         []int64{},
			Labels: []*EventLabel{
				{
					ID:        25,
					Title:     "Afterpod",
					Color:     "#3e8068",
					CreatedAt: "2019-06-05T14:32:20.211Z",
					UpdatedAt: "2019-06-05T14:32:20.211Z",
					Type:      "GroupLabel",
					GroupID:   4,
				},
				{
					ID:        86,
					Title:     "Element",
					Color:     "#231afe",
					ProjectID: 4,
					CreatedAt: "2019-06-05T14:32:20.637Z",
					UpdatedAt: "2019-06-05T14:32:20.637Z",
					Type:      "ProjectLabel",
				},
			},
		},
	}
	assert.Equal(t, expectedEvent, event)
}

func TestIssueEventUnmarshal(t *testing.T) {
	t.Parallel()
	jsonObject := loadFixture(t, "testdata/webhooks/issue.json")

	var event *IssueEvent
	err := json.Unmarshal(jsonObject, &event)
	require.NoError(t, err)

	apiLabel := &EventLabel{
		ID:          206,
		Title:       "API",
		Color:       "#ffffff",
		ProjectID:   14,
		CreatedAt:   "2013-12-03T17:15:43Z",
		UpdatedAt:   "2013-12-03T17:15:43Z",
		Description: "API related issues",
		Type:        "ProjectLabel",
		GroupID:     41,
	}

	expectedEvent := &IssueEvent{
		ObjectKind: "issue",
		EventType:  "issue",
		User: &EventUser{
			ID:        1,
			Name:      "Administrator",
			Username:  "root",
			AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=40\u0026d=identicon",
			Email:     "admin@example.com",
		},
		Project: IssueEventProject{
			ID:                1,
			Name:              "GitLab Test",
			Description:       "Aut reprehenderit ut est.",
			WebURL:            "http://example.com/gitlabhq/gitlab-test",
			GitSSHURL:         "git@example.com:gitlabhq/gitlab-test.git",
			GitHTTPURL:        "http://example.com/gitlabhq/gitlab-test.git",
			Namespace:         "GitlabHQ",
			PathWithNamespace: "gitlabhq/gitlab-test",
			DefaultBranch:     "master",
			Homepage:          "http://example.com/gitlabhq/gitlab-test",
			URL:               "http://example.com/gitlabhq/gitlab-test.git",
			SSHURL:            "git@example.com:gitlabhq/gitlab-test.git",
			HTTPURL:           "http://example.com/gitlabhq/gitlab-test.git",
		},
		Repository: &Repository{
			Name:        "Gitlab Test",
			URL:         "http://example.com/gitlabhq/gitlab-test.git",
			Description: "Aut reprehenderit ut est.",
			Homepage:    "http://example.com/gitlabhq/gitlab-test",
		},
		ObjectAttributes: IssueEventObjectAttributes{
			ID:               301,
			Title:            "New API: create/update/delete file",
			AssigneeIDs:      []int64{51},
			AssigneeID:       51,
			AuthorID:         51,
			ProjectID:        14,
			CreatedAt:        "2013-12-03T17:15:43Z",
			UpdatedAt:        "2013-12-03T17:15:43Z",
			UpdatedByID:      1,
			Description:      "Create new API for manipulations with repository",
			StateID:          StateIDOpen,
			DiscussionLocked: true,
			Weight:           10,
			IID:              23,
			URL:              "http://example.com/diaspora/issues/23",
			State:            "opened",
			Action:           "open",
			Severity:         "high",
			EscalationStatus: "triggered",
			EscalationPolicy: IssueEventObjectAttributesEscalationPolicy{
				ID:   18,
				Name: "Engineering On-call",
			},
			Labels: []*EventLabel{apiLabel},
		},
		Assignee: &EventUser{
			Name:      "User1",
			Username:  "user1",
			AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=40\u0026d=identicon",
		},
		Assignees: &[]EventUser{
			{
				Name:      "User1",
				Username:  "user1",
				AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=40\u0026d=identicon",
			},
		},
		Labels: []*EventLabel{apiLabel},
		Changes: IssueEventChanges{
			UpdatedByID: EventChangesUpdatedByID{
				Current: 1,
			},
			UpdatedAt: EventChangesUpdatedAt{
				Previous: "2017-09-15 16:50:55 UTC",
				Current:  "2017-09-15 16:52:00 UTC",
			},
			ClosedAt: IssueEventChangesClosedAt{
				Previous: "2017-09-15 16:54:55 UTC",
				Current:  "2017-09-15 16:56:00 UTC",
			},
			StateID: EventChangesStateID{
				Previous: StateIDNone,
				Current:  StateIDOpen,
			},
			Labels: EventChangesLabels{
				Previous: []*EventLabel{apiLabel},
				Current: []*EventLabel{
					{
						ID:          205,
						Title:       "Platform",
						Color:       "#123123",
						ProjectID:   14,
						CreatedAt:   "2013-12-03T17:15:43Z",
						UpdatedAt:   "2013-12-03T17:15:43Z",
						Description: "Platform related issues",
						Type:        "ProjectLabel",
						GroupID:     41,
					},
				},
			},
			Description: EventChangesDescription{
				Current: "New description",
			},
			Title: EventChangesTitle{
				Current: "New title",
			},
			TotalTimeSpent: IssueEventChangesTotalTimeSpent{
				Previous: 8100,
				Current:  9900,
			},
		},
	}
	assert.Equal(t, expectedEvent, event)
}

func TestMergeCommentEventUnmarshal(t *testing.T) {
	t.Parallel()
	jsonObject := loadFixture(t, "testdata/webhooks/note_merge_request.json")

	var event *MergeCommentEvent
	err := json.Unmarshal(jsonObject, &event)
	require.NoError(t, err)

	gitlabOrgRepo := &Repository{
		Name:              "Gitlab Test",
		Description:       "Aut reprehenderit ut est.",
		WebURL:            "http://example.com/gitlab-org/gitlab-test",
		GitSSHURL:         "git@example.com:gitlab-org/gitlab-test.git",
		GitHTTPURL:        "http://example.com/gitlab-org/gitlab-test.git",
		Namespace:         "Gitlab Org",
		PathWithNamespace: "gitlab-org/gitlab-test",
		DefaultBranch:     "master",
		Homepage:          "http://example.com/gitlab-org/gitlab-test",
		URL:               "http://example.com/gitlab-org/gitlab-test.git",
		SSHURL:            "git@example.com:gitlab-org/gitlab-test.git",
		HTTPURL:           "http://example.com/gitlab-org/gitlab-test.git",
	}

	expectedEvent := &MergeCommentEvent{
		ObjectKind: "note",
		EventType:  "note",
		User: &EventUser{
			ID:        1,
			Name:      "Administrator",
			Username:  "root",
			AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=40\u0026d=identicon",
			Email:     "admin@example.com",
		},
		ProjectID: 5,
		Project: MergeCommentEventProject{
			ID:                5,
			Name:              "Gitlab Test",
			Description:       "Aut reprehenderit ut est.",
			WebURL:            "http://example.com/gitlab-org/gitlab-test",
			GitSSHURL:         "git@example.com:gitlab-org/gitlab-test.git",
			GitHTTPURL:        "http://example.com/gitlab-org/gitlab-test.git",
			Namespace:         "Gitlab Org",
			PathWithNamespace: "gitlab-org/gitlab-test",
			DefaultBranch:     "master",
			Homepage:          "http://example.com/gitlab-org/gitlab-test",
			URL:               "http://example.com/gitlab-org/gitlab-test.git",
			SSHURL:            "git@example.com:gitlab-org/gitlab-test.git",
			HTTPURL:           "http://example.com/gitlab-org/gitlab-test.git",
		},
		Repository: &Repository{
			Name:        "Gitlab Test",
			URL:         "http://localhost/gitlab-org/gitlab-test.git",
			Description: "Aut reprehenderit ut est.",
			Homepage:    "http://example.com/gitlab-org/gitlab-test",
		},
		ObjectAttributes: MergeCommentEventObjectAttributes{
			ID:           1244,
			Note:         "This MR needs work.",
			NoteableType: "MergeRequest",
			AuthorID:     1,
			CreatedAt:    "2015-05-17 18:21:36 UTC",
			UpdatedAt:    "2015-05-17 18:21:36 UTC",
			ProjectID:    5,
			NoteableID:   7,
			Action:       CommentEventActionCreate,
			URL:          "http://example.com/gitlab-org/gitlab-test/merge_requests/1#note_1244",
		},
		MergeRequest: MergeCommentEventMergeRequest{
			ID:              7,
			TargetBranch:    "markdown",
			SourceBranch:    "master",
			SourceProjectID: 5,
			AuthorID:        8,
			AssigneeID:      28,
			AssigneeIDs:     []int64{28},
			ReviewerIDs:     []int64{13},
			Title:           "Tempora et eos debitis quae laborum et.",
			CreatedAt:       "2015-03-01 20:12:53 UTC",
			UpdatedAt:       "2015-03-21 18:27:27 UTC",
			MilestoneID:     11,
			State:           "opened",
			MergeStatus:     "cannot_be_merged",
			TargetProjectID: 5,
			IID:             1,
			Description:     "Et voluptas corrupti assumenda temporibus. Architecto cum animi eveniet amet asperiores. Vitae numquam voluptate est natus sit et ad id.",
			Labels: []*EventLabel{
				{
					ID:        206,
					Title:     "Afterpod",
					Color:     "#3e8068",
					CreatedAt: "2019-06-05T14:32:20.211Z",
					UpdatedAt: "2019-06-05T14:32:20.211Z",
					Type:      "GroupLabel",
					GroupID:   4,
				},
				{
					ID:        86,
					Title:     "Element",
					Color:     "#231afe",
					ProjectID: 4,
					CreatedAt: "2019-06-05T14:32:20.637Z",
					UpdatedAt: "2019-06-05T14:32:20.637Z",
					Type:      "ProjectLabel",
				},
			},
			Source: gitlabOrgRepo,
			Target: gitlabOrgRepo,
			LastCommit: EventMergeRequestLastCommit{
				ID:      "562e173be03b8ff2efb05345d12df18815438a4b",
				Message: "Merge branch 'another-branch' into 'master'\n\nCheck in this test\n",
				Title:   "Merge branch 'another-branch' into 'master'",
				Timestamp: func() *time.Time {
					ts, err := time.Parse(time.RFC3339, "2015-04-08T21:00:25-07:00")
					require.NoError(t, err)
					return &ts
				}(),
				URL: "http://example.com/gitlab-org/gitlab-test/commit/562e173be03b8ff2efb05345d12df18815438a4b",
				Author: EventCommitAuthor{
					Name:  "John Smith",
					Email: "john@example.com",
				},
			},
			Assignee: &EventUser{
				Name:      "User1",
				Username:  "user1",
				AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=40\u0026d=identicon",
			},
			DetailedMergeStatus: "checking",
			URL:                 "http://example.com/gitlab-org/gitlab-test/-/merge_requests/1",
		},
	}
	assert.Equal(t, expectedEvent, event)
}

func TestMergeEventUnmarshal(t *testing.T) {
	t.Parallel()
	jsonObject := loadFixture(t, "testdata/webhooks/merge_request.json")

	var event *MergeEvent
	err := json.Unmarshal(jsonObject, &event)
	require.NoError(t, err)

	awesomeRepo := &Repository{
		Name:              "Awesome Project",
		Description:       "Aut reprehenderit ut est.",
		WebURL:            "http://example.com/awesome_space/awesome_project",
		GitSSHURL:         "git@example.com:awesome_space/awesome_project.git",
		GitHTTPURL:        "http://example.com/awesome_space/awesome_project.git",
		Namespace:         "Awesome Space",
		PathWithNamespace: "awesome_space/awesome_project",
		DefaultBranch:     "master",
		Homepage:          "http://example.com/awesome_space/awesome_project",
		URL:               "http://example.com/awesome_space/awesome_project.git",
		SSHURL:            "git@example.com:awesome_space/awesome_project.git",
		HTTPURL:           "http://example.com/awesome_space/awesome_project.git",
	}

	apiLabel := &EventLabel{
		ID:          206,
		Title:       "API",
		Color:       "#ffffff",
		ProjectID:   14,
		CreatedAt:   "2013-12-03T17:15:43Z",
		UpdatedAt:   "2013-12-03T17:15:43Z",
		Description: "API related issues",
		Type:        "ProjectLabel",
		GroupID:     41,
	}

	user1 := &EventUser{
		ID:        1,
		Name:      "User1",
		Username:  "user1",
		AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=40\u0026d=identicon",
	}

	expectedEvent := &MergeEvent{
		ObjectKind: "merge_request",
		EventType:  "merge_request",
		User: &EventUser{
			ID:        1,
			Name:      "User1",
			Username:  "user1",
			AvatarURL: "http://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=40\u0026d=identicon",
			Email:     "user1@example.com",
		},
		Project: MergeEventProject{
			ID:                1,
			Name:              "Gitlab Test",
			Description:       "Aut reprehenderit ut est.",
			WebURL:            "http://example.com/gitlabhq/gitlab-test",
			GitSSHURL:         "git@example.com:gitlabhq/gitlab-test.git",
			GitHTTPURL:        "http://example.com/gitlabhq/gitlab-test.git",
			Namespace:         "GitlabHQ",
			PathWithNamespace: "gitlabhq/gitlab-test",
			DefaultBranch:     "master",
			Homepage:          "http://example.com/gitlabhq/gitlab-test",
			URL:               "http://example.com/gitlabhq/gitlab-test.git",
			SSHURL:            "git@example.com:gitlabhq/gitlab-test.git",
			HTTPURL:           "http://example.com/gitlabhq/gitlab-test.git",
		},
		Repository: &Repository{
			Name:        "Gitlab Test",
			URL:         "http://example.com/gitlabhq/gitlab-test.git",
			Description: "Aut reprehenderit ut est.",
			Homepage:    "http://example.com/gitlabhq/gitlab-test",
		},
		ObjectAttributes: MergeEventObjectAttributes{
			ID:                          99,
			TargetBranch:                "master",
			SourceBranch:                "ms-viewport",
			SourceProjectID:             14,
			AuthorID:                    51,
			AssigneeID:                  1,
			AssigneeIDs:                 []int64{1},
			ReviewerIDs:                 []int64{1},
			Title:                       "MS-Viewport",
			CreatedAt:                   "2013-12-03T17:23:34Z",
			UpdatedAt:                   "2013-12-03T17:23:34Z",
			LastEditedAt:                "2023-03-27 00:03:05 UTC",
			LastEditedByID:              51,
			StateID:                     StateIDOpen,
			State:                       "opened",
			BlockingDiscussionsResolved: true,
			FirstContribution:           true,
			MergeStatus:                 "unchecked",
			TargetProjectID:             14,
			IID:                         1,
			URL:                         "http://example.com/diaspora/merge_requests/1",
			Source:                      awesomeRepo,
			Target:                      awesomeRepo,
			LastCommit: EventMergeRequestLastCommit{
				ID:      "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
				Message: "fixed readme",
				Title:   "MR Title",
				Timestamp: func() *time.Time {
					ts, err := time.Parse(time.RFC3339, "2012-01-03T23:36:29+02:00")
					require.NoError(t, err)
					return &ts
				}(),
				URL: "http://example.com/awesome_space/awesome_project/commits/da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
				Author: EventCommitAuthor{
					Name:  "GitLab dev user",
					Email: "gitlabdev@dv6700.(none)",
				},
			},
			HumanTotalTimeSpent: "30m",
			HumanTimeChange:     "30m",
			HumanTimeEstimate:   "1h",
			Labels:              []*EventLabel{apiLabel},
			Action:              "open",
			DetailedMergeStatus: "mergeable",
			System:              true,
			SystemAction:        "code_owner_approvals_reset_on_push",
		},
		Labels: []*EventLabel{apiLabel},
		Changes: MergeEventChanges{
			UpdatedByID: EventChangesUpdatedByID{
				Current: 1,
			},
			UpdatedAt: EventChangesUpdatedAt{
				Previous: "2017-09-15 16:50:55 UTC",
				Current:  "2017-09-15 16:52:00 UTC",
			},
			StateID: EventChangesStateID{
				Previous: StateIDLocked,
				Current:  StateIDMerged,
			},
			Labels: EventChangesLabels{
				Previous: []*EventLabel{apiLabel},
				Current: []*EventLabel{
					{
						ID:          205,
						Title:       "Platform",
						Color:       "#123123",
						ProjectID:   14,
						CreatedAt:   "2013-12-03T17:15:43Z",
						UpdatedAt:   "2013-12-03T17:15:43Z",
						Description: "Platform related issues",
						Type:        "ProjectLabel",
						GroupID:     41,
					},
				},
			},
		},
		Assignees: []*EventUser{user1},
		Reviewers: []*EventUser{user1},
	}
	assert.Equal(t, expectedEvent, event)
}

func TestMemberEventUnmarshal(t *testing.T) {
	t.Parallel()
	jsonObject := loadFixture(t, "testdata/webhooks/member.json")

	var event *MemberEvent
	err := json.Unmarshal(jsonObject, &event)
	require.NoError(t, err)

	createdAt, err := time.Parse(time.RFC3339, "2020-12-11T04:57:22Z")
	require.NoError(t, err)
	updatedAt, err := time.Parse(time.RFC3339, "2020-12-11T04:57:22Z")
	require.NoError(t, err)
	expiresAt, err := time.Parse(time.RFC3339, "2020-12-14T00:00:00Z")
	require.NoError(t, err)

	expectedEvent := &MemberEvent{
		CreatedAt:    &createdAt,
		UpdatedAt:    &updatedAt,
		GroupName:    "webhook-test",
		GroupPath:    "webhook-test",
		GroupID:      100,
		UserUsername: "user1",
		UserName:     "User1",
		UserEmail:    "testuser@webhooktest.com",
		UserID:       64,
		GroupAccess:  "Guest",
		ExpiresAt:    &expiresAt,
		EventName:    "user_add_to_group",
	}
	assert.Equal(t, expectedEvent, event)
}

func TestMergeEventUnmarshalFromGroup(t *testing.T) {
	t.Parallel()
	jsonObject := loadFixture(t, "testdata/webhooks/group_merge_request.json")

	var event *MergeEvent
	err := json.Unmarshal(jsonObject, &event)
	require.NoError(t, err)

	exmProjectRepo := &Repository{
		Name:              "example-project",
		WebURL:            "http://example.com/exm-namespace/example-project",
		GitSSHURL:         "git@example.com:exm-namespace/example-project.git",
		GitHTTPURL:        "http://example.com/exm-namespace/example-project.git",
		Namespace:         "exm-namespace",
		Visibility:        PublicVisibility,
		PathWithNamespace: "exm-namespace/example-project",
		DefaultBranch:     "master",
		Homepage:          "http://example.com/exm-namespace/example-project",
		URL:               "git@example.com:exm-namespace/example-project.git",
		SSHURL:            "git@example.com:exm-namespace/example-project.git",
		HTTPURL:           "http://example.com/exm-namespace/example-project.git",
	}

	expectedEvent := &MergeEvent{
		ObjectKind: "merge_request",
		User: &EventUser{
			ID:        42,
			Name:      "User1",
			Username:  "user1",
			AvatarURL: "http://www.gravatar.com/avatar/d22738dc40839e3d95fca77ca3eac067?s=80\u0026d=identicon",
			Email:     "user1@mail.com",
		},
		Project: MergeEventProject{
			Name:              "example-project",
			WebURL:            "http://example.com/exm-namespace/example-project",
			GitSSHURL:         "git@example.com:exm-namespace/example-project.git",
			GitHTTPURL:        "http://example.com/exm-namespace/example-project.git",
			Namespace:         "exm-namespace",
			PathWithNamespace: "exm-namespace/example-project",
			DefaultBranch:     "master",
			Homepage:          "http://example.com/exm-namespace/example-project",
			URL:               "git@example.com:exm-namespace/example-project.git",
			SSHURL:            "git@example.com:exm-namespace/example-project.git",
			HTTPURL:           "http://example.com/exm-namespace/example-project.git",
			Visibility:        PublicVisibility,
		},
		ObjectAttributes: MergeEventObjectAttributes{
			ID:             15917,
			MergeParams:    &MergeParams{},
			MergeCommitSHA: "ac3ca1559bc39abf963586372eff7f8fdded646e",
			Source:         exmProjectRepo,
			Target:         exmProjectRepo,
			LastCommit: EventMergeRequestLastCommit{
				ID:      "61b6a0d35dbaf915760233b637622e383d3cc9ec",
				Message: "commit message",
				Timestamp: func() *time.Time {
					ts, err := time.Parse(time.RFC3339, "2016-12-01T15:07:53+02:00")
					require.NoError(t, err)
					return &ts
				}(),
				URL: "http://example.com/exm-namespace/example-project/commit/61b6a0d35dbaf915760233b637622e383d3cc9ec",
				Author: EventCommitAuthor{
					Name:  "Test User",
					Email: "test.user@mail.com",
				},
			},
			URL:    "http://example.com/exm-namespace/example-project/merge_requests/1402",
			Action: "merge",
		},
		Repository: &Repository{
			Name:        "example-project",
			URL:         "git@example.com:exm-namespace/example-project.git",
			Description: "",
			Homepage:    "http://example.com/exm-namespace/example-project",
		},
	}
	assert.Equal(t, expectedEvent, event)
}

func TestPipelineEventUnmarshal(t *testing.T) {
	t.Parallel()
	jsonObject := loadFixture(t, "testdata/webhooks/pipeline.json")

	var event *PipelineEvent
	err := json.Unmarshal(jsonObject, &event)
	require.NoError(t, err)

	user1 := &EventUser{
		ID:        42,
		Name:      "User1",
		Username:  "user1",
		Email:     "user1@example.com",
		AvatarURL: "http://www.gravatar.com/avatar/e32bd13e2add097461cb96824b7a829c?s=80\u0026d=identicon",
	}

	expectedEvent := &PipelineEvent{
		ObjectKind: "pipeline",
		ObjectAttributes: PipelineEventObjectAttributes{
			ID:             31,
			IID:            123,
			Ref:            "master",
			SHA:            "bcbb5ec396a2c0f828686f14fac9b80b780504f2",
			BeforeSHA:      "bcbb5ec396a2c0f828686f14fac9b80b780504f2",
			Source:         "merge_request_event",
			Status:         "success",
			DetailedStatus: "passed",
			Stages:         []string{"build", "test", "deploy"},
			CreatedAt:      "2016-08-12 15:23:28 UTC",
			FinishedAt:     "2016-08-12 15:26:29 UTC",
			Duration:       63,
			QueuedDuration: 12,
			Variables: []PipelineEventObjectAttributesVariable{
				{Key: "NESTOR_PROD_ENVIRONMENT", Value: "us-west-1"},
			},
		},
		MergeRequest: PipelineEventMergeRequest{
			ID:                  1,
			IID:                 1,
			Title:               "Test",
			SourceBranch:        "test",
			SourceProjectID:     1,
			TargetBranch:        "master",
			TargetProjectID:     1,
			State:               "opened",
			MergeRequestStatus:  "can_be_merged",
			DetailedMergeStatus: "mergeable",
			URL:                 "http://192.168.64.1:3005/gitlab-org/gitlab-test/merge_requests/1",
		},
		User: user1,
		Project: PipelineEventProject{
			ID:                1,
			Name:              "Gitlab Test",
			Description:       "Atque in sunt eos similique dolores voluptatem.",
			WebURL:            "http://192.168.64.1:3005/gitlab-org/gitlab-test",
			GitSSHURL:         "git@192.168.64.1:gitlab-org/gitlab-test.git",
			GitHTTPURL:        "http://192.168.64.1:3005/gitlab-org/gitlab-test.git",
			Namespace:         "Gitlab Org",
			PathWithNamespace: "gitlab-org/gitlab-test",
			DefaultBranch:     "master",
		},
		Commit: PipelineEventCommit{
			ID:      "bcbb5ec396a2c0f828686f14fac9b80b780504f2",
			Message: "test\n",
			Timestamp: func() *time.Time {
				ts, err := time.Parse(time.RFC3339, "2016-08-12T17:23:21+02:00")
				require.NoError(t, err)
				return &ts
			}(),
			URL: "http://example.com/gitlab-org/gitlab-test/commit/bcbb5ec396a2c0f828686f14fac9b80b780504f2",
			Author: EventCommitAuthor{
				Name:  "User",
				Email: "user@gitlab.com",
			},
		},
		SourcePipeline: EventSourcePipeline{
			Project: EventSourcePipelineProject{
				ID:                41,
				WebURL:            "https://gitlab.example.com/gitlab-org/upstream-project",
				PathWithNamespace: "gitlab-org/upstream-project",
			},
			PipelineID: 30,
			JobID:      3401,
		},
		Builds: []PipelineEventBuild{
			{
				ID:             380,
				Stage:          "deploy",
				Name:           "production",
				Status:         "skipped",
				CreatedAt:      "2016-08-12 15:23:28 UTC",
				Duration:       17.1,
				QueuedDuration: 3.5,
				FailureReason:  "script_failure",
				When:           "manual",
				Manual:         true,
				AllowFailure:   true,
				User:           user1,
				Runner: PipelineEventBuildRunner{
					ID:          42,
					Description: "shared-runners-manager-1.gitlab.com",
					RunnerType:  "instance_type",
					Active:      true,
					IsShared:    true,
					Tags:        []string{"docker", "gce"},
				},
				Environment: EventEnvironment{
					Name:           "production",
					Action:         "start",
					DeploymentTier: "production",
				},
			},
			{
				ID:             377,
				Stage:          "test",
				Name:           "test-image",
				Status:         "success",
				CreatedAt:      "2016-08-12 15:23:28 UTC",
				StartedAt:      "2016-08-12 15:26:12 UTC",
				Duration:       17.0,
				QueuedDuration: 196.0,
				When:           "on_success",
				User:           user1,
				Runner: PipelineEventBuildRunner{
					ID:          380987,
					Description: "shared-runners-manager-6.gitlab.com",
					Active:      true,
					IsShared:    true,
				},
			},
			{
				ID:             378,
				Stage:          "test",
				Name:           "test-build",
				Status:         "success",
				CreatedAt:      "2016-08-12 15:23:28 UTC",
				StartedAt:      "2016-08-12 15:26:12 UTC",
				FinishedAt:     "2016-08-12 15:26:29 UTC",
				Duration:       17.0,
				QueuedDuration: 196.0,
				When:           "on_success",
				User:           user1,
				Runner: PipelineEventBuildRunner{
					ID:          380987,
					Description: "shared-runners-manager-6.gitlab.com",
					Active:      true,
					IsShared:    true,
				},
			},
			{
				ID:             376,
				Stage:          "build",
				Name:           "build-image",
				Status:         "success",
				CreatedAt:      "2016-08-12 15:23:28 UTC",
				StartedAt:      "2016-08-12 15:24:56 UTC",
				FinishedAt:     "2016-08-12 15:25:26 UTC",
				Duration:       17.0,
				QueuedDuration: 196.0,
				When:           "on_success",
				User:           user1,
				Runner: PipelineEventBuildRunner{
					ID:          380987,
					Description: "shared-runners-manager-6.gitlab.com",
					Active:      true,
					IsShared:    true,
				},
			},
			{
				ID:             379,
				Stage:          "deploy",
				Name:           "staging",
				Status:         "created",
				CreatedAt:      "2016-08-12 15:23:28 UTC",
				Duration:       17.0,
				QueuedDuration: 196.0,
				When:           "on_success",
				User:           user1,
			},
		},
	}
	assert.Equal(t, expectedEvent, event)
}

func TestProjectResourceAccessTokenEventUnmarshal(t *testing.T) {
	t.Parallel()
	jsonObject := loadFixture(t, "testdata/webhooks/resource_access_token_project.json")
	var event *ProjectResourceAccessTokenEvent
	err := json.Unmarshal(jsonObject, &event)
	require.NoError(t, err)

	require.NotNil(t, event)

	expiresAt, err := ParseISOTime("2024-01-26")
	require.NoError(t, err)

	expected := &ProjectResourceAccessTokenEvent{
		ObjectKind: "access_token",
		EventName:  "expiring_access_token",
	}

	expected.ObjectAttributes.ID = 25
	expected.ObjectAttributes.UserID = 90
	expected.ObjectAttributes.Name = "acd"
	expected.ObjectAttributes.CreatedAt = "2024-01-24 16:27:40 UTC"
	expected.ObjectAttributes.ExpiresAt = &expiresAt

	expected.Project.ID = 7
	expected.Project.Name = "Flight"
	expected.Project.Description = "Eum dolore maxime atque reprehenderit voluptatem."
	expected.Project.WebURL = "https://example.com/flightjs/Flight"
	expected.Project.AvatarURL = ""
	expected.Project.GitSSHURL = "ssh://git@example.com/flightjs/Flight.git"
	expected.Project.GitHTTPURL = "https://example.com/flightjs/Flight.git"
	expected.Project.Namespace = "Flightjs"
	expected.Project.VisibilityLevel = 0
	expected.Project.PathWithNamespace = "flightjs/Flight"
	expected.Project.DefaultBranch = "master"
	expected.Project.CIConfigPath = ""
	expected.Project.Homepage = "https://example.com/flightjs/Flight"
	expected.Project.URL = "ssh://git@example.com/flightjs/Flight.git"
	expected.Project.SSHURL = "ssh://git@example.com/flightjs/Flight.git"
	expected.Project.HTTPURL = "https://example.com/flightjs/Flight.git"

	assert.Equal(t, expected, event)
}

func TestPushEventUnmarshal(t *testing.T) {
	t.Parallel()
	jsonObject := loadFixture(t, "testdata/webhooks/push.json")
	var event *PushEvent
	err := json.Unmarshal(jsonObject, &event)
	require.NoError(t, err)

	expectedEvent := &PushEvent{
		ObjectKind:   "push",
		EventName:    "push",
		Before:       "95790bf891e76fee5e1747ab589903a6a1f80f22",
		After:        "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
		Ref:          "refs/heads/master",
		RefProtected: true,
		CheckoutSHA:  "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
		UserID:       4,
		UserName:     "John Smith",
		UserUsername: "jsmith",
		UserEmail:    "john@example.com",
		UserAvatar:   "https://s.gravatar.com/avatar/d4c74594d841139328695756648b6bd6?s=8://s.gravatar.com/avatar/d4c74594d841139328695756648b6bd6?s=80",
		ProjectID:    15,
		Project: PushEventProject{
			ID:                15,
			Name:              "Diaspora",
			WebURL:            "http://example.com/mike/diaspora",
			GitSSHURL:         "git@example.com:mike/diaspora.git",
			GitHTTPURL:        "http://example.com/mike/diaspora.git",
			Namespace:         "Mike",
			PathWithNamespace: "mike/diaspora",
			DefaultBranch:     "master",
			Homepage:          "http://example.com/mike/diaspora",
			URL:               "git@example.com:mike/diaspora.git",
			SSHURL:            "git@example.com:mike/diaspora.git",
			HTTPURL:           "http://example.com/mike/diaspora.git",
		},
		Repository: &Repository{
			Name:        "Diaspora",
			URL:         "git@example.com:mike/diaspora.git",
			Description: "",
			Homepage:    "http://example.com/mike/diaspora",
			GitHTTPURL:  "http://example.com/mike/diaspora.git",
			GitSSHURL:   "git@example.com:mike/diaspora.git",
		},
		Commits: []*PushEventCommit{
			{
				ID:      "b6568db1bc1dcd7f8b4d5a946b0b91f9dacd7327",
				Message: "Merge branch 'some-feature' into 'master'\n\nRelease v1.0.0\n\nSee merge request jsmith/example!1",
				Title:   "Merge branch 'some-feature' into 'master'",
				Timestamp: func() *time.Time {
					ts, err := time.Parse(time.RFC3339, "2011-12-12T14:27:31+02:00")
					require.NoError(t, err)
					return &ts
				}(),
				URL: "http://example.com/mike/diaspora/commit/b6568db1bc1dcd7f8b4d5a946b0b91f9dacd7327",
				Author: EventCommitAuthor{
					Name:  "Jordi Mallach",
					Email: "jordi@softcatala.org",
				},
				Added:    []string{"CHANGELOG"},
				Modified: []string{"app/controller/application.rb"},
				Removed:  []string{},
			},
			{
				ID:      "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
				Message: "fixed readme\n",
				Title:   "fixed readme",
				Timestamp: func() *time.Time {
					ts, err := time.Parse(time.RFC3339, "2012-01-03T23:36:29+02:00")
					require.NoError(t, err)
					return &ts
				}(),
				URL: "http://example.com/mike/diaspora/commit/da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
				Author: EventCommitAuthor{
					Name:  "GitLab dev user",
					Email: "gitlabdev@dv6700.(none)",
				},
				Added:    []string{"CHANGELOG"},
				Modified: []string{"app/controller/application.rb"},
				Removed:  []string{},
			},
		},
		TotalCommitsCount: 4,
	}
	assert.Equal(t, expectedEvent, event)
}

func TestReleaseEventUnmarshal(t *testing.T) {
	t.Parallel()
	jsonObject := loadFixture(t, "testdata/webhooks/release.json")

	var event *ReleaseEvent
	err := json.Unmarshal(jsonObject, &event)
	require.NoError(t, err)

	avatarURL := "https://www.gravatar.com/avatar/e64c7d89f26bd1972efa854d13d7dd61?s=80&d=identicon"

	expectedEvent := &ReleaseEvent{
		ID:          8273642,
		CreatedAt:   "2021-02-25 21:23:34 UTC",
		Description: "Release!",
		Name:        "1.0.0",
		Tag:         "1.0.0",
		ReleasedAt:  "2021-02-25 21:23:34 UTC",
		ObjectKind:  "release",
		Project: ReleaseEventProject{
			ID:                327622,
			Name:              "Project Name",
			Description:       "",
			WebURL:            "http://example.com/exm-namespace/example-project",
			AvatarURL:         &avatarURL,
			GitSSHURL:         "git@gitlab.com:exm-namespace/example-project.git",
			GitHTTPURL:        "http://example.com/exm-namespace/example-project.git",
			Namespace:         "exm-namespace",
			VisibilityLevel:   0,
			PathWithNamespace: "exm-namespace/example-project",
			DefaultBranch:     "master",
			CIConfigPath:      "",
			Homepage:          "http://example.com/exm-namespace/example-project",
			URL:               "git@gitlab.com:exm-namespace/example-project.git",
			SSHURL:            "git@gitlab.com:exm-namespace/example-project.git",
			HTTPURL:           "http://example.com/exm-namespace/example-project.git",
		},
		URL:    "http://example.com/exm-namespace/example-project/-/releases/1.0.0",
		Action: "create",
		Assets: ReleaseEventAssets{
			Count: 4,
			Links: []ReleaseEventAssetsLink{
				{
					ID:       1,
					External: true,
					LinkType: "other",
					Name:     "Changelog",
					URL:      "https://example.net/changelog",
				},
			},
			Sources: []ReleaseEventAssetsSource{
				{
					Format: "zip",
					URL:    "http://example.com/exm-namespace/example-project/-/archive/1.0.0/example-project-1.0.0.zip",
				},
				{
					Format: "tar.gz",
					URL:    "http://example.com/exm-namespace/example-project/-/archive/1.0.0/example-project-1.0.0.tar.gz",
				},
				{
					Format: "tar.bz2",
					URL:    "http://example.com/exm-namespace/example-project/-/archive/1.0.0/example-project-1.0.0.tar.bz2",
				},
				{
					Format: "tar",
					URL:    "http://example.com/exm-namespace/example-project/-/archive/1.0.0/example-project-1.0.0.tar",
				},
			},
		},
		Commit: ReleaseEventCommit{
			ID:        "2626dbdb936782b5c54816b1c6d45b1279303c6d",
			Message:   "Merge branch 'example-branch' into 'master'\n\nCheck in this test",
			Title:     "Merge branch 'example-branch' into 'master'",
			Timestamp: "2021-02-25T21:21:58+00:00",
			URL:       "http://example.com/exm-namespace/example-project/-/commit/2626dbdb936782b5c54816b1c6d45b1279303c6d",
			Author: EventCommitAuthor{
				Name:  "User",
				Email: "user@gitlab.com",
			},
		},
	}
	assert.Equal(t, expectedEvent, event)
}

func TestSubGroupEventUnmarshal(t *testing.T) {
	t.Parallel()
	jsonObject := loadFixture(t, "testdata/webhooks/subgroup.json")

	var event *SubGroupEvent
	err := json.Unmarshal(jsonObject, &event)
	require.NoError(t, err)

	createdAt, err := time.Parse(time.RFC3339, "2022-01-24T14:23:59Z")
	require.NoError(t, err)
	updatedAt, err := time.Parse(time.RFC3339, "2022-01-24T14:23:59Z")
	require.NoError(t, err)

	expectedEvent := &SubGroupEvent{
		CreatedAt:      &createdAt,
		UpdatedAt:      &updatedAt,
		EventName:      "subgroup_create",
		Name:           "SubGroup 1",
		Path:           "subgroup-1",
		FullPath:       "group-1/subgroup-1",
		GroupID:        2,
		ParentGroupID:  1,
		ParentName:     "Group 1",
		ParentPath:     "group-1",
		ParentFullPath: "group-1",
	}
	assert.Equal(t, expectedEvent, event)
}

func TestTagEventUnmarshal(t *testing.T) {
	t.Parallel()
	jsonObject := loadFixture(t, "testdata/webhooks/tag_push.json")
	var event *TagEvent
	err := json.Unmarshal(jsonObject, &event)
	require.NoError(t, err)

	expectedEvent := &TagEvent{
		ObjectKind:   "tag_push",
		EventName:    "tag_push",
		Before:       "0000000000000000000000000000000000000000",
		After:        "82b3d5ae55f7080f1e6022629cdb57bfae7cccc7",
		Ref:          "refs/tags/v1.0.0",
		CheckoutSHA:  "82b3d5ae55f7080f1e6022629cdb57bfae7cccc7",
		UserID:       1,
		UserName:     "John Smith",
		UserUsername: "jsmith",
		UserAvatar:   "https://s.gravatar.com/avatar/d4c74594d841139328695756648b6bd6?s=8://s.gravatar.com/avatar/d4c74594d841139328695756648b6bd6?s=80",
		ProjectID:    1,
		Project: TagEventProject{
			ID:                1,
			Name:              "Example",
			WebURL:            "http://example.com/jsmith/example",
			GitSSHURL:         "git@example.com:jsmith/example.git",
			GitHTTPURL:        "http://example.com/jsmith/example.git",
			Namespace:         "Jsmith",
			PathWithNamespace: "jsmith/example",
			DefaultBranch:     "master",
			Homepage:          "http://example.com/jsmith/example",
			URL:               "git@example.com:jsmith/example.git",
			SSHURL:            "git@example.com:jsmith/example.git",
			HTTPURL:           "http://example.com/jsmith/example.git",
		},
		Repository: &Repository{
			Name:        "Example",
			URL:         "ssh://git@example.com/jsmith/example.git",
			Description: "",
			Homepage:    "http://example.com/jsmith/example",
			GitHTTPURL:  "http://example.com/jsmith/example.git",
			GitSSHURL:   "git@example.com:jsmith/example.git",
		},
		Commits: []*TagEventCommit{
			{
				ID:      "82b3d5ae55f7080f1e6022629cdb57bfae7cccc7",
				Message: "Merge branch 'some-feature' into 'master'\n\nRelease v1.0.0\n\nSee merge request jsmith/example!1",
				Title:   "Merge branch 'some-feature' into 'master'",
				Timestamp: func() *time.Time {
					ts, err := time.Parse(time.RFC3339, "2012-01-03T23:36:29+02:00")
					require.NoError(t, err)
					return &ts
				}(),
				URL: "http://example.com/jsmith/example/commit/82b3d5ae55f7080f1e6022629cdb57bfae7cccc7",
				Author: EventCommitAuthor{
					Name:  "John Smith",
					Email: "johnsmith@example.com",
				},
				Added:    []string{"CHANGELOG"},
				Modified: []string{"UPGRADE.md"},
				Removed:  []string{},
			},
		},
		TotalCommitsCount: 1,
	}
	assert.Equal(t, expectedEvent, event)
}

func TestSnippetCommentEventUnmarshal(t *testing.T) {
	t.Parallel()

	// GIVEN a snippet comment event JSON payload
	jsonObject := loadFixture(t, "testdata/webhooks/note_snippet.json")

	// WHEN the JSON is unmarshaled
	var event *SnippetCommentEvent
	err := json.Unmarshal(jsonObject, &event)

	// THEN no error should occur and the event should be populated correctly
	assert.NoError(t, err)
	assert.NotNil(t, event)

	assert.Equal(t, "note", event.ObjectKind)
	assert.Equal(t, "note", event.EventType)
	assert.Equal(t, int64(5), event.ProjectID)

	// User assertions
	assert.Equal(t, int64(42), event.User.ID)
	assert.Equal(t, "User1", event.User.Name)
	assert.Equal(t, "user1", event.User.Username)

	// Project assertions
	assert.Equal(t, "Gitlab Test", event.Project.Name)
	assert.Equal(t, "gitlab-org/gitlab-test", event.Project.PathWithNamespace)

	// Repository assertions
	assert.Equal(t, "Gitlab Test", event.Repository.Name)

	// Object attributes assertions
	assert.Equal(t, int64(1245), event.ObjectAttributes.ID)
	assert.Equal(t, "Is this snippet doing what it's supposed to be doing?", event.ObjectAttributes.Note)
	assert.Equal(t, "Snippet", event.ObjectAttributes.NoteableType)
	assert.Equal(t, int64(1), event.ObjectAttributes.AuthorID)
	assert.Equal(t, int64(53), event.ObjectAttributes.NoteableID)
	assert.Equal(t, CommentEventActionCreate, event.ObjectAttributes.Action)
	assert.Equal(t, "http://example.com/gitlab-org/gitlab-test/snippets/53#note_1245", event.ObjectAttributes.URL)

	// Snippet assertions
	assert.NotNil(t, event.Snippet)
	assert.Equal(t, int64(53), event.Snippet.ID)
	assert.Equal(t, "test", event.Snippet.Title)
	assert.Equal(t, "puts 'Hello world'", event.Snippet.Content)
	assert.Equal(t, int64(1), event.Snippet.AuthorID)
	assert.Equal(t, int64(5), event.Snippet.ProjectID)
	assert.Equal(t, "test.rb", event.Snippet.Filename)
	assert.Equal(t, "ProjectSnippet", event.Snippet.Type)
	assert.Equal(t, int64(0), event.Snippet.VisibilityLevel)
	assert.Equal(t, "Prints 'Hello world'", event.Snippet.Description)
	assert.False(t, event.Snippet.Secret)
	assert.False(t, event.Snippet.RepositoryReadOnly)
}

func TestWikiPageEventUnmarshal(t *testing.T) {
	t.Parallel()

	// GIVEN a wiki page event JSON payload
	jsonObject := loadFixture(t, "testdata/webhooks/wiki_page.json")

	// WHEN the JSON is unmarshaled
	var event *WikiPageEvent
	err := json.Unmarshal(jsonObject, &event)

	// THEN no error should occur and the event should be populated correctly
	assert.NoError(t, err)
	assert.NotNil(t, event)

	assert.Equal(t, "wiki_page", event.ObjectKind)

	// User assertions
	assert.Equal(t, "User1", event.User.Name)
	assert.Equal(t, "user1", event.User.Username)

	// Project assertions
	assert.Equal(t, "awesome-project", event.Project.Name)
	assert.Equal(t, "This is awesome", event.Project.Description)
	assert.Equal(t, "root/awesome-project", event.Project.PathWithNamespace)
	assert.Equal(t, "master", event.Project.DefaultBranch)

	// Wiki assertions
	assert.Equal(t, "http://example.com/root/awesome-project/wikis/home", event.Wiki.WebURL)
	assert.Equal(t, "git@example.com:root/awesome-project.wiki.git", event.Wiki.GitSSHURL)
	assert.Equal(t, "http://example.com/root/awesome-project.wiki.git", event.Wiki.GitHTTPURL)
	assert.Equal(t, "root/awesome-project.wiki", event.Wiki.PathWithNamespace)
	assert.Equal(t, "master", event.Wiki.DefaultBranch)

	// Object attributes assertions
	assert.Equal(t, "Awesome", event.ObjectAttributes.Title)
	assert.Equal(t, "awesome content goes here", event.ObjectAttributes.Content)
	assert.Equal(t, "markdown", event.ObjectAttributes.Format)
	assert.Equal(t, "adding an awesome page to the wiki", event.ObjectAttributes.Message)
	assert.Equal(t, "awesome", event.ObjectAttributes.Slug)
	assert.Equal(t, "http://example.com/root/awesome-project/wikis/awesome", event.ObjectAttributes.URL)
	assert.Equal(t, "create", event.ObjectAttributes.Action)
	assert.Equal(t, "http://example.com/root/awesome-project/wikis/awesome/diff", event.ObjectAttributes.DiffURL)
}

func TestEmojiEventUnmarshal(t *testing.T) {
	t.Parallel()

	// GIVEN an emoji event JSON payload
	jsonObject := loadFixture(t, "testdata/webhooks/emoji_issue.json")

	// WHEN the JSON is unmarshaled
	var event *EmojiEvent
	err := json.Unmarshal(jsonObject, &event)

	// THEN no error should occur and the event should be populated correctly
	assert.NoError(t, err)
	assert.NotNil(t, event)

	assert.Equal(t, "emoji", event.ObjectKind)
	assert.Equal(t, "award", event.EventType)
	assert.Equal(t, int64(7), event.ProjectID)

	// User assertions
	assert.Equal(t, int64(1), event.User.ID)
	assert.Equal(t, "Administrator", event.User.Name)
	assert.Equal(t, "root", event.User.Username)
	assert.Equal(t, "admin@example.com", event.User.Email)

	// Project assertions
	assert.NotNil(t, event.Project)
	assert.Equal(t, int64(7), event.Project.ID)
	assert.Equal(t, "Flight", event.Project.Name)
	assert.Equal(t, "flightjs/Flight", event.Project.PathWithNamespace)
	assert.Equal(t, int64(0), event.Project.VisibilityLevel)

	// Object attributes assertions
	assert.Equal(t, int64(42), event.ObjectAttributes.ID)
	assert.Equal(t, int64(1), event.ObjectAttributes.UserID)
	assert.Equal(t, "thumbsup", event.ObjectAttributes.Name)
	assert.Equal(t, "Issue", event.ObjectAttributes.AwardableType)
	assert.Equal(t, int64(123), event.ObjectAttributes.AwardableID)
	assert.Equal(t, "award", event.ObjectAttributes.Action)
	assert.Equal(t, "https://example.com/flightjs/Flight/-/issues/1", event.ObjectAttributes.AwardedOnURL)

	// Issue assertions
	assert.NotNil(t, event.Issue)
	assert.Equal(t, int64(123), event.Issue.ID)
	assert.Equal(t, int64(1), event.Issue.IID)
	assert.Equal(t, int64(7), event.Issue.ProjectID)
	assert.Equal(t, "Test Issue", event.Issue.Title)
	assert.Equal(t, "This is a test issue", event.Issue.Description)
	assert.Equal(t, StateIDOpen, event.Issue.StateID)
	assert.Equal(t, "opened", event.Issue.State)
	assert.Equal(t, "unknown", event.Issue.Severity)
	assert.False(t, event.Issue.Confidential)
}

func TestEmojiEventUnmarshalMergeRequest(t *testing.T) {
	t.Parallel()

	// GIVEN an emoji event JSON payload on a merge request
	jsonObject := loadFixture(t, "testdata/webhooks/emoji_merge_request.json")

	// WHEN the JSON is unmarshaled
	var event *EmojiEvent
	err := json.Unmarshal(jsonObject, &event)

	// THEN no error should occur and the event should be populated correctly
	assert.NoError(t, err)
	assert.NotNil(t, event)

	assert.Equal(t, "emoji", event.ObjectKind)
	assert.Equal(t, "award", event.EventType)
	assert.Equal(t, "MergeRequest", event.ObjectAttributes.AwardableType)

	// MergeRequest assertions
	assert.NotNil(t, event.MergeRequest)
	assert.Equal(t, int64(123), event.MergeRequest.ID)
	assert.Equal(t, int64(1), event.MergeRequest.IID)
	assert.Equal(t, "Test Merge Request", event.MergeRequest.Title)
	assert.Equal(t, "opened", event.MergeRequest.State)

	// Issue and Snippet should be nil
	assert.Nil(t, event.Issue)
	assert.Nil(t, event.ProjectSnippet)
	assert.Nil(t, event.Note)
	assert.Nil(t, event.Commit)
}

func TestEmojiEventUnmarshalSnippet(t *testing.T) {
	t.Parallel()

	// GIVEN an emoji event JSON payload on a snippet
	jsonObject := loadFixture(t, "testdata/webhooks/emoji_snippet.json")

	// WHEN the JSON is unmarshaled
	var event *EmojiEvent
	err := json.Unmarshal(jsonObject, &event)

	// THEN no error should occur and the event should be populated correctly
	assert.NoError(t, err)
	assert.NotNil(t, event)

	assert.Equal(t, "emoji", event.ObjectKind)
	assert.Equal(t, "award", event.EventType)
	assert.Equal(t, "Snippet", event.ObjectAttributes.AwardableType)

	// Snippet assertions
	assert.NotNil(t, event.ProjectSnippet)
	assert.Equal(t, int64(456), event.ProjectSnippet.ID)
	assert.Equal(t, "Test Snippet", event.ProjectSnippet.Title)
	assert.Equal(t, "ProjectSnippet", event.ProjectSnippet.Type)

	// Issue and MergeRequest should be nil
	assert.Nil(t, event.Issue)
	assert.Nil(t, event.MergeRequest)
	assert.Nil(t, event.Note)
	assert.Nil(t, event.Commit)
}

func TestEmojiEventUnmarshalNoteIssue(t *testing.T) {
	t.Parallel()

	// GIVEN an emoji event JSON payload on a note on an issue
	jsonObject := loadFixture(t, "testdata/webhooks/emoji_note_issue.json")

	// WHEN the JSON is unmarshaled
	var event *EmojiEvent
	err := json.Unmarshal(jsonObject, &event)

	// THEN no error should occur and the event should be populated correctly
	assert.NoError(t, err)
	assert.NotNil(t, event)

	assert.Equal(t, "emoji", event.ObjectKind)
	assert.Equal(t, "award", event.EventType)
	assert.Equal(t, "Note", event.ObjectAttributes.AwardableType)

	// Note assertions
	assert.NotNil(t, event.Note)
	assert.Equal(t, int64(789), event.Note.ID)
	assert.Equal(t, "Issue", event.Note.NoteableType)
	assert.Equal(t, int64(123), event.Note.NoteableID)

	// Issue assertions
	assert.NotNil(t, event.Issue)
	assert.Equal(t, int64(123), event.Issue.ID)

	// MergeRequest and Snippet should be nil
	assert.Nil(t, event.MergeRequest)
	assert.Nil(t, event.ProjectSnippet)
	assert.Nil(t, event.Commit)
}

func TestEmojiEventUnmarshalNoteMergeRequest(t *testing.T) {
	t.Parallel()

	// GIVEN an emoji event JSON payload on a note on a merge request
	jsonObject := loadFixture(t, "testdata/webhooks/emoji_note_merge_request.json")

	// WHEN the JSON is unmarshaled
	var event *EmojiEvent
	err := json.Unmarshal(jsonObject, &event)

	// THEN no error should occur and the event should be populated correctly
	assert.NoError(t, err)
	assert.NotNil(t, event)

	assert.Equal(t, "emoji", event.ObjectKind)
	assert.Equal(t, "award", event.EventType)
	assert.Equal(t, "Note", event.ObjectAttributes.AwardableType)

	// Note assertions
	assert.NotNil(t, event.Note)
	assert.Equal(t, int64(790), event.Note.ID)
	assert.Equal(t, "MergeRequest", event.Note.NoteableType)

	// MergeRequest assertions
	assert.NotNil(t, event.MergeRequest)
	assert.Equal(t, int64(123), event.MergeRequest.ID)

	// Issue and Snippet should be nil
	assert.Nil(t, event.Issue)
	assert.Nil(t, event.ProjectSnippet)
	assert.Nil(t, event.Commit)
}

func TestEmojiEventUnmarshalNoteSnippet(t *testing.T) {
	t.Parallel()

	// GIVEN an emoji event JSON payload on a note on a snippet
	jsonObject := loadFixture(t, "testdata/webhooks/emoji_note_snippet.json")

	// WHEN the JSON is unmarshaled
	var event *EmojiEvent
	err := json.Unmarshal(jsonObject, &event)

	// THEN no error should occur and the event should be populated correctly
	assert.NoError(t, err)
	assert.NotNil(t, event)

	assert.Equal(t, "emoji", event.ObjectKind)
	assert.Equal(t, "award", event.EventType)
	assert.Equal(t, "Note", event.ObjectAttributes.AwardableType)

	// Note assertions
	assert.NotNil(t, event.Note)
	assert.Equal(t, int64(791), event.Note.ID)
	assert.Equal(t, "Snippet", event.Note.NoteableType)

	// Snippet assertions
	assert.NotNil(t, event.ProjectSnippet)
	assert.Equal(t, int64(456), event.ProjectSnippet.ID)

	// Issue and MergeRequest should be nil
	assert.Nil(t, event.Issue)
	assert.Nil(t, event.MergeRequest)
	assert.Nil(t, event.Commit)
}

func TestEmojiEventUnmarshalNoteCommit(t *testing.T) {
	t.Parallel()

	// GIVEN an emoji event JSON payload on a note on a commit
	jsonObject := loadFixture(t, "testdata/webhooks/emoji_note_commit.json")

	// WHEN the JSON is unmarshaled
	var event *EmojiEvent
	err := json.Unmarshal(jsonObject, &event)

	// THEN no error should occur and the event should be populated correctly
	assert.NoError(t, err)
	assert.NotNil(t, event)

	assert.Equal(t, "emoji", event.ObjectKind)
	assert.Equal(t, "award", event.EventType)
	assert.Equal(t, "Note", event.ObjectAttributes.AwardableType)

	// Note assertions
	assert.NotNil(t, event.Note)
	assert.Equal(t, int64(792), event.Note.ID)
	assert.Equal(t, "Commit", event.Note.NoteableType)
	assert.NotNil(t, event.Note.CommitID)
	assert.Equal(t, "cfe32cf61b73a0d5e9f13e774abde7ff789b1660", *event.Note.CommitID)

	// Commit assertions
	assert.NotNil(t, event.Commit)
	assert.Equal(t, "cfe32cf61b73a0d5e9f13e774abde7ff789b1660", event.Commit.ID)
	assert.Equal(t, "Add submodule", event.Commit.Title)

	// Issue, MergeRequest and Snippet should be nil
	assert.Nil(t, event.Issue)
	assert.Nil(t, event.MergeRequest)
	assert.Nil(t, event.ProjectSnippet)
}

func TestMilestoneWebhookEventUnmarshal(t *testing.T) {
	t.Parallel()

	// GIVEN a milestone webhook event JSON payload
	jsonObject := loadFixture(t, "testdata/webhooks/milestone_project.json")

	// WHEN the JSON is unmarshaled
	var event *MilestoneWebhookEvent
	err := json.Unmarshal(jsonObject, &event)

	// THEN no error should occur and the event should be populated correctly
	assert.NoError(t, err)
	assert.NotNil(t, event)

	assert.Equal(t, "milestone", event.ObjectKind)
	assert.Equal(t, "milestone", event.EventType)
	assert.Equal(t, "create", event.Action)

	// Project assertions (project milestone)
	assert.NotNil(t, event.Project)
	assert.Equal(t, int64(7), event.Project.ID)
	assert.Equal(t, "Flight", event.Project.Name)
	assert.Equal(t, "flightjs/Flight", event.Project.PathWithNamespace)
	assert.Equal(t, int64(0), event.Project.VisibilityLevel)

	// Group should be nil for project milestone
	assert.Nil(t, event.Group)

	// Object attributes assertions
	assert.Equal(t, int64(42), event.ObjectAttributes.ID)
	assert.Equal(t, int64(1), event.ObjectAttributes.IID)
	assert.Equal(t, "v1.0.0", event.ObjectAttributes.Title)
	assert.Equal(t, "First major release milestone", event.ObjectAttributes.Description)
	assert.Equal(t, "active", event.ObjectAttributes.State)
	assert.Equal(t, int64(7), event.ObjectAttributes.ProjectID)
	assert.Nil(t, event.ObjectAttributes.GroupID)

	// Date assertions
	dueDate, err := ParseISOTime("2024-03-01")
	assert.NoError(t, err)
	assert.Equal(t, &dueDate, event.ObjectAttributes.DueDate)

	startDate, err := ParseISOTime("2024-01-01")
	assert.NoError(t, err)
	assert.Equal(t, &startDate, event.ObjectAttributes.StartDate)
}

func TestMilestoneWebhookEventUnmarshalGroup(t *testing.T) {
	t.Parallel()

	// GIVEN a group milestone webhook event JSON payload
	jsonObject := loadFixture(t, "testdata/webhooks/milestone_group.json")

	// WHEN the JSON is unmarshaled
	var event *MilestoneWebhookEvent
	err := json.Unmarshal(jsonObject, &event)

	// THEN no error should occur and the event should be populated correctly
	assert.NoError(t, err)
	assert.NotNil(t, event)

	assert.Equal(t, "milestone", event.ObjectKind)
	assert.Equal(t, "milestone", event.EventType)
	assert.Equal(t, "create", event.Action)

	// Project should be empty for group milestone
	assert.Empty(t, event.Project)

	// Group assertions (group milestone)
	assert.NotNil(t, event.Group)
	assert.Equal(t, int64(35), event.Group.GroupID)
	assert.Equal(t, "Flightjs", event.Group.GroupName)
	assert.Equal(t, "flightjs", event.Group.GroupPath)
	assert.Equal(t, "flightjs", event.Group.FullPath)

	// Object attributes assertions
	assert.Equal(t, int64(42), event.ObjectAttributes.ID)
	assert.Equal(t, int64(1), event.ObjectAttributes.IID)
	assert.Equal(t, "v1.0.0", event.ObjectAttributes.Title)
	assert.Equal(t, int64(35), *event.ObjectAttributes.GroupID)
	assert.Equal(t, int64(0), event.ObjectAttributes.ProjectID)

	// Date assertions
	dueDate, err := ParseISOTime("2024-03-01")
	assert.NoError(t, err)
	assert.Equal(t, &dueDate, event.ObjectAttributes.DueDate)

	startDate, err := ParseISOTime("2024-01-01")
	assert.NoError(t, err)
	assert.Equal(t, &startDate, event.ObjectAttributes.StartDate)
}

func TestProjectWebhookEventUnmarshal(t *testing.T) {
	t.Parallel()

	// GIVEN a project webhook event JSON payload
	jsonObject := loadFixture(t, "testdata/webhooks/project.json")

	// WHEN the JSON is unmarshaled
	var event *ProjectWebhookEvent
	err := json.Unmarshal(jsonObject, &event)

	// THEN no error should occur and the event should be populated correctly
	assert.NoError(t, err)
	assert.NotNil(t, event)

	assert.Equal(t, "project_create", event.EventName)
	assert.Equal(t, "2024-01-24 16:27:40 UTC", event.CreatedAt)
	assert.Equal(t, "2024-01-24 16:27:40 UTC", event.UpdatedAt)
	assert.Equal(t, "Flight", event.Name)
	assert.Equal(t, "flight", event.Path)
	assert.Equal(t, "flightjs/flight", event.PathWithNamespace)
	assert.Equal(t, int64(7), event.ProjectID)
	assert.Equal(t, int64(35), event.ProjectNamespaceID)
	assert.Equal(t, "private", event.ProjectVisibility)
	assert.Empty(t, event.OldPathWithNamespace)

	// Owners assertions
	assert.Len(t, event.Owners, 1)
	assert.Equal(t, "Administrator", event.Owners[0].Name)
	assert.Equal(t, "admin@example.com", event.Owners[0].Email)
}

func TestVulnerabilityEventUnmarshal(t *testing.T) {
	t.Parallel()

	// GIVEN a vulnerability event JSON payload
	jsonObject := loadFixture(t, "testdata/webhooks/vulnerability.json")

	// WHEN the JSON is unmarshaled
	var event *VulnerabilityEvent
	err := json.Unmarshal(jsonObject, &event)

	// THEN no error should occur and the event should be populated correctly
	assert.NoError(t, err)
	assert.NotNil(t, event)

	assert.Equal(t, "vulnerability", event.ObjectKind)

	// Object attributes assertions
	assert.Equal(t, int64(42), event.ObjectAttributes.ID)
	assert.Equal(t, "https://example.com/flightjs/Flight/-/security/vulnerabilities/42", event.ObjectAttributes.URL)
	assert.Equal(t, "Potential SQL Injection", event.ObjectAttributes.Title)
	assert.Equal(t, "detected", event.ObjectAttributes.State)
	assert.Equal(t, int64(7), event.ObjectAttributes.ProjectID)
	assert.Equal(t, "high", event.ObjectAttributes.Severity)
	assert.False(t, event.ObjectAttributes.SeverityOverridden)
	assert.Equal(t, "sast", event.ObjectAttributes.ReportType)
	assert.Equal(t, "high", event.ObjectAttributes.Confidence)
	assert.False(t, event.ObjectAttributes.ConfidenceOverridden)
	assert.Equal(t, int64(1), event.ObjectAttributes.ConfirmedByID)
	assert.False(t, event.ObjectAttributes.AutoResolved)
	assert.False(t, event.ObjectAttributes.ResolvedOnDefaultBranch)

	// Location assertions
	assert.Equal(t, "app/models/user.rb", event.ObjectAttributes.Location.File)
	assert.Equal(t, "pg", event.ObjectAttributes.Location.Dependency.Package.Name)
	assert.Equal(t, "1.2.3", event.ObjectAttributes.Location.Dependency.Version)

	// CVSS assertions
	assert.Len(t, event.ObjectAttributes.CVSS, 1)
	assert.Equal(t, "CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H", event.ObjectAttributes.CVSS[0].Vector)
	assert.Equal(t, "GitLab", event.ObjectAttributes.CVSS[0].Vendor)

	// Identifiers assertions
	assert.Len(t, event.ObjectAttributes.Identifiers, 1)
	assert.Equal(t, "CVE-2024-1234", event.ObjectAttributes.Identifiers[0].Name)
	assert.Equal(t, "CVE-2024-1234", event.ObjectAttributes.Identifiers[0].ExternalID)
	assert.Equal(t, "cve", event.ObjectAttributes.Identifiers[0].ExternalType)

	// Issues assertions
	assert.Len(t, event.ObjectAttributes.Issues, 1)
	assert.Equal(t, "Fix SQL Injection vulnerability", event.ObjectAttributes.Issues[0].Title)
	assert.Equal(t, "https://example.com/flightjs/Flight/-/issues/10", event.ObjectAttributes.Issues[0].URL)
}

// parseDuration unit tests

func TestParseDuration_RawNumber(t *testing.T) {
	t.Parallel()
	got, err := parseDuration([]byte(`17.1`))
	require.NoError(t, err)
	assert.Equal(t, 17.1, got)
}

func TestParseDuration_QuotedNumber(t *testing.T) {
	t.Parallel()
	got, err := parseDuration([]byte(`"17.1"`))
	require.NoError(t, err)
	assert.Equal(t, 17.1, got)
}

func TestParseDuration_Zero(t *testing.T) {
	t.Parallel()
	got, err := parseDuration([]byte(`0`))
	require.NoError(t, err)
	assert.Equal(t, 0.0, got)
}

func TestParseDuration_Null(t *testing.T) {
	t.Parallel()
	got, err := parseDuration([]byte(`null`))
	require.NoError(t, err)
	assert.Equal(t, 0.0, got)
}

func TestParseDuration_EmptyQuotedString(t *testing.T) {
	t.Parallel()
	got, err := parseDuration([]byte(`""`))
	require.NoError(t, err)
	assert.Equal(t, 0.0, got)
}

func TestParseDuration_EmptyBytes(t *testing.T) {
	t.Parallel()
	got, err := parseDuration([]byte(``))
	require.NoError(t, err)
	assert.Equal(t, 0.0, got)
}

func TestParseDuration_InvalidString(t *testing.T) {
	t.Parallel()
	_, err := parseDuration([]byte(`"not-a-number"`))
	require.Error(t, err)
}

// JobEvent marshal/unmarshal tests for quoted durations

func TestJobEventUnmarshal_QuotedDurations(t *testing.T) {
	t.Parallel()

	raw := `{
		"object_kind": "build",
		"build_id": 1,
		"build_duration": "49.5",
		"build_queued_duration": "0.19"
	}`

	var event JobEvent
	require.NoError(t, json.Unmarshal([]byte(raw), &event))
	assert.Equal(t, 49.5, event.BuildDuration)
	assert.Equal(t, 0.19, event.BuildQueuedDuration)
}

func TestJobEventUnmarshal_NullDurations(t *testing.T) {
	t.Parallel()

	raw := `{
		"object_kind": "build",
		"build_id": 1,
		"build_duration": null,
		"build_queued_duration": null
	}`

	var event JobEvent
	require.NoError(t, json.Unmarshal([]byte(raw), &event))
	assert.Equal(t, 0.0, event.BuildDuration)
	assert.Equal(t, 0.0, event.BuildQueuedDuration)
}

func TestJobEventMarshal_ProducesNumeric(t *testing.T) {
	t.Parallel()

	event := JobEvent{
		BuildDuration:       49.5,
		BuildQueuedDuration: 0.19,
	}

	data, err := json.Marshal(&event)
	require.NoError(t, err)

	var raw map[string]any
	require.NoError(t, json.Unmarshal(data, &raw))
	assert.Equal(t, 49.5, raw["build_duration"])
	assert.Equal(t, 0.19, raw["build_queued_duration"])
}

// PipelineEventBuild marshal/unmarshal tests for quoted durations

func TestPipelineEventBuildUnmarshal_QuotedDurations(t *testing.T) {
	t.Parallel()

	raw := `{
		"id": 99,
		"stage": "test",
		"name": "unit",
		"status": "success",
		"duration": "17.1",
		"queued_duration": "3.5"
	}`

	var build PipelineEventBuild
	require.NoError(t, json.Unmarshal([]byte(raw), &build))
	assert.Equal(t, 17.1, build.Duration)
	assert.Equal(t, 3.5, build.QueuedDuration)
}

func TestPipelineEventBuildUnmarshal_NullDurations(t *testing.T) {
	t.Parallel()

	raw := `{
		"id": 99,
		"stage": "test",
		"name": "unit",
		"status": "success",
		"duration": null,
		"queued_duration": null
	}`

	var build PipelineEventBuild
	require.NoError(t, json.Unmarshal([]byte(raw), &build))
	assert.Equal(t, 0.0, build.Duration)
	assert.Equal(t, 0.0, build.QueuedDuration)
}

func TestPipelineEventBuildMarshal_ProducesNumeric(t *testing.T) {
	t.Parallel()

	build := PipelineEventBuild{
		ID:             99,
		Duration:       17.1,
		QueuedDuration: 3.5,
	}

	data, err := json.Marshal(&build)
	require.NoError(t, err)

	var raw map[string]any
	require.NoError(t, json.Unmarshal(data, &raw))
	assert.Equal(t, 17.1, raw["duration"])
	assert.Equal(t, 3.5, raw["queued_duration"])
}

// PipelineEventObjectAttributes marshal/unmarshal tests for quoted durations

func TestPipelineEventObjectAttributesUnmarshal_QuotedDurations(t *testing.T) {
	t.Parallel()

	raw := `{
		"id": 1,
		"iid": 1,
		"ref": "main",
		"status": "success",
		"duration": "63",
		"queued_duration": "12"
	}`

	var attrs PipelineEventObjectAttributes
	require.NoError(t, json.Unmarshal([]byte(raw), &attrs))
	assert.Equal(t, int64(63), attrs.Duration)
	assert.Equal(t, int64(12), attrs.QueuedDuration)
}

func TestPipelineEventObjectAttributesUnmarshal_NullDurations(t *testing.T) {
	t.Parallel()

	raw := `{
		"id": 1,
		"iid": 1,
		"ref": "main",
		"status": "success",
		"duration": null,
		"queued_duration": null
	}`

	var attrs PipelineEventObjectAttributes
	require.NoError(t, json.Unmarshal([]byte(raw), &attrs))
	assert.Equal(t, int64(0), attrs.Duration)
	assert.Equal(t, int64(0), attrs.QueuedDuration)
}

func TestPipelineEventObjectAttributesMarshal_ProducesNumeric(t *testing.T) {
	t.Parallel()

	attrs := PipelineEventObjectAttributes{
		Duration:       63,
		QueuedDuration: 12,
	}

	data, err := json.Marshal(&attrs)
	require.NoError(t, err)

	var raw map[string]any
	require.NoError(t, json.Unmarshal(data, &raw))
	assert.Equal(t, float64(63), raw["duration"])
	assert.Equal(t, float64(12), raw["queued_duration"])
}

// Error case tests for UnmarshalJSON on affected structs

func TestJobEventUnmarshal_InvalidJSON(t *testing.T) {
	t.Parallel()
	raw := `{not valid json`
	var event JobEvent
	require.Error(t, json.Unmarshal([]byte(raw), &event))
}

func TestJobEventUnmarshal_InvalidDuration(t *testing.T) {
	t.Parallel()
	raw := `{
		"object_kind": "build",
		"build_id": 1,
		"build_duration": "not-a-number",
		"build_queued_duration": "0.19"
	}`
	var event JobEvent
	require.Error(t, json.Unmarshal([]byte(raw), &event))
}

func TestJobEventUnmarshal_InvalidQueuedDuration(t *testing.T) {
	t.Parallel()
	raw := `{
		"object_kind": "build",
		"build_id": 1,
		"build_duration": "49.5",
		"build_queued_duration": "not-a-number"
	}`
	var event JobEvent
	require.Error(t, json.Unmarshal([]byte(raw), &event))
}

func TestPipelineEventBuildUnmarshal_InvalidJSON(t *testing.T) {
	t.Parallel()
	raw := `{not valid json`
	var build PipelineEventBuild
	require.Error(t, json.Unmarshal([]byte(raw), &build))
}

func TestPipelineEventBuildUnmarshal_InvalidDuration(t *testing.T) {
	t.Parallel()
	raw := `{
		"id": 99,
		"stage": "test",
		"name": "unit",
		"status": "success",
		"duration": "not-a-number",
		"queued_duration": "3.5"
	}`
	var build PipelineEventBuild
	require.Error(t, json.Unmarshal([]byte(raw), &build))
}

func TestPipelineEventBuildUnmarshal_InvalidQueuedDuration(t *testing.T) {
	t.Parallel()
	raw := `{
		"id": 99,
		"stage": "test",
		"name": "unit",
		"status": "success",
		"duration": "17.1",
		"queued_duration": "not-a-number"
	}`
	var build PipelineEventBuild
	require.Error(t, json.Unmarshal([]byte(raw), &build))
}

func TestPipelineEventObjectAttributesUnmarshal_InvalidJSON(t *testing.T) {
	t.Parallel()
	raw := `{not valid json`
	var attrs PipelineEventObjectAttributes
	require.Error(t, json.Unmarshal([]byte(raw), &attrs))
}

func TestPipelineEventObjectAttributesUnmarshal_InvalidDuration(t *testing.T) {
	t.Parallel()
	raw := `{
		"id": 1,
		"iid": 1,
		"ref": "main",
		"status": "success",
		"duration": "not-a-number",
		"queued_duration": "12"
	}`
	var attrs PipelineEventObjectAttributes
	require.Error(t, json.Unmarshal([]byte(raw), &attrs))
}

func TestPipelineEventObjectAttributesUnmarshal_InvalidQueuedDuration(t *testing.T) {
	t.Parallel()
	raw := `{
		"id": 1,
		"iid": 1,
		"ref": "main",
		"status": "success",
		"duration": "63",
		"queued_duration": "not-a-number"
	}`
	var attrs PipelineEventObjectAttributes
	require.Error(t, json.Unmarshal([]byte(raw), &attrs))
}
