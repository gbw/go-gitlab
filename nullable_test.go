//
// Copyright 2024 oapi-codegen
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
// This code was adapted from https://github.com/oapi-codegen/nullable
//

package gitlab

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

type Obj struct {
	Foo Nullable[string] `json:"foo,omitempty"` // note "omitempty" is important for fields that are optional
}

func TestNullable(t *testing.T) {
	t.Parallel()

	// --- parsing from json and serializing back to JSON
	t.Run("JSON Unmarshal/Marshal", func(t *testing.T) {
		t.Parallel()

		t.Run("Case where there is an actual value", func(t *testing.T) {
			t.Parallel()
			data := `{"foo":"bar"}`
			// deserialize from json
			myObj := parse(data, t)
			require.Equal(t, Obj{Foo: Nullable[string]{true: "bar"}}, myObj)
			require.False(t, myObj.Foo.IsNull())
			require.True(t, myObj.Foo.IsSpecified())
			value, err := myObj.Foo.Get()
			require.NoError(t, err)
			require.Equal(t, "bar", value)
			require.Equal(t, "bar", myObj.Foo.MustGet())
			// serialize back to json: leads to the same data
			require.Equal(t, data, serialize(myObj, t))
		})

		t.Run("Case where no value is specified", func(t *testing.T) {
			t.Parallel()
			data := `{}`
			// deserialize from json
			myObj := parse(data, t)
			require.Equal(t, Obj{Foo: nil}, myObj)
			require.False(t, myObj.Foo.IsNull())
			require.False(t, myObj.Foo.IsSpecified())
			_, err := myObj.Foo.Get()
			require.ErrorContains(t, err, "value is not specified")
			// serialize back to json: leads to the same data
			require.Equal(t, data, serialize(myObj, t))
		})

		t.Run("Case where the specified value is explicitly null", func(t *testing.T) {
			t.Parallel()
			data := `{"foo":null}`
			// deserialize from json
			myObj := parse(data, t)
			require.Equal(t, Obj{Foo: Nullable[string]{false: ""}}, myObj)
			require.True(t, myObj.Foo.IsNull())
			require.True(t, myObj.Foo.IsSpecified())
			_, err := myObj.Foo.Get()
			require.ErrorContains(t, err, "value is null")
			require.Panics(t, func() { myObj.Foo.MustGet() })
			// serialize back to json: leads to the same data
			require.Equal(t, data, serialize(myObj, t))
		})
	})

	// --- building objects from a Go client
	t.Run("Go Client Object", func(t *testing.T) {
		t.Parallel()

		t.Run("Case where there is an actual value", func(t *testing.T) {
			t.Parallel()
			myObj := Obj{}
			myObj.Foo.Set("bar")
			require.JSONEq(t, `{"foo":"bar"}`, serialize(myObj, t))
		})

		t.Run("Case where the value should be unspecified", func(t *testing.T) {
			t.Parallel()
			myObj := Obj{}
			// do nothing: unspecified by default
			require.JSONEq(t, `{}`, serialize(myObj, t))
			// explicitly mark unspecified
			myObj.Foo.SetUnspecified()
			require.JSONEq(t, `{}`, serialize(myObj, t))
		})

		t.Run("Case where the value should be null", func(t *testing.T) {
			t.Parallel()
			myObj := Obj{}
			myObj.Foo.SetNull()
			require.JSONEq(t, `{"foo":null}`, serialize(myObj, t))
		})
	})
}

func parse(data string, t *testing.T) Obj {
	var myObj Obj
	err := json.Unmarshal([]byte(data), &myObj)
	require.NoError(t, err)
	return myObj
}

func serialize(o Obj, t *testing.T) string {
	data, err := json.Marshal(o)
	require.NoError(t, err)
	return string(data)
}
