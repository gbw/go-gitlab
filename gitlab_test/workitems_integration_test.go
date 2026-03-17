//go:build integration

package gitlab_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	gitlab "gitlab.com/gitlab-org/api/client-go/v2"
)

func TestWorkItemLifeCycle(t *testing.T) {
	t.Parallel()

	client := SetupIntegrationClient(t)
	SkipIfNotLicensed(t, client)

	// GIVEN a test group
	group := CreateTestGroup(t, client)

	// STEP 1: Create a work item
	// WHEN creating a new work item with initial options
	createOpt := gitlab.CreateWorkItemOptions{
		Title:        "Integration Test Work Item",
		Description:  gitlab.Ptr("Initial description"),
		HealthStatus: gitlab.Ptr("onTrack"),
		Color:        gitlab.Ptr("green"),
	}

	createdWI, _, err := client.WorkItems.CreateWorkItem(group.FullPath, gitlab.WorkItemTypeEpic, &createOpt)
	require.NoError(t, err, "CreateWorkItem failed")
	require.NotNil(t, createdWI)

	// clean up in case test fails too early
	t.Cleanup(func() {
		_, err := client.WorkItems.DeleteWorkItem(group.FullPath, createdWI.IID, gitlab.WithContext(context.Background()))
		if err != nil && errors.Is(err, gitlab.ErrNotFound) {
			return
		}
		require.NoError(t, err, "Failed to delete test work item in cleanup")
	})

	// THEN the work item should have the provided fields set correctly
	assert.Equal(t, "Integration Test Work Item", createdWI.Title, "Field: Title")
	assert.Equal(t, "Initial description", createdWI.Description, "Field: Description")
	assert.Equal(t, "onTrack", deref(t, createdWI.HealthStatus), "Field: HealthStatus")
	assert.Equal(t, "#008000", deref(t, createdWI.Color), "Field: Color")

	// STEP 2: Get the work item
	// WHEN retrieving the work item by full path and IID
	gotWI, _, err := client.WorkItems.GetWorkItem(group.FullPath, createdWI.IID)
	require.NoError(t, err, "GetWorkItem failed")
	require.NotNil(t, gotWI)

	// THEN the retrieved work item should have the same provided fields
	assert.Equal(t, createdWI.Title, gotWI.Title, "Field: Title")
	assert.Equal(t, createdWI.Description, gotWI.Description, "Field: Description")
	assert.Equal(t, deref(t, createdWI.HealthStatus), deref(t, gotWI.HealthStatus), "Field: HealthStatus")
	assert.Equal(t, deref(t, createdWI.Color), deref(t, gotWI.Color), "Field: Color")

	// STEP 3: Update the work item
	// WHEN updating the work item with new values
	updateOpt := gitlab.UpdateWorkItemOptions{
		Title:        gitlab.Ptr("Updated Work Item Title"),
		Description:  gitlab.Ptr("Updated description"),
		HealthStatus: gitlab.Ptr("needsAttention"),
		Color:        gitlab.Ptr("red"),
	}

	updatedWI, _, err := client.WorkItems.UpdateWorkItem(group.FullPath, createdWI.IID, &updateOpt)
	require.NoError(t, err, "UpdateWorkItem failed")
	require.NotNil(t, updatedWI)

	// THEN the work item should have the updated fields set correctly
	assert.Equal(t, "Updated Work Item Title", updatedWI.Title, "Field: Title")
	assert.Equal(t, "Updated description", updatedWI.Description, "Field: Description")
	assert.Equal(t, "needsAttention", deref(t, updatedWI.HealthStatus), "Field: HealthStatus")
	assert.Equal(t, "#FF0000", deref(t, updatedWI.Color), "Field: Color")

	// STEP 4: Get the work item again
	// WHEN retrieving the work item after update
	finalWI, _, err := client.WorkItems.GetWorkItem(group.FullPath, createdWI.IID)
	require.NoError(t, err, "GetWorkItem after update failed")
	require.NotNil(t, finalWI)

	// THEN the retrieved work item should have the same updated fields
	assert.Equal(t, updatedWI.Title, finalWI.Title, "Field: Title")
	assert.Equal(t, updatedWI.Description, finalWI.Description, "Field: Description")
	assert.Equal(t, deref(t, updatedWI.HealthStatus), deref(t, finalWI.HealthStatus), "Field: HealthStatus")
	assert.Equal(t, deref(t, updatedWI.Color), deref(t, finalWI.Color), "Field: Color")

	// STEP 5: Delete the work item
	// WHEN deleting the work item
	_, err = client.WorkItems.DeleteWorkItem(group.FullPath, createdWI.IID)
	require.NoError(t, err, "DeleteWorkItem failed")

	// THEN the work item should no longer be retrievable
	_, _, err = client.WorkItems.GetWorkItem(group.FullPath, createdWI.IID)
	require.ErrorIs(t, err, gitlab.ErrNotFound)
}

func deref(t *testing.T, ptr *string) string {
	t.Helper()

	if ptr == nil {
		t.Fatal("pointer is nil")
	}

	return *ptr
}
