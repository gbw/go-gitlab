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
	"reflect"
	"testing"
)

func TestGetEpicNote(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/epics/4329/notes/3", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":3,"type":null,"body":"foo bar","attachment":null,"system":false,"noteable_id":4392,"noteable_type":"Epic","resolvable":false,"noteable_iid":null}`)
	})

	note, _, err := client.Notes.GetEpicNote("1", 4329, 3, nil)
	if err != nil {
		t.Fatal(err)
	}

	want := &Note{
		ID:           3,
		Body:         "foo bar",
		Attachment:   "",
		Title:        "",
		FileName:     "",
		System:       false,
		NoteableID:   4392,
		NoteableType: "Epic",
	}

	if !reflect.DeepEqual(note, want) {
		t.Errorf("Notes.GetEpicNote want %#v, got %#v", note, want)
	}
}

func TestGetMergeRequestNote(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/merge_requests/4329/notes/3", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"id":3,"type":"DiffNote","body":"foo bar","attachment":null,"system":false,"noteable_id":4392,"noteable_type":"Epic","resolvable":false,"noteable_iid":null}`)
	})

	note, _, err := client.Notes.GetMergeRequestNote("1", 4329, 3, nil)
	if err != nil {
		t.Fatal(err)
	}

	want := &Note{
		ID:           3,
		Type:         DiffNote,
		Body:         "foo bar",
		System:       false,
		NoteableID:   4392,
		NoteableType: "Epic",
	}

	if !reflect.DeepEqual(note, want) {
		t.Errorf("Notes.GetEpicNote want %#v, got %#v", note, want)
	}
}

func TestCreateNote(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/issues/1/notes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"id": 1, "body": "Body of note", "author" : {"id" : 1, "name": "snehal", "username": "snehal", "state": "active", "email": "snehal@example.com"}}`)
	})

	createNoteOptions := &CreateIssueNoteOptions{
		Body: Ptr("Body of note"),
	}

	note, _, err := client.Notes.CreateIssueNote("1", 1, createNoteOptions)
	if err != nil {
		t.Fatal(err)
	}

	want := &Note{
		ID:   1,
		Body: "Body of note",
		Author: struct {
			ID        int    "json:\"id\""
			Username  string "json:\"username\""
			Email     string "json:\"email\""
			Name      string "json:\"name\""
			State     string "json:\"state\""
			AvatarURL string "json:\"avatar_url\""
			WebURL    string "json:\"web_url\""
		}{
			ID: 1, Username: "snehal", Name: "snehal", Email: "snehal@example.com", State: "active", AvatarURL: "", WebURL: "",
		},
		Internal: false,
	}

	if !reflect.DeepEqual(want, note) {
		t.Errorf("Notes.CreateNote returned %+v, want %+v", note, want)
	}
}

func TestCreateInternalNote(t *testing.T) {
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/issues/1/notes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"id": 1, "body": "Body of internal note", "internal": true}`)
	})

	createNoteOptions := &CreateIssueNoteOptions{
		Body: Ptr("Body of internal note"),
	}

	note, _, err := client.Notes.CreateIssueNote("1", 1, createNoteOptions)
	if err != nil {
		t.Fatal(err)
	}

	want := &Note{
		ID:       1,
		Body:     "Body of internal note",
		Internal: true,
	}

	if !reflect.DeepEqual(want, note) {
		t.Errorf("Notes.CreateNote returned %+v, want %+v", note, want)
	}
}
