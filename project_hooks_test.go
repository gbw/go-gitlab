//
// Copyright 2021, Eric Stevens
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
	"github.com/stretchr/testify/require"
)

func TestListProjectHooks(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/hooks", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `[
	{
		"id": 1,
		"url": "http://example.com/hook",
		"name": "This is the name of an example hook",
		"description": "This is the description of an example hook",
		"confidential_note_events": true,
		"project_id": 1,
		"push_events": true,
		"push_events_branch_filter": "main",
		"issues_events": true,
		"confidential_issues_events": true,
		"merge_requests_events": true,
		"tag_push_events": true,
		"note_events": true,
		"job_events": true,
		"pipeline_events": true,
		"wiki_page_events": true,
		"deployment_events": true,
		"releases_events": true,
		"emoji_events": true,
		"enable_ssl_verification": true,
		"alert_status": "executable",
		"created_at": "2024-10-13T13:37:00Z",
		"resource_access_token_events": true,
		"resource_deploy_token_events": true,
		"custom_webhook_template": "my custom template",
		"custom_headers": [
			{"key": "Authorization"},
			{"key": "OtherHeader"}
		]
	}
]`)
	})

	hooks, resp, err := client.Projects.ListProjectHooks(1, nil)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	createdAt := time.Date(2024, time.October, 13, 13, 37, 0, 0, time.UTC)
	want := []*ProjectHook{{
		ID:                        1,
		URL:                       "http://example.com/hook",
		Name:                      "This is the name of an example hook",
		Description:               "This is the description of an example hook",
		ConfidentialNoteEvents:    true,
		ProjectID:                 1,
		PushEvents:                true,
		PushEventsBranchFilter:    "main",
		IssuesEvents:              true,
		ConfidentialIssuesEvents:  true,
		MergeRequestsEvents:       true,
		TagPushEvents:             true,
		NoteEvents:                true,
		JobEvents:                 true,
		PipelineEvents:            true,
		WikiPageEvents:            true,
		DeploymentEvents:          true,
		ReleasesEvents:            true,
		EmojiEvents:               true,
		EnableSSLVerification:     true,
		CreatedAt:                 &createdAt,
		AlertStatus:               "executable",
		ResourceAccessTokenEvents: true,
		ResourceDeployTokenEvents: true,
		CustomWebhookTemplate:     "my custom template",
		CustomHeaders: []*HookCustomHeader{
			{
				Key: "Authorization",
			},
			{
				Key: "OtherHeader",
			},
		},
	}}

	assert.Equal(t, want, hooks)
}

func TestAddProjectHook(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/hooks", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		testBodyJSON(t, r, map[string]any{
			"url":                          "http://example.com/hook",
			"emoji_events":                 true,
			"signing_token":                "whsec_dGVzdA==",
			"feature_flag_events":          true,
			"milestone_events":             true,
			"resource_deploy_token_events": true,
		})
		fmt.Fprint(w, `
{
	"id": 1,
	"url": "http://example.com/hook",
	"project_id": 1,
	"push_events": true,
	"issues_events": true,
	"merge_requests_events": true,
	"tag_push_events": true,
	"emoji_events": true,
	"enable_ssl_verification": true,
	"resource_deploy_token_events": true,
	"created_at": "2024-10-13T13:37:00Z"
}`)
	})

	opt := &AddProjectHookOptions{
		URL:                       Ptr("http://example.com/hook"),
		EmojiEvents:               Ptr(true),
		SigningToken:              Ptr("whsec_dGVzdA=="),
		FeatureFlagEvents:         Ptr(true),
		MilestoneEvents:           Ptr(true),
		ResourceDeployTokenEvents: Ptr(true),
	}

	hook, resp, err := client.Projects.AddProjectHook(1, opt)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	createdAt := time.Date(2024, time.October, 13, 13, 37, 0, 0, time.UTC)
	want := &ProjectHook{
		ID:                        1,
		URL:                       "http://example.com/hook",
		ProjectID:                 1,
		PushEvents:                true,
		IssuesEvents:              true,
		MergeRequestsEvents:       true,
		TagPushEvents:             true,
		EmojiEvents:               true,
		EnableSSLVerification:     true,
		ResourceDeployTokenEvents: true,
		CreatedAt:                 &createdAt,
	}

	assert.Equal(t, want, hook)
}

func TestEditProjectHook(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/hooks/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		testBodyJSON(t, r, map[string]any{
			"push_events":                  false,
			"emoji_events":                 true,
			"signing_token":                "whsec_dGVzdA==",
			"feature_flag_events":          true,
			"milestone_events":             true,
			"resource_deploy_token_events": true,
		})
		fmt.Fprint(w, `
{
	"id": 1,
	"url": "http://example.com/hook",
	"project_id": 1,
	"push_events": false,
	"issues_events": true,
	"merge_requests_events": true,
	"tag_push_events": true,
	"emoji_events": true,
	"enable_ssl_verification": true,
	"resource_deploy_token_events": true,
	"created_at": "2024-10-13T13:37:00Z"
}`)
	})

	opt := &EditProjectHookOptions{
		PushEvents:                Ptr(false),
		EmojiEvents:               Ptr(true),
		SigningToken:              Ptr("whsec_dGVzdA=="),
		FeatureFlagEvents:         Ptr(true),
		MilestoneEvents:           Ptr(true),
		ResourceDeployTokenEvents: Ptr(true),
	}

	hook, resp, err := client.Projects.EditProjectHook(1, 1, opt)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	createdAt := time.Date(2024, time.October, 13, 13, 37, 0, 0, time.UTC)
	want := &ProjectHook{
		ID:                        1,
		URL:                       "http://example.com/hook",
		ProjectID:                 1,
		PushEvents:                false,
		IssuesEvents:              true,
		MergeRequestsEvents:       true,
		TagPushEvents:             true,
		EmojiEvents:               true,
		EnableSSLVerification:     true,
		ResourceDeployTokenEvents: true,
		CreatedAt:                 &createdAt,
	}

	assert.Equal(t, want, hook)
}

// Test that the "CustomWebhookTemplate" serializes properly
func TestAddProjectHook_CustomTemplateStuff(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/hooks",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPost)
			testBodyJSON(t, r, map[string]any{
				"custom_webhook_template": `{"example":"{{object_kind}}"}`,
				"custom_headers": []any{
					map[string]any{
						"key":   "Authorization",
						"value": "stuff",
					},
					map[string]any{
						"key":   "Favorite-Pet",
						"value": "Cats",
					},
				},
			})

			w.WriteHeader(http.StatusCreated)

			fmt.Fprint(w, `{
				"custom_webhook_template": "testValue",
				"custom_headers": [
					{
						"key": "Authorization"
					},
					{
						"key": "Favorite-Pet"
					}
				]
			}`)
		},
	)

	hook, resp, err := client.Projects.AddProjectHook(1, &AddProjectHookOptions{
		CustomWebhookTemplate: Ptr(`{"example":"{{object_kind}}"}`),
		CustomHeaders: &[]*HookCustomHeader{
			{
				Key:   "Authorization",
				Value: "stuff",
			},
			{
				Key:   "Favorite-Pet",
				Value: "Cats",
			},
		},
	})

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	assert.Equal(t, "testValue", hook.CustomWebhookTemplate)
	assert.Len(t, hook.CustomHeaders, 2)
}

// Test that the "CustomWebhookTemplate" serializes properly when editing
func TestEditProjectHook_CustomTemplateStuff(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/hooks/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodPut)
			testBodyJSON(t, r, map[string]any{
				"custom_webhook_template": `{"example":"{{object_kind}}"}`,
				"custom_headers": []any{
					map[string]any{
						"key":   "Authorization",
						"value": "stuff",
					},
					map[string]any{
						"key":   "Favorite-Pet",
						"value": "Cats",
					},
				},
			})

			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, `{
				"custom_webhook_template": "testValue",
				"custom_headers": [
					{
						"key": "Authorization"
					},
					{
						"key": "Favorite-Pet"
					}
				]}`)
		},
	)

	hook, resp, err := client.Projects.EditProjectHook(1, 1, &EditProjectHookOptions{
		CustomWebhookTemplate: Ptr(`{"example":"{{object_kind}}"}`),
		CustomHeaders: &[]*HookCustomHeader{
			{
				Key:   "Authorization",
				Value: "stuff",
			},
			{
				Key:   "Favorite-Pet",
				Value: "Cats",
			},
		},
	})

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "testValue", hook.CustomWebhookTemplate)
	assert.Len(t, hook.CustomHeaders, 2)
}

func TestDeleteProjectHook(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/hooks/1",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodDelete)
			w.WriteHeader(http.StatusNoContent)
		},
	)

	resp, err := client.Projects.DeleteProjectHook(1, 1)

	assert.NoError(t, err)

	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestTriggerTestProjectHook(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/hooks/1/test/push_events", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{"message":"201 Created"}`)
	})

	mux.HandleFunc("/api/v4/projects/1/hooks/1/test/invalid_trigger", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"error": "trigger does not have a valid value"}`)
	})

	tests := []struct {
		name       string
		projectID  any
		hookID     int64
		event      ProjectHookEvent
		wantErr    bool
		wantStatus int
		wantErrMsg string
	}{
		{
			name:       "Valid trigger",
			projectID:  1,
			hookID:     1,
			event:      ProjectHookEventPush,
			wantErr:    false,
			wantStatus: http.StatusCreated,
		},
		{
			name:       "Invalid project ID",
			projectID:  "invalid",
			hookID:     1,
			event:      ProjectHookEventPush,
			wantErr:    true,
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "Invalid trigger type",
			projectID:  1,
			hookID:     1,
			event:      "invalid_trigger",
			wantErr:    true,
			wantStatus: http.StatusBadRequest,
			wantErrMsg: "trigger does not have a valid value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			resp, err := client.Projects.TriggerTestProjectHook(tt.projectID, tt.hookID, tt.event)

			if tt.wantErr {
				require.Error(t, err)
				if tt.wantStatus != 0 {
					require.Equal(t, tt.wantStatus, resp.StatusCode)
				}
				if tt.wantErrMsg != "" {
					require.Contains(t, err.Error(), tt.wantErrMsg)
				}
			} else {
				require.NoError(t, err)
				require.NotNil(t, resp)
				require.Equal(t, tt.wantStatus, resp.StatusCode)
			}
		})
	}
}

func TestSetProjectWebhookURLVariable(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/hooks/2/url_variables/TEST_KEY", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
	})

	opt := &SetProjectWebhookURLVariableOptions{
		Value: Ptr("testValue"),
	}
	resp, err := client.Projects.SetProjectWebhookURLVariable(1, 2, "TEST_KEY", opt)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestDeleteProjectWebhookURLVariable(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	mux.HandleFunc("/api/v4/projects/1/hooks/2/url_variables/TEST_KEY", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
	})

	resp, err := client.Projects.DeleteProjectWebhookURLVariable(1, 2, "TEST_KEY")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func TestGetProjectWebhookHeader(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// Removed most of the arguments to keep test slim
	mux.HandleFunc("/api/v4/projects/1/hooks/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		fmt.Fprint(w, `{
			"id": 1,
			"custom_webhook_template": "{\"event\":\"{{object_kind}}\"}",
			"token_present": true,
			"signing_token_present": true,
			"custom_headers": [
			  {
				"key": "Authorization"
			  },
			  {
				"key": "OtherKey"
			  }
			]
		  }`)
	})

	hook, resp, err := client.Projects.GetProjectHook(1, 1)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	want := &ProjectHook{
		ID:                    1,
		CustomWebhookTemplate: "{\"event\":\"{{object_kind}}\"}",
		TokenPresent:          true,
		SigningTokenPresent:   true,
		CustomHeaders: []*HookCustomHeader{
			{
				Key: "Authorization",
			},
			{
				Key: "OtherKey",
			},
		},
	}

	assert.Equal(t, want, hook)
}

func TestSetProjectWebhookHeader(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// Removed most of the arguments to keep test slim
	mux.HandleFunc("/api/v4/projects/1/hooks/1/custom_headers/Authorization", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPut)
		testBodyJSON(t, r, map[string]any{
			"value": "testValue",
		})

		w.WriteHeader(http.StatusNoContent)

		fmt.Fprint(w, ``)
	})

	resp, err := client.Projects.SetProjectCustomHeader(1, 1, "Authorization", &SetHookCustomHeaderOptions{Value: Ptr("testValue")})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestDeleteProjectWebhookHeader(t *testing.T) {
	t.Parallel()
	mux, client := setup(t)

	// Removed most of the arguments to keep test slim
	mux.HandleFunc("/api/v4/projects/1/hooks/1/custom_headers/Authorization", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodDelete)
		w.WriteHeader(http.StatusNoContent)
		fmt.Fprint(w, ``)
	})

	resp, err := client.Projects.DeleteProjectCustomHeader(1, 1, "Authorization")
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}
