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
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListPipelineTriggers(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/triggers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
			{
				"id": 10,
				"description": "my trigger",
				"created_at": "2016-01-07T09:53:58.235Z",
				"last_used": null,
				"token": "6d056f63e50fe6f8c5f8f4aa10edb7",
				"updated_at": "2016-01-07T09:53:58.235Z",
				"owner": null
			}
		]`)
	})

	pipelines, resp, err := client.PipelineTriggers.ListPipelineTriggers(1, nil)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	want := []*PipelineTrigger{
		{
			ID:          10,
			Description: "my trigger",
			CreatedAt:   Ptr(time.Date(2016, 1, 7, 9, 53, 58, 235000000, time.UTC)),
			Token:       "6d056f63e50fe6f8c5f8f4aa10edb7",
			UpdatedAt:   Ptr(time.Date(2016, 1, 7, 9, 53, 58, 235000000, time.UTC)),
		},
	}
	assert.Equal(t, want, pipelines)
}

func TestGetPipelineTrigger(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/triggers/10", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"id": 10,
			"description": "my trigger",
			"created_at": "2016-01-07T09:53:58.235Z",
			"last_used": null,
			"token": "6d056f63e50fe6f8c5f8f4aa10edb7",
			"updated_at": "2016-01-07T09:53:58.235Z",
			"owner": null
		}`)
	})

	pipeline, resp, err := client.PipelineTriggers.GetPipelineTrigger(1, 10)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	want := &PipelineTrigger{
		ID:          10,
		Description: "my trigger",
		CreatedAt:   Ptr(time.Date(2016, 1, 7, 9, 53, 58, 235000000, time.UTC)),
		Token:       "6d056f63e50fe6f8c5f8f4aa10edb7",
		UpdatedAt:   Ptr(time.Date(2016, 1, 7, 9, 53, 58, 235000000, time.UTC)),
	}
	assert.Equal(t, want, pipeline)
}

func TestAddPipelineTrigger(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/triggers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
			"id": 10,
			"description": "my trigger",
			"created_at": "2016-01-07T09:53:58.235Z",
			"last_used": null,
			"token": "6d056f63e50fe6f8c5f8f4aa10edb7",
			"updated_at": "2016-01-07T09:53:58.235Z",
			"owner": null
		}`)
	})

	opt := &AddPipelineTriggerOptions{Description: Ptr("my trigger")}
	pipeline, resp, err := client.PipelineTriggers.AddPipelineTrigger(1, opt)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	want := &PipelineTrigger{
		ID:          10,
		Description: "my trigger",
		CreatedAt:   Ptr(time.Date(2016, 1, 7, 9, 53, 58, 235000000, time.UTC)),
		Token:       "6d056f63e50fe6f8c5f8f4aa10edb7",
		UpdatedAt:   Ptr(time.Date(2016, 1, 7, 9, 53, 58, 235000000, time.UTC)),
	}
	assert.Equal(t, want, pipeline)
}

func TestEditPipelineTrigger(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/triggers/10", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprint(w, `{
			"id": 10,
			"description": "my trigger",
			"created_at": "2016-01-07T09:53:58.235Z",
			"last_used": null,
			"token": "6d056f63e50fe6f8c5f8f4aa10edb7",
			"updated_at": "2016-01-07T09:53:58.235Z",
			"owner": null
		}`)
	})

	opt := &EditPipelineTriggerOptions{Description: Ptr("my trigger")}
	pipeline, resp, err := client.PipelineTriggers.EditPipelineTrigger(1, 10, opt)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	want := &PipelineTrigger{
		ID:          10,
		Description: "my trigger",
		CreatedAt:   Ptr(time.Date(2016, 1, 7, 9, 53, 58, 235000000, time.UTC)),
		Token:       "6d056f63e50fe6f8c5f8f4aa10edb7",
		UpdatedAt:   Ptr(time.Date(2016, 1, 7, 9, 53, 58, 235000000, time.UTC)),
	}
	assert.Equal(t, want, pipeline)
}

func TestDeletePipelineTrigger(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/triggers/10", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.PipelineTriggers.DeletePipelineTrigger(1, 10)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestRunPipelineTrigger(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/trigger/pipeline", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"id":1, "status":"pending"}`)
	})

	opt := &RunPipelineTriggerOptions{Ref: Ptr("master")}
	pipeline, resp, err := client.PipelineTriggers.RunPipelineTrigger(1, opt)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	want := &Pipeline{ID: 1, Status: "pending"}
	assert.Equal(t, want, pipeline)
}
