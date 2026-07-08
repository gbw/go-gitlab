package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSecurityScanProfiles_AttachSecurityScanProfile(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testGraphQLRequestBody(t, r, GraphQLQuery{
			Query: securityScanProfileAttachMutation,
			Variables: map[string]any{
				"input": map[string]any{
					"securityScanProfileId": SecurityScanProfileGID("dependency_scanning_post_processing"),
					"projectIds":            []string{"gid://gitlab/Project/1"},
					"groupIds":              []string{},
				},
			},
		})
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `
			{
				"data": {
					"securityScanProfileAttach": {
						"errors": []
					}
				}
			}
		`)
	})

	opt := &AttachSecurityScanProfileOptions{
		SecurityScanProfileID: SecurityScanProfileGID("dependency_scanning_post_processing"),
		ProjectIDs:            []int64{1},
	}
	_, err := client.SecurityScanProfiles.AttachSecurityScanProfile(opt)
	require.NoError(t, err)
}

func TestSecurityScanProfiles_AttachSecurityScanProfile_nilOpt(t *testing.T) {
	t.Parallel()

	_, client := setup(t)

	_, err := client.SecurityScanProfiles.AttachSecurityScanProfile(nil)
	assert.ErrorContains(t, err, "opt is required")
}

func TestSecurityScanProfiles_AttachSecurityScanProfile_errors(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testGraphQLRequestBody(t, r, GraphQLQuery{
			Query: securityScanProfileAttachMutation,
			Variables: map[string]any{
				"input": map[string]any{
					"securityScanProfileId": SecurityScanProfileGID("dependency_scanning_post_processing"),
					"projectIds":            []string{"gid://gitlab/Project/1"},
					"groupIds":              []string{},
				},
			},
		})
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `
			{
				"data": {
					"securityScanProfileAttach": {
						"errors": ["The resource that you are attempting to access does not exist"]
					}
				}
			}
		`)
	})

	opt := &AttachSecurityScanProfileOptions{
		SecurityScanProfileID: SecurityScanProfileGID("dependency_scanning_post_processing"),
		ProjectIDs:            []int64{1},
	}
	_, err := client.SecurityScanProfiles.AttachSecurityScanProfile(opt)
	assert.ErrorContains(t, err, "does not exist")
}

func TestSecurityScanProfiles_AttachSecurityScanProfile_topLevelErrors(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	// A disabled feature flag or missing permission comes back as a top-level
	// GraphQL error with HTTP 200 and a null data payload.
	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testGraphQLRequestBody(t, r, GraphQLQuery{
			Query: securityScanProfileAttachMutation,
			Variables: map[string]any{
				"input": map[string]any{
					"securityScanProfileId": SecurityScanProfileGID("dependency_scanning_post_processing"),
					"projectIds":            []string{"gid://gitlab/Project/1"},
					"groupIds":              []string{},
				},
			},
		})
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `
			{
				"errors": [
					{"message": "The resource that you are attempting to access does not exist or you don't have permission to perform this action"}
				],
				"data": {"securityScanProfileAttach": null}
			}
		`)
	})

	opt := &AttachSecurityScanProfileOptions{
		SecurityScanProfileID: SecurityScanProfileGID("dependency_scanning_post_processing"),
		ProjectIDs:            []int64{1},
	}
	_, err := client.SecurityScanProfiles.AttachSecurityScanProfile(opt)
	assert.ErrorContains(t, err, "don't have permission")
}

func TestSecurityScanProfiles_DetachSecurityScanProfile(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testGraphQLRequestBody(t, r, GraphQLQuery{
			Query: securityScanProfileDetachMutation,
			Variables: map[string]any{
				"input": map[string]any{
					"securityScanProfileId": SecurityScanProfileGID("dependency_scanning_post_processing"),
					"projectIds":            []string{"gid://gitlab/Project/1"},
					"groupIds":              []string{},
				},
			},
		})
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `
			{
				"data": {
					"securityScanProfileDetach": {
						"errors": []
					}
				}
			}
		`)
	})

	opt := &DetachSecurityScanProfileOptions{
		SecurityScanProfileID: SecurityScanProfileGID("dependency_scanning_post_processing"),
		ProjectIDs:            []int64{1},
	}
	_, err := client.SecurityScanProfiles.DetachSecurityScanProfile(opt)
	require.NoError(t, err)
}

func TestSecurityScanProfiles_DetachSecurityScanProfile_nilOpt(t *testing.T) {
	t.Parallel()

	_, client := setup(t)

	_, err := client.SecurityScanProfiles.DetachSecurityScanProfile(nil)
	assert.ErrorContains(t, err, "opt is required")
}

func TestSecurityScanProfiles_DetachSecurityScanProfile_errors(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testGraphQLRequestBody(t, r, GraphQLQuery{
			Query: securityScanProfileDetachMutation,
			Variables: map[string]any{
				"input": map[string]any{
					"securityScanProfileId": SecurityScanProfileGID("dependency_scanning_post_processing"),
					"projectIds":            []string{"gid://gitlab/Project/1"},
					"groupIds":              []string{},
				},
			},
		})
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `
			{
				"data": {
					"securityScanProfileDetach": {
						"errors": ["Record not found"]
					}
				}
			}
		`)
	})

	opt := &DetachSecurityScanProfileOptions{
		SecurityScanProfileID: SecurityScanProfileGID("dependency_scanning_post_processing"),
		ProjectIDs:            []int64{1},
	}
	_, err := client.SecurityScanProfiles.DetachSecurityScanProfile(opt)
	assert.ErrorContains(t, err, "Record not found")
}

func TestSecurityScanProfiles_ListProjectScanProfileStatuses(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testGraphQLRequestBody(t, r, GraphQLQuery{
			Query: scanProfileStatusesQuery,
			Variables: map[string]any{
				"fullPath": "mygroup/myproject",
			},
		})
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `
			{
				"data": {
					"project": {
						"scanProfileStatuses": [
							{
								"status": "ACTIVE",
								"scanProfile": {
									"id": "gid://gitlab/Security::ScanProfile/dependency_scanning_post_processing",
									"name": "Dependency Scanning Auto-Remediation (default)",
									"scanType": "DEPENDENCY_SCANNING_POST_PROCESSING"
								}
							}
						]
					}
				}
			}
		`)
	})

	statuses, _, err := client.SecurityScanProfiles.ListProjectScanProfileStatuses("mygroup/myproject")
	require.NoError(t, err)

	want := []ScanProfileStatus{
		{
			Status: "ACTIVE",
			ScanProfile: ScanProfile{
				ID:       "gid://gitlab/Security::ScanProfile/dependency_scanning_post_processing",
				Name:     "Dependency Scanning Auto-Remediation (default)",
				ScanType: "DEPENDENCY_SCANNING_POST_PROCESSING",
			},
		},
	}
	assert.Equal(t, want, statuses)
}

func TestSecurityScanProfiles_ListProjectScanProfileStatuses_projectNotFound(t *testing.T) {
	t.Parallel()

	mux, client := setup(t)

	mux.HandleFunc("/api/graphql", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testGraphQLRequestBody(t, r, GraphQLQuery{
			Query: scanProfileStatusesQuery,
			Variables: map[string]any{
				"fullPath": "nonexistent/project",
			},
		})
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"data": {"project": null}}`)
	})

	_, _, err := client.SecurityScanProfiles.ListProjectScanProfileStatuses("nonexistent/project")
	assert.ErrorIs(t, err, ErrNotFound)
}
