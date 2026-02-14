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

func TestListCustomUserAttributes(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/users/2/custom_attributes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"key":"testkey1", "value":"testvalue1"}, {"key":"testkey2", "value":"testvalue2"}]`)
	})

	customAttributes, _, err := client.CustomAttribute.ListCustomUserAttributes(2)
	require.NoError(t, err)

	want := []*CustomAttribute{{Key: "testkey1", Value: "testvalue1"}, {Key: "testkey2", Value: "testvalue2"}}
	assert.Equal(t, want, customAttributes)
}

func TestListCustomGroupAttributes(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/2/custom_attributes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"key":"testkey1", "value":"testvalue1"}, {"key":"testkey2", "value":"testvalue2"}]`)
	})

	customAttributes, _, err := client.CustomAttribute.ListCustomGroupAttributes(2)
	require.NoError(t, err)

	want := []*CustomAttribute{{Key: "testkey1", Value: "testvalue1"}, {Key: "testkey2", Value: "testvalue2"}}
	assert.Equal(t, want, customAttributes)
}

func TestListCustomProjectAttributes(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/2/custom_attributes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[{"key":"testkey1", "value":"testvalue1"}, {"key":"testkey2", "value":"testvalue2"}]`)
	})

	customAttributes, _, err := client.CustomAttribute.ListCustomProjectAttributes(2)
	require.NoError(t, err)

	want := []*CustomAttribute{{Key: "testkey1", Value: "testvalue1"}, {Key: "testkey2", Value: "testvalue2"}}
	assert.Equal(t, want, customAttributes)
}

func TestGetCustomUserAttribute(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/users/2/custom_attributes/testkey1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"key":"testkey1", "value":"testvalue1"}`)
	})

	customAttribute, _, err := client.CustomAttribute.GetCustomUserAttribute(2, "testkey1")
	require.NoError(t, err)

	want := &CustomAttribute{Key: "testkey1", Value: "testvalue1"}
	assert.Equal(t, want, customAttribute)
}

func TestGetCustomGroupAttribute(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/2/custom_attributes/testkey1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"key":"testkey1", "value":"testvalue1"}`)
	})

	customAttribute, _, err := client.CustomAttribute.GetCustomGroupAttribute(2, "testkey1")
	require.NoError(t, err)

	want := &CustomAttribute{Key: "testkey1", Value: "testvalue1"}
	assert.Equal(t, want, customAttribute)
}

func TestGetCustomProjectAttribute(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/2/custom_attributes/testkey1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{"key":"testkey1", "value":"testvalue1"}`)
	})

	customAttribute, _, err := client.CustomAttribute.GetCustomProjectAttribute(2, "testkey1")
	require.NoError(t, err)

	want := &CustomAttribute{Key: "testkey1", Value: "testvalue1"}
	assert.Equal(t, want, customAttribute)
}

func TestSetCustomUserAttribute(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/users/2/custom_attributes/testkey1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprint(w, `{"key":"testkey1", "value":"testvalue1"}`)
	})

	customAttribute, _, err := client.CustomAttribute.SetCustomUserAttribute(2, CustomAttribute{
		Key:   "testkey1",
		Value: "testvalue1",
	})
	require.NoError(t, err)

	want := &CustomAttribute{Key: "testkey1", Value: "testvalue1"}
	assert.Equal(t, want, customAttribute)
}

func TestSetCustomGroupAttribute(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/2/custom_attributes/testkey1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		fmt.Fprint(w, `{"key":"testkey1", "value":"testvalue1"}`)
	})

	customAttribute, _, err := client.CustomAttribute.SetCustomGroupAttribute(2, CustomAttribute{
		Key:   "testkey1",
		Value: "testvalue1",
	})
	require.NoError(t, err)

	want := &CustomAttribute{Key: "testkey1", Value: "testvalue1"}
	assert.Equal(t, want, customAttribute)
}

func TestDeleteCustomUserAttribute(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/users/2/custom_attributes/testkey1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusAccepted)
	})

	resp, err := client.CustomAttribute.DeleteCustomUserAttribute(2, "testkey1")
	require.NoError(t, err)

	assert.Equal(t, http.StatusAccepted, resp.StatusCode)
}

func TestDeleteCustomGroupAttribute(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/groups/2/custom_attributes/testkey1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusAccepted)
	})

	resp, err := client.CustomAttribute.DeleteCustomGroupAttribute(2, "testkey1")
	require.NoError(t, err)

	assert.Equal(t, http.StatusAccepted, resp.StatusCode)
}

func TestDeleteCustomProjectAttribute(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/2/custom_attributes/testkey1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusAccepted)
	})

	resp, err := client.CustomAttribute.DeleteCustomProjectAttribute(2, "testkey1")
	require.NoError(t, err)

	assert.Equal(t, http.StatusAccepted, resp.StatusCode)
}
