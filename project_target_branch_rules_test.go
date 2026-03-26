package gitlab

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTargetBranchRules_ListProjectTargetBranchRules(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `
			{
				"data": {
					"project": {
						"targetBranchRules": {
							"nodes": [
								{
									"id": "gid://gitlab/Projects::TargetBranchRule/1",
									"name": "feature/*",
									"targetBranch": "develop",
									"createdAt": "2024-01-15T10:00:00Z"
								},
								{
									"id": "gid://gitlab/Projects::TargetBranchRule/2",
									"name": "hotfix/*",
									"targetBranch": "main",
									"createdAt": "2024-01-16T10:00:00Z"
								}
							]
						}
					}
				}
			}
		`)
	})

	rules, _, err := client.Projects.ListProjectTargetBranchRules("mygroup/myproject")
	require.NoError(t, err)

	t1, err := time.Parse(time.RFC3339, "2024-01-15T10:00:00Z")
	require.NoError(t, err)
	t2, err := time.Parse(time.RFC3339, "2024-01-16T10:00:00Z")
	require.NoError(t, err)

	want := []TargetBranchRule{
		{ID: 1, Name: "feature/*", TargetBranch: "develop", CreatedAt: t1},
		{ID: 2, Name: "hotfix/*", TargetBranch: "main", CreatedAt: t2},
	}
	assert.Equal(t, want, rules)
}

func TestTargetBranchRules_ListProjectTargetBranchRules_projectNotFound(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"data": {"project": null}}`)
	})

	_, _, err := client.Projects.ListProjectTargetBranchRules("nonexistent/project")
	assert.ErrorIs(t, err, ErrNotFound)
}

func TestTargetBranchRules_CreateTargetBranchRule(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `
			{
				"data": {
					"projectTargetBranchRuleCreate": {
						"targetBranchRule": {
							"id": "gid://gitlab/Projects::TargetBranchRule/42",
							"name": "release/*",
							"targetBranch": "stable",
							"createdAt": "2024-03-01T08:00:00Z"
						},
						"errors": []
					}
				}
			}
		`)
	})

	opt := &CreateTargetBranchRuleOptions{
		Name:         "release/*",
		TargetBranch: "stable",
	}
	rule, _, err := client.Projects.CreateTargetBranchRule(1, opt)
	require.NoError(t, err)

	ts, err := time.Parse(time.RFC3339, "2024-03-01T08:00:00Z")
	require.NoError(t, err)

	want := &TargetBranchRule{
		ID:           42,
		Name:         "release/*",
		TargetBranch: "stable",
		CreatedAt:    ts,
	}
	assert.Equal(t, want, rule)
}

func TestTargetBranchRules_CreateTargetBranchRule_nilOpt(t *testing.T) {
	t.Parallel()

	_, client := setup(t)

	_, _, err := client.Projects.CreateTargetBranchRule(1, nil)
	assert.ErrorContains(t, err, "opt is required")
}

func TestTargetBranchRules_CreateTargetBranchRule_errors(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `
			{
				"data": {
					"projectTargetBranchRuleCreate": {
						"targetBranchRule": null,
						"errors": ["Name has already been taken"]
					}
				}
			}
		`)
	})

	opt := &CreateTargetBranchRuleOptions{Name: "release/*", TargetBranch: "stable"}
	_, _, err := client.Projects.CreateTargetBranchRule(1, opt)
	assert.ErrorContains(t, err, "Name has already been taken")
}

func TestTargetBranchRules_DeleteTargetBranchRule(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `
			{
				"data": {
					"projectTargetBranchRuleDestroy": {
						"errors": []
					}
				}
			}
		`)
	})

	_, err := client.Projects.DeleteTargetBranchRule(42)
	assert.NoError(t, err)
}

func TestTargetBranchRules_DeleteTargetBranchRule_errors(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `
			{
				"data": {
					"projectTargetBranchRuleDestroy": {
						"errors": ["Record not found"]
					}
				}
			}
		`)
	})

	_, err := client.Projects.DeleteTargetBranchRule(999)
	assert.ErrorContains(t, err, "Record not found")
}
