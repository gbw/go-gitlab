//go:build integration

package gitlab_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	gitlab "gitlab.com/gitlab-org/api/client-go/v2"
)

// Integration tests for the Project Hooks API.
// These tests require a GitLab instance running on localhost:8095.
// They also require a valid admin token in GITLAB_TOKEN environment variable.

func Test_ProjectHooksListProjectHooks_Integration(t *testing.T) {
	t.Parallel()
	client := SetupIntegrationClient(t)

	project := CreateTestProject(t, client)
	hook, err := CreateTestProjectHook(t, project.ID, client)
	require.NoError(t, err, "Failed to create test hook")

	hooks, _, err := client.Projects.ListProjectHooks(project.ID, nil)
	require.NoError(t, err, "Failed to list project hooks")

	assert.NotNil(t, hooks)
	assert.GreaterOrEqual(t, len(hooks), 1)
	assert.Equal(t, hook.ID, hooks[0].ID)
}

func Test_ProjectHooksGetProjectHook_Integration(t *testing.T) {
	t.Parallel()
	client := SetupIntegrationClient(t)

	project := CreateTestProject(t, client)
	hook, err := CreateTestProjectHook(t, project.ID, client)
	require.NoError(t, err, "Failed to create test hook")

	retrievedHook, _, err := client.Projects.GetProjectHook(project.ID, hook.ID)
	require.NoError(t, err, "Failed to get project hook")

	assert.Equal(t, hook.ID, retrievedHook.ID)
	assert.Equal(t, hook.URL, retrievedHook.URL)
	assert.True(t, retrievedHook.PushEvents)
}

func Test_ProjectHooksAddProjectHook_Integration(t *testing.T) {
	t.Parallel()
	client := SetupIntegrationClient(t)

	project := CreateTestProject(t, client)

	suffix := time.Now().UnixNano()
	hookURL := fmt.Sprintf("https://example.com/%d", suffix)

	hook, _, err := client.Projects.AddProjectHook(project.ID, &gitlab.AddProjectHookOptions{
		URL:                       &hookURL,
		PushEvents:                gitlab.Ptr(true),
		PushEventsBranchFilter:    gitlab.Ptr("main"),
		IssuesEvents:              gitlab.Ptr(true),
		ConfidentialIssuesEvents:  gitlab.Ptr(true),
		MergeRequestsEvents:       gitlab.Ptr(true),
		TagPushEvents:             gitlab.Ptr(true),
		NoteEvents:                gitlab.Ptr(true),
		ConfidentialNoteEvents:    gitlab.Ptr(true),
		JobEvents:                 gitlab.Ptr(true),
		PipelineEvents:            gitlab.Ptr(true),
		WikiPageEvents:            gitlab.Ptr(true),
		DeploymentEvents:          gitlab.Ptr(true),
		ReleasesEvents:            gitlab.Ptr(true),
		EmojiEvents:               gitlab.Ptr(true),
		FeatureFlagEvents:         gitlab.Ptr(true),
		MilestoneEvents:           gitlab.Ptr(true),
		ResourceAccessTokenEvents: gitlab.Ptr(true),
		ResourceDeployTokenEvents: gitlab.Ptr(true),
		EnableSSLVerification:     gitlab.Ptr(true),
		Token:                     gitlab.Ptr("secret-token"),
	})
	require.NoError(t, err, "Failed to add project hook")

	t.Cleanup(func() {
		_, err := client.Projects.DeleteProjectHook(project.ID, hook.ID)
		require.NoError(t, err, "Failed to delete test project hook")
	})

	assert.NotZero(t, hook.ID)
	assert.Equal(t, hookURL, hook.URL)
	assert.True(t, hook.PushEvents)
	assert.Equal(t, "main", hook.PushEventsBranchFilter)
	assert.True(t, hook.IssuesEvents)
	assert.True(t, hook.ConfidentialIssuesEvents)
	assert.True(t, hook.MergeRequestsEvents)
	assert.True(t, hook.TagPushEvents)
	assert.True(t, hook.NoteEvents)
	assert.True(t, hook.ConfidentialNoteEvents)
	assert.True(t, hook.JobEvents)
	assert.True(t, hook.PipelineEvents)
	assert.True(t, hook.WikiPageEvents)
	assert.True(t, hook.DeploymentEvents)
	assert.True(t, hook.ReleasesEvents)
	assert.True(t, hook.EmojiEvents)
	assert.True(t, hook.FeatureFlagEvents)
	assert.True(t, hook.MilestoneEvents)
	assert.True(t, hook.ResourceAccessTokenEvents)
	assert.True(t, hook.ResourceDeployTokenEvents)
	assert.True(t, hook.EnableSSLVerification)
}

func Test_ProjectHooksEditProjectHook_Integration(t *testing.T) {
	t.Parallel()
	client := SetupIntegrationClient(t)

	project := CreateTestProject(t, client)
	hook, err := CreateTestProjectHook(t, project.ID, client)
	require.NoError(t, err, "Failed to create test hook")

	suffix := time.Now().UnixNano()
	hookURL := fmt.Sprintf("https://example.com/%d", suffix)

	updatedHook, _, err := client.Projects.EditProjectHook(project.ID, hook.ID, &gitlab.EditProjectHookOptions{
		URL:                       &hookURL,
		PushEvents:                gitlab.Ptr(true),
		PushEventsBranchFilter:    gitlab.Ptr("main"),
		IssuesEvents:              gitlab.Ptr(true),
		ConfidentialIssuesEvents:  gitlab.Ptr(true),
		MergeRequestsEvents:       gitlab.Ptr(true),
		TagPushEvents:             gitlab.Ptr(true),
		NoteEvents:                gitlab.Ptr(true),
		ConfidentialNoteEvents:    gitlab.Ptr(true),
		JobEvents:                 gitlab.Ptr(true),
		PipelineEvents:            gitlab.Ptr(true),
		WikiPageEvents:            gitlab.Ptr(true),
		DeploymentEvents:          gitlab.Ptr(true),
		ReleasesEvents:            gitlab.Ptr(true),
		EmojiEvents:               gitlab.Ptr(true),
		FeatureFlagEvents:         gitlab.Ptr(true),
		MilestoneEvents:           gitlab.Ptr(true),
		ResourceAccessTokenEvents: gitlab.Ptr(true),
		ResourceDeployTokenEvents: gitlab.Ptr(true),
		EnableSSLVerification:     gitlab.Ptr(true),
		Token:                     gitlab.Ptr("secret-token"),
	})
	require.NoError(t, err, "Failed to edit project hook")

	assert.NotZero(t, updatedHook.ID)
	assert.Equal(t, hookURL, updatedHook.URL)
	assert.True(t, updatedHook.PushEvents)
	assert.Equal(t, "main", updatedHook.PushEventsBranchFilter)
	assert.True(t, updatedHook.IssuesEvents)
	assert.True(t, updatedHook.ConfidentialIssuesEvents)
	assert.True(t, updatedHook.MergeRequestsEvents)
	assert.True(t, updatedHook.TagPushEvents)
	assert.True(t, updatedHook.NoteEvents)
	assert.True(t, updatedHook.ConfidentialNoteEvents)
	assert.True(t, updatedHook.JobEvents)
	assert.True(t, updatedHook.PipelineEvents)
	assert.True(t, updatedHook.WikiPageEvents)
	assert.True(t, updatedHook.DeploymentEvents)
	assert.True(t, updatedHook.ReleasesEvents)
	assert.True(t, updatedHook.EmojiEvents)
	assert.True(t, updatedHook.FeatureFlagEvents)
	assert.True(t, updatedHook.MilestoneEvents)
	assert.True(t, updatedHook.ResourceAccessTokenEvents)
	assert.True(t, updatedHook.ResourceDeployTokenEvents)
	assert.True(t, updatedHook.EnableSSLVerification)
}

func Test_ProjectHooksDeleteProjectHook_Integration(t *testing.T) {
	t.Parallel()
	client := SetupIntegrationClient(t)

	project := CreateTestProject(t, client)
	hook, err := CreateTestProjectHook(t, project.ID, client)
	require.NoError(t, err, "Failed to create test hook")

	_, err = client.Projects.DeleteProjectHook(project.ID, hook.ID)
	require.NoError(t, err, "Failed to delete project hook")
}
