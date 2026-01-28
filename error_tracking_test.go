//
// Copyright 2022, Ryan Glab <ryan.j.glab@gmail.com>
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

func TestGetErrorTracking(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/error_tracking/settings", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"active": true,
			"project_name": "sample sentry project",
			"sentry_external_url": "https://sentry.io/myawesomeproject/project",
			"api_url": "https://sentry.io/api/1/projects/myawesomeproject/project",
			"integrated": false
		}`)
	})

	et, _, err := client.ErrorTracking.GetErrorTrackingSettings(1)
	require.NoError(t, err)

	want := &ErrorTrackingSettings{
		Active:            true,
		ProjectName:       "sample sentry project",
		SentryExternalURL: "https://sentry.io/myawesomeproject/project",
		APIURL:            "https://sentry.io/api/1/projects/myawesomeproject/project",
		Integrated:        false,
	}

	assert.Equal(t, want, et)
}

func TestDisableErrorTracking(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/error_tracking/settings", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		fmt.Fprint(w, `{
			"active": false,
			"project_name": "sample sentry project",
			"sentry_external_url": "https://sentry.io/myawesomeproject/project",
			"api_url": "https://sentry.io/api/1/projects/myawesomeproject/project",
			"integrated": false
		}`)
	})

	et, _, err := client.ErrorTracking.EnableDisableErrorTracking(
		1,
		&EnableDisableErrorTrackingOptions{
			Active:     Ptr(false),
			Integrated: Ptr(false),
		},
	)
	require.NoError(t, err)

	want := &ErrorTrackingSettings{
		Active:            false,
		ProjectName:       "sample sentry project",
		SentryExternalURL: "https://sentry.io/myawesomeproject/project",
		APIURL:            "https://sentry.io/api/1/projects/myawesomeproject/project",
		Integrated:        false,
	}

	assert.Equal(t, want, et)
}

func TestListErrorTrackingClientKeys(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/error_tracking/client_keys", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
			{
				"id": 1,
				"active": true,
				"public_key": "glet_aa77551d849c083f76d0bc545ed053a3",
				"sentry_dsn": "https://glet_aa77551d849c083f76d0bc545ed053a3@gitlab.example.com/api/v4/error_tracking/collector/5"
			}
		]`)
	})

	cks, _, err := client.ErrorTracking.ListClientKeys(1, &ListClientKeysOptions{
		ListOptions: ListOptions{
			Page:    1,
			PerPage: 10,
		},
	})
	require.NoError(t, err)

	want := []*ErrorTrackingClientKey{{
		ID:        1,
		Active:    true,
		PublicKey: "glet_aa77551d849c083f76d0bc545ed053a3",
		SentryDsn: "https://glet_aa77551d849c083f76d0bc545ed053a3@gitlab.example.com/api/v4/error_tracking/collector/5",
	}}

	assert.Equal(t, want, cks)
}

func TestCreateClientKey(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/error_tracking/client_keys", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
			"id": 1,
			"active": true,
			"public_key": "glet_aa77551d849c083f76d0bc545ed053a3",
			"sentry_dsn": "https://glet_aa77551d849c083f76d0bc545ed053a3@gitlab.example.com/api/v4/error_tracking/collector/5"
		}`)
	})

	ck, _, err := client.ErrorTracking.CreateClientKey(1)
	require.NoError(t, err)

	want := &ErrorTrackingClientKey{
		ID:        1,
		Active:    true,
		PublicKey: "glet_aa77551d849c083f76d0bc545ed053a3",
		SentryDsn: "https://glet_aa77551d849c083f76d0bc545ed053a3@gitlab.example.com/api/v4/error_tracking/collector/5",
	}

	assert.Equal(t, want, ck)
}

func TestDeleteClientKey(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/error_tracking/client_keys/3", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		testURL(t, r, "/api/v4/projects/1/error_tracking/client_keys/3")
	})

	_, err := client.ErrorTracking.DeleteClientKey(1, 3)
	require.NoError(t, err)
}
