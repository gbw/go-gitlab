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

// Integration tests for the System Hooks API.
// These tests require a GitLab instance running on localhost:8095.
// They also require a valid admin token in GITLAB_TOKEN environment variable.

func createTestSystemHook(t *testing.T, client *gitlab.Client) *gitlab.Hook {
	t.Helper()

	suffix := time.Now().UnixNano()
	hookURL := fmt.Sprintf("https://example.com/%d", suffix)

	hook, _, err := client.SystemHooks.AddHook(&gitlab.AddHookOptions{
		URL:        &hookURL,
		PushEvents: gitlab.Ptr(true),
	})
	require.NoError(t, err, "Failed to create test system hook")

	t.Cleanup(func() {
		_, err := client.SystemHooks.DeleteHook(hook.ID)
		if err != nil && err.Error() == "404 Not Found" {
			t.Logf("System hook %d already deleted", hook.ID)
			return
		}
		require.NoError(t, err, "Failed to delete test system hook")
	})

	return hook
}

func Test_SystemHooksListHooks_Integration(t *testing.T) {
	client := SetupIntegrationClient(t)

	hook := createTestSystemHook(t, client)

	hooks, _, err := client.SystemHooks.ListHooks()
	require.NoError(t, err, "Failed to list system hooks")

	assert.NotNil(t, hooks)
	assert.GreaterOrEqual(t, len(hooks), 1)

	var found bool
	for _, h := range hooks {
		if h.ID == hook.ID {
			found = true
			break
		}
	}
	assert.True(t, found, "Created hook not found in list")
}

func Test_SystemHooksGetHook_Integration(t *testing.T) {
	client := SetupIntegrationClient(t)

	hook := createTestSystemHook(t, client)

	retrieved, _, err := client.SystemHooks.GetHook(hook.ID)
	require.NoError(t, err, "Failed to get system hook")

	assert.Equal(t, hook.ID, retrieved.ID)
	assert.Equal(t, hook.URL, retrieved.URL)
	assert.True(t, retrieved.PushEvents)
}

func Test_SystemHooksAddHook_Integration(t *testing.T) {
	client := SetupIntegrationClient(t)

	suffix := time.Now().UnixNano()
	hookURL := fmt.Sprintf("https://example.com/%d", suffix)

	hook, _, err := client.SystemHooks.AddHook(&gitlab.AddHookOptions{
		URL:                    &hookURL,
		PushEvents:             gitlab.Ptr(true),
		TagPushEvents:          gitlab.Ptr(true),
		MergeRequestsEvents:    gitlab.Ptr(true),
		RepositoryUpdateEvents: gitlab.Ptr(true),
		EnableSSLVerification:  gitlab.Ptr(true),
		Token:                  gitlab.Ptr("secret-token"),
	})
	require.NoError(t, err, "Failed to add system hook")

	t.Cleanup(func() {
		_, err := client.SystemHooks.DeleteHook(hook.ID)
		require.NoError(t, err, "Failed to delete test system hook")
	})

	assert.NotZero(t, hook.ID)
	assert.Equal(t, hookURL, hook.URL)
	assert.True(t, hook.PushEvents)
	assert.True(t, hook.TagPushEvents)
	assert.True(t, hook.MergeRequestsEvents)
	assert.True(t, hook.RepositoryUpdateEvents)
	assert.True(t, hook.EnableSSLVerification)
}

func Test_SystemHooksEditHook_Integration(t *testing.T) {
	client := SetupIntegrationClient(t)

	hook := createTestSystemHook(t, client)

	suffix := time.Now().UnixNano()
	hookURL := fmt.Sprintf("https://example.com/%d", suffix)

	updated, _, err := client.SystemHooks.EditHook(hook.ID, &gitlab.EditHookOptions{
		URL:                    &hookURL,
		PushEvents:             gitlab.Ptr(true),
		TagPushEvents:          gitlab.Ptr(true),
		MergeRequestsEvents:    gitlab.Ptr(true),
		RepositoryUpdateEvents: gitlab.Ptr(true),
		EnableSSLVerification:  gitlab.Ptr(true),
		Token:                  gitlab.Ptr("secret-token"),
	})
	require.NoError(t, err, "Failed to edit system hook")

	assert.Equal(t, hook.ID, updated.ID)
	assert.Equal(t, hookURL, updated.URL)
	assert.True(t, updated.PushEvents)
	assert.True(t, updated.TagPushEvents)
	assert.True(t, updated.MergeRequestsEvents)
	assert.True(t, updated.RepositoryUpdateEvents)
	assert.True(t, updated.EnableSSLVerification)
}

func Test_SystemHooksDeleteHook_Integration(t *testing.T) {
	client := SetupIntegrationClient(t)

	hook := createTestSystemHook(t, client)

	_, err := client.SystemHooks.DeleteHook(hook.ID)
	require.NoError(t, err, "Failed to delete system hook")
}
