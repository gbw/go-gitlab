//go:build integration

package gitlab_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

// This file contains helper functions that are useful for
// writing tests. This includes a helper to create a client
// related to acceptance tests.
func SetupIntegrationClient(t *testing.T) *gitlab.Client {
	t.Helper()

	// Get the token from environment
	token := os.Getenv("GITLAB_TOKEN")
	if token == "" {
		t.Skip("GITLAB_TOKEN environment variable not set")
	}

	// Get the baseUrl from environment. If it's not set, default
	// to the local setup.
	baseURL := os.Getenv("GITLAB_BASE_URL")
	if baseURL == "" {
		baseURL = "https://localhost:8095/api/v4"
	}

	// Return a client with the base URL and the token.
	client, err := gitlab.NewClient(token, gitlab.WithBaseURL(baseURL))
	require.NoError(t, err, "failed to create GitLab Client for BaseURL "+baseURL)

	return client
}

// Skips the given test after the client is configured when running in CE
// This is required to ensure that integration testing functions that require
// an EE instance don't fail
func SkipIfRunningCE(t *testing.T, client *gitlab.Client) {
	t.Helper()

	// Check if we're running in CE context
	isEE, err := IsRunningInEEContext(t, client)
	require.NoError(t, err, "Failed to determine GitLab edition")

	// Skip the test if running on CE
	if !isEE {
		t.Skip("Skipping test - requires GitLab Enterprise Edition")
	}
}

// Global variable to cache the result of EE evaluation for all the tests
var isEE *bool

// function calls gitlab server metadata API once and caches the result
// to determine if license model is enterprise or not
func IsRunningInEEContext(t *testing.T, client *gitlab.Client) (bool, error) {
	t.Helper()

	if isEE != nil {
		return *isEE, nil
	}
	metadata, _, err := client.Metadata.GetMetadata()
	if err != nil {
		return false, err
	}

	// Cache the results for later.
	// Note - if run on versions earlier to 15.5, it will error since
	// this key wasn't returned. With we're 3 major versions later, this
	// seems like a safe assumption.
	isEE = &metadata.Enterprise
	return *isEE, err
}

// CreateTestUser creates a test user with a random username and email.
// The user is automatically cleaned up when the test finishes.
func CreateTestUser(t *testing.T, client *gitlab.Client) (*gitlab.User, error) {
	t.Helper()

	// Generate random username and email
	suffix := time.Now().UnixNano()

	username := fmt.Sprintf("testuser%d", suffix)
	email := fmt.Sprintf("testuser%d@example.com", suffix)
	name := fmt.Sprintf("Test User %d", suffix)

	// Create the user
	user, _, err := client.Users.CreateUser(&gitlab.CreateUserOptions{
		Username: &username,
		Email:    &email,
		Name:     &name,
		// Required field - must be fairly random or GitLab won't allow it
		// nosemgrep - testing password
		Password:         gitlab.Ptr("f0hYXux#yy2CFypKq!aV"),
		SkipConfirmation: gitlab.Ptr(true), // Skip email confirmation
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Clean up the user when the test finishes
	t.Cleanup(func() {
		_, err := client.Users.DeleteUser(user.ID)
		require.NoError(t, err, "Failed to delete test user")
	})

	return user, nil
}

// CreateTestProject creates a test Project with a random name. It'll be set to `public` to
// help with other testing pieces.
//
// The user is automatically cleaned up when the test finishes.
func CreateTestProject(t *testing.T, client *gitlab.Client) *gitlab.Project {
	t.Helper()

	suffix := time.Now().UnixNano()
	return CreateTestProjectWithOptions(t, client, &gitlab.CreateProjectOptions{
		Name:       gitlab.Ptr(fmt.Sprintf("project%d", suffix)),
		Visibility: gitlab.Ptr(gitlab.PublicVisibility),
	})
}

// CreateTestProjectWithOptions creates a test Project with the provided options.
//
// The user is automatically cleaned up when the test finishes.
func CreateTestProjectWithOptions(t *testing.T, client *gitlab.Client, opts *gitlab.CreateProjectOptions) *gitlab.Project {
	t.Helper()

	project, _, err := client.Projects.CreateProject(opts, gitlab.WithContext(context.Background()))
	if err != nil {
		t.Fatalf("Failed to create project: %v", err)
	}

	// Add a cleanup function
	t.Cleanup(func() {
		_, _ = client.Projects.DeleteProject(project.ID, nil, gitlab.WithContext(context.Background()))
	})

	return project
}

// CreateTestGroup creates a test group with a random name.
// The group is automatically cleaned up when the test finishes.
func CreateTestGroup(t *testing.T, client *gitlab.Client) (*gitlab.Group, error) {
	t.Helper()

	// Generate random name
	suffix := time.Now().UnixNano()
	name := fmt.Sprintf("testgroup%d", suffix)

	// Create the group
	group, _, err := client.Groups.CreateGroup(&gitlab.CreateGroupOptions{
		Name:       &name,
		Path:       &name,
		Visibility: gitlab.Ptr(gitlab.PublicVisibility),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create group: %w", err)
	}

	// Clean up the group when the test finishes
	t.Cleanup(func() {
		_, err := client.Groups.DeleteGroup(group.ID, &gitlab.DeleteGroupOptions{})
		require.NoError(t, err, "Failed to delete test group")
	})

	return group, nil
}

// CreateTestGroupHook creates a test group hook with a random url.
// The group hook is automatically cleaned up when the test finishes.
func CreateTestGroupHook(t *testing.T, gid any, client *gitlab.Client) (*gitlab.GroupHook, error) {
	t.Helper()

	// Generate random name
	suffix := time.Now().UnixNano()
	url := fmt.Sprintf("https://example.com/%d", suffix)

	// Create the group
	hook, _, err := client.Groups.AddGroupHook(gid, &gitlab.AddGroupHookOptions{
		URL:        &url,
		PushEvents: gitlab.Ptr(true),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create group hook: %w", err)
	}

	// Clean up the group hook when the test finishes
	t.Cleanup(func() {
		_, err := client.Groups.DeleteGroupHook(gid, hook.ID)
		if err != nil && err.Error() == "404 Not Found" {
			t.Logf("Group hook %d already deleted", hook.ID)
			return
		}
		require.NoError(t, err, "Failed to delete test group hook")
	})

	return hook, nil
}
