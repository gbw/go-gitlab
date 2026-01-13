//go:build integration

package gitlab_test

import (
	"testing"

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