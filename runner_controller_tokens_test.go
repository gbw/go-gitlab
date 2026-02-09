package gitlab

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListRunnerControllerTokens(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/runner_controllers/1/tokens", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
			{
				"id": 1,
				"description": "Token 1",
				"created_at": "2020-02-14T00:00:00.000Z",
				"updated_at": "2020-02-15T00:00:00.000Z"
			},
			{
				"id": 2,
				"description": "Token 2",
				"created_at": "2020-03-15T00:00:00.000Z",
				"updated_at": "2020-03-16T00:00:00.000Z"
			}
		]`)
	})

	tokens, _, err := client.RunnerControllerTokens.ListRunnerControllerTokens(1, &ListRunnerControllerTokensOptions{})
	assert.NoError(t, err)

	want := []*RunnerControllerToken{
		{
			ID:          1,
			Description: "Token 1",
			CreatedAt:   Ptr(time.Date(2020, time.February, 14, 0, 0, 0, 0, time.UTC)),
			UpdatedAt:   Ptr(time.Date(2020, time.February, 15, 0, 0, 0, 0, time.UTC)),
		},
		{
			ID:          2,
			Description: "Token 2",
			CreatedAt:   Ptr(time.Date(2020, time.March, 15, 0, 0, 0, 0, time.UTC)),
			UpdatedAt:   Ptr(time.Date(2020, time.March, 16, 0, 0, 0, 0, time.UTC)),
		},
	}
	assert.Equal(t, want, tokens)
}

func TestGetRunnerControllerToken(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/runner_controllers/1/tokens/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"id": 1,
			"description": "Test Token",
			"created_at": "2020-02-14T00:00:00.000Z",
			"updated_at": "2020-02-15T00:00:00.000Z"
		}`)
	})

	token, _, err := client.RunnerControllerTokens.GetRunnerControllerToken(1, 1)
	assert.NoError(t, err)

	want := &RunnerControllerToken{
		ID:          1,
		Description: "Test Token",
		CreatedAt:   Ptr(time.Date(2020, time.February, 14, 0, 0, 0, 0, time.UTC)),
		UpdatedAt:   Ptr(time.Date(2020, time.February, 15, 0, 0, 0, 0, time.UTC)),
	}
	assert.Equal(t, want, token)
}

func TestCreateRunnerControllerToken(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/runner_controllers/1/tokens", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testBodyJSON(t, r, map[string]string{
			"description": "New Token",
		})
		fmt.Fprint(w, `{
			"id": 3,
			"description": "New Token",
			"token": "glrct-abc123def456",
			"created_at": "2020-04-16T00:00:00.000Z",
			"updated_at": "2020-04-16T00:00:00.000Z"
		}`)
	})

	opt := &CreateRunnerControllerTokenOptions{
		Description: Ptr("New Token"),
	}
	token, _, err := client.RunnerControllerTokens.CreateRunnerControllerToken(1, opt)
	assert.NoError(t, err)

	want := &RunnerControllerToken{
		ID:          3,
		Description: "New Token",
		Token:       "glrct-abc123def456",
		CreatedAt:   Ptr(time.Date(2020, time.April, 16, 0, 0, 0, 0, time.UTC)),
		UpdatedAt:   Ptr(time.Date(2020, time.April, 16, 0, 0, 0, 0, time.UTC)),
	}
	assert.Equal(t, want, token)
}

func TestRotateRunnerControllerToken(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a runner controller token exists
	mux.HandleFunc("/api/v4/runner_controllers/1/tokens/1/rotate", func(w http.ResponseWriter, r *http.Request) {
		// WHEN the rotate endpoint is called with POST
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
			"id": 1,
			"runner_controller_id": 1,
			"description": "Rotated Token",
			"token": "glrct-rotated123",
			"created_at": "2020-02-14T00:00:00.000Z",
			"updated_at": "2020-05-20T00:00:00.000Z"
		}`)
	})

	token, _, err := client.RunnerControllerTokens.RotateRunnerControllerToken(1, 1)
	assert.NoError(t, err)

	// THEN the rotated token is returned with a new token value
	want := &RunnerControllerToken{
		ID:                 1,
		RunnerControllerID: 1,
		Description:        "Rotated Token",
		Token:              "glrct-rotated123",
		CreatedAt:          Ptr(time.Date(2020, time.February, 14, 0, 0, 0, 0, time.UTC)),
		UpdatedAt:          Ptr(time.Date(2020, time.May, 20, 0, 0, 0, 0, time.UTC)),
	}
	assert.Equal(t, want, token)
}

func TestRevokeRunnerControllerToken(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/runner_controllers/1/tokens/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := client.RunnerControllerTokens.RevokeRunnerControllerToken(1, 1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}
