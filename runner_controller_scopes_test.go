package gitlab

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListRunnerControllerScopes(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a runner controller with an instance-level scope exists
	mux.HandleFunc("/api/v4/runner_controllers/1/scopes", func(w http.ResponseWriter, r *http.Request) {
		// WHEN listing scopes for the runner controller
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"instance_level_scopings": [
				{
					"created_at": "2026-01-01T00:00:00.000Z",
					"updated_at": "2026-01-01T00:00:00.000Z"
				}
			],
			"runner_level_scopings": []
		}`)
	})

	scopes, _, err := client.RunnerControllerScopes.ListRunnerControllerScopes(1)
	assert.NoError(t, err)

	// THEN the scopes are returned with instance-level scopings
	want := &RunnerControllerScopes{
		InstanceLevelScopings: []*RunnerControllerInstanceLevelScoping{
			{
				CreatedAt: Ptr(time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC)),
				UpdatedAt: Ptr(time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC)),
			},
		},
		RunnerLevelScopings: []*RunnerControllerRunnerLevelScoping{},
	}
	assert.Equal(t, want, scopes)
}

func TestListRunnerControllerScopes_Empty(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a runner controller with no scopes exists
	mux.HandleFunc("/api/v4/runner_controllers/1/scopes", func(w http.ResponseWriter, r *http.Request) {
		// WHEN listing scopes for the runner controller
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"instance_level_scopings": [],
			"runner_level_scopings": []
		}`)
	})

	scopes, _, err := client.RunnerControllerScopes.ListRunnerControllerScopes(1)
	assert.NoError(t, err)

	// THEN empty scopings arrays are returned
	want := &RunnerControllerScopes{
		InstanceLevelScopings: []*RunnerControllerInstanceLevelScoping{},
		RunnerLevelScopings:   []*RunnerControllerRunnerLevelScoping{},
	}
	assert.Equal(t, want, scopes)
}

func TestAddRunnerControllerInstanceScope(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a runner controller without an instance-level scope exists
	mux.HandleFunc("/api/v4/runner_controllers/1/scopes/instance", func(w http.ResponseWriter, r *http.Request) {
		// WHEN adding an instance-level scope
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{
			"created_at": "2026-01-01T00:00:00.000Z",
			"updated_at": "2026-01-01T00:00:00.000Z"
		}`)
	})

	scoping, resp, err := client.RunnerControllerScopes.AddRunnerControllerInstanceScope(1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	// THEN the created instance-level scoping is returned
	want := &RunnerControllerInstanceLevelScoping{
		CreatedAt: Ptr(time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC)),
		UpdatedAt: Ptr(time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC)),
	}
	assert.Equal(t, want, scoping)
}

func TestRemoveRunnerControllerInstanceScope(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a runner controller with an instance-level scope exists
	mux.HandleFunc("/api/v4/runner_controllers/1/scopes/instance", func(w http.ResponseWriter, r *http.Request) {
		// WHEN removing the instance-level scope
		testMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := client.RunnerControllerScopes.RemoveRunnerControllerInstanceScope(1)
	assert.NoError(t, err)

	// THEN 204 No Content is returned
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestListRunnerControllerScopes_WithRunnerLevelScopings(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a runner controller with runner-level scopings exists
	mux.HandleFunc("/api/v4/runner_controllers/1/scopes", func(w http.ResponseWriter, r *http.Request) {
		// WHEN listing scopes for the runner controller
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"instance_level_scopings": [],
			"runner_level_scopings": [
				{
					"runner_id": 5,
					"created_at": "2026-01-01T00:00:00.000Z",
					"updated_at": "2026-01-01T00:00:00.000Z"
				},
				{
					"runner_id": 10,
					"created_at": "2026-01-02T00:00:00.000Z",
					"updated_at": "2026-01-02T00:00:00.000Z"
				}
			]
		}`)
	})

	scopes, _, err := client.RunnerControllerScopes.ListRunnerControllerScopes(1)
	assert.NoError(t, err)

	// THEN the scopes are returned with runner-level scopings
	want := &RunnerControllerScopes{
		InstanceLevelScopings: []*RunnerControllerInstanceLevelScoping{},
		RunnerLevelScopings: []*RunnerControllerRunnerLevelScoping{
			{
				RunnerID:  5,
				CreatedAt: Ptr(time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC)),
				UpdatedAt: Ptr(time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC)),
			},
			{
				RunnerID:  10,
				CreatedAt: Ptr(time.Date(2026, time.January, 2, 0, 0, 0, 0, time.UTC)),
				UpdatedAt: Ptr(time.Date(2026, time.January, 2, 0, 0, 0, 0, time.UTC)),
			},
		},
	}
	assert.Equal(t, want, scopes)
}

func TestAddRunnerControllerRunnerScope(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a runner controller without a runner scope for runner 5 exists
	mux.HandleFunc("/api/v4/runner_controllers/1/scopes/runners/5", func(w http.ResponseWriter, r *http.Request) {
		// WHEN adding a runner scope for runner 5
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{
			"runner_id": 5,
			"created_at": "2026-01-01T00:00:00.000Z",
			"updated_at": "2026-01-01T00:00:00.000Z"
		}`)
	})

	scoping, resp, err := client.RunnerControllerScopes.AddRunnerControllerRunnerScope(1, 5)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	// THEN the created runner-level scoping is returned
	want := &RunnerControllerRunnerLevelScoping{
		RunnerID:  5,
		CreatedAt: Ptr(time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC)),
		UpdatedAt: Ptr(time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC)),
	}
	assert.Equal(t, want, scoping)
}

func TestRemoveRunnerControllerRunnerScope(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// GIVEN a runner controller with a runner scope for runner 5 exists
	mux.HandleFunc("/api/v4/runner_controllers/1/scopes/runners/5", func(w http.ResponseWriter, r *http.Request) {
		// WHEN removing the runner scope for runner 5
		testMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := client.RunnerControllerScopes.RemoveRunnerControllerRunnerScope(1, 5)
	assert.NoError(t, err)

	// THEN 204 No Content is returned
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}
