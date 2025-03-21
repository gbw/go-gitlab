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

func TestTagsService_ListTags(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/repository/tags", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
      {
        "name": "1.0.0",
        "message": "test",
        "target": "fffff",
        "protected": false
      },{
        "name": "1.0.1",
        "protected": true
      }
    ]`)
	})

	opt := &ListTagsOptions{ListOptions: ListOptions{Page: 2, PerPage: 3}}

	tags, _, err := client.Tags.ListTags(1, opt)
	if err != nil {
		t.Errorf("Tags.ListTags returned error: %v", err)
	}

	want := []*Tag{
		{
			Name:      "1.0.0",
			Message:   "test",
			Target:    "fffff",
			Protected: false,
		},
		{
			Name:      "1.0.1",
			Protected: true,
		},
	}
	if !reflect.DeepEqual(want, tags) {
		t.Errorf("Tags.ListTags returned %+v, want %+v", tags, want)
	}
}

func TestTagsService_CreateReleaseNote(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/repository/tags/1.0.0/release",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			fmt.Fprint(w, `{"tag_name": "1.0.0", "description": "Amazing release. Wow"}`)
		})

	opt := &CreateReleaseNoteOptions{Description: Ptr("Amazing release. Wow")}

	release, _, err := client.Tags.CreateReleaseNote(1, "1.0.0", opt)
	if err != nil {
		t.Errorf("Tags.CreateRelease returned error: %v", err)
	}

	want := &ReleaseNote{TagName: "1.0.0", Description: "Amazing release. Wow"}
	if !reflect.DeepEqual(want, release) {
		t.Errorf("Tags.CreateRelease returned %+v, want %+v", release, want)
	}
}

func TestTagsService_UpdateReleaseNote(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/repository/tags/1.0.0/release",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPut)
			fmt.Fprint(w, `{"tag_name": "1.0.0", "description": "Amazing release. Wow!"}`)
		})

	opt := &UpdateReleaseNoteOptions{Description: Ptr("Amazing release. Wow!")}

	release, _, err := client.Tags.UpdateReleaseNote(1, "1.0.0", opt)
	if err != nil {
		t.Errorf("Tags.UpdateRelease returned error: %v", err)
	}

	want := &ReleaseNote{TagName: "1.0.0", Description: "Amazing release. Wow!"}
	if !reflect.DeepEqual(want, release) {
		t.Errorf("Tags.UpdateRelease returned %+v, want %+v", release, want)
	}
}
