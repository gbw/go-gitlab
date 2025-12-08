//go:build integration

package gitlab_test

import (
	"context"
	"fmt"
	"os"
	"strings"
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

// SkipIfNotLicensed skips the test if the GitLab instance doesn't have
// a Premium or Ultimate license. This is required to ensure that integration
// tests requiring licensed features don't fail on unlicensed instances.
func SkipIfNotLicensed(t *testing.T, client *gitlab.Client) {
	t.Helper()

	// Check if we have a valid license
	isLicensed, err := HasPremiumOrUltimateLicense(t, client)
	require.NoError(t, err, "Failed to determine GitLab license status")

	// Skip the test if not licensed
	if !isLicensed {
		t.Skip("Skipping test - requires GitLab Premium or Ultimate license")
	}
}

// Global variable to cache the license check result for all tests
var isLicensed *bool

// HasPremiumOrUltimateLicense calls the GitLab License API once and caches
// the result to determine if the instance has a Premium or Ultimate plan.
func HasPremiumOrUltimateLicense(t *testing.T, client *gitlab.Client) (bool, error) {
	t.Helper()

	if isLicensed != nil {
		return *isLicensed, nil
	}

	license, _, err := client.License.GetLicense()
	if err != nil {
		// If we can't get the license (e.g., no license installed), treat as unlicensed
		isLicensed = gitlab.Ptr(false)
		return false, nil
	}

	// Check if the plan is Premium or Ultimate (case insensitive, though it should always be lowercase)
	plan := strings.ToLower(license.Plan)
	result := (plan == "premium" || plan == "ultimate")

	isLicensed = &result
	return result, nil
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
