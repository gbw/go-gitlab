//
// Copyright 2021, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchService_Users(t *testing.T) {
	t.Parallel()

	t.Run("successful search", func(t *testing.T) {
		t.Parallel()
		mux, client := setup(t)

		mux.HandleFunc("/api/v4/search", func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			mustWriteHTTPResponse(t, w, "testdata/search_users.json")
		})

		opts := &SearchOptions{ListOptions: ListOptions{PerPage: 2}}
		users, resp, err := client.Search.Users("doe", opts)
		assert.NoError(t, err)
		assert.NotNil(t, resp)

		want := []*User{{
			ID:        1,
			Username:  "user1",
			Name:      "John Doe1",
			State:     "active",
			AvatarURL: "http://www.gravatar.com/avatar/c922747a93b40d1ea88262bf1aebee62?s=80&d=identicon",
			WebURL:    "http://localhost/user1",
		}}
		assert.Equal(t, want, users)
	})

	t.Run("error handling", func(t *testing.T) {
		t.Parallel()
		mux, client := setup(t)

		mux.HandleFunc("/api/v4/search", func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, `{"message": "Internal Server Error"}`)
		})

		opts := &SearchOptions{ListOptions: ListOptions{PerPage: 20}}
		users, resp, err := client.Search.Users("doe", opts)

		assert.Error(t, err)
		assert.Nil(t, users)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("empty search term", func(t *testing.T) {
		t.Parallel()
		mux, client := setup(t)

		mux.HandleFunc("/api/v4/search", func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, `[]`)
		})

		opts := &SearchOptions{ListOptions: ListOptions{PerPage: 20}}
		users, resp, err := client.Search.Users("", opts)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Empty(t, users)
	})
	t.Run("pagination - second page", func(t *testing.T) {
		t.Parallel()
		mux, client := setup(t)

		mux.HandleFunc("/api/v4/search", func(w http.ResponseWriter, r *http.Request) {
			testParam(t, r, "page", "2")
			testParam(t, r, "per_page", "1")
			testMethod(t, r, http.MethodGet)
			mustWriteHTTPResponse(t, w, "testdata/search_users_pagination.json")
		})

		opts := &SearchOptions{ListOptions: ListOptions{Page: 2, PerPage: 1}}
		users, _, err := client.Search.Users("doe", opts)

		assert.NoError(t, err)
		assert.Len(t, users, 1)
		assert.Equal(t, int64(2), users[0].ID)
	})
}

func TestSearchService_UsersByGroup(t *testing.T) {
	t.Parallel()

	t.Run("valid group ID - returns expected users", func(t *testing.T) {
		t.Parallel()
		mux, client := setup(t)

		mux.HandleFunc("/api/v4/groups/3/-/search", func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)

			query := r.URL.Query()
			if got := query.Get("search"); got != "doe" {
				t.Errorf("expected query 'search=doe', got 'search=%s'", got)
			}

			mustWriteHTTPResponse(t, w, "testdata/search_users.json")
		})

		users, _, err := client.Search.UsersByGroup("3", "doe", nil)

		assert.NoError(t, err)

		want := []*User{{
			ID:        1,
			Username:  "user1",
			Name:      "John Doe1",
			State:     "active",
			AvatarURL: "http://www.gravatar.com/avatar/c922747a93b40d1ea88262bf1aebee62?s=80&d=identicon",
			WebURL:    "http://localhost/user1",
		}}
		assert.Equal(t, want, users)
	})

	t.Run("invalid group ID - returns 404", func(t *testing.T) {
		t.Parallel()
		mux, client := setup(t)

		mux.HandleFunc("/api/v4/groups/invalid/-/search", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, `{"message": "404 Group Not Found"}`)
		})

		users, resp, err := client.Search.UsersByGroup("invalid", "doe", &SearchOptions{})
		assert.Error(t, err)
		assert.Nil(t, users)
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})
}

func TestSearchService_UsersByProject(t *testing.T) {
	t.Parallel()

	t.Run("valid project ID", func(t *testing.T) {
		t.Parallel()
		mux, client := setup(t)

		mux.HandleFunc("/api/v4/projects/6/-/search", func(w http.ResponseWriter, r *http.Request) {
			testParam(t, r, "search", "doe")
			testParam(t, r, "scope", "users")
			testMethod(t, r, http.MethodGet)
			mustWriteHTTPResponse(t, w, "testdata/search_users.json")
		})
		opts := &SearchOptions{ListOptions: ListOptions{PerPage: 2}}
		users, _, err := client.Search.UsersByProject("6", "doe", opts)

		assert.NoError(t, err)

		want := []*User{{
			ID:        1,
			Username:  "user1",
			Name:      "John Doe1",
			State:     "active",
			AvatarURL: "http://www.gravatar.com/avatar/c922747a93b40d1ea88262bf1aebee62?s=80&d=identicon",
			WebURL:    "http://localhost/user1",
		}}
		assert.Equal(t, want, users)
	})

	t.Run("invalid project ID", func(t *testing.T) {
		t.Parallel()
		mux, client := setup(t)

		mux.HandleFunc("/api/v4/projects/invalid/-/search", func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, `{"message": "404 Project Not Found"}`)
		})

		opts := &SearchOptions{ListOptions: ListOptions{PerPage: 2}}
		users, resp, err := client.Search.UsersByProject("invalid", "doe", opts)

		assert.Error(t, err)
		assert.Nil(t, users)
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})
}

func TestSearchService_Projects(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a search query for projects
	// WHEN searching for projects
	mux.HandleFunc("/api/v4/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testParam(t, r, "scope", "projects")
		testParam(t, r, "search", "gitlab")
		fmt.Fprint(w, `[{"id": 1, "name": "gitlab-ce"}]`)
	})

	opts := &SearchOptions{}

	// THEN matching projects should be returned
	projects, resp, err := client.Search.Projects("gitlab", opts)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, projects, 1)
	assert.Equal(t, int64(1), projects[0].ID)
	assert.Equal(t, "gitlab-ce", projects[0].Name)
}

func TestSearchService_ProjectsByGroup(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a group and a search query
	// WHEN searching for projects in the group
	mux.HandleFunc("/api/v4/groups/1/-/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testParam(t, r, "scope", "projects")
		testParam(t, r, "search", "test")
		fmt.Fprint(w, `[{"id": 2, "name": "test-project"}]`)
	})

	opts := &SearchOptions{}

	// THEN matching projects in the group should be returned
	projects, resp, err := client.Search.ProjectsByGroup(1, "test", opts)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, projects, 1)
	assert.Equal(t, int64(2), projects[0].ID)
}

func TestSearchService_Issues(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a search query for issues
	// WHEN searching for issues
	mux.HandleFunc("/api/v4/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testParam(t, r, "scope", "issues")
		testParam(t, r, "search", "bug")
		fmt.Fprint(w, `[{"id": 1, "title": "Bug in login"}]`)
	})

	opts := &SearchOptions{}

	// THEN matching issues should be returned
	issues, resp, err := client.Search.Issues("bug", opts)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, issues, 1)
	assert.Equal(t, int64(1), issues[0].ID)
}

func TestSearchService_IssuesByGroup(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a group and a search query
	// WHEN searching for issues in the group
	mux.HandleFunc("/api/v4/groups/1/-/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testParam(t, r, "scope", "issues")
		testParam(t, r, "search", "feature")
		fmt.Fprint(w, `[{"id": 2, "title": "New feature request"}]`)
	})

	opts := &SearchOptions{}

	// THEN matching issues in the group should be returned
	issues, resp, err := client.Search.IssuesByGroup(1, "feature", opts)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, issues, 1)
}

func TestSearchService_IssuesByProject(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a project and a search query
	// WHEN searching for issues in the project
	mux.HandleFunc("/api/v4/projects/1/-/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testParam(t, r, "scope", "issues")
		testParam(t, r, "search", "critical")
		fmt.Fprint(w, `[{"id": 3, "title": "Critical bug"}]`)
	})

	opts := &SearchOptions{}

	// THEN matching issues in the project should be returned
	issues, resp, err := client.Search.IssuesByProject(1, "critical", opts)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, issues, 1)
}

func TestSearchService_MergeRequests(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a search query for merge requests
	// WHEN searching for merge requests
	mux.HandleFunc("/api/v4/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testParam(t, r, "scope", "merge_requests")
		testParam(t, r, "search", "fix")
		fmt.Fprint(w, `[{"id": 1, "title": "Fix authentication"}]`)
	})

	opts := &SearchOptions{}

	// THEN matching merge requests should be returned
	mrs, resp, err := client.Search.MergeRequests("fix", opts)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, mrs, 1)
}

func TestSearchService_MergeRequestsByGroup(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a group and a search query
	// WHEN searching for merge requests in the group
	mux.HandleFunc("/api/v4/groups/1/-/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testParam(t, r, "scope", "merge_requests")
		testParam(t, r, "search", "refactor")
		fmt.Fprint(w, `[{"id": 2, "title": "Refactor code"}]`)
	})

	opts := &SearchOptions{}

	// THEN matching merge requests in the group should be returned
	mrs, resp, err := client.Search.MergeRequestsByGroup(1, "refactor", opts)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, mrs, 1)
}

func TestSearchService_MergeRequestsByProject(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a project and a search query
	// WHEN searching for merge requests in the project
	mux.HandleFunc("/api/v4/projects/1/-/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testParam(t, r, "scope", "merge_requests")
		testParam(t, r, "search", "update")
		fmt.Fprint(w, `[{"id": 3, "title": "Update dependencies"}]`)
	})

	opts := &SearchOptions{}

	// THEN matching merge requests in the project should be returned
	mrs, resp, err := client.Search.MergeRequestsByProject(1, "update", opts)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, mrs, 1)
}

func TestSearchService_Milestones(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a search query for milestones
	// WHEN searching for milestones
	mux.HandleFunc("/api/v4/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testParam(t, r, "scope", "milestones")
		testParam(t, r, "search", "v1.0")
		fmt.Fprint(w, `[{"id": 1, "title": "v1.0"}]`)
	})

	opts := &SearchOptions{}

	// THEN matching milestones should be returned
	milestones, resp, err := client.Search.Milestones("v1.0", opts)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, milestones, 1)
}

func TestSearchService_MilestonesByGroup(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a group and a search query
	// WHEN searching for milestones in the group
	mux.HandleFunc("/api/v4/groups/1/-/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testParam(t, r, "scope", "milestones")
		testParam(t, r, "search", "release")
		fmt.Fprint(w, `[{"id": 2, "title": "Release 2.0"}]`)
	})

	opts := &SearchOptions{}

	// THEN matching milestones in the group should be returned
	milestones, resp, err := client.Search.MilestonesByGroup(1, "release", opts)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, milestones, 1)
}

func TestSearchService_MilestonesByProject(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a project and a search query
	// WHEN searching for milestones in the project
	mux.HandleFunc("/api/v4/projects/1/-/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testParam(t, r, "scope", "milestones")
		testParam(t, r, "search", "sprint")
		fmt.Fprint(w, `[{"id": 3, "title": "Sprint 1"}]`)
	})

	opts := &SearchOptions{}

	// THEN matching milestones in the project should be returned
	milestones, resp, err := client.Search.MilestonesByProject(1, "sprint", opts)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, milestones, 1)
}

func TestSearchService_SnippetTitles(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a search query for snippet titles
	// WHEN searching for snippets
	mux.HandleFunc("/api/v4/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testParam(t, r, "scope", "snippet_titles")
		testParam(t, r, "search", "example")
		fmt.Fprint(w, `[{"id": 1, "title": "Example snippet"}]`)
	})

	opts := &SearchOptions{}

	// THEN matching snippets should be returned
	snippets, resp, err := client.Search.SnippetTitles("example", opts)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, snippets, 1)
}

func TestSearchService_NotesByProject(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a project and a search query
	// WHEN searching for notes in the project
	mux.HandleFunc("/api/v4/projects/1/-/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testParam(t, r, "scope", "notes")
		testParam(t, r, "search", "comment")
		fmt.Fprint(w, `[{"id": 1, "body": "This is a comment"}]`)
	})

	opts := &SearchOptions{}

	// THEN matching notes in the project should be returned
	notes, resp, err := client.Search.NotesByProject(1, "comment", opts)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, notes, 1)
}

func TestSearchService_WikiBlobs(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a search query for wiki blobs
	// WHEN searching for wiki blobs
	mux.HandleFunc("/api/v4/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testParam(t, r, "scope", "wiki_blobs")
		testParam(t, r, "search", "documentation")
		fmt.Fprint(w, `[{"basename": "home", "data": "documentation content"}]`)
	})

	opts := &SearchOptions{}

	// THEN matching wiki blobs should be returned
	wikis, resp, err := client.Search.WikiBlobs("documentation", opts)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, wikis, 1)
}

func TestSearchService_WikiBlobsByGroup(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a group and a search query
	// WHEN searching for wiki blobs in the group
	mux.HandleFunc("/api/v4/groups/1/-/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testParam(t, r, "scope", "wiki_blobs")
		testParam(t, r, "search", "guide")
		fmt.Fprint(w, `[{"basename": "guide", "data": "guide content"}]`)
	})

	opts := &SearchOptions{}

	// THEN matching wiki blobs in the group should be returned
	wikis, resp, err := client.Search.WikiBlobsByGroup(1, "guide", opts)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, wikis, 1)
}

func TestSearchService_WikiBlobsByProject(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a project and a search query
	// WHEN searching for wiki blobs in the project
	mux.HandleFunc("/api/v4/projects/1/-/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testParam(t, r, "scope", "wiki_blobs")
		testParam(t, r, "search", "tutorial")
		fmt.Fprint(w, `[{"basename": "tutorial", "data": "tutorial content"}]`)
	})

	opts := &SearchOptions{}

	// THEN matching wiki blobs in the project should be returned
	wikis, resp, err := client.Search.WikiBlobsByProject(1, "tutorial", opts)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, wikis, 1)
}

func TestSearchService_Commits(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a search query for commits
	// WHEN searching for commits
	mux.HandleFunc("/api/v4/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testParam(t, r, "scope", "commits")
		testParam(t, r, "search", "fix")
		fmt.Fprint(w, `[{"id": "abc123", "message": "fix: bug"}]`)
	})

	opts := &SearchOptions{}

	// THEN matching commits should be returned
	commits, resp, err := client.Search.Commits("fix", opts)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, commits, 1)
}

func TestSearchService_CommitsByGroup(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a group and a search query
	// WHEN searching for commits in the group
	mux.HandleFunc("/api/v4/groups/1/-/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testParam(t, r, "scope", "commits")
		testParam(t, r, "search", "feature")
		fmt.Fprint(w, `[{"id": "def456", "message": "feat: new feature"}]`)
	})

	opts := &SearchOptions{}

	// THEN matching commits in the group should be returned
	commits, resp, err := client.Search.CommitsByGroup(1, "feature", opts)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, commits, 1)
}

func TestSearchService_CommitsByProject(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a project and a search query
	// WHEN searching for commits in the project
	mux.HandleFunc("/api/v4/projects/1/-/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testParam(t, r, "scope", "commits")
		testParam(t, r, "search", "refactor")
		fmt.Fprint(w, `[{"id": "ghi789", "message": "refactor: code cleanup"}]`)
	})

	opts := &SearchOptions{}

	// THEN matching commits in the project should be returned
	commits, resp, err := client.Search.CommitsByProject(1, "refactor", opts)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, commits, 1)
}

func TestSearchService_Blobs(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a search query for blobs
	// WHEN searching for blobs
	mux.HandleFunc("/api/v4/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testParam(t, r, "scope", "blobs")
		testParam(t, r, "search", "function")
		fmt.Fprint(w, `[{"basename": "main.go", "data": "function code"}]`)
	})

	opts := &SearchOptions{}

	// THEN matching blobs should be returned
	blobs, resp, err := client.Search.Blobs("function", opts)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, blobs, 1)
}

func TestSearchService_BlobsByGroup(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a group and a search query
	// WHEN searching for blobs in the group
	mux.HandleFunc("/api/v4/groups/1/-/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testParam(t, r, "scope", "blobs")
		testParam(t, r, "search", "class")
		fmt.Fprint(w, `[{"basename": "user.rb", "data": "class User"}]`)
	})

	opts := &SearchOptions{}

	// THEN matching blobs in the group should be returned
	blobs, resp, err := client.Search.BlobsByGroup(1, "class", opts)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, blobs, 1)
}

func TestSearchService_BlobsByProject(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a project and a search query
	// WHEN searching for blobs in the project
	mux.HandleFunc("/api/v4/projects/1/-/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testParam(t, r, "scope", "blobs")
		testParam(t, r, "search", "interface")
		fmt.Fprint(w, `[{"basename": "api.go", "data": "interface API"}]`)
	})

	opts := &SearchOptions{}

	// THEN matching blobs in the project should be returned
	blobs, resp, err := client.Search.BlobsByProject(1, "interface", opts)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, blobs, 1)
}
