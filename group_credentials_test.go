package gitlab

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListGroupPersonalAccessTokens(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/manage/personal_access_tokens", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprintf(w, `[
			{
				"id": 1,
				"name": "test-token",
				"revoked": false,
				"created_at": "2023-01-01T00:00:00Z",
				"description": "Test token for group",
				"scopes": ["api", "read_user"],
				"user_id": 10,
				"last_used_at": "2023-01-15T10:30:00Z",
				"active": true,
				"expires_at": "2024-01-01"
			},
			{
				"id": 2,
				"name": "another-token",
				"revoked": true,
				"created_at": "2023-02-01T00:00:00Z",
				"description": "Another test token",
				"scopes": ["read_api"],
				"user_id": 20,
				"last_used_at": null,
				"active": false,
				"expires_at": "2024-02-01"
			}
		]`)
	})

	createdAt1 := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)
	createdAt2 := time.Date(2023, time.February, 1, 0, 0, 0, 0, time.UTC)
	lastUsedAt := time.Date(2023, time.January, 15, 10, 30, 0, 0, time.UTC)
	expiresAt1 := ISOTime(time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC))
	expiresAt2 := ISOTime(time.Date(2024, time.February, 1, 0, 0, 0, 0, time.UTC))

	expected := []*GroupPersonalAccessToken{
		{
			ID:          1,
			Name:        "test-token",
			Revoked:     false,
			CreatedAt:   &createdAt1,
			Description: "Test token for group",
			Scopes:      []string{"api", "read_user"},
			UserID:      10,
			LastUsedAt:  &lastUsedAt,
			Active:      true,
			ExpiresAt:   &expiresAt1,
		},
		{
			ID:          2,
			Name:        "another-token",
			Revoked:     true,
			CreatedAt:   &createdAt2,
			Description: "Another test token",
			Scopes:      []string{"read_api"},
			UserID:      20,
			LastUsedAt:  nil,
			Active:      false,
			ExpiresAt:   &expiresAt2,
		},
	}

	tokens, resp, err := client.GroupCredentials.ListGroupPersonalAccessTokens(1, nil)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, expected, tokens)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
