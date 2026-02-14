package gitlab_test

import (
	"fmt"

	gitlab "gitlab.com/gitlab-org/api/client-go"
)

// Example_fileUpload demonstrates uploading a file to a GitLab repository.
func Example_fileUpload() {
	// Note: The setupFileUploadMock() function below is ONLY for the example purpose
	// and has nothing to do with how a user will use client-go.
	// In production, you would use a real authenticated GitLab client.
	client, server := setupFileUploadMock()
	defer server.Close()

	// Create a new file in the repository
	opts := &gitlab.CreateFileOptions{
		Branch:        gitlab.Ptr("main"),
		Content:       gitlab.Ptr("# My Project\n\nDocumentation for this project."),
		CommitMessage: gitlab.Ptr("Add README"),
	}

	file, _, _ := client.RepositoryFiles.CreateFile(
		"my-group/my-project",
		"README.md",
		opts,
	)

	fmt.Printf("Created %s on branch %s\n", file.FilePath, file.Branch)

	// Output:
	// Created README.md on branch main
}
