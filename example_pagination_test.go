package gitlab_test

import (
	"fmt"

	gitlab "gitlab.com/gitlab-org/api/client-go"
)

// Example_pagination demonstrates standard offset-based pagination.
// This automatically handles pagination using page numbers.
func Example_pagination() {
	// Note: The setupPaginationMock() function below is ONLY for the example purpose
	// and has nothing to do with how a user will use client-go.
	// In production, you would use a real authenticated GitLab client.
	client, server := setupPaginationMock(false)
	defer server.Close()

	// Configure pagination options
	opts := &gitlab.ListProjectsOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 10,
			Page:    1,
		},
		Owned: gitlab.Ptr(true),
	}

	// Scan2 automatically handles pagination
	count := 0
	for range gitlab.Must(gitlab.Scan2(func(p gitlab.PaginationOptionFunc) ([]*gitlab.Project, *gitlab.Response, error) {
		return client.Projects.ListProjects(opts, p)
	})) {
		count++
	}

	fmt.Printf("Iterated over %d projects\n", count)

	// Output: Iterated over 2 projects
}

// Example_keysetPagination demonstrates keyset-based pagination.
// Keyset pagination is more efficient for large datasets and prevents
// duplicates when data changes during pagination.
func Example_keysetPagination() {
	// Note: The setupPaginationMock() function below is ONLY for the example purpose
	// and has nothing to do with how a user will use client-go.
	// In production, you would use a real authenticated GitLab client.
	client, server := setupPaginationMock(true)
	defer server.Close()

	// Configure keyset pagination
	opts := &gitlab.ListProjectsOptions{
		ListOptions: gitlab.ListOptions{
			OrderBy:    "id",
			Pagination: "keyset",
			PerPage:    5,
			Sort:       "asc",
		},
		Owned: gitlab.Ptr(true),
	}

	// Scan2 works with both pagination types
	count := 0
	for range gitlab.Must(gitlab.Scan2(func(p gitlab.PaginationOptionFunc) ([]*gitlab.Project, *gitlab.Response, error) {
		return client.Projects.ListProjects(opts, p)
	})) {
		count++
	}

	fmt.Printf("Iterated over %d projects\n", count)

	// Output: Iterated over 2 projects
}
