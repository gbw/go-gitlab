// Copyright 2021, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gitlab

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProjectTokenAccessSettings(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// Handle project ID 1, and print a result of access settings
	mux.HandleFunc("/api/v4/projects/1/job_token_scope", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		// Print on the response
		fmt.Fprint(w, `{"inbound_enabled":true}`)
	})

	want := &JobTokenAccessSettings{
		InboundEnabled: true,
	}

	settings, _, err := client.JobTokenScope.GetProjectJobTokenAccessSettings(1)

	assert.NoError(t, err)
	assert.Equal(t, want, settings)
}

func TestPatchProjectJobTokenAccessSettings(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/job_token_scope", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		testBodyJSON(t, r, map[string]any{
			"enabled": false,
		})

		// Ensure we provide the proper response
		w.WriteHeader(http.StatusNoContent)

		// Print an empty body, since that's what the API provides.
		fmt.Fprint(w, "")
	})

	resp, err := client.JobTokenScope.PatchProjectJobTokenAccessSettings(
		1,
		&PatchProjectJobTokenAccessSettingsOptions{
			Enabled: false,
		},
	)
	assert.NoError(t, err)
	assert.Equal(t, 204, resp.StatusCode)
}

// This tests that when calling the GetProjectJobTokenInboundAllowList, we get a
// list of projects back properly. There isn't a "deep" test with every attribute
// specified, because the object returned is a *Project object, which is already
// tested in project.go.
func TestGetProjectJobTokenInboundAllowList(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// Handle project ID 1, and print a result of two projects
	mux.HandleFunc("/api/v4/projects/1/job_token_scope/allowlist", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		// Print on the response
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	want := []*Project{{ID: 1}, {ID: 2}}
	projects, _, err := client.JobTokenScope.GetProjectJobTokenInboundAllowList(
		1,
		&GetJobTokenInboundAllowListOptions{},
	)

	assert.NoError(t, err)
	assert.Equal(t, want, projects)
}

func TestAddProjectToJobScopeAllowList(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/job_token_scope/allowlist", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testBodyJSON(t, r, map[string]any{
			"target_project_id": 2.0,
		})

		// Ensure we provide the proper response
		w.WriteHeader(http.StatusCreated)

		// Print on the response with the proper target project
		fmt.Fprintf(w, `{
			"source_project_id": 1,
			"target_project_id": 2
		}`)
	})

	want := &JobTokenInboundAllowItem{
		SourceProjectID: 1,
		TargetProjectID: 2,
	}

	addTokenResponse, resp, err := client.JobTokenScope.AddProjectToJobScopeAllowList(
		1,
		&JobTokenInboundAllowOptions{TargetProjectID: Ptr(int64(2))},
	)
	assert.NoError(t, err)
	assert.Equal(t, want, addTokenResponse)
	assert.Equal(t, 201, resp.StatusCode)
}

func TestRemoveProjectFromJobScopeAllowList(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/job_token_scope/allowlist/2", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)

		// Ensure we provide the proper response
		w.WriteHeader(http.StatusNoContent)

		// Print an empty body, since that's what the API provides.
		fmt.Fprint(w, "")
	})

	resp, err := client.JobTokenScope.RemoveProjectFromJobScopeAllowList(1, 2)
	assert.NoError(t, err)
	assert.Equal(t, 204, resp.StatusCode)
}

// This tests that when calling the GetJobTokenAllowlistGroups, we get a list
// of groups back. There isn't a "deep" test with every attribute specified,
// because the object returned is a *Group object, which is already tested in
// groups.go.
func TestGetJobTokenAllowlistGroups(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// Handle project ID 1, and print a result of two groups
	mux.HandleFunc("/api/v4/projects/1/job_token_scope/groups_allowlist", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		// Print on the response
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	want := []*Group{{ID: 1}, {ID: 2}}
	groups, _, err := client.JobTokenScope.GetJobTokenAllowlistGroups(
		1,
		&GetJobTokenAllowlistGroupsOptions{},
	)

	assert.NoError(t, err)
	assert.Equal(t, want, groups)
}

func TestAddGroupToJobTokenAllowlist(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/job_token_scope/groups_allowlist", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testBodyJSON(t, r, map[string]any{
			"target_group_id": 2.0,
		})

		// Ensure we provide the proper response
		w.WriteHeader(http.StatusCreated)

		// Print on the response with the proper target group
		fmt.Fprintf(w, `{
			"source_project_id": 1,
			"target_group_id": 2
		}`)
	})

	want := &JobTokenAllowlistItem{
		SourceProjectID: 1,
		TargetGroupID:   2,
	}

	addTokenResponse, resp, err := client.JobTokenScope.AddGroupToJobTokenAllowlist(
		1,
		&AddGroupToJobTokenAllowlistOptions{TargetGroupID: Ptr(int64(2))},
	)
	assert.NoError(t, err)
	assert.Equal(t, want, addTokenResponse)
	assert.Equal(t, 201, resp.StatusCode)
}

func TestRemoveGroupFromJobTokenAllowlist(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/job_token_scope/groups_allowlist/2", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)

		// Ensure we provide the proper response
		w.WriteHeader(http.StatusNoContent)

		// Print an empty body, since that's what the API provides.
		fmt.Fprint(w, "")
	})

	resp, err := client.JobTokenScope.RemoveGroupFromJobTokenAllowlist(1, 2)
	assert.NoError(t, err)
	assert.Equal(t, 204, resp.StatusCode)
}
