//go:build integration

package gitlab_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	gitlab "gitlab.com/gitlab-org/api/client-go/v2"
)

func TestCreateWorkItem(t *testing.T) {
	t.Parallel()

	client := SetupIntegrationClient(t)
	SkipIfNotLicensed(t, client)

	// GIVEN a test group
	group := CreateTestGroup(t, client)

	// AND a work item creation options
	opt := gitlab.CreateWorkItemOptions{
		Title:       "Test Work Item",
		Description: gitlab.Ptr("This is a test work item"),
		Weight:      gitlab.Ptr(int64(100)),
	}

	// WHEN creating a new work item with the given options
	wi, err := CreateTestWorkItem(t, client, group.FullPath, gitlab.WorkItemTypeEpic, &opt)
	require.NoError(t, err, "CreateWorkItem failed")

	// THEN the work item should be created successfully
	assert.NotNil(t, wi)

	// AND all provided fields should be set correctly
	assert.Equal(t, "Test Work Item", wi.Title)
	assert.Equal(t, "This is a test work item", wi.Description)
	// assert.Equal(t, int64(100), wi.Weight)
}
