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

	mux.HandleFunc("/api/v4/applications",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			fmt.Fprint(w, `
{
    "id":1,
    "application_name":"testApplication"
}`)
		},
	)

	opt := &CreateApplicationOptions{
		Name: Ptr("testApplication"),
	}
	app, _, err := client.Applications.CreateApplication(opt)
	require.NoError(t, err)

	want := &Application{
		ID:              1,
		ApplicationName: "testApplication",
	}
	assert.Equal(t, want, app)
}

func TestListApplications(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/applications",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			fmt.Fprint(w, `[
    {"id":1},
    {"id":2}
]`)
		},
	)

	apps, _, err := client.Applications.ListApplications(&ListApplicationsOptions{})
	require.NoError(t, err)

	want := []*Application{
		{ID: 1},
		{ID: 2},
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
