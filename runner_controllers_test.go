package gitlab

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListRunnerControllers(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/runner_controllers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
			{
				"id": 1,
				"description": "Controller 1",
				"state": "enabled",
				"created_at": "2020-02-14T00:00:00.000Z",
				"updated_at": "2020-02-15T00:00:00.000Z"
			},
			{
				"id": 2,
				"description": "Controller 2",
				"state": "disabled",
				"created_at": "2020-03-14T00:00:00.000Z",
				"updated_at": "2020-03-15T00:00:00.000Z"
			}
		]`)
	})

	controllers, _, err := client.RunnerControllers.ListRunnerControllers(&ListRunnerControllersOptions{})
	assert.NoError(t, err)

	want := []*RunnerController{
		{
			ID:          1,
			Description: "Controller 1",
			State:       RunnerControllerStateEnabled,
			CreatedAt:   Ptr(time.Date(2020, time.February, 14, 0, 0, 0, 0, time.UTC)),
			UpdatedAt:   Ptr(time.Date(2020, time.February, 15, 0, 0, 0, 0, time.UTC)),
		},
		{
			ID:          2,
			Description: "Controller 2",
			State:       RunnerControllerStateDisabled,
			CreatedAt:   Ptr(time.Date(2020, time.March, 14, 0, 0, 0, 0, time.UTC)),
			UpdatedAt:   Ptr(time.Date(2020, time.March, 15, 0, 0, 0, 0, time.UTC)),
		},
	}
	assert.Equal(t, want, controllers)
}

func TestGetRunnerController(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/runner_controllers/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"id": 1,
			"description": "Test Controller",
			"state": "enabled",
			"created_at": "2020-02-14T00:00:00.000Z",
			"updated_at": "2020-02-15T00:00:00.000Z"
		}`)
	})

	controller, _, err := client.RunnerControllers.GetRunnerController(1)
	assert.NoError(t, err)

	want := &RunnerController{
		ID:          1,
		Description: "Test Controller",
		State:       RunnerControllerStateEnabled,
		CreatedAt:   Ptr(time.Date(2020, time.February, 14, 0, 0, 0, 0, time.UTC)),
		UpdatedAt:   Ptr(time.Date(2020, time.February, 15, 0, 0, 0, 0, time.UTC)),
	}
	assert.Equal(t, want, controller)
}

func TestCreateRunnerController(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/runner_controllers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testBodyJSON(t, r, map[string]any{
			"description": "New Controller",
			"state":       "dry_run",
		})
		fmt.Fprint(w, `{
			"id": 3,
			"description": "New Controller",
			"state": "dry_run",
			"created_at": "2020-04-14T00:00:00.000Z",
			"updated_at": "2020-04-14T00:00:00.000Z"
		}`)
	})

	opt := &CreateRunnerControllerOptions{
		Description: Ptr("New Controller"),
		State:       Ptr(RunnerControllerStateDryRun),
	}
	controller, _, err := client.RunnerControllers.CreateRunnerController(opt)
	assert.NoError(t, err)

	want := &RunnerController{
		ID:          3,
		Description: "New Controller",
		State:       RunnerControllerStateDryRun,
		CreatedAt:   Ptr(time.Date(2020, time.April, 14, 0, 0, 0, 0, time.UTC)),
		UpdatedAt:   Ptr(time.Date(2020, time.April, 14, 0, 0, 0, 0, time.UTC)),
	}
	assert.Equal(t, want, controller)
}

func TestUpdateRunnerController(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/runner_controllers/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		testBodyJSON(t, r, map[string]any{
			"description": "Updated Controller",
			"state":       "disabled",
		})
		fmt.Fprint(w, `{
			"id": 1,
			"description": "Updated Controller",
			"state": "disabled",
			"created_at": "2020-02-14T00:00:00.000Z",
			"updated_at": "2020-05-15T00:00:00.000Z"
		}`)
	})

	opt := &UpdateRunnerControllerOptions{
		Description: Ptr("Updated Controller"),
		State:       Ptr(RunnerControllerStateDisabled),
	}
	controller, _, err := client.RunnerControllers.UpdateRunnerController(1, opt)
	assert.NoError(t, err)

	want := &RunnerController{
		ID:          1,
		Description: "Updated Controller",
		State:       RunnerControllerStateDisabled,
		CreatedAt:   Ptr(time.Date(2020, time.February, 14, 0, 0, 0, 0, time.UTC)),
		UpdatedAt:   Ptr(time.Date(2020, time.May, 15, 0, 0, 0, 0, time.UTC)),
	}
	assert.Equal(t, want, controller)
}

func TestDeleteRunnerController(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/runner_controllers/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
	})

	resp, err := client.RunnerControllers.DeleteRunnerController(1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}
