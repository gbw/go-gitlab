package gitlab_test

import (
	"encoding/json"
	"os"

	gitlab "gitlab.com/gitlab-org/api/client-go"
)

func ExampleScanAndCollectN() {
	// Create a client (this would normally use your GitLab instance URL and token)
	client, err := gitlab.NewAuthSourceClient(
		gitlab.AccessTokenAuthSource{"your-token"},
		gitlab.WithBaseURL("https://gitlab.example.com/api/v4"),
	)
	if err != nil {
		// Handle the error
		panic(err)
	}

	opts := &gitlab.ListProjectsOptions{}

	pager := func(pageOpt gitlab.PaginationOptionFunc) ([]*gitlab.Project, *gitlab.Response, error) {
		// Call ListProjects with pageOpt to retrieve the next page
		return client.Projects.ListProjects(opts, pageOpt)
	}

	// Retrieve at most 42 projects
	const limit = 42

	projects, err := gitlab.ScanAndCollectN(pager, limit)
	if err != nil {
		// Handle the error
		panic(err)
	}

	// Use the slice — here we serialize it to JSON, but you could sort it, pass it to another function, etc.
	// Note: if you want to iterate over items, use gitlab.Scan2() instead
	if err := json.NewEncoder(os.Stdout).Encode(projects); err != nil {
		panic(err)
	}
}
