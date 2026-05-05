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
	"net/http"
	"time"
)

type (
	// OrbitServiceInterface handles communication with the Orbit
	// (GitLab Knowledge Graph) endpoints of the GitLab API.
	//
	// Note: This API is experimental and may change or be removed
	// in future versions.
	//
	// GitLab API docs: https://docs.gitlab.com/api/orbit/
	OrbitServiceInterface interface {
		// GetStatus returns Orbit (GitLab Knowledge Graph) cluster
		// health.
		//
		// Note: This API is experimental and may change or be
		// removed in future versions.
		//
		// GitLab API docs:
		// https://docs.gitlab.com/api/orbit/#get-status
		GetStatus(opt *GetOrbitStatusOptions, options ...RequestOptionFunc) (*OrbitStatus, *Response, error)

		// GetSchema returns the Orbit graph ontology (domains, nodes,
		// edges).
		//
		// Note: This API is experimental and may change or be
		// removed in future versions.
		//
		// GitLab API docs:
		// https://docs.gitlab.com/api/orbit/#get-schema
		GetSchema(opt *GetOrbitSchemaOptions, options ...RequestOptionFunc) (*OrbitSchema, *Response, error)

		// GetTools returns the Orbit MCP tool manifest.
		//
		// Note: This API is experimental and may change or be
		// removed in future versions.
		//
		// GitLab API docs:
		// https://docs.gitlab.com/api/orbit/#get-tools
		GetTools(options ...RequestOptionFunc) (*OrbitTools, *Response, error)

		// Query executes an Orbit (Knowledge Graph) query.
		//
		// Note: This API is experimental and may change or be
		// removed in future versions.
		//
		// GitLab API docs:
		// https://docs.gitlab.com/api/orbit/#post-query
		Query(opt *OrbitQueryRequest, options ...RequestOptionFunc) (*OrbitQueryResult, *Response, error)

		// GetGraphStatus returns the indexing status of the Knowledge
		// Graph for a namespace or project.
		//
		// Note: This API is experimental and may change or be
		// removed in future versions.
		//
		// GitLab API docs:
		// https://docs.gitlab.com/api/orbit/#get-graph-status
		GetGraphStatus(opt *GetGraphStatusOptions, options ...RequestOptionFunc) (*OrbitGraphStatus, *Response, error)
	}

	// OrbitService handles communication with the Orbit (GitLab
	// Knowledge Graph) endpoints of the GitLab API. Orbit is the
	// product name for the Knowledge Graph (GKG) service.
	//
	// All endpoints are user-scoped (not project-scoped) and gated
	// behind the `knowledge_graph` feature flag. When the flag is off
	// every `/api/v4/orbit/*` endpoint returns HTTP 404.
	//
	// Note: This API is experimental and may change or be removed
	// in future versions. Response shapes may evolve before Orbit
	// reaches general availability. Where the response schema is
	// format-dependent (`raw` vs `llm`) or too deep to type fully,
	// fields are exposed as json.RawMessage so callers can render or
	// forward them verbatim.
	//
	// GitLab API docs: https://docs.gitlab.com/api/orbit/
	OrbitService struct {
		client *Client
	}
)

var _ OrbitServiceInterface = (*OrbitService)(nil)

// GetOrbitStatusOptions represents the available GetStatus() options.
//
// GitLab API docs: https://docs.gitlab.com/api/orbit/#get-status
type GetOrbitStatusOptions struct {
	// ResponseFormat selects the response shape: "raw" (structured
	// JSON, default) or "llm" (compact text for LLM consumption).
	ResponseFormat *OrbitResponseFormatValue `url:"response_format,omitempty"`
}

// OrbitStatus represents the Orbit cluster health response returned
// by `GET /api/v4/orbit/status`.
//
// When ResponseFormat is "raw" (default), the structured fields
// (Status, Timestamp, Version, Components) are populated and
// FormattedText is empty. When ResponseFormat is "llm", only
// FormattedText is populated and the structured fields are absent.
//
// GitLab API docs: https://docs.gitlab.com/api/orbit/#get-status
type OrbitStatus struct {
	FormattedText string                  `json:"formatted_text,omitempty"`
	Status        string                  `json:"status,omitempty"`
	Timestamp     string                  `json:"timestamp,omitempty"`
	Version       string                  `json:"version,omitempty"`
	Components    []*OrbitStatusComponent `json:"components,omitempty"`
}

// OrbitStatusComponent represents a single Orbit subsystem in the
// status response.
//
// GitLab API docs: https://docs.gitlab.com/api/orbit/#get-status
type OrbitStatusComponent struct {
	Name     string               `json:"name,omitempty"`
	Status   string               `json:"status,omitempty"`
	Replicas *OrbitStatusReplicas `json:"replicas,omitempty"`
	Metrics  json.RawMessage      `json:"metrics,omitempty"`
}

// OrbitStatusReplicas represents the ready/desired replica counts of
// a single Orbit subsystem.
//
// GitLab API docs: https://docs.gitlab.com/api/orbit/#get-status
type OrbitStatusReplicas struct {
	Ready   int64 `json:"ready"`
	Desired int64 `json:"desired"`
}

// GetOrbitSchemaOptions represents the available GetSchema() options.
//
// GitLab API docs: https://docs.gitlab.com/api/orbit/#get-schema
type GetOrbitSchemaOptions struct {
	// Expand is the list of node names to expand with full properties
	// and relationships. The serialized query parameter is comma-joined,
	// matching the API convention: `?expand=User,Project,MergeRequest`.
	Expand *[]string `url:"expand,comma,omitempty" json:"expand,omitempty"`

	// Format selects the response format: `raw` (structured JSON) or
	// `llm` (TOON text optimized for LLM consumption). Defaults to
	// `raw` server-side when omitted.
	Format *OrbitResponseFormatValue `url:"format,omitempty" json:"format,omitempty"`
}

// OrbitSchema represents the Orbit graph ontology returned by
// `GET /api/v4/orbit/schema` with `format=raw`.
//
// Node entries are exposed as json.RawMessage because their shape
// depends on the `expand` parameter — without `expand` each entry
// carries only summary fields; with `expand` matched nodes include
// `properties`, `style`, and incoming/outgoing edge lists, while
// unmatched nodes remain summary-only in the same response.
//
// GitLab API docs: https://docs.gitlab.com/api/orbit/#get-schema
type OrbitSchema struct {
	SchemaVersion string               `json:"schema_version,omitempty"`
	Domains       []*OrbitSchemaDomain `json:"domains,omitempty"`
	Nodes         []json.RawMessage    `json:"nodes,omitempty"`
	Edges         []*OrbitSchemaEdge   `json:"edges,omitempty"`
}

// OrbitSchemaDomain represents a logical grouping of nodes in the
// graph ontology.
//
// GitLab API docs: https://docs.gitlab.com/api/orbit/#get-schema
type OrbitSchemaDomain struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	NodeNames   []string `json:"node_names,omitempty"`
}

// OrbitSchemaEdge represents an edge type in the graph ontology and
// the source/target node combinations it connects.
//
// GitLab API docs: https://docs.gitlab.com/api/orbit/#get-schema
type OrbitSchemaEdge struct {
	Name        string                    `json:"name,omitempty"`
	Description string                    `json:"description,omitempty"`
	Variants    []*OrbitSchemaEdgeVariant `json:"variants,omitempty"`
}

// OrbitSchemaEdgeVariant represents a single source/target pair for
// an edge type.
//
// GitLab API docs: https://docs.gitlab.com/api/orbit/#get-schema
type OrbitSchemaEdgeVariant struct {
	SourceType string `json:"source_type,omitempty"`
	TargetType string `json:"target_type,omitempty"`
}

// OrbitTools represents the MCP tool manifest returned by
// `GET /api/v4/orbit/tools`. The Rust service is the source of truth
// for tool metadata; Rails passes the response through verbatim.
//
// On the wire the response body is a bare JSON array of tool
// definitions; the typed wrapper here exposes a named Tools field for
// convenience and decodes the bare array via a custom UnmarshalJSON.
//
// GitLab API docs: https://docs.gitlab.com/api/orbit/#get-tools
type OrbitTools struct {
	Tools []*OrbitTool `json:"tools,omitempty"`
}

// UnmarshalJSON decodes the bare JSON array returned by
// `/api/v4/orbit/tools` into the Tools slice.
func (t *OrbitTools) UnmarshalJSON(data []byte) error {
	var tools []*OrbitTool
	if err := json.Unmarshal(data, &tools); err != nil {
		return err
	}
	t.Tools = tools
	return nil
}

// MarshalJSON encodes OrbitTools as a bare JSON array, matching the
// wire format served by `/api/v4/orbit/tools`.
func (t OrbitTools) MarshalJSON() ([]byte, error) {
	if t.Tools == nil {
		return []byte("[]"), nil
	}
	return json.Marshal(t.Tools)
}

// OrbitTool represents a single MCP tool definition.
//
// GitLab API docs: https://docs.gitlab.com/api/orbit/#get-tools
type OrbitTool struct {
	Name        string          `json:"name,omitempty"`
	Description string          `json:"description,omitempty"`
	Parameters  json.RawMessage `json:"parameters,omitempty"`
}

// OrbitQueryRequest represents the JSON body sent to
// `POST /api/v4/orbit/query`.
//
// Query is the structured DSL query object. It is left as
// json.RawMessage so callers (and `glab orbit query`) can pass the
// authoritative DSL through unchanged — the live JSON Schema is
// available from `GET /api/v4/orbit/tools` and is the source of truth
// for query shape.
//
// ResponseFormat selects the response shape: "raw" (structured JSON)
// or "llm" (compact text optimized for LLM consumption).
//
// GitLab API docs: https://docs.gitlab.com/api/orbit/#post-query
type OrbitQueryRequest struct {
	Query          json.RawMessage           `json:"query"`
	ResponseFormat *OrbitResponseFormatValue `json:"response_format,omitempty"`
}

// OrbitQueryResult represents the response returned by
// `POST /api/v4/orbit/query`.
//
// The exact shape of Result is opaque and depends on the query type
// and the requested response_format; treat it as a raw JSON payload
// to render or forward verbatim. RawQueryStrings entries are
// themselves JSON-encoded objects (not raw SQL) — each describes a
// compiled query stage produced by the Knowledge Graph engine.
//
// GitLab API docs: https://docs.gitlab.com/api/orbit/#post-query
type OrbitQueryResult struct {
	Result          json.RawMessage `json:"result,omitempty"`
	QueryType       string          `json:"query_type,omitempty"`
	RawQueryStrings []string        `json:"raw_query_strings,omitempty"`
	RowCount        int64           `json:"row_count,omitempty"`
}

// GetStatus returns Orbit cluster health from
// `GET /api/v4/orbit/status`.
//
// Note: This API is experimental and may change or be removed in
// future versions.
//
// GitLab API docs: https://docs.gitlab.com/api/orbit/#get-status
func (s *OrbitService) GetStatus(opt *GetOrbitStatusOptions, options ...RequestOptionFunc) (*OrbitStatus, *Response, error) {
	return do[*OrbitStatus](s.client,
		withMethod(http.MethodGet),
		withPath("orbit/status"),
		withAPIOpts(opt),
		withRequestOpts(options...),
	)
}

// GetSchema returns the Orbit graph ontology from
// `GET /api/v4/orbit/schema`.
//
// Note: This API is experimental and may change or be removed in
// future versions.
//
// GitLab API docs: https://docs.gitlab.com/api/orbit/#get-schema
func (s *OrbitService) GetSchema(opt *GetOrbitSchemaOptions, options ...RequestOptionFunc) (*OrbitSchema, *Response, error) {
	return do[*OrbitSchema](s.client,
		withMethod(http.MethodGet),
		withPath("orbit/schema"),
		withAPIOpts(opt),
		withRequestOpts(options...),
	)
}

// GetTools returns the Orbit MCP tool manifest from
// `GET /api/v4/orbit/tools`.
//
// Note: This API is experimental and may change or be removed in
// future versions.
//
// GitLab API docs: https://docs.gitlab.com/api/orbit/#get-tools
func (s *OrbitService) GetTools(options ...RequestOptionFunc) (*OrbitTools, *Response, error) {
	return do[*OrbitTools](s.client,
		withMethod(http.MethodGet),
		withPath("orbit/tools"),
		withRequestOpts(options...),
	)
}

// Query executes an Orbit (Knowledge Graph) query against
// `POST /api/v4/orbit/query`.
//
// The request body is a typed OrbitQueryRequest whose Query member is
// passed through as raw JSON; this preserves the authoritative DSL
// shape served by `GET /api/v4/orbit/tools` without forcing a static
// type that would lag the server. NewRequest serializes the body and
// sets `Content-Type: application/json` automatically — there is no
// need to manage the header manually.
//
// Note: This API is experimental and may change or be removed in
// future versions.
//
// GitLab API docs: https://docs.gitlab.com/api/orbit/#post-query
func (s *OrbitService) Query(opt *OrbitQueryRequest, options ...RequestOptionFunc) (*OrbitQueryResult, *Response, error) {
	return do[*OrbitQueryResult](s.client,
		withMethod(http.MethodPost),
		withPath("orbit/query"),
		withAPIOpts(opt),
		withRequestOpts(options...),
	)
}

// GetGraphStatusOptions represents the available GetGraphStatus()
// options. Exactly one of NamespaceID, ProjectID, or FullPath must be
// provided; the server returns 400 otherwise.
//
// GitLab API docs:
// https://docs.gitlab.com/api/orbit/#get-graph-status
type GetGraphStatusOptions struct {
	// NamespaceID is the ID of the namespace (group) to query.
	NamespaceID *int64 `url:"namespace_id,omitempty"`

	// ProjectID is the ID of the project to query.
	ProjectID *int64 `url:"project_id,omitempty"`

	// FullPath is the full path of a project or group
	// (e.g. "gitlab-org/gitlab").
	FullPath *string `url:"full_path,omitempty"`

	// ResponseFormat selects the response shape: "raw" (structured
	// JSON, default) or "llm" (compact text for LLM consumption).
	ResponseFormat *OrbitResponseFormatValue `url:"response_format,omitempty"`
}

// OrbitGraphStatus represents the indexing status response returned by
// `GET /api/v4/orbit/graph_status`.
//
// When ResponseFormat is "raw" (default), the structured fields
// (Projects, Domains, Indexing) are populated and FormattedText is
// empty. When ResponseFormat is "llm", only FormattedText is populated
// and the structured fields are absent.
//
// GitLab API docs:
// https://docs.gitlab.com/api/orbit/#get-graph-status
type OrbitGraphStatus struct {
	FormattedText string                    `json:"formatted_text,omitempty"`
	Projects      *OrbitGraphStatusProjects `json:"projects,omitempty"`
	Domains       []*OrbitGraphStatusDomain `json:"domains,omitempty"`
	Indexing      *OrbitGraphStatusIndexing `json:"indexing,omitempty"`
}

// OrbitGraphStatusProjects holds the indexed vs total project counts
// for the queried namespace or project.
//
// GitLab API docs:
// https://docs.gitlab.com/api/orbit/#get-graph-status
type OrbitGraphStatusProjects struct {
	Indexed    int64 `json:"indexed"`
	TotalKnown int64 `json:"total_known"`
}

// OrbitGraphStatusDomain represents a single domain (e.g. "SDLC") in
// the graph status response and its node-level item counts.
//
// GitLab API docs:
// https://docs.gitlab.com/api/orbit/#get-graph-status
type OrbitGraphStatusDomain struct {
	Name  string                        `json:"name,omitempty"`
	Items []*OrbitGraphStatusDomainItem `json:"items,omitempty"`
}

// OrbitGraphStatusDomainItem represents the count of a single node
// type within a domain.
//
// GitLab API docs:
// https://docs.gitlab.com/api/orbit/#get-graph-status
type OrbitGraphStatusDomainItem struct {
	Name  string `json:"name,omitempty"`
	Count int64  `json:"count"`
}

// OrbitGraphStatusIndexing holds the indexing pipeline state and
// timing information for the queried namespace or project.
//
// Possible State values: "not_indexed", "backfilling", "indexed",
// "error", "indexing", "unknown".
//
// GitLab API docs:
// https://docs.gitlab.com/api/orbit/#get-graph-status
type OrbitGraphStatusIndexing struct {
	State           string     `json:"state,omitempty"`
	LastStartedAt   *time.Time `json:"last_started_at,omitempty"`
	LastCompletedAt *time.Time `json:"last_completed_at,omitempty"`
	LastDurationMs  *int64     `json:"last_duration_ms,omitempty"`
	LastError       *string    `json:"last_error,omitempty"`
}

// GetGraphStatus returns the Knowledge Graph indexing status for a
// namespace or project from `GET /api/v4/orbit/graph_status`.
//
// Exactly one of opt.NamespaceID, opt.ProjectID, or opt.FullPath must
// be set; the server returns 400 when none or more than one is
// provided.
//
// Note: This API is experimental and may change or be removed in
// future versions.
//
// GitLab API docs:
// https://docs.gitlab.com/api/orbit/#get-graph-status
func (s *OrbitService) GetGraphStatus(opt *GetGraphStatusOptions, options ...RequestOptionFunc) (*OrbitGraphStatus, *Response, error) {
	return do[*OrbitGraphStatus](s.client,
		withMethod(http.MethodGet),
		withPath("orbit/graph_status"),
		withAPIOpts(opt),
		withRequestOpts(options...),
	)
}
