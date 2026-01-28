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

func TestGetEpic(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/7/epics/8", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":8, "title": "Incredible idea", "description": "This is a test epic", "author" : {"id" : 26, "name": "jramsay"}}`)
	})

	epic, _, err := client.Epics.GetEpic("7", 8)
	require.NoError(t, err)

	want := &Epic{
		ID:          8,
		Title:       "Incredible idea",
		Description: "This is a test epic",
		Author:      &EpicAuthor{ID: 26, Name: "jramsay"},
	}

	assert.Equal(t, want, epic)
}

func TestDeleteEpic(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/7/epics/8", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.Epics.DeleteEpic("7", 8)
	require.NoError(t, err)
}

func TestListGroupEpics(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/7/epics", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testURL(t, r, "/api/v4/groups/7/epics?author_id=26&state=opened")
		fmt.Fprint(w, `[{"id":8, "title": "Incredible idea", "description": "This is a test epic", "author" : {"id" : 26, "name": "jramsay"}}]`)
	})

	listGroupEpics := &ListGroupEpicsOptions{
		AuthorID: Ptr(int64(26)),
		State:    Ptr("opened"),
	}

	epics, _, err := client.Epics.ListGroupEpics("7", listGroupEpics)
	require.NoError(t, err)

	want := []*Epic{{
		ID:          8,
		Title:       "Incredible idea",
		Description: "This is a test epic",
		Author:      &EpicAuthor{ID: 26, Name: "jramsay"},
	}}

	assert.Equal(t, want, epics)
}

func TestCreateEpic(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/7/epics", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"id":8, "title": "Incredible idea", "description": "This is a test epic", "author" : {"id" : 26, "name": "jramsay"}}`)
	})

	createEpicOptions := &CreateEpicOptions{
		Title:       Ptr("Incredible idea"),
		Description: Ptr("This is a test epic"),
	}

	epic, _, err := client.Epics.CreateEpic("7", createEpicOptions)
	require.NoError(t, err)

	want := &Epic{
		ID:          8,
		Title:       "Incredible idea",
		Description: "This is a test epic",
		Author:      &EpicAuthor{ID: 26, Name: "jramsay"},
	}

	assert.Equal(t, want, epic)
}

func TestUpdateEpic(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/7/epics/8", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprint(w, `{"id":8, "title": "Incredible idea", "description": "This is a test epic", "author" : {"id" : 26, "name": "jramsay"}}`)
	})

	updateEpicOptions := &UpdateEpicOptions{
		Title:       Ptr("Incredible idea"),
		Description: Ptr("This is a test epic"),
	}

	epic, _, err := client.Epics.UpdateEpic("7", 8, updateEpicOptions)
	require.NoError(t, err)

	want := &Epic{
		ID:          8,
		Title:       "Incredible idea",
		Description: "This is a test epic",
		Author:      &EpicAuthor{ID: 26, Name: "jramsay"},
	}

	assert.Equal(t, want, epic)
}

func TestGetEpicLinks(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/7/epics/8/epics", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"id":9, "title": "Child epic", "description": "This is a child epic", "author" : {"id" : 27, "name": "asmith"}}]`)
	})

	epics, _, err := client.Epics.GetEpicLinks("7", 8)
	require.NoError(t, err)

	want := []*Epic{{
		ID:          9,
		Title:       "Child epic",
		Description: "This is a child epic",
		Author:      &EpicAuthor{ID: 27, Name: "asmith"},
	}}

	assert.Equal(t, want, epics)
}

func TestGetEpicWithDateTimeFields(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/7/epics/10", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"id": 10,
			"title": "Epic with datetime fields",
			"description": "Testing ISO 8601 datetime parsing",
			"author": {"id": 26, "name": "jramsay"},
			"start_date": "2026-06-13T00:00:00+00:00",
			"due_date": "2026-12-31T23:59:59+00:00",
			"start_date_fixed": "2026-06-13T00:00:00+00:00",
			"due_date_fixed": "2026-12-31T23:59:59+00:00",
			"start_date_from_milestones": "2026-06-01T00:00:00+00:00",
			"due_date_from_milestones": "2026-12-31T00:00:00+00:00"
		}`)
	})

	epic, _, err := client.Epics.GetEpic("7", 10)
	require.NoError(t, err)

	// Verify all datetime fields were parsed successfully
	assert.NotNil(t, epic.StartDate, "Expected StartDate to be non-nil")
	assert.NotNil(t, epic.DueDate, "Expected DueDate to be non-nil")
	assert.NotNil(t, epic.StartDateFixed, "Expected StartDateFixed to be non-nil")
	assert.NotNil(t, epic.DueDateFixed, "Expected DueDateFixed to be non-nil")
	assert.NotNil(t, epic.StartDateFromMilestones, "Expected StartDateFromMilestones to be non-nil")
	assert.NotNil(t, epic.DueDateFromMilestones, "Expected DueDateFromMilestones to be non-nil")

	want := &Epic{
		ID:                      10,
		Title:                   "Epic with datetime fields",
		Description:             "Testing ISO 8601 datetime parsing",
		Author:                  &EpicAuthor{ID: 26, Name: "jramsay"},
		StartDate:               epic.StartDate,
		DueDate:                 epic.DueDate,
		StartDateFixed:          epic.StartDateFixed,
		DueDateFixed:            epic.DueDateFixed,
		StartDateFromMilestones: epic.StartDateFromMilestones,
		DueDateFromMilestones:   epic.DueDateFromMilestones,
	}

	assert.Equal(t, want, epic)
}
