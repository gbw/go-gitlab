//
// Copyright 2026, Jimmy Spagnola
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
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListProjectServiceAccounts(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/service_accounts", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
			{
				"id": 57,
				"username": "service_account_project_345_6018816a18e515214e0c34c2b33523fc",
				"name": "Service account user",
				"email": "service_account_project_345_abc@noreply.gitlab.example.com"
			}
		]`)
	})

	serviceAccounts, resp, err := client.Projects.ListProjectServiceAccounts(1, nil)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	want := []*ProjectServiceAccount{
		{
			ID:       57,
			Username: "service_account_project_345_6018816a18e515214e0c34c2b33523fc",
			Name:     "Service account user",
			Email:    "service_account_project_345_abc@noreply.gitlab.example.com",
		},
	}
	assert.Equal(t, want, serviceAccounts)
}

func TestCreateProjectServiceAccount(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/service_accounts", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
			"id": 57,
			"username": "service_account_project_345_6018816a18e515214e0c34c2b33523fc",
			"name": "Service account user",
			"email": "custom_email@example.com"
		}`)
	})

	sa, resp, err := client.Projects.CreateProjectServiceAccount(1, &CreateProjectServiceAccountOptions{
		Name:     Ptr("Service account user"),
		Username: Ptr("service_account_project_345_6018816a18e515214e0c34c2b33523fc"),
		Email:    Ptr("custom_email@example.com"),
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	want := &ProjectServiceAccount{
		ID:       57,
		Username: "service_account_project_345_6018816a18e515214e0c34c2b33523fc",
		Name:     "Service account user",
		Email:    "custom_email@example.com",
	}
	assert.Equal(t, want, sa)
}

func TestUpdateProjectServiceAccount(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/service_accounts/57", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		fmt.Fprint(w, `{
			"id": 57,
			"username": "service_account_project_345_6018816a18e515214e0c34c2b33523fc",
			"name": "Updated Service Account",
			"email": "service_account_project_345_abc@noreply.gitlab.example.com",
			"unconfirmed_email": "custom_email@example.com"
		}`)
	})

	sa, resp, err := client.Projects.UpdateProjectServiceAccount(1, 57, &UpdateProjectServiceAccountOptions{
		Name:  Ptr("Updated Service Account"),
		Email: Ptr("custom_email@example.com"),
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	want := &ProjectServiceAccount{
		ID:               57,
		Username:         "service_account_project_345_6018816a18e515214e0c34c2b33523fc",
		Name:             "Updated Service Account",
		Email:            "service_account_project_345_abc@noreply.gitlab.example.com",
		UnconfirmedEmail: "custom_email@example.com",
	}
	assert.Equal(t, want, sa)
}

func TestDeleteProjectServiceAccount(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/service_accounts/57", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.Projects.DeleteProjectServiceAccount(1, 57, &DeleteProjectServiceAccountOptions{HardDelete: Ptr(true)})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestListProjectServiceAccountPersonalAccessTokens(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/service_accounts/57/personal_access_tokens", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{
			"id": 6,
			"name": "service_account_token",
			"revoked": false,
			"created_at": "2023-06-13T07:47:13.000Z",
			"description": "A token",
			"scopes": ["api"],
			"user_id": 71,
			"last_used_at": null,
			"active": true,
			"expires_at": "2024-06-12",
			"token": "random_token"
		}]`)
	})

	pats, resp, err := client.Projects.ListProjectServiceAccountPersonalAccessTokens(1, 57, &ListProjectServiceAccountPersonalAccessTokensOptions{
		State: Ptr("active"),
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	createdAt := time.Date(2023, time.June, 13, 7, 47, 13, 0, time.UTC)
	expiresAt := ISOTime(time.Date(2024, time.June, 12, 0, 0, 0, 0, time.UTC))

	want := []*PersonalAccessToken{{
		ID:          6,
		Name:        "service_account_token",
		Revoked:     false,
		CreatedAt:   &createdAt,
		Description: "A token",
		Scopes:      []string{"api"},
		UserID:      71,
		Active:      true,
		ExpiresAt:   &expiresAt,
		Token:       "random_token",
	}}
	assert.Equal(t, want, pats)
}

func TestCreateProjectServiceAccountPersonalAccessToken(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/service_accounts/57/personal_access_tokens", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
			"id": 6,
			"name": "service_account_token",
			"revoked": false,
			"created_at": "2023-06-13T07:47:13.000Z",
			"description": "A token",
			"scopes": ["api"],
			"user_id": 71,
			"last_used_at": null,
			"active": true,
			"expires_at": "2024-06-12",
			"token": "random_token"
		}`)
	})

	expireTime, err := ParseISOTime("2024-06-12")
	require.NoError(t, err)

	pat, resp, err := client.Projects.CreateProjectServiceAccountPersonalAccessToken(1, 57, &CreateProjectServiceAccountPersonalAccessTokenOptions{
		Scopes:    Ptr([]string{"api"}),
		Name:      Ptr("service_account_token"),
		ExpiresAt: Ptr(expireTime),
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	createdAt := time.Date(2023, time.June, 13, 7, 47, 13, 0, time.UTC)
	expiresAt := ISOTime(time.Date(2024, time.June, 12, 0, 0, 0, 0, time.UTC))

	want := &PersonalAccessToken{
		ID:          6,
		Name:        "service_account_token",
		Revoked:     false,
		CreatedAt:   &createdAt,
		Description: "A token",
		Scopes:      []string{"api"},
		UserID:      71,
		Active:      true,
		ExpiresAt:   &expiresAt,
		Token:       "random_token",
	}
	assert.Equal(t, want, pat)
}

func TestRevokeProjectServiceAccountPersonalAccessToken(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/service_accounts/57/personal_access_tokens/6", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.Projects.RevokeProjectServiceAccountPersonalAccessToken(1, 57, 6)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestRotateProjectServiceAccountPersonalAccessToken(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/service_accounts/57/personal_access_tokens/6/rotate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
			"id": 7,
			"name": "service_account_token",
			"revoked": false,
			"created_at": "2023-06-13T07:54:49.000Z",
			"description": "A token",
			"scopes": ["api"],
			"user_id": 71,
			"last_used_at": null,
			"active": true,
			"expires_at": "2025-06-20",
			"token": "random_token_2"
		}`)
	})

	createdAt := time.Date(2023, time.June, 13, 7, 54, 49, 0, time.UTC)
	expiresAt := ISOTime(time.Date(2025, time.June, 20, 0, 0, 0, 0, time.UTC))

	pat, resp, err := client.Projects.RotateProjectServiceAccountPersonalAccessToken(1, 57, 6, &RotateProjectServiceAccountPersonalAccessTokenOptions{
		ExpiresAt: &expiresAt,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	want := &PersonalAccessToken{
		ID:          7,
		Name:        "service_account_token",
		Revoked:     false,
		CreatedAt:   &createdAt,
		Description: "A token",
		Scopes:      []string{"api"},
		UserID:      71,
		Active:      true,
		ExpiresAt:   &expiresAt,
		Token:       "random_token_2",
	}
	assert.Equal(t, want, pat)
}
