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
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListPersonalAccessTokensWithUserFilter(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/personal_access_tokens", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/list_personal_access_tokens_with_user_filter.json")
	})

	personalAccessTokens, _, err := client.PersonalAccessTokens.ListPersonalAccessTokens(
		&ListPersonalAccessTokensOptions{
			UserID:        Ptr(int64(1)),
			ExpiresBefore: Ptr(ISOTime(time.Date(2022, time.March, 22, 0, 0, 0, 0, time.UTC))),
			ListOptions:   ListOptions{Page: 1, PerPage: 10},
		},
	)
	require.NoError(t, err)

	expiresAt1 := ISOTime(time.Date(2022, time.March, 21, 0, 0, 0, 0, time.UTC))
	expiresAt2 := ISOTime(time.Date(2022, time.March, 20, 0, 0, 0, 0, time.UTC))

	want := []*PersonalAccessToken{
		{
			ID:          1,
			Name:        "test 1",
			Description: "Describe test 1",
			Revoked:     true,
			CreatedAt:   mustParseTime("2020-02-20T14:58:56.238Z"),
			Scopes:      []string{"api"},
			UserID:      1,
			LastUsedAt:  mustParseTime("2021-04-20T16:31:39.105Z"),
			Active:      false,
			ExpiresAt:   &expiresAt1,
		},
		{
			ID:          2,
			Name:        "test 2",
			Description: "Describe test 2",
			Revoked:     false,
			CreatedAt:   mustParseTime("2022-03-20T03:56:18.968Z"),
			Scopes:      []string{"api", "read_user"},
			UserID:      1,
			LastUsedAt:  nil,
			Active:      false,
			ExpiresAt:   &expiresAt2,
		},
	}

	assert.Equal(t, want, personalAccessTokens)
}

func TestListPersonalAccessTokensNoUserFilter(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/personal_access_tokens", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/list_personal_access_tokens_without_user_filter.json")
	})

	personalAccessTokens, _, err := client.PersonalAccessTokens.ListPersonalAccessTokens(
		&ListPersonalAccessTokensOptions{ListOptions: ListOptions{Page: 1, PerPage: 10}},
	)
	require.NoError(t, err)

	expiresAt1 := ISOTime(time.Date(2022, time.March, 21, 0, 0, 0, 0, time.UTC))
	expiresAt2 := ISOTime(time.Date(2022, time.March, 20, 0, 0, 0, 0, time.UTC))

	want := []*PersonalAccessToken{
		{
			ID:          1,
			Name:        "test 1",
			Description: "Describe test 1",
			Revoked:     true,
			CreatedAt:   mustParseTime("2020-02-20T14:58:56.238Z"),
			Scopes:      []string{"api"},
			UserID:      1,
			LastUsedAt:  mustParseTime("2021-04-20T16:31:39.105Z"),
			Active:      false,
			ExpiresAt:   &expiresAt1,
		},
		{
			ID:          2,
			Name:        "test 2",
			Description: "Describe test 2",
			Revoked:     false,
			CreatedAt:   mustParseTime("2022-03-20T03:56:18.968Z"),
			Scopes:      []string{"api", "read_user"},
			UserID:      2,
			LastUsedAt:  nil,
			Active:      false,
			ExpiresAt:   &expiresAt2,
		},
	}

	assert.Equal(t, want, personalAccessTokens)
}

func TestGetSinglePersonalAccessTokenByID(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/personal_access_tokens/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/list_personal_access_tokens_single.json")
	})

	token, _, err := client.PersonalAccessTokens.GetSinglePersonalAccessTokenByID(1)
	require.NoError(t, err)

	want := &PersonalAccessToken{
		ID:          1,
		Name:        "Test Token",
		Description: "Describe test token",
		Revoked:     false,
		CreatedAt:   mustParseTime("2020-07-23T14:31:47.729Z"),
		Scopes:      []string{"api"},
		UserID:      1,
		LastUsedAt:  mustParseTime("2021-10-06T17:58:37.550Z"),
		Active:      true,
	}

	assert.Equal(t, want, token)
}

func TestGetSinglePersonalAccessToken(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/personal_access_tokens/self", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		mustWriteHTTPResponse(t, w, "testdata/list_personal_access_tokens_single.json")
	})

	token, _, err := client.PersonalAccessTokens.GetSinglePersonalAccessToken()
	require.NoError(t, err)

	want := &PersonalAccessToken{
		ID:          1,
		Name:        "Test Token",
		Description: "Describe test token",
		Revoked:     false,
		CreatedAt:   mustParseTime("2020-07-23T14:31:47.729Z"),
		Scopes:      []string{"api"},
		UserID:      1,
		LastUsedAt:  mustParseTime("2021-10-06T17:58:37.550Z"),
		Active:      true,
	}

	assert.Equal(t, want, token)
}

func TestRotatePersonalAccessToken(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)
	mux.HandleFunc("/api/v4/personal_access_tokens/42/rotate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		mustWriteHTTPResponse(t, w, "testdata/rotate_personal_access_token.json")
	})

	expiration := ISOTime(time.Date(2023, time.August, 15, 0, 0, 0, 0, time.UTC))
	opts := &RotatePersonalAccessTokenOptions{ExpiresAt: &expiration}
	rotatedToken, _, err := client.PersonalAccessTokens.RotatePersonalAccessToken(42, opts)
	require.NoError(t, err)

	want := &PersonalAccessToken{
		ID:          42,
		UserID:      1337,
		Name:        "Rotated Token",
		Description: "Describe rotated token",
		Scopes:      []string{"api"},
		ExpiresAt:   &expiration,
		CreatedAt:   mustParseTime("2023-08-01T15:00:00.000Z"),
		Active:      true,
		Revoked:     false,
		Token:       "s3cr3t",
	}

	assert.Equal(t, want, rotatedToken)
}

func TestRotatePersonalAccessTokenByID(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)
	mux.HandleFunc("/api/v4/personal_access_tokens/42/rotate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		mustWriteHTTPResponse(t, w, "testdata/rotate_personal_access_token.json")
	})

	expiration := ISOTime(time.Date(2023, time.August, 15, 0, 0, 0, 0, time.UTC))
	opts := &RotatePersonalAccessTokenOptions{ExpiresAt: &expiration}
	rotatedToken, _, err := client.PersonalAccessTokens.RotatePersonalAccessTokenByID(42, opts)
	require.NoError(t, err)

	want := &PersonalAccessToken{
		ID:          42,
		UserID:      1337,
		Name:        "Rotated Token",
		Description: "Describe rotated token",
		Scopes:      []string{"api"},
		ExpiresAt:   &expiration,
		CreatedAt:   mustParseTime("2023-08-01T15:00:00.000Z"),
		Active:      true,
		Revoked:     false,
		Token:       "s3cr3t",
	}

	assert.Equal(t, want, rotatedToken)
}

func TestRotatePersonalAccessTokenSelf(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)
	mux.HandleFunc("/api/v4/personal_access_tokens/self/rotate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		mustWriteHTTPResponse(t, w, "testdata/rotate_personal_access_token.json")
	})

	expiration := ISOTime(time.Date(2023, time.August, 15, 0, 0, 0, 0, time.UTC))
	opts := &RotatePersonalAccessTokenOptions{ExpiresAt: &expiration}
	rotatedToken, _, err := client.PersonalAccessTokens.RotatePersonalAccessTokenSelf(opts)
	require.NoError(t, err)

	want := &PersonalAccessToken{
		ID:          42,
		UserID:      1337,
		Name:        "Rotated Token",
		Description: "Describe rotated token",
		Scopes:      []string{"api"},
		ExpiresAt:   &expiration,
		CreatedAt:   mustParseTime("2023-08-01T15:00:00.000Z"),
		Active:      true,
		Revoked:     false,
		Token:       "s3cr3t",
	}

	assert.Equal(t, want, rotatedToken)
}

func TestRevokePersonalAccessToken(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/personal_access_tokens/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.PersonalAccessTokens.RevokePersonalAccessToken(1)
	require.NoError(t, err)
}

func TestRevokePersonalAccessTokenByID(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/personal_access_tokens/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.PersonalAccessTokens.RevokePersonalAccessTokenByID(1)
	require.NoError(t, err)
}

func TestRevokePersonalAccessTokenSelf(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/personal_access_tokens/self", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.PersonalAccessTokens.RevokePersonalAccessTokenSelf()
	require.NoError(t, err)
}
