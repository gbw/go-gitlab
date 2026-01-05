//go:build integration

package gitlab_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

// Integration tests for the Groups API.
// These tests require a GitLab instance running on localhost:8095.
// They also require a valid admin token in GITLAB_TOKEN environment variable.

func Test_GroupsGetGroup_MergeSettings_Integration(t *testing.T) {
	// GIVEN a GitLab client and a test group
	client := SetupIntegrationClient(t)
	group := CreateTestGroup(t, client)

	// WHEN retrieving the group
	retrievedGroup, _, err := client.Groups.GetGroup(group.ID, nil)
	require.NoError(t, err, "Failed to get group")

	// THEN the merge settings fields should be present
	assert.NotNil(t, retrievedGroup)
	// The fields should have default values (false for most merge restrictions)
	assert.False(t, retrievedGroup.OnlyAllowMergeIfPipelineSucceeds)
	assert.False(t, retrievedGroup.AllowMergeOnSkippedPipeline)
	assert.False(t, retrievedGroup.OnlyAllowMergeIfAllDiscussionsAreResolved)
}

func Test_GroupsUpdateGroup_MergeSettings_Integration(t *testing.T) {
	// GIVEN a GitLab client and a test group
	client := SetupIntegrationClient(t)
	group := CreateTestGroup(t, client)

	// WHEN updating the group with merge settings enabled
	updatedGroup, _, err := client.Groups.UpdateGroup(group.ID, &gitlab.UpdateGroupOptions{
		OnlyAllowMergeIfPipelineSucceeds:          gitlab.Ptr(true),
		AllowMergeOnSkippedPipeline:               gitlab.Ptr(true),
		OnlyAllowMergeIfAllDiscussionsAreResolved: gitlab.Ptr(true),
	})
	require.NoError(t, err, "Failed to update group")

	// THEN the merge settings should be updated
	assert.True(t, updatedGroup.OnlyAllowMergeIfPipelineSucceeds)
	assert.True(t, updatedGroup.AllowMergeOnSkippedPipeline)
	assert.True(t, updatedGroup.OnlyAllowMergeIfAllDiscussionsAreResolved)

	// AND WHEN retrieving the group again
	retrievedGroup, _, err := client.Groups.GetGroup(group.ID, nil)
	require.NoError(t, err, "Failed to get group")

	// THEN the merge settings should persist
	assert.True(t, retrievedGroup.OnlyAllowMergeIfPipelineSucceeds)
	assert.True(t, retrievedGroup.AllowMergeOnSkippedPipeline)
	assert.True(t, retrievedGroup.OnlyAllowMergeIfAllDiscussionsAreResolved)
}

func Test_GroupsUpdateGroup_MergeSettings_Disable_Integration(t *testing.T) {
	// GIVEN a GitLab client and a test group with merge settings enabled
	client := SetupIntegrationClient(t)
	group := CreateTestGroup(t, client)

	// Enable merge settings first
	_, _, err := client.Groups.UpdateGroup(group.ID, &gitlab.UpdateGroupOptions{
		OnlyAllowMergeIfPipelineSucceeds:          gitlab.Ptr(true),
		AllowMergeOnSkippedPipeline:               gitlab.Ptr(true),
		OnlyAllowMergeIfAllDiscussionsAreResolved: gitlab.Ptr(true),
	})
	require.NoError(t, err, "Failed to enable merge settings")

	// WHEN updating the group to disable merge settings
	updatedGroup, _, err := client.Groups.UpdateGroup(group.ID, &gitlab.UpdateGroupOptions{
		OnlyAllowMergeIfPipelineSucceeds:          gitlab.Ptr(false),
		AllowMergeOnSkippedPipeline:               gitlab.Ptr(false),
		OnlyAllowMergeIfAllDiscussionsAreResolved: gitlab.Ptr(false),
	})
	require.NoError(t, err, "Failed to update group")

	// THEN the merge settings should be disabled
	assert.False(t, updatedGroup.OnlyAllowMergeIfPipelineSucceeds)
	assert.False(t, updatedGroup.AllowMergeOnSkippedPipeline)
	assert.False(t, updatedGroup.OnlyAllowMergeIfAllDiscussionsAreResolved)
}
