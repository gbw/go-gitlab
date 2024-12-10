# GitLab client-go (former `github.com/xanzy/go-gitlab`)

A GitLab API client enabling Go programs to interact with GitLab in a simple and uniform way

## Coverage

This API client package covers most of the existing GitLab API calls and is updated regularly
to add new and/or missing endpoints. Currently, the following services are supported:

- [x] Applications
- [x] Award Emojis
- [x] Branches
- [x] Broadcast Messages
- [x] Commits
- [x] Container Registry
- [x] Custom Attributes
- [x] Deploy Keys
- [x] Deployments
- [x] Discussions (threaded comments)
- [x] Environments
- [x] Epic Issues
- [x] Epics
- [x] Error Tracking
- [x] Events
- [x] Feature Flags
- [x] Geo Nodes
- [x] Generic Packages
- [x] GitLab CI Config Templates
- [x] Gitignores Templates
- [x] Group Access Requests
- [x] Group Issue Boards
- [x] Group Members
- [x] Group Milestones
- [x] Group Wikis
- [x] Group-Level Variables
- [x] Groups
- [x] Instance Clusters
- [x] Invites
- [x] Issue Boards
- [x] Issues
- [x] Jobs
- [x] Keys
- [x] Labels
- [x] License
- [x] Markdown
- [x] Merge Request Approvals
- [x] Merge Requests
- [x] Namespaces
- [x] Notes (comments)
- [x] Notification Settings
- [x] Open Source License Templates
- [x] Packages
- [x] Pages
- [x] Pages Domains
- [x] Personal Access Tokens
- [x] Pipeline Schedules
- [x] Pipeline Triggers
- [x] Pipelines
- [x] Plan limits
- [x] Project Access Requests
- [x] Project Badges
- [x] Project Clusters
- [x] Project Import/export
- [x] Project Members
- [x] Project Milestones
- [x] Project Repository Storage Moves
- [x] Project Snippets
- [x] Project Vulnerabilities
- [x] Project-Level Variables
- [x] Projects (including setting Webhooks)
- [x] Protected Branches
- [x] Protected Environments
- [x] Protected Tags
- [x] Repositories
- [x] Repository Files
- [x] Repository Submodules
- [x] Runners
- [x] Search
- [x] Services
- [x] Settings
- [x] Sidekiq Metrics
- [x] System Hooks
- [x] Tags
- [x] Todos
- [x] Topics
- [x] Users
- [x] Validate CI Configuration
- [x] Version
- [x] Wikis

## Usage

```go
import "gitlab.com/gitlab-org/api/client-go"
```

Construct a new GitLab client, then use the various services on the client to
access different parts of the GitLab API. For example, to list all
users:

```go
git, err := gitlab.NewClient("yourtokengoeshere")
if err != nil {
  log.Fatalf("Failed to create client: %v", err)
}
users, _, err := git.Users.ListUsers(&gitlab.ListUsersOptions{})
```

There are a few `With...` option functions that can be used to customize
the API client. For example, to set a custom base URL:

```go
git, err := gitlab.NewClient("yourtokengoeshere", gitlab.WithBaseURL("https://git.mydomain.com/api/v4"))
if err != nil {
  log.Fatalf("Failed to create client: %v", err)
}
users, _, err := git.Users.ListUsers(&gitlab.ListUsersOptions{})
```

Some API methods have optional parameters that can be passed. For example,
to list all projects for user "svanharmelen":

```go
git := gitlab.NewClient("yourtokengoeshere")
opt := &gitlab.ListProjectsOptions{Search: gitlab.Ptr("svanharmelen")}
projects, _, err := git.Projects.ListProjects(opt)
```

### Examples

The [examples](https://gitlab.com/gitlab-org/api/client-go/tree/master/examples) directory
contains a couple for clear examples, of which one is partially listed here as well:

```go
package main

import (
	"log"

	"gitlab.com/gitlab-org/api/client-go"
)

func main() {
	git, err := gitlab.NewClient("yourtokengoeshere")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Create new project
	p := &gitlab.CreateProjectOptions{
		Name:                     gitlab.Ptr("My Project"),
		Description:              gitlab.Ptr("Just a test project to play with"),
		MergeRequestsAccessLevel: gitlab.Ptr(gitlab.EnabledAccessControl),
		SnippetsAccessLevel:      gitlab.Ptr(gitlab.EnabledAccessControl),
		Visibility:               gitlab.Ptr(gitlab.PublicVisibility),
	}
	project, _, err := git.Projects.CreateProject(p)
	if err != nil {
		log.Fatal(err)
	}

	// Add a new snippet
	s := &gitlab.CreateProjectSnippetOptions{
		Title:           gitlab.Ptr("Dummy Snippet"),
		FileName:        gitlab.Ptr("snippet.go"),
		Content:         gitlab.Ptr("package main...."),
		Visibility:      gitlab.Ptr(gitlab.PublicVisibility),
	}
	_, _, err = git.ProjectSnippets.CreateSnippet(project.ID, s)
	if err != nil {
		log.Fatal(err)
	}
}
```

For complete usage of go-gitlab, see the full [package docs](https://godoc.org/gitlab.com/gitlab-org/api/client-go).

## Contributing

Contributions are always welcome. For more information, check out the
[contributing guide](https://gitlab.com/gitlab-org/api/client-go/-/blob/main/CONTRIBUTING.md).

## Maintenance

This is a community maintained project. If you have a paid GitLab subscription,
please note that this project is not packaged as a part of GitLab, and falls outside
of the scope of support.

For more information, see GitLab's
[Statement of Support](https://about.gitlab.com/support/statement-of-support.html).
Please fill out an issue in this projects issue tracker and someone from the community
will respond as soon as they are available to help you.

### Known GitLab Projects using this package

- [GitLab Terraform Provider](https://gitlab.com/gitlab-org/terraform-provider-gitlab)
  maintained by the community with support from ~"group::environments"
- [GitLab CLI (`glab`)](https://gitlab.com/gitlab-org/cli)
  maintained by ~"group::code review"
