//go:build integration

package gitlab_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

// Integration tests for the Projects API.
// These tests require a GitLab instance running on localhost:8095.
// They also require a valid admin token in GITLAB_TOKEN environment variable.

// Test_ProjectPullMirror_Integration tests the ConfigureProjectPullMirror function to validate
// that the mirror configuration attributes come back properly
func Test_ProjectPullMirror_Integration(t *testing.T) {
	// GIVEN a GitLab instance running on localhost:8095
	client := SetupIntegrationClient(t)

	// Requires Premium/Ultimate EE
	SkipIfNotLicensed(t, client)

	// Create a test project for the pull mirror,
	// And a project that will be mirrored
	project := CreateTestProject(t, client)
	projectToMirror := CreateTestProject(t, client)

	// When you create a pull mirror for the project
	mirror, _, err := client.Projects.ConfigureProjectPullMirror(project.ID, &gitlab.ConfigureProjectPullMirrorOptions{
		Enabled:                          gitlab.Ptr(true),
		URL:                              &projectToMirror.HTTPURLToRepo,
		MirrorTriggerBuilds:              gitlab.Ptr(true),
		OnlyMirrorProtectedBranches:      gitlab.Ptr(true),
		MirrorOverwritesDivergedBranches: gitlab.Ptr(true),
	})
	require.NoError(t, err)

	// Then the attributes are populated
	assert.Equal(t, true, mirror.MirrorTriggerBuilds)
	assert.Equal(t, true, mirror.OnlyMirrorProtectedBranches)
	assert.Equal(t, true, mirror.MirrorOverwritesDivergedBranches)
	assert.Equal(t, projectToMirror.HTTPURLToRepo, mirror.URL)
}

func Test_ProjectListProjectHooks_Integration(t *testing.T) {
	t.Parallel()

	client := SetupIntegrationClient(t)

	project := CreateTestProject(t, client)
	hook, err := CreateTestProjectHook(t, project.ID, client)
	require.NoError(t, err, "Failed to create test hook")

	hooks, _, err := client.Projects.ListProjectHooks(project.ID, &gitlab.ListProjectHooksOptions{})
	require.NoError(t, err, "Failed to list project hooks")

	assert.NotNil(t, hooks)
	assert.GreaterOrEqual(t, len(hooks), 1)
	assert.Equal(t, hook.ID, hooks[0].ID)
}

func Test_ProjectGetProjectHook_Integration(t *testing.T) {
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

func Test_ProjectAddProjectHook(t *testing.T) {
	t.Parallel()

	client := SetupIntegrationClient(t)

	project := CreateTestProject(t, client)
	suffix := time.Now().UnixNano()
	name := fmt.Sprintf("testhook%d", suffix)
	description := fmt.Sprintf("Test Hook %d", suffix)
	hookURL := fmt.Sprintf("https://example.com/%d", suffix)

	hook, _, err := client.Projects.AddProjectHook(project.ID, &gitlab.AddProjectHookOptions{
		URL:                  &hookURL,
		Name:                 &name,
		Description:          &description,
		VulnerabilityEvents:  gitlab.Ptr(true),
		BranchFilterStrategy: gitlab.Ptr("all_branches"),
	})
	require.NoError(t, err, "Failed to add project hook")

	assert.NotZero(t, hook.ID)
	assert.Equal(t, hookURL, hook.URL)
	assert.Equal(t, name, hook.Name)
	assert.Equal(t, description, hook.Description)
	assert.True(t, hook.VulnerabilityEvents)
	assert.Equal(t, "all_branches", hook.BranchFilterStrategy)
}

func Test_ProjectEditProjectHook(t *testing.T) {
	t.Parallel()

	client := SetupIntegrationClient(t)

	project := CreateTestProject(t, client)
	hook, err := CreateTestProjectHook(t, project.ID, client)
	require.NoError(t, err, "Failed to create test hook")
	suffix := time.Now().UnixNano()
	name := fmt.Sprintf("testhook%d", suffix)
	description := fmt.Sprintf("Test Hook %d", suffix)
	hookURL := fmt.Sprintf("https://example.com/%d", suffix)

	updatedHook, _, err := client.Projects.EditProjectHook(project.ID, hook.ID, &gitlab.EditProjectHookOptions{
		URL:                  &hookURL,
		Name:                 &name,
		Description:          &description,
		VulnerabilityEvents:  gitlab.Ptr(true),
		BranchFilterStrategy: gitlab.Ptr("all_branches"),
	})
	require.NoError(t, err, "Failed to edit project hook")

	assert.NotZero(t, updatedHook.ID)
	assert.Equal(t, hookURL, updatedHook.URL)
	assert.Equal(t, name, updatedHook.Name)
	assert.Equal(t, description, updatedHook.Description)
	assert.True(t, updatedHook.VulnerabilityEvents)
	assert.Equal(t, "all_branches", updatedHook.BranchFilterStrategy)
}

func Test_ProjectsMaxArtifactsSize_Integration(t *testing.T) {
	// GIVEN a GitLab client and a test project
	client := SetupIntegrationClient(t)
	project := CreateTestProject(t, client)

	// WHEN editing the project to set MaxArtifactsSize to 150 MB
	updatedProject, _, err := client.Projects.EditProject(project.ID, &gitlab.EditProjectOptions{
		MaxArtifactsSize: gitlab.Ptr(int64(150)), // 150 MB
	})
	require.NoError(t, err, "Failed to update project MaxArtifactsSize")

	// THEN the setting should be reflected in the update response
	assert.Equal(t, int64(150), updatedProject.MaxArtifactsSize)

	// AND WHEN retrieving the project again
	retrievedProject, _, err := client.Projects.GetProject(project.ID, nil)
	require.NoError(t, err, "Failed to retrieve project after update")

	// THEN MaxArtifactsSize should persist
	assert.Equal(t, int64(150), retrievedProject.MaxArtifactsSize)
}
