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

func Test_GroupsMaxArtifactsSize_Integration(t *testing.T) {
	// GIVEN a GitLab client and a test group
	client := SetupIntegrationClient(t)
	group := CreateTestGroup(t, client)

	// WHEN updating the group to set MaxArtifactsSize to 100 MB
	updatedGroup, _, err := client.Groups.UpdateGroup(group.ID, &gitlab.UpdateGroupOptions{
		MaxArtifactsSize: gitlab.Ptr(int64(100)), // 100 MB
	})
	require.NoError(t, err, "Failed to update group MaxArtifactsSize")

	// THEN the setting should be reflected in the update response
	assert.Equal(t, int64(100), updatedGroup.MaxArtifactsSize)

	// AND WHEN retrieving the group again
	retrievedGroup, _, err := client.Groups.GetGroup(group.ID, nil)
	require.NoError(t, err, "Failed to retrieve group after update")

	// THEN MaxArtifactsSize should persist
	assert.Equal(t, int64(100), retrievedGroup.MaxArtifactsSize)
}

func Test_GroupProtectedBranches_Integration(t *testing.T) {
	// GIVEN a GitLab client and a test group
	client := SetupIntegrationClient(t)
	group := CreateTestGroup(t, client)

	// Define branch name
	branchName := "main"

	// WHEN protecting a branch
	protectedBranch, _, err := client.GroupProtectedBranches.ProtectRepositoryBranches(group.ID, &gitlab.ProtectGroupRepositoryBranchesOptions{
		Name:             gitlab.Ptr(branchName),
		PushAccessLevel:  gitlab.Ptr(gitlab.MaintainerPermissions),
		MergeAccessLevel: gitlab.Ptr(gitlab.MaintainerPermissions),
	})
	require.NoError(t, err, "Failed to protect branch")

	// THEN the branch should be protected
	assert.Equal(t, branchName, protectedBranch.Name)

	// WHEN listing protected branches
	branches, _, err := client.GroupProtectedBranches.ListProtectedBranches(group.ID, nil)
	require.NoError(t, err, "Failed to list protected branches")

	// THEN the protected branch should be in the list
	found := false
	for _, b := range branches {
		if b.Name == branchName {
			found = true
			break
		}
	}
	assert.True(t, found, "Protected branch not found in list")

	// WHEN getting the protected branch
	gotBranch, _, err := client.GroupProtectedBranches.GetProtectedBranch(group.ID, branchName)
	require.NoError(t, err, "Failed to get protected branch")

	// THEN it should match
	assert.Equal(t, branchName, gotBranch.Name)

	// WHEN updating the protected branch
	updatedBranch, _, err := client.GroupProtectedBranches.UpdateProtectedBranch(group.ID, branchName, &gitlab.UpdateGroupProtectedBranchOptions{
		AllowForcePush: gitlab.Ptr(true),
	})
	require.NoError(t, err, "Failed to update protected branch")

	// THEN the update should be reflected
	assert.True(t, updatedBranch.AllowForcePush)

	// WHEN unprotecting the branch
	_, err = client.GroupProtectedBranches.UnprotectRepositoryBranches(group.ID, branchName)
	require.NoError(t, err, "Failed to unprotect branch")

	// THEN getting the branch should fail (404)
	_, resp, err := client.GroupProtectedBranches.GetProtectedBranch(group.ID, branchName)
	assert.Error(t, err)
	assert.Equal(t, 404, resp.StatusCode)
}

func Test_GroupsCreateGroup_CodeOwnerApprovalRequired_Integration(t *testing.T) {
	// GIVEN a GitLab client
	client := SetupIntegrationClient(t)

	// WHEN creating a group with CodeOwnerApprovalRequired enabled
	group, _, err := client.Groups.CreateGroup(&gitlab.CreateGroupOptions{
		Name: gitlab.Ptr("code-owner-create-test"),
		Path: gitlab.Ptr("code-owner-create-test"),
		DefaultBranchProtectionDefaults: &gitlab.DefaultBranchProtectionDefaultsOptions{
			CodeOwnerApprovalRequired: gitlab.Ptr(true),
		},
	})
	require.NoError(t, err, "Failed to create group")

	// THEN the setting should be enabled
	require.NotNil(t, group.DefaultBranchProtectionDefaults)
	assert.True(t, group.DefaultBranchProtectionDefaults.CodeOwnerApprovalRequired)

	// AND WHEN retrieving the group
	retrievedGroup, _, err := client.Groups.GetGroup(group.ID, nil)
	require.NoError(t, err)

	// THEN it should persist
	require.NotNil(t, retrievedGroup.DefaultBranchProtectionDefaults)
	assert.True(t, retrievedGroup.DefaultBranchProtectionDefaults.CodeOwnerApprovalRequired)
}

func Test_GroupsUpdateGroup_CodeOwnerApprovalRequired_Integration(t *testing.T) {
	// GIVEN a GitLab client and a test group
	client := SetupIntegrationClient(t)
	group := CreateTestGroup(t, client)

	// WHEN updating CodeOwnerApprovalRequired to true
	updatedGroup, _, err := client.Groups.UpdateGroup(group.ID, &gitlab.UpdateGroupOptions{
		DefaultBranchProtectionDefaults: &gitlab.DefaultBranchProtectionDefaultsOptions{
			CodeOwnerApprovalRequired: gitlab.Ptr(true),
		},
	})
	require.NoError(t, err)

	// THEN the update response should reflect the change
	require.NotNil(t, updatedGroup.DefaultBranchProtectionDefaults)
	assert.True(t, updatedGroup.DefaultBranchProtectionDefaults.CodeOwnerApprovalRequired)

	// AND WHEN retrieving the group again
	retrievedGroup, _, err := client.Groups.GetGroup(group.ID, nil)
	require.NoError(t, err)

	// THEN the setting should persist
	require.NotNil(t, retrievedGroup.DefaultBranchProtectionDefaults)
	assert.True(t, retrievedGroup.DefaultBranchProtectionDefaults.CodeOwnerApprovalRequired)
}
