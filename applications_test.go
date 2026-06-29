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
	"github.com/stretchr/testify/require"
)

func TestCreateApplication(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a server that returns a created application with scopes
	mux.HandleFunc("/api/v4/applications",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			fmt.Fprint(w, `
{
    "id":1,
    "application_name":"testApplication",
    "scopes":["api","read_user"]
}`)
		},
	)

	// WHEN creating an application with scopes
	opt := &CreateApplicationOptions{
		Name:   Ptr("testApplication"),
		Scopes: Ptr("api read_user"),
	}
	app, _, err := client.Applications.CreateApplication(opt)
	require.NoError(t, err)

	// THEN the returned application contains the scopes
	want := &Application{
		ID:              1,
		ApplicationName: "testApplication",
		Scopes:          []string{"api", "read_user"},
	}
	assert.Equal(t, want, app)
}

func TestListApplications(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a server that returns applications with scopes
	mux.HandleFunc("/api/v4/applications",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w, `[
    {"id":1,"scopes":["api"]},
    {"id":2,"scopes":["read_user","email"]}
]`)
		},
	)

	// WHEN listing applications
	apps, _, err := client.Applications.ListApplications(&ListApplicationsOptions{})
	require.NoError(t, err)

	// THEN the returned applications contain their scopes
	want := []*Application{
		{ID: 1, Scopes: []string{"api"}},
		{ID: 2, Scopes: []string{"read_user", "email"}},
	}
	assert.Equal(t, want, apps)
}

func TestDeleteApplication(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/applications/4",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodDelete)
			w.WriteHeader(http.StatusAccepted)
		},
	)

	resp, err := client.Applications.DeleteApplication(4)
	require.NoError(t, err)

	assert.Equal(t, http.StatusAccepted, resp.StatusCode)
}

func TestRenewApplicationSecret(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a server that returns an application with a renewed secret
	mux.HandleFunc("/api/v4/applications/1/renew-secret",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			fmt.Fprint(w, `
{
    "id":1,
    "application_id":"abc123",
    "application_name":"testApplication",
    "secret":"newSecret",
    "callback_url":"http://example.com",
    "confidential":true,
    "scopes":["api"]
}`)
		},
	)

	// WHEN renewing the secret for an application
	app, _, err := client.Applications.RenewApplicationSecret(1)
	require.NoError(t, err)

	// THEN the returned application contains the new secret
	want := &Application{
		ID:              1,
		ApplicationID:   "abc123",
		ApplicationName: "testApplication",
		Secret:          "newSecret",
		CallbackURL:     "http://example.com",
		Confidential:    true,
		Scopes:          []string{"api"},
	}
	assert.Equal(t, want, app)
}
