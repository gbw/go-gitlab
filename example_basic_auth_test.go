package gitlab_test

import "fmt"

// Example_basicAuth demonstrates authenticating with username and password.
// This uses OAuth2 password credentials flow under the hood.
func Example_basicAuth() {
	// Note: The setupBasicAuthMock() function below is ONLY for the example purpose
	// and has nothing to do with how a user will use client-go.
	// In production, you would authenticate against a real GitLab instance.
	client, server := setupBasicAuthMock()
	defer server.Close()

	// Use the authenticated client
	projects, _, _ := client.Projects.ListProjects(nil)
	fmt.Printf("Found %d projects\n", len(projects))

	// Output: Found 2 projects
}
