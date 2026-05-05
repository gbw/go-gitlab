//
// Copyright 2026, GitLab Inc.
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
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOrbitService_GetStatus(t *testing.T) {
	t.Parallel()
	// GIVEN an Orbit cluster reporting healthy status with three components
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/orbit/status", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"status": "healthy",
			"timestamp": "2026-04-28T12:00:00Z",
			"version": "0.5.0",
			"components": [
				{"name": "clickhouse", "status": "healthy", "replicas": {"ready": 3, "desired": 3}, "metrics": {"kind": "Deployment"}},
				{"name": "indexer", "status": "healthy", "replicas": {"ready": 2, "desired": 2}, "metrics": {}},
				{"name": "webserver", "status": "healthy", "replicas": {"ready": 0, "desired": 0}, "metrics": {}}
			]
		}`)
	})

	// WHEN GetStatus is called with no options
	status, resp, err := client.Orbit.GetStatus(nil)

	// THEN the typed response matches the server payload
	require.NoError(t, err)
	require.NotNil(t, resp)

	want := &OrbitStatus{
		Status:    "healthy",
		Timestamp: "2026-04-28T12:00:00Z",
		Version:   "0.5.0",
		Components: []*OrbitStatusComponent{
			{
				Name:     "clickhouse",
				Status:   "healthy",
				Replicas: &OrbitStatusReplicas{Ready: 3, Desired: 3},
				Metrics:  json.RawMessage(`{"kind": "Deployment"}`),
			},
			{
				Name:     "indexer",
				Status:   "healthy",
				Replicas: &OrbitStatusReplicas{Ready: 2, Desired: 2},
				Metrics:  json.RawMessage(`{}`),
			},
			{
				Name:     "webserver",
				Status:   "healthy",
				Replicas: &OrbitStatusReplicas{Ready: 0, Desired: 0},
				Metrics:  json.RawMessage(`{}`),
			},
		},
	}
	assert.Equal(t, want, status)
}

func TestOrbitService_GetStatus_FeatureFlagOff(t *testing.T) {
	t.Parallel()
	// GIVEN the knowledge_graph feature flag is disabled, the API returns 404
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/orbit/status", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{"error":"404 Not Found"}`)
	})

	// WHEN GetStatus is called with no options
	status, resp, err := client.Orbit.GetStatus(nil)

	// THEN the caller receives the underlying response so it can map 404 to a structured error
	require.Error(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	assert.Nil(t, status)
}

func TestOrbitService_GetStatus_LLMFormat(t *testing.T) {
	t.Parallel()
	// GIVEN the orbit/status endpoint returns the compact LLM text when
	// response_format=llm is requested
	mux, client := setup(t)

	var gotQuery string
	mux.HandleFunc("/api/v4/orbit/status", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		gotQuery = r.URL.RawQuery
		fmt.Fprint(w, `{"formatted_text":"status: healthy\nversion: \"0.5.0\""}`)
	})

	// WHEN GetStatus is called with response_format=llm
	status, _, err := client.Orbit.GetStatus(&GetOrbitStatusOptions{
		ResponseFormat: Ptr(OrbitResponseFormatLLM),
	})

	// THEN the response_format parameter is forwarded and FormattedText is populated
	require.NoError(t, err)
	assert.Contains(t, gotQuery, "response_format=llm")
	require.NotNil(t, status)
	assert.Equal(t, "status: healthy\nversion: \"0.5.0\"", status.FormattedText)
	assert.Empty(t, status.Status, "structured fields must be absent in llm response")
}

func TestOrbitService_GetSchema(t *testing.T) {
	t.Parallel()
	// GIVEN a schema response with two domains and summary node entries
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/orbit/schema", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		// AND no expand or format query params
		assert.Empty(t, r.URL.RawQuery)
		fmt.Fprint(w, `{
			"schema_version": "1.0",
			"domains": [
				{"name": "core", "description": "Core entities", "node_names": ["User", "Project"]},
				{"name": "plan", "description": "Planning", "node_names": ["Issue", "Epic"]}
			],
			"nodes": [
				{"name": "User", "domain": "core"},
				{"name": "Project", "domain": "core"}
			],
			"edges": [
				{"name": "AUTHORED", "description": "Authorship", "variants": [
					{"source_type": "User", "target_type": "MergeRequest"},
					{"source_type": "User", "target_type": "Issue"}
				]}
			]
		}`)
	})

	// WHEN GetSchema is called with no options
	schema, resp, err := client.Orbit.GetSchema(nil)

	// THEN domains and edges decode strongly-typed; nodes remain raw
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, "1.0", schema.SchemaVersion)
	assert.Len(t, schema.Domains, 2)
	assert.Equal(t, "core", schema.Domains[0].Name)
	assert.Equal(t, []string{"User", "Project"}, schema.Domains[0].NodeNames)
	assert.Len(t, schema.Nodes, 2)
	assert.JSONEq(t, `{"name": "User", "domain": "core"}`, string(schema.Nodes[0]))
	require.Len(t, schema.Edges, 1)
	assert.Equal(t, "AUTHORED", schema.Edges[0].Name)
	assert.Equal(t, "Authorship", schema.Edges[0].Description)
	require.Len(t, schema.Edges[0].Variants, 2)
	assert.Equal(t, "User", schema.Edges[0].Variants[0].SourceType)
	assert.Equal(t, "MergeRequest", schema.Edges[0].Variants[0].TargetType)
}

func TestOrbitService_GetSchema_WithExpand(t *testing.T) {
	t.Parallel()
	// GIVEN an Orbit schema endpoint that records the request URL
	mux, client := setup(t)

	var gotURL string
	mux.HandleFunc("/api/v4/orbit/schema", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		gotURL = r.URL.RawQuery
		fmt.Fprint(w, `{"schema_version": "1.0"}`)
	})

	// WHEN GetSchema is called with multiple expand nodes and llm format
	expand := []string{"User", "Project", "MergeRequest"}
	_, _, err := client.Orbit.GetSchema(&GetOrbitSchemaOptions{
		Expand: &expand,
		Format: Ptr(OrbitResponseFormatLLM),
	})

	// THEN expand is comma-joined per API convention and format is set
	require.NoError(t, err)
	assert.Contains(t, gotURL, "expand=User%2CProject%2CMergeRequest")
	assert.Contains(t, gotURL, "format=llm")
}

func TestOrbitService_GetTools(t *testing.T) {
	t.Parallel()
	// GIVEN a tools response with two MCP tool definitions
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/orbit/tools", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		// The wire format is a bare JSON array, not a wrapped object.
		fmt.Fprint(w, `[
			{
				"name": "query_graph",
				"description": "Execute graph queries",
				"parameters": {"type": "object", "properties": {"query": {}}}
			},
			{
				"name": "get_graph_schema",
				"description": "List Knowledge Graph schema",
				"parameters": {"type": "object"}
			}
		]`)
	})

	// WHEN GetTools is called
	tools, resp, err := client.Orbit.GetTools()

	// THEN both tools are decoded with their parameter schemas preserved as raw JSON
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Len(t, tools.Tools, 2)
	assert.Equal(t, "query_graph", tools.Tools[0].Name)
	assert.JSONEq(t,
		`{"type": "object", "properties": {"query": {}}}`,
		string(tools.Tools[0].Parameters),
	)
	assert.Equal(t, "get_graph_schema", tools.Tools[1].Name)
}

func TestOrbitService_Query(t *testing.T) {
	t.Parallel()
	// GIVEN an Orbit query endpoint that echoes back a structured result
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/orbit/query", func(w http.ResponseWriter, r *http.Request) {
		// AND the request must be POST with a JSON content type
		testMethod(t, r, http.MethodPost)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		// AND the body matches the typed request, with `query` passed through unchanged
		want := map[string]any{
			"query": map[string]any{
				"query_type": "traversal",
				"node":       map[string]any{"id": "p", "entity": "Project"},
				"limit":      float64(5),
			},
			"response_format": "raw",
		}
		var got map[string]any
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&got))
		assert.Equal(t, want, got)

		fmt.Fprint(w, `{
			"result": [
				{"_id": "1", "_type": "Project", "name": "alpha"},
				{"_id": "2", "_type": "Project", "name": "beta"}
			],
			"query_type": "traversal",
			"raw_query_strings": ["SELECT ..."],
			"row_count": 2
		}`)
	})

	// WHEN Query is called with a typed request whose query is opaque JSON
	queryDSL := json.RawMessage(`{
		"query_type": "traversal",
		"node": {"id": "p", "entity": "Project"},
		"limit": 5
	}`)
	result, resp, err := client.Orbit.Query(&OrbitQueryRequest{
		Query:          queryDSL,
		ResponseFormat: Ptr(OrbitResponseFormatRaw),
	})

	// THEN the result envelope decodes and `result` stays as raw JSON
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, "traversal", result.QueryType)
	assert.Equal(t, []string{"SELECT ..."}, result.RawQueryStrings)
	assert.Equal(t, int64(2), result.RowCount)
	assert.JSONEq(t,
		`[{"_id": "1", "_type": "Project", "name": "alpha"}, {"_id": "2", "_type": "Project", "name": "beta"}]`,
		string(result.Result),
	)
}

func TestOrbitService_Query_LLMFormat(t *testing.T) {
	t.Parallel()
	// GIVEN response_format="llm" returns `result` as a JSON-encoded GOON string
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/orbit/query", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		fmt.Fprint(w, `{
			"result": "@goon{v:1}\nnodes:\n  Project[1]{id,name}:\n    1,alpha",
			"query_type": "traversal",
			"row_count": 1
		}`)
	})

	// WHEN Query is called with response_format=llm
	result, _, err := client.Orbit.Query(&OrbitQueryRequest{
		Query:          json.RawMessage(`{"query_type": "traversal"}`),
		ResponseFormat: Ptr(OrbitResponseFormatLLM),
	})

	// THEN the JSON-encoded string is preserved verbatim in Result
	require.NoError(t, err)
	var resultStr string
	require.NoError(t, json.Unmarshal(result.Result, &resultStr))
	assert.Equal(t, "@goon{v:1}\nnodes:\n  Project[1]{id,name}:\n    1,alpha", resultStr)
	assert.Equal(t, int64(1), result.RowCount)
}

func TestOrbitService_Query_NamespaceForbidden(t *testing.T) {
	t.Parallel()
	// GIVEN the user has the FF on but no Knowledge Graph enabled namespaces
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/orbit/query", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, `{"message":"403 No Knowledge Graph enabled namespaces available"}`)
	})

	// WHEN Query is called
	result, resp, err := client.Orbit.Query(&OrbitQueryRequest{
		Query: json.RawMessage(`{"query_type": "traversal"}`),
	})

	// THEN the caller receives the response so it can surface a structured error
	require.Error(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusForbidden, resp.StatusCode)
	assert.Nil(t, result)
}

func TestOrbitService_GetGraphStatus_ByNamespaceID(t *testing.T) {
	t.Parallel()
	// GIVEN an orbit graph_status endpoint returning indexed status
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/orbit/graph_status", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		// AND the request carries the namespace_id query parameter
		assert.Equal(t, "42", r.URL.Query().Get("namespace_id"))
		fmt.Fprint(w, `{
			"projects": {"indexed": 5, "total_known": 10},
			"domains": [
				{
					"name": "SDLC",
					"items": [
						{"name": "MergeRequest", "count": 42}
					]
				}
			],
			"indexing": {
				"state": "indexed",
				"last_started_at": "2026-04-15T00:00:00Z",
				"last_completed_at": "2026-04-15T01:00:00Z",
				"last_duration_ms": 3600000,
				"last_error": null
			}
		}`)
	})

	// WHEN GetGraphStatus is called with a namespace_id
	namespaceID := int64(42)
	status, resp, err := client.Orbit.GetGraphStatus(&GetGraphStatusOptions{
		NamespaceID: &namespaceID,
	})

	// THEN the response decodes into the typed struct
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, status)

	require.NotNil(t, status.Projects)
	assert.Equal(t, int64(5), status.Projects.Indexed)
	assert.Equal(t, int64(10), status.Projects.TotalKnown)

	require.Len(t, status.Domains, 1)
	assert.Equal(t, "SDLC", status.Domains[0].Name)
	require.Len(t, status.Domains[0].Items, 1)
	assert.Equal(t, "MergeRequest", status.Domains[0].Items[0].Name)
	assert.Equal(t, int64(42), status.Domains[0].Items[0].Count)

	require.NotNil(t, status.Indexing)
	assert.Equal(t, "indexed", status.Indexing.State)
	require.NotNil(t, status.Indexing.LastStartedAt)
	assert.Equal(t, "2026-04-15T00:00:00Z", status.Indexing.LastStartedAt.UTC().Format("2006-01-02T15:04:05Z"))
	require.NotNil(t, status.Indexing.LastCompletedAt)
	assert.Equal(t, "2026-04-15T01:00:00Z", status.Indexing.LastCompletedAt.UTC().Format("2006-01-02T15:04:05Z"))
	require.NotNil(t, status.Indexing.LastDurationMs)
	assert.Equal(t, int64(3600000), *status.Indexing.LastDurationMs)
	assert.Nil(t, status.Indexing.LastError)
}

func TestOrbitService_GetGraphStatus_ByProjectID(t *testing.T) {
	t.Parallel()
	// GIVEN an orbit graph_status endpoint that records the query parameters
	mux, client := setup(t)

	var gotQuery string
	mux.HandleFunc("/api/v4/orbit/graph_status", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		gotQuery = r.URL.RawQuery
		fmt.Fprint(w, `{
			"projects": {"indexed": 1, "total_known": 1},
			"domains": [],
			"indexing": {"state": "indexed"}
		}`)
	})

	// WHEN GetGraphStatus is called with a project_id
	projectID := int64(99)
	status, _, err := client.Orbit.GetGraphStatus(&GetGraphStatusOptions{
		ProjectID: &projectID,
	})

	// THEN the project_id query param is forwarded and the response decodes
	require.NoError(t, err)
	assert.Contains(t, gotQuery, "project_id=99")
	require.NotNil(t, status)
	assert.Equal(t, "indexed", status.Indexing.State)
}

func TestOrbitService_GetGraphStatus_ByFullPath(t *testing.T) {
	t.Parallel()
	// GIVEN an orbit graph_status endpoint that records the query parameters
	mux, client := setup(t)

	var gotQuery string
	mux.HandleFunc("/api/v4/orbit/graph_status", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		gotQuery = r.URL.RawQuery
		fmt.Fprint(w, `{
			"projects": {"indexed": 3, "total_known": 3},
			"domains": [],
			"indexing": {"state": "indexed"}
		}`)
	})

	// WHEN GetGraphStatus is called with a full_path
	fullPath := "gitlab-org/gitlab"
	status, _, err := client.Orbit.GetGraphStatus(&GetGraphStatusOptions{
		FullPath: &fullPath,
	})

	// THEN the full_path query param is forwarded (URL-encoded) and the response decodes
	require.NoError(t, err)
	assert.Contains(t, gotQuery, "full_path=gitlab-org%2Fgitlab")
	require.NotNil(t, status)
	assert.Equal(t, "indexed", status.Indexing.State)
}

func TestOrbitService_GetGraphStatus_LLMFormat(t *testing.T) {
	t.Parallel()
	// GIVEN the orbit/graph_status endpoint returns only formatted_text
	// when response_format=llm is requested (structured fields are absent)
	mux, client := setup(t)

	var gotQuery string
	mux.HandleFunc("/api/v4/orbit/graph_status", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		gotQuery = r.URL.RawQuery
		fmt.Fprint(w, `{"formatted_text":"projects:\n  indexed: 2\n  total_known: 2\nindexing:\n  state: indexed"}`)
	})

	// WHEN GetGraphStatus is called with response_format=llm
	namespaceID := int64(7)
	status, _, err := client.Orbit.GetGraphStatus(&GetGraphStatusOptions{
		NamespaceID:    &namespaceID,
		ResponseFormat: Ptr(OrbitResponseFormatLLM),
	})

	// THEN the response_format parameter is forwarded and FormattedText is populated
	require.NoError(t, err)
	assert.Contains(t, gotQuery, "response_format=llm")
	require.NotNil(t, status)
	assert.Equal(t, "projects:\n  indexed: 2\n  total_known: 2\nindexing:\n  state: indexed", status.FormattedText)
	assert.Nil(t, status.Projects, "structured fields must be absent in llm response")
	assert.Nil(t, status.Indexing, "structured fields must be absent in llm response")
}

func TestOrbitService_GetGraphStatus_Forbidden(t *testing.T) {
	t.Parallel()
	// GIVEN the user has no Knowledge Graph enabled namespaces
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/orbit/graph_status", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, `{"message":"403 No Knowledge Graph enabled namespaces available"}`)
	})

	// WHEN GetGraphStatus is called
	namespaceID := int64(1)
	status, resp, err := client.Orbit.GetGraphStatus(&GetGraphStatusOptions{
		NamespaceID: &namespaceID,
	})

	// THEN a 403 error is returned and the response is available for inspection
	require.Error(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusForbidden, resp.StatusCode)
	assert.Nil(t, status)
}

func TestOrbitService_GetGraphStatus_ServiceUnavailable(t *testing.T) {
	t.Parallel()
	// GIVEN the GKG gRPC service is unavailable
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/orbit/graph_status", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprint(w, `{"message":"503 Service Unavailable"}`)
	})

	// WHEN GetGraphStatus is called
	namespaceID := int64(1)
	status, resp, err := client.Orbit.GetGraphStatus(&GetGraphStatusOptions{
		NamespaceID: &namespaceID,
	})

	// THEN a 503 error is returned and the response is available for inspection
	require.Error(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusServiceUnavailable, resp.StatusCode)
	assert.Nil(t, status)
}
