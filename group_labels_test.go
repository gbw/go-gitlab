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

func TestCreateGroupGroupLabel(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/labels", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{"id":1, "name": "MyGroupLabel", "color" : "#11FF22"}`)
	})

	l := &CreateGroupLabelOptions{
		Name:  Ptr("MyGroupLabel"),
		Color: Ptr("#11FF22"),
	}
	label, _, err := client.GroupLabels.CreateGroupLabel("1", l)
	if err != nil {
		t.Fatal(err)
	}
	want := &GroupLabel{ID: 1, Name: "MyGroupLabel", Color: "#11FF22"}
	if !reflect.DeepEqual(want, label) {
		t.Errorf("GroupLabels.CreateGroupLabel returned %+v, want %+v", label, want)
	}
}

func TestDeleteGroupLabelByID(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/labels/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.GroupLabels.DeleteGroupLabel("1", "1", nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteGroupLabelByName(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/labels/MyGroupLabel", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	_, err := client.GroupLabels.DeleteGroupLabel("1", "MyGroupLabel", nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateGroupLabel(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/labels/MyGroupLabel", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprint(w, `{"id":1, "name": "NewLabel", "color" : "#11FF23" , "description":"This is updated label"}`)
	})

	l := &UpdateGroupLabelOptions{
		NewName:     Ptr("NewLabel"),
		Color:       Ptr("#11FF23"),
		Description: Ptr("This is updated label"),
	}

	label, resp, err := client.GroupLabels.UpdateGroupLabel("1", "MyGroupLabel", l)

	if resp == nil {
		t.Fatal(err)
	}
	if err != nil {
		t.Fatal(err)
	}

	want := &GroupLabel{ID: 1, Name: "NewLabel", Color: "#11FF23", Description: "This is updated label"}

	if !reflect.DeepEqual(want, label) {
		t.Errorf("GroupLabels.UpdateGroupLabel returned %+v, want %+v", label, want)
	}
}

func TestSubscribeToGroupLabel(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/labels/5/subscribe", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{  "id" : 5, "name" : "kind/bug", "color" : "#d9534f", "description": "Bug reported by user", "open_issues_count": 1, "closed_issues_count": 0, "open_merge_requests_count": 1, "subscribed": true,"priority": null}`)
	})

	label, _, err := client.GroupLabels.SubscribeToGroupLabel("1", "5")
	if err != nil {
		t.Fatal(err)
	}
	want := &GroupLabel{ID: 5, Name: "kind/bug", Color: "#d9534f", Description: "Bug reported by user", OpenIssuesCount: 1, ClosedIssuesCount: 0, OpenMergeRequestsCount: 1, Subscribed: true}
	if !reflect.DeepEqual(want, label) {
		t.Errorf("GroupLabels.SubscribeToGroupLabel returned %+v, want %+v", label, want)
	}
}

func TestUnsubscribeFromGroupLabel(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/labels/5/unsubscribe", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
	})

	_, err := client.GroupLabels.UnsubscribeFromGroupLabel("1", "5")
	if err != nil {
		t.Fatal(err)
	}
}

func TestListGroupLabels(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/labels", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{  "id" : 5, "name" : "kind/bug", "color" : "#d9534f", "description": "Bug reported by user", "open_issues_count": 1, "closed_issues_count": 0, "open_merge_requests_count": 1, "subscribed": true,"priority": null}]`)
	})

	o := &ListGroupLabelsOptions{
		ListOptions: ListOptions{
			Page:    1,
			PerPage: 10,
		},
	}
	label, _, err := client.GroupLabels.ListGroupLabels("1", o)
	if err != nil {
		t.Log(err.Error() == "invalid ID type 1.1, the ID must be an int or a string")
	}
	want := []*GroupLabel{{ID: 5, Name: "kind/bug", Color: "#d9534f", Description: "Bug reported by user", OpenIssuesCount: 1, ClosedIssuesCount: 0, OpenMergeRequestsCount: 1, Subscribed: true}}
	if !reflect.DeepEqual(want, label) {
		t.Errorf("GroupLabels.ListGroupLabels returned %+v, want %+v", label, want)
	}
}

func TestGetGroupLabel(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/1/labels/5", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{  "id" : 5, "name" : "kind/bug", "color" : "#d9534f", "description": "Bug reported by user", "open_issues_count": 1, "closed_issues_count": 0, "open_merge_requests_count": 1, "subscribed": true,"priority": null}`)
	})

	label, _, err := client.GroupLabels.GetGroupLabel("1", 5)
	if err != nil {
		t.Log(err)
	}

	want := &GroupLabel{ID: 5, Name: "kind/bug", Color: "#d9534f", Description: "Bug reported by user", OpenIssuesCount: 1, ClosedIssuesCount: 0, OpenMergeRequestsCount: 1, Subscribed: true}
	if !reflect.DeepEqual(want, label) {
		t.Errorf("GroupLabels.GetGroupLabel returned %+v, want %+v", label, want)
	}
}
